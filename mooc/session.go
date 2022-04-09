package mooc

import (
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"

	"MoocDownload/mooc/download"
	"MoocDownload/mooc/js"
	"MoocDownload/mooc/model"
)

type MoocSession struct {
	Session  *resty.Client
	Cookie   string
	Token    string
	MemberId string
	Tid      string
}

// 检查当前cookie状态
func (This *MoocSession) CheckStatus() bool {
	var StatusStruct model.Status
	This.Session.Header.Add("edu-script-token", This.Token)
	_, err := This.Session.R().SetQueryParams(map[string]string{
		"csrfKey": This.Token,
	}).SetFormData(map[string]string{
		"memberId": This.MemberId,
	}).SetResult(&StatusStruct).Post("https://www.icourse163.org/web/j/memberBean.getMocMemberPersonalDtoById.rpc")
	This.Session.Header.Del("edu-script-token")
	if err != nil {
		return false
	}
	if StatusStruct.Code == 0 {
		color.Cyan.Printf("\n欢迎您:%s\n", StatusStruct.Result.NickName)
		if len(StatusStruct.Result.Description) != 0 {
			color.Blue.Printf("\n你的座右铭:%s\n", StatusStruct.Result.Description)
		}
		return true
	} else {
		color.Red.Println("\ncookie已失效，请更新您的cookie!\n")
		time.Sleep(3 * time.Second)
		os.Exit(1)
		return false
	}
}

func (This *MoocSession) GetLastLearnedMocTermDto(tid string) string {
	This.Session.Header.Add("edu-script-token", This.Token)
	res, err := This.Session.R().SetQueryParams(map[string]string{
		"csrfKey": This.Token,
	}).SetFormData(map[string]string{
		"termId": tid,
	}).Post("https://www.icourse163.org/web/j/courseBean.getLastLearnedMocTermDto.rpc")
	This.Session.Header.Del("edu-script-token")
	if err != nil {
		panic(err)
	}
	jsonStr := res.String()
	return jsonStr
}

func (This *MoocSession) GetSignatureVideoId(UnitId int, contentType string) (int, string) {
	res, _ := This.Session.R().SetQueryParams(map[string]string{
		"csrfKey": This.Token,
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
func (This *MoocSession) Video(UnitId int, unitName, chapterNamePath, contentType string) {
	var VodVideoStruct model.VodVideo
	videoId, signature := This.GetSignatureVideoId(UnitId, contentType)
	_, _ = This.Session.R().SetQueryParams(map[string]string{
		"videoId":    strconv.Itoa(videoId),
		"signature":  signature,
		"clientType": "1",
	}).SetResult(&VodVideoStruct).Get("https://vod.study.163.com/eds/api/v1/vod/video")

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
		videoToken := js.Token(k)
		res1, _ := This.Session.R().SetQueryParams(map[string]string{
			"token": videoToken,
			"t":     strconv.FormatInt(time.Now().UnixMilli(), 10),
		}).Get(videoUrl)
		tsList, KeyByte, _ := download.VipGetTsKey(res1.String(), videoId, contentType)
		// secondaryEncrypt 为 true 代表 key也进行了加密
		download.VipVideo1(tsList, KeyByte, unitName, chapterNamePath)
	} else if contentType == "7" {
		videoToken := js.Token(k)
		res1, _ := This.Session.R().SetQueryParams(map[string]string{
			"token": videoToken,
			"t":     strconv.FormatInt(time.Now().UnixMilli(), 10),
		}).Get(videoUrl)
		re := regexp.MustCompile("(http.*/)")
		baseUrl := re.FindStringSubmatch(videoUrl)[1]
		if !strings.Contains(videoUrl, "https") {
			baseUrl = strings.Replace(baseUrl, "http", "https", -1)
		}
		tsList, KeyByte, IV := download.VipGetTsKey(res1.String(), videoId, contentType)
		for i, j := range tsList {
			tsList[i] = baseUrl + j
		}
		download.VipVideo7(tsList, KeyByte, IV, contentType, unitName, chapterNamePath)
	} else {
		switch format {
		case "hls":
			res0, _ := This.Session.R().Get(videoUrl)
			re := regexp.MustCompile("(http.*/)")
			baseUrl := re.FindStringSubmatch(videoUrl)[1]
			if !strings.Contains(videoUrl, "https") {
				baseUrl = strings.Replace(baseUrl, "http", "https", -1)
			}
			tsList := download.FreeGetTs(res0.String())
			for i, j := range tsList {
				tsList[i] = baseUrl + j
			}
			download.FreeVideo(tsList, unitName, chapterNamePath)
		case "mp4":
			download.Mp4Flv(videoUrl, unitName, chapterNamePath, false)
		case "flv":
			download.Mp4Flv(videoUrl, unitName, chapterNamePath, true)
		}

	}
}

// 文本资料下载
func (This *MoocSession) Courseware(ContentId int, UnitId int, path string) {
	res, _ := This.Session.R().SetFormData(map[string]string{
		"callCount":       "1",
		"scriptSessionId": "${scriptSessionId}190",
		"httpSessionId":   This.Token,
		"c0-scriptName":   "CourseBean",
		"c0-methodName":   "getLessonUnitLearnVo",
		"c0-id":           "0",
		"c0-param0":       "number:" + strconv.Itoa(ContentId),
		"c0-param1":       "number:3",
		"c0-param2":       "number:0",
		"c0-param3":       "number:" + strconv.Itoa(UnitId),
		"batchId":         strconv.Itoa(int(time.Now().UnixMilli())),
	}).Post("https://www.icourse163.org/dwr/call/plaincall/CourseBean.getLessonUnitLearnVo.dwr")
	cmp := regexp.MustCompile("textUrl:\"(http.*?)\"")
	textUrl := cmp.FindAllStringSubmatch(res.String(), 1)[0][1]
	download.Text(textUrl, path)
}
