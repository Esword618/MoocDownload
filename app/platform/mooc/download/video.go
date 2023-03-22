package download

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/Esword618/MoocDownload/app/global"
	"github.com/Esword618/MoocDownload/app/platform/mooc/js"
	"github.com/Esword618/MoocDownload/app/platform/mooc/model"
	"github.com/Esword618/MoocDownload/app/utils"
	"github.com/Esword618/MoocDownload/app/video"
	"github.com/Esword618/MoocDownload/config"
	"github.com/Esword618/MoocDownload/crypt"
	"github.com/imroc/req/v3"
	"github.com/panjf2000/ants/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
	"os"
	"regexp"
	"strconv"
	"sync"
)

type MoocDownload struct {
	Ctx    context.Context
	Client *req.Client
	MJ     *js.MoocJs
}

func NewMD() *MoocDownload {
	//client := req.C().EnableDumpAll()
	client := req.C()
	client.SetCommonHeaders(map[string]string{
		"origin":     "https://www.icourse163.org",
		"referer":    "https://www.icourse163.org/",
		"authority":  "mooc2vod.stu.126.net",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
	})
	mj := js.NewMoocJs()
	return &MoocDownload{
		Client: client,
		MJ:     mj,
	}
}

// VipVideo1 付费视频
func (md *MoocDownload) VipVideo1(TsList []string, KeyByte []byte, chapterNamePath string, Unit model.UnitStruct) {
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	utils.PathExistsAndMK(temPath)
	var wg *sync.WaitGroup
	wg = new(sync.WaitGroup)
	pool, _ := ants.NewPool(config.PoolSize)
	defer pool.Release()
	total := len(TsList)
	wg.Add(total)
	for index, TsUrl := range TsList {
		_ = pool.Submit(md.VipDecryptTs1(KeyByte, index, chapterNamePath, TsUrl, wg))

		// 进度
		//runtime.LogDebug(global.GlobalCtx, TsUrl)
		percentage := int(100 * (float64(index+1) / float64(total)))
		Unit.Progress.Percentage = percentage
		Unit.Progress.Name = Unit.UnitName
		Unit.Progress.Uuid = Unit.Uuid
		Unit.Progress.Status = "warning"
		runtime.EventsEmit(global.GlobalCtx, Unit.Uuid, Unit.Progress)
		//fmt.Println(Unit.Uuid, Progress)
	}
	wg.Wait()
	length := len(Unit.UnitName)
	Unit.UnitName = Unit.UnitName[:length-4]
	video.MergeTsFileListToSingleMp4(len(TsList), Unit.UnitName, chapterNamePath)
	//fmt.Println(Unit.UnitName + ".mp4 done\n")
}

// VipDecryptTs1 付费视频 contentType = 1
func (md *MoocDownload) VipDecryptTs1(KeyByte []byte, index int, chapterNamePath string, TsUrl string, wg *sync.WaitGroup) func() {
	return func() {
		defer wg.Done()
		res, _ := md.Client.R().Get(TsUrl)
		encrypter := res.Bytes()
		path := fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, index)
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		iv := utils.Iv(index)
		Byte := crypt.CBCDecrypter(encrypter, KeyByte, iv)

		file.Write(Byte)
	}
}

// VipVideo7 付费视频
func (md *MoocDownload) VipVideo7(TsList []string, key []byte, IV int, contentType, chapterNamePath string, Unit model.UnitStruct) {
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	utils.PathExistsAndMK(temPath)
	var wg *sync.WaitGroup
	wg = new(sync.WaitGroup)
	pool, _ := ants.NewPool(config.PoolSize)
	defer pool.Release()
	total := len(TsList)
	wg.Add(total)
	for index, TsUrl := range TsList {
		_ = pool.Submit(md.VipDecryptTs7(chapterNamePath, TsUrl, key, index, IV, wg))

		// 进度
		percentage := int(100 * (float64(index+1) / float64(total)))
		Unit.Progress.Percentage = percentage
		Unit.Progress.Name = Unit.UnitName
		Unit.Progress.Uuid = Unit.Uuid
		Unit.Progress.Status = "warning"
		runtime.EventsEmit(global.GlobalCtx, Unit.Uuid, Unit.Progress)
		//fmt.Println(Unit.Uuid, Progress)
	}
	wg.Wait()
	length := len(Unit.UnitName)
	Unit.UnitName = Unit.UnitName[:length-4]
	video.MergeTsFileListToSingleMp4(len(TsList), Unit.UnitName, chapterNamePath)
	//fmt.Println(Unit.UnitName + ".mp4 done\n")
}

// VipDecryptTs7 付费视频解密下载
func (md *MoocDownload) VipDecryptTs7(chapterNamePath string, TsUrl string, key []byte, index int, Iv int, wg *sync.WaitGroup) func() {
	return func() {
		defer wg.Done()
		res, _ := md.Client.R().Get(TsUrl)
		encrypter := res.Bytes()
		path := fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, index)
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		iv := utils.Iv(Iv)
		// Key, _ := hex.DecodeString(key)
		Byte := crypt.CBCDecrypter(encrypter, key, iv)
		file.Write(Byte)
	}
}

