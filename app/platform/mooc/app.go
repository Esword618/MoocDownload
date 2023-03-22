package mooc

import (
	"encoding/json"
	"fmt"
	"github.com/Esword618/MoocDownload/app/global"
	"github.com/Esword618/MoocDownload/app/platform/mooc/download"
	"github.com/Esword618/MoocDownload/app/platform/mooc/model"
	"github.com/Esword618/MoocDownload/app/utils"
	"github.com/Esword618/MoocDownload/config"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"github.com/sourcegraph/conc/pool"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"regexp"
)

type Mooc struct {
	Session    *req.Client
	Cookie     string
	Token      string
	MemberId   string
	Tid        string
	InfoStruct model.CourseStruct
	Md         *download.MoocDownload
	// 下载信息
	DownloadInfo model.CourseStruct
}

func NewMooc() *Mooc {
	md := download.NewMD()
	return &Mooc{
		Md: md,
	}
}

// ParseLink 解析链接
func (m *Mooc) ParseLink(link string) model.CourseStruct {
	re := regexp.MustCompile("tid=(\\d+)")
	tid := re.FindStringSubmatch(link)[1]
	m.Tid = tid
	jsonStr := m.getLastLearnedMocTermDto(tid)
	CourseStruct := handleJsonStr(jsonStr)
	m.InfoStruct = CourseStruct
	runtime.EventsEmit(global.GlobalCtx, "parse_link_result", CourseStruct)
	// 打印json
	//body, err := json.MarshalIndent(CourseStruct, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", body)
	//fmt.Println("--------------------------------")
	return CourseStruct
}

// Download 下载
func (m *Mooc) Download(downloadInfo model.CourseStruct) {
	p := pool.New().WithMaxGoroutines(config.BatchSize)
	for _, unit := range downloadInfo.Units {
		filePath := fmt.Sprintf("%s\\%s\\%s\\%s", downloadInfo.CourseName, unit.ChapterName, unit.LessonName, unit.UnitName)
		exist, _ := utils.PathExist(filePath)
		if exist {
			unit.Progress.Percentage = 100
			unit.Progress.Name = unit.UnitName
			unit.Progress.Uuid = unit.Uuid
			unit.Progress.Status = "warning"
			runtime.EventsEmit(global.GlobalCtx, unit.Uuid, unit.Progress)
		}
		if !exist {
			unitContentType := unit.UnitContentType
			var path string
			switch unitContentType {
			case 1:
				p.Go(func() {
					// 加密 and 不加密 视频
					path = fmt.Sprintf("%s\\%s\\%s", downloadInfo.CourseName, unit.ChapterName, unit.LessonName)
					m.video(unit, path, "1")
				})
			case 3:
				p.Go(func() {
					// 资料
					path = fmt.Sprintf("%s\\%s\\%s", downloadInfo.CourseName, unit.ChapterName, unit.LessonName)
					m.courseware(unit, path)
				})
			case 5:
				// 试卷
				path = fmt.Sprintf("%s\\%s\\%s\\%s.pdf", downloadInfo.CourseName, unit.ChapterName, unit.LessonName, unit.UnitName)
			case 7:
				p.Go(func() {
					// 加密视频
					path = fmt.Sprintf("%s\\%s\\%s", downloadInfo.CourseName, unit.ChapterName, unit.LessonName)
					m.video(unit, path, "7")
				})
			}
		}
	}
	p.Wait()
}

func (m *Mooc) UpdateCookie() {
	cookieStr := viper.GetString("mooc.cookie")
	CookieMap := utils.CookieToMap(cookieStr)
	Token := CookieMap["NTESSTUDYSI"]
	NeteaseWdaUid := CookieMap["NETEASE_WDA_UID"]
	cmp := regexp.MustCompile("(\\d+)#")
	MemberId := cmp.FindStringSubmatch(NeteaseWdaUid)[1]
	// 调试
	//session := req.C().EnableDumpAll()
	session := req.C()
	session.SetCommonHeaders(map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36",
		"origin":     "https://www.icourse163.org",
		"cookie":     cookieStr,
	})
	m.Token = Token
	m.MemberId = MemberId
	m.Session = session
	m.Cookie = cookieStr
}

