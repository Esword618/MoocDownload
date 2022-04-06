package mooc

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
	"github.com/spf13/viper"

	"MoocDownload/mooc/model"
	"MoocDownload/mooc/utils"
)

var moocSession *MoocSession

// https://www.icourse163.org/learn/kaopei-1003293002?tid=1463199446#/learn/announce
// https://www.icourse163.org/learn/HIT-1002533005?tid=1467082464#/learn/announce
// Link = "https://www.icourse163.org/learn/HIT-1002533005?tid=1467082464#/learn/announce"

func MoocMain() {
	cookieStr := utils.ReadCookie()
	color.Red.Printf("\n请把链接粘贴到处:")
	var Link string
	fmt.Scanln(&Link)
	Client := resty.New()
	re := regexp.MustCompile("tid=(\\d+)")
	tid := re.FindStringSubmatch(Link)[1]
	CookieMap := utils.CookieToMap(cookieStr)
	token := CookieMap["NTESSTUDYSI"]
	NETEASE_WDA_UID := CookieMap["NETEASE_WDA_UID"]
	cmp := regexp.MustCompile("(\\d+)#")
	memberId := cmp.FindStringSubmatch(NETEASE_WDA_UID)[1]
	Client.SetHeaders(map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36",
		"origin":     "https://www.icourse163.org",
		"cookie":     cookieStr,
	})
	moocSession = &MoocSession{
		Session:  Client,
		Tid:      tid,
		Token:    token,
		MemberId: memberId,
	}
	B := moocSession.CheckStatus()
	if !B {
		color.Red.Println("cookie 已失效,请重新登录！")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}
	jsonStr := moocSession.GetLastLearnedMocTermDto(tid)
	InfoStruct := utils.HandleJsonStr(jsonStr)
	_download(InfoStruct)
}

func _download(InfoStruct model.MyMocTermDto) {
	basePath, _ := os.Getwd()
	courseName := InfoStruct.CourseName
	courseName = utils.RemoveInvalidChar(courseName)
	courseNamePath := fmt.Sprintf("%s\\download\\%s", basePath, courseName)
	utils.PathExists(courseNamePath)
	videoBool := viper.GetInt("download.video")
	coursewareBool := viper.GetInt("download.courseware")
	paperBool := viper.GetInt("download.paper")
	for _, chapter := range InfoStruct.Chapters {
		chapterName := chapter.ChapterName
		chapterName = utils.RemoveInvalidChar(chapterName)
		chapterNamePath := fmt.Sprintf("%s\\%s", courseNamePath, chapterName)
		utils.PathExists(chapterNamePath)
		temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
		utils.PathExists(temPath)
		for _, unit := range chapter.MyUnits {
			contentType := unit.ContentType
			UnitId := unit.UnitId
			ContentId := unit.ContentId
			unitName := unit.UnitName
			unitName = utils.RemoveInvalidChar(unitName)
			switch contentType {
			case 1:
				if videoBool == 1 {
					path := fmt.Sprintf("%s\\%s.mp4", chapterNamePath, unitName)
					_, err := os.Stat(path)
					if err != nil {
						moocSession.Video(UnitId, unitName, chapterNamePath, "1")
					}
				}

			case 3:
				if coursewareBool == 1 {
					path := fmt.Sprintf("%s\\%s.pdf", chapterNamePath, unitName)
					_, err := os.Stat(path)
					if err != nil {
						moocSession.Courseware(ContentId, UnitId, path)
					}
				}
			case 5:
				if paperBool == 1 {

				}
			case 7:
				if videoBool == 1 {
					path := fmt.Sprintf("%s\\%s.mp4", chapterNamePath, unitName)
					_, err := os.Stat(path)
					if err != nil {
						moocSession.Video(UnitId, unitName, chapterNamePath, "7")
					}
				}
			}
		}
		err := os.RemoveAll(temPath)
		if err != nil {
			panic(err)
		}
	}
}