// FreeVideo 公开课视频
func (md *MoocDownload) FreeVideo(TsList []string, chapterNamePath string, Unit model.UnitStruct) {
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	utils.PathExistsAndMK(temPath)
	var wg *sync.WaitGroup
	wg = new(sync.WaitGroup)
	pool, _ := ants.NewPool(config.PoolSize)
	defer pool.Release()
	total := len(TsList)
	wg.Add(total)
	for index, TsUrl := range TsList {
		_ = pool.Submit(md.FreeTs(chapterNamePath, TsUrl, index, wg))

		// 进度 first
		percentage := int(100 * (float64(index+1) / float64(total)))
		Unit.Progress.Percentage = percentage
		Unit.Progress.Name = Unit.UnitName
		Unit.Progress.Uuid = Unit.Uuid
		Unit.Progress.Status = "warning"
		runtime.EventsEmit(global.GlobalCtx, Unit.Uuid, Unit.Progress)
		//fmt.Println(Unit.Uuid, percentage)
	}
	wg.Wait()
	length := len(Unit.UnitName)
	Unit.UnitName = Unit.UnitName[:length-4]
	video.MergeTsFileListToSingleMp4(len(TsList), Unit.UnitName, chapterNamePath)
}

// FreeTs 公开课视频下载
func (md *MoocDownload) FreeTs(chapterNamePath string, TsUrl string, index int, wg *sync.WaitGroup) func() {
	return func() {
		defer wg.Done()
		res, _ := md.Client.R().Get(TsUrl)
		VideoByte := res.Bytes()
		path := fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, index)
		target, _ := os.Create(path)
		target.Write(VideoByte)
		target.Close()
	}
}

func (md *MoocDownload) FreeGetTs(M3u8Str string) []string {
	tsCmp := regexp.MustCompile("[a-zA-Z0-9].*?.ts")
	// 获取ts列表
	tsList := tsCmp.FindAllString(M3u8Str, -1)
	return tsList
}

// Mp4Flv 合成视频
func (md *MoocDownload) Mp4Flv(videoUrl, chapterNamePath string, Isflv2mp4 bool, Unit model.UnitStruct) {
	length := len(Unit.UnitName)
	Unit.UnitName = Unit.UnitName[:length-4]
	var filename string
	if Isflv2mp4 {
		filename = Unit.UnitName + ".flv"
	} else {
		filename = Unit.UnitName + ".mp4"
	}
	concurrencyN := config.PoolSize
	resume := config.Resume
	NewDownloader(concurrencyN, resume).Download(videoUrl, chapterNamePath, filename)
	if Isflv2mp4 {
		flvPath := chapterNamePath + "\\" + filename
		mp4Path := chapterNamePath + "\\" + Unit.UnitName + ".mp4"
		video.Flv2Mp4(mp4Path, flvPath)
	}
	fmt.Println(filename + " done\n")
}

// VipGetTsKey 获取ts解密的key
func (md *MoocDownload) VipGetTsKey(encryptStr string, videoId int, contentType string) ([]string, []byte, int) {
	if contentType == "1" {
		videoId_ := strconv.Itoa(videoId)
		m3u8 := md.MJ.SecondaryDecrypt(encryptStr, videoId_, false)
		tsCmp := regexp.MustCompile("http.*ts")
		// 获取ts列表
		tsList := tsCmp.FindAllString(m3u8, -1)
		// 获取key
		keyCmp := regexp.MustCompile(`URI="(.*?)"`)
		keyUrl := keyCmp.FindStringSubmatch(m3u8)[1]
		// 用req库与requests的1.0.42版本使用报错，暂时使用requests 1.0.4版本
		//res, _ := md.Client.R().Get(keyUrl)
		//key := js.SecondaryDecrypt(res.String(), videoId_, true)
		headers := url.NewHeaders()
		headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
		headers.Set("origin", "https://www.icourse163.org")
		headers.Set("referer", "https://www.icourse163.org/")
		headers.Set("authority", "mooc2vod.stu.126.net")
		res, _ := requests.Get(keyUrl, &url.Request{Headers: headers})
		key := md.MJ.SecondaryDecrypt(res.Text, videoId_, true)

		KeyByte, _ := hex.DecodeString(key)
		return tsList, KeyByte, 0
	} else {
		m3u8 := encryptStr
		tsCmp := regexp.MustCompile("[a-zA-Z0-9].*?.ts")
		// 获取ts列表
		tsList := tsCmp.FindAllString(m3u8, -1)
		// 获取key
		keyCmp := regexp.MustCompile(`URI="(.*?)"`)
		keyUrl := keyCmp.FindStringSubmatch(m3u8)[1]

		IVcmp := regexp.MustCompile("IV=(\\d+.*?\\d+)")
		IVStr := IVcmp.FindStringSubmatch(m3u8)[1]
		IV, _ := strconv.ParseInt(IVStr, 16, 10)
		res, _ := md.Client.R().Get(keyUrl)
		KeyByte := res.Bytes()
		return tsList, KeyByte, int(IV)
	}
}