// 处理下载的信息
func (m *Mooc) HandledownloadInfo(uuidLIst []string) model.CourseStruct {
	var downloadInfo model.CourseStruct
	downloadInfo.CourseName = m.InfoStruct.CourseName
	for _, unit := range m.InfoStruct.Units {
		b := utils.StringInArray(unit.Uuid, uuidLIst)
		if b {
			downloadInfo.Units = append(downloadInfo.Units, unit)
		}
	}
	//m.DownloadInfo = downloadInfo
	return downloadInfo
	//body, err := json.MarshalIndent(downloadInfo, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", body)
	//fmt.Println("--------------------------------")
}

// handleJsonStr 处理数据
func handleJsonStr(jsonStr string) model.CourseStruct {
	var STRUT model.LastLearnedMocTermDto
	var InfoStruct model.CourseStruct
	var uuidList []string
	var err = json.Unmarshal([]byte(jsonStr), &STRUT)
	if err != nil {
		panic(err)
	}
	courseName := STRUT.Result.MocTermDto.CourseName
	courseName = utils.RemoveInvalidChar(courseName)

	InfoStruct.CourseName = courseName

	chapters := STRUT.Result.MocTermDto.Chapters
	for _, chapter := range chapters {

		var myUnit model.UnitStruct

		chapterName := chapter.Name
		chapterName = utils.RemoveInvalidChar(chapterName)

		myUnit.ChapterName = chapterName

		ChapterContentType := chapter.ContentType
		myUnit.ChapterContentType = ChapterContentType

		if ChapterContentType == 2 {
			// 有 bug
			myUnit.LessonName = chapterName
			myUnit.UnitName = chapterName + ".docx"
			myUnit.ContentId = 0
			myUnit.UnitId = chapter.ID
			myUnit.UnitContentType = ChapterContentType
			// 生成基于 时间 的uuid
			u1, _ := uuid.NewUUID()
			myUnit.Uuid = u1.String()
			uuidList = append(uuidList, u1.String())
			InfoStruct.Units = append(InfoStruct.Units, myUnit)
		} else {

			lessons := chapter.Lessons

			for _, lesson := range lessons {

				LessonName := lesson.Name
				LessonName = utils.RemoveInvalidChar(LessonName)
				myUnit.LessonName = LessonName

				units := lesson.Units

				for _, unit := range units {
					unitContentType := unit.ContentType
					UnitName := unit.Name
					UnitName = utils.RemoveInvalidChar(UnitName)

					//myUnit.UnitName = UnitName
					myUnit.UnitContentType = unitContentType

					switch unitContentType {
					case 1:
						myUnit.UnitName = UnitName + ".mp4"
						myUnit.ContentId = 0
						myUnit.UnitId = unit.ID
					case 3:
						myUnit.UnitName = UnitName + ".pdf"
						myUnit.ContentId = unit.ContentID
						myUnit.UnitId = unit.ID
					case 5:
						myUnit.UnitName = UnitName + ".docx"
						myUnit.ContentId = unit.ContentID
						myUnit.UnitId = 0
					case 7:
						myUnit.UnitName = UnitName + ".mp4"
						myUnit.ContentId = 0
						myUnit.UnitId = unit.ID
					}
					// 生成基于 时间 的uuid
					u1, _ := uuid.NewUUID()
					myUnit.Uuid = u1.String()
					uuidList = append(uuidList, u1.String())
					InfoStruct.Units = append(InfoStruct.Units, myUnit)
				}
			}
		}
	}
	InfoStruct.UuidList = uuidList
	return InfoStruct
}
