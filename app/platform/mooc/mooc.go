package mooc

import (
	"encoding/json"
	"github.com/Esword618/MoocDownload/app/platform/mooc/model"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 检测 cookie
func (m *Mooc) CheckStatus() bool {
	var StatusStruct model.Status
	m.Session.Headers.Add("edu-script-token", m.Token)
	_, err := m.Session.R().SetQueryParams(map[string]string{
		"csrfKey": m.Token,
	}).SetFormData(map[string]string{
		"memberId": m.MemberId,
	}).SetSuccessResult(&StatusStruct).Post("https://www.icourse163.org/web/j/memberBean.getMocMemberPersonalDtoById.rpc")
	m.Session.Headers.Del("edu-script-token")
	log.Println(StatusStruct)
	if err != nil {
		return false
	}
	if StatusStruct.Code == 0 {
		//color.Cyan.Printf("\n欢迎您:%s\n", StatusStruct.Result.NickName)
		if len(StatusStruct.Result.Description) != 0 {
			//color.Blue.Printf("\n你的座右铭:%s\n", StatusStruct.Result.Description)
		}
		return true
	} else {
		//color.Red.Println("\ncookie已失效，请更新您的cookie!\n")
		return false
	}
}

// 获取课程内容
func (m *Mooc) getLastLearnedMocTermDto(tid string) string {
	m.Session.Headers.Add("edu-script-token", m.Token)
	res, err := m.Session.R().SetQueryParams(map[string]string{
		"csrfKey": m.Token,
	}).SetFormData(map[string]string{
		"termId": tid,
	}).Post("https://www.icourse163.org/web/j/courseBean.getLastLearnedMocTermDto.rpc")
	m.Session.Headers.Del("edu-script-token")
	if err != nil {
		panic(err)
	}
	jsonStr := res.String()
	return jsonStr
}

func (m *Mooc) getSignatureVideoId(UnitId int, contentType string) (int, string) {
	res, _ := m.Session.R().SetQueryParams(map[string]string{
		"csrfKey": m.Token,
	}).SetFormData(map[string]string{
		"bizId":       strconv.Itoa(UnitId),
		"bizType":     "1",
		"contentType": contentType,
	}).Post("https://www.icourse163.org/web/j/resourceRpcBean.getResourceToken.rpc")
	var VideoStruct model.Video
	var err = json.Unmarshal([]byte(res.String()), &VideoStruct)
	if err != nil {
		panic(err)
	}
	signature := VideoStruct.Result.VideoSignDto.Signature
	videoId := VideoStruct.Result.VideoSignDto.VideoID
	return videoId, signature
}

// 视频下载
func (m *Mooc) video(Unit model.UnitStruct, chapterNamePath, contentType string) {
	var VodVideoStruct model.VodVideo
	videoId, signature := m.getSignatureVideoId(Unit.UnitId, contentType)
	_, _ = m.Session.R().SetQueryParams(map[string]string{
		"videoId":    strconv.Itoa(videoId),
		"signature":  signature,
		"clientType": "1",
	}).SetSuccessResult(&VodVideoStruct).Get("https://vod.study.163.com/eds/api/v1/vod/video")

	var videoUrl string
	var k string
	var secondaryEncrypt bool
	var format string
	Count := len(VodVideoStruct.Result.Videos)
	if Count%3 == 0 {
		videos := VodVideoStruct.Result.Videos
		video := videos[2]
		format = video.Format
		videoUrl = video.VideoURL
		k = video.K
		secondaryEncrypt = video.SecondaryEncrypt
	} else {
		videos := VodVideoStruct.Result.Videos
		video := videos[Count%3-1]
		format = video.Format
		videoUrl = video.VideoURL
		k = video.K
		secondaryEncrypt = video.SecondaryEncrypt
	}
	if secondaryEncrypt && contentType == "1" {
		//fmt.Println(Unit.UnitName)
		videoToken, _ := m.Md.MJ.TokenCrypt(k, 1)
		res1, _ := m.Session.R().SetQueryParams(map[string]string{
			"token": videoToken,
			"t":     strconv.FormatInt(time.Now().UnixMilli(), 10),
		}).Get(videoUrl)

		tsList, KeyByte, _ := m.Md.VipGetTsKey(res1.String(), videoId, contentType)

		// secondaryEncrypt 为 true 代表 key也进行了加密
		m.Md.VipVideo1(tsList, KeyByte, chapterNamePath, Unit)
	} else if contentType == "7" {
		videoToken, _ := m.Md.MJ.TokenCrypt(k, 1)
		res1, _ := m.Session.R().SetQueryParams(map[string]string{
			"token": videoToken,
			"t":     strconv.FormatInt(time.Now().UnixMilli(), 10),
		}).Get(videoUrl)
		re := regexp.MustCompile("(http.*/)")
		baseUrl := re.FindStringSubmatch(videoUrl)[1]
		if !strings.Contains(videoUrl, "https") {
			baseUrl = strings.Replace(baseUrl, "http", "https", -1)
		}
		tsList, KeyByte, IV := m.Md.VipGetTsKey(res1.String(), videoId, contentType)
		for i, j := range tsList {
			tsList[i] = baseUrl + j
		}
		m.Md.VipVideo7(tsList, KeyByte, IV, contentType, chapterNamePath, Unit)
	} else {
		switch format {
		case "hls":
			res0, _ := m.Session.R().Get(videoUrl)
			re := regexp.MustCompile("(http.*/)")
			baseUrl := re.FindStringSubmatch(videoUrl)[1]
			if !strings.Contains(videoUrl, "https") {
				baseUrl = strings.Replace(baseUrl, "http", "https", -1)
			}
			tsList := m.Md.FreeGetTs(res0.String())
			for i, j := range tsList {
				tsList[i] = baseUrl + j
			}
			m.Md.FreeVideo(tsList, chapterNamePath, Unit)
		case "mp4":
			m.Md.Mp4Flv(videoUrl, chapterNamePath, false, Unit)
		case "flv":
			m.Md.Mp4Flv(videoUrl, chapterNamePath, true, Unit)
		}
	}
}

// 文本资料下载
func (m *Mooc) courseware(Unit model.UnitStruct, chapterNamePath string) {
	res, _ := m.Session.R().SetFormData(map[string]string{
		"callCount":       "1",
		"scriptSessionId": "${scriptSessionId}190",
		"httpSessionId":   m.Token,
		"c0-scriptName":   "CourseBean",
		"c0-methodName":   "getLessonUnitLearnVo",
		"c0-id":           "0",
		"c0-param0":       "number:" + strconv.Itoa(Unit.ContentId),
		"c0-param1":       "number:3",
		"c0-param2":       "number:0",
		"c0-param3":       "number:" + strconv.Itoa(Unit.UnitId),
		"batchId":         strconv.Itoa(int(time.Now().UnixMilli())),
	}).Post("https://www.icourse163.org/dwr/call/plaincall/CourseBean.getLessonUnitLearnVo.dwr")
	cmp := regexp.MustCompile("textUrl:\"(http.*?)\"")
	textUrl := cmp.FindAllStringSubmatch(res.String(), 1)[0][1]
	length := len(Unit.UnitName)
	Unit.UnitName = Unit.UnitName[:length-4]
	m.Md.Text(textUrl, Unit, chapterNamePath)
}
