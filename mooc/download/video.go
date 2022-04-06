package download

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"

	"MoocDownload/crypt"
	"MoocDownload/mooc/js"
	"MoocDownload/mooc/utils"
)

func VipDecryptTs1(KeyByte []byte, index int, chapterNamePath string, TsUrl string, wg *sync.WaitGroup) func() {
	return func() {
		defer wg.Done()
		// client := resty.New()
		// r,_ := client.R().SetHeaders(map[string]string{
		//	"origin":"https://www.icourse163.org",
		//	"referer": "https://www.icourse163.org/",
		//	"authority": "mooc2vod.stu.126.net",
		//	"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		// }).Get(TsUrl)
		headers := url.NewHeaders()
		headers.Set("origin", "https://www.icourse163.org")
		headers.Set("referer", "https://www.icourse163.org/")
		headers.Set("authority", "mooc2vod.stu.126.net")
		headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
		req := url.NewRequest()
		req.Headers = headers
		r, _ := requests.Get(TsUrl, req)
		encrypter := r.Content
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

// 付费视频解密下载
func VipDecryptTs7(chapterNamePath string, TsUrl string, key []byte, index int, Iv int, wg *sync.WaitGroup) func() {
	return func() {
		// fmt.Println(index, "<---->", TsUrl)
		defer wg.Done()
		client := *resty.New()
		client.SetHeaders(map[string]string{
			"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
			"origin":     "https://www.icourse163.org",
			"referer":    "https://www.icourse163.org/",
			"authority":  "mooc2vod.stu.126.net",
		})
		res, _ := client.R().Get(TsUrl)
		encrypter := res.Body()
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

// 公开课视频下载
func FreeTs(chapterNamePath string, TsUrl string, index int, wg *sync.WaitGroup) func() {
	return func() {
		defer wg.Done()
		headers := url.NewHeaders()
		headers.Set("origin", "https://www.icourse163.org")
		headers.Set("referer", "https://www.icourse163.org/")
		headers.Set("authority", "mooc2vod.stu.126.net")
		headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
		req := url.NewRequest()
		req.Headers = headers
		r, _ := requests.Get(TsUrl, req)
		VideoByte := r.Content
		path := fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, index)
		target, _ := os.Create(path)
		target.Write(VideoByte)
		target.Close()
	}
}

func VipGetTsKey(encryptStr string, videoId int, contentType string) ([]string, []byte, int) {
	if contentType == "1" {
		videoId_ := strconv.Itoa(videoId)
		m3u8 := js.M3u8(encryptStr, videoId_)
		tsCmp := regexp.MustCompile("http.*?ts")
		// 获取ts列表
		tsList := tsCmp.FindAllString(m3u8, -1)
		// 获取key
		keyCmp := regexp.MustCompile(`URI="(.*?)"`)
		keyUrl := keyCmp.FindStringSubmatch(m3u8)[1]
		fmt.Println(keyUrl)
		color.Red.Printf("===================")
		headers := url.NewHeaders()
		headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
		headers.Set("origin", "https://www.icourse163.org")
		headers.Set("referer", "https://www.icourse163.org/")
		headers.Set("authority", "mooc2vod.stu.126.net")
		res, _ := requests.Get(keyUrl, &url.Request{Headers: headers})
		text := res.Text
		key := js.Key(text, videoId_)
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

		client := resty.New()
		client.SetHeaders(map[string]string{
			"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
			"origin":     "https://www.icourse163.org",
			"referer":    "https://www.icourse163.org/",
			"authority":  "mooc2vod.stu.126.net",
		})
		res, _ := client.R().Get(keyUrl)
		KeyByte := res.Body()
		return tsList, KeyByte, int(IV)
	}

}

func FreeGetTs(M3u8Str string) []string {
	tsCmp := regexp.MustCompile("[a-zA-Z0-9].*?.ts")
	// 获取ts列表
	tsList := tsCmp.FindAllString(M3u8Str, -1)
	return tsList
}

// 付费视频
func VipVideo1(TsList []string, KeyByte []byte, unitName string, chapterNamePath string) {
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	utils.PathExists(temPath)
	var wg sync.WaitGroup
	pool, _ := ants.NewPool(15)
	defer pool.Release()

	barP := mpb.New(mpb.WithWidth(60), mpb.WithWaitGroup(&wg))

	total := len(TsList)
	name := fmt.Sprintf("%s.mp4 :", unitName)
	// create a single bar, which will inherit container's width
	bar := barP.New(int64(total),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),

		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name), C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	wg.Add(total)
	for index, TsUrl := range TsList {
		_ = pool.Submit(VipDecryptTs1(KeyByte, index, chapterNamePath, TsUrl, &wg))
		bar.Increment()
	}

	barP.Wait()
	wg.Wait()
	MergeTs(len(TsList), unitName, chapterNamePath)
	fmt.Println(unitName + ".mp4 done\n")
	// err := os.RemoveAll(temPath)
	// if err != nil {
	// 	panic(err)
	// }
}

// 付费视频
func VipVideo7(TsList []string, key []byte, IV int, contentType, unitName string, chapterNamePath string) {
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	utils.PathExists(temPath)
	var wg sync.WaitGroup
	concurrencyN := viper.GetInt("download.concurrencyn")
	pool, _ := ants.NewPool(concurrencyN)
	defer pool.Release()

	barP := mpb.New(mpb.WithWaitGroup(&wg))

	total := len(TsList)
	name := fmt.Sprintf("%s.mp4 :", unitName)
	// create a single bar, which will inherit container's width
	bar := barP.New(int64(total),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),

		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name), C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	wg.Add(total)
	for index, TsUrl := range TsList {
		_ = pool.Submit(VipDecryptTs7(chapterNamePath, TsUrl, key, index, IV, &wg))
		bar.Increment()
	}
	barP.Wait()
	wg.Wait()
	MergeTs(len(TsList), unitName, chapterNamePath)
	fmt.Println(unitName + ".mp4 done\n")
	// err := os.RemoveAll(temPath)
	// if err != nil {
	// 	panic(err)
	// }
}

// 公开课视频
func FreeVideo(TsList []string, unitName string, chapterNamePath string) {
	// temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	// utils.PathExists(temPath)

	var wg sync.WaitGroup
	pool, _ := ants.NewPool(15)
	defer pool.Release()

	barP := mpb.New(mpb.WithWidth(60), mpb.WithWaitGroup(&wg))

	total := len(TsList)
	name := fmt.Sprintf("%s.mp4 :", unitName)
	// create a single bar, which will inherit container's width
	bar := barP.New(int64(total),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name), C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	wg.Add(total)
	for index, TsUrl := range TsList {
		// wg.Add(1)
		_ = pool.Submit(FreeTs(chapterNamePath, TsUrl, index, &wg))
		bar.Increment()
	}
	barP.Wait()
	wg.Wait()
	MergeTs(len(TsList), unitName, chapterNamePath)
}

// 合成视频
func MergeTs(count int, name string, chapterNamePath string) {
	ffmpeg, lookErr := exec.LookPath("ffmpeg")
	if lookErr != nil {
		color.Red.Println("ffmpeg 无法调用！请稍后重试！")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}
	if count <= 200 {
		var concatStr string
		for i := 0; i < count; i++ {
			if i != count-1 {
				concatStr += fmt.Sprintf("%s\\tem\\%d.ts|", chapterNamePath, i)
			} else {
				concatStr += fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, i)
			}
		}
		args := []string{
			"-y",
			"-i",
			fmt.Sprintf("concat:%s", concatStr),
			"-acodec",
			"copy",
			"-vcodec",
			"copy",
			"-absf",
			"aac_adtstoasc",
			fmt.Sprintf("%s\\%s.mp4", chapterNamePath, name),
		}
		cmd := exec.Command(ffmpeg, args...)
		_, err := cmd.Output()
		if err != nil {
			panic(err)
			// fmt.Println(err)
		}
	} else {
		color.Blue.Println("视频过大，正在处理中")
		TsList := []string{}
		for i := 0; i < count; i++ {
			TsList = append(TsList, chapterNamePath+"\\tem\\"+strconv.Itoa(i)+".ts")
		}
		tsCount := 30
		n := math.Ceil(float64(count / tsCount))
		mp4List := [][]string{}
		for i := 0.0; i <= n; i++ {
			if i == n {
				mp4List = append(mp4List, TsList[int(i*float64(tsCount)):])
			} else {
				mp4List = append(mp4List, TsList[int(i*float64(tsCount)):int((i+1)*float64(tsCount))])
			}
		}
		mp4PathList := []string{}
		for i, j := range mp4List {
			concatStr := strings.Join(j, "|")
			args := []string{
				"-y",
				"-i",
				fmt.Sprintf("concat:%s", concatStr),
				"-acodec",
				"copy",
				"-vcodec",
				"copy",
				"-absf",
				"aac_adtstoasc",
				fmt.Sprintf("%s\\tem\\%d.mp4", chapterNamePath, i),
			}
			cmd := exec.Command(ffmpeg, args...)
			_, err := cmd.Output()
			if err != nil {
				// panic(err)
				fmt.Println(err)
			}
			args = []string{
				"-y",
				"-i",
				fmt.Sprintf("%s\\tem\\%d.mp4", chapterNamePath, i),
				"-vcodec",
				"copy",
				"-acodec",
				"copy",
				"-vbsf",
				"h264_mp4toannexb",
				fmt.Sprintf("%s\\tem\\%v.ts", chapterNamePath, i),
			}
			mp4PathList = append(mp4PathList, fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, i))
			cmd = exec.Command(ffmpeg, args...)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
		}
		concatStr := strings.Join(mp4PathList, "|")
		args := []string{
			"-y",
			"-i",
			fmt.Sprintf("concat:%s", concatStr),
			"-acodec",
			"copy",
			"-vcodec",
			"copy",
			"-absf",
			"aac_adtstoasc",
			fmt.Sprintf("%s\\%s.mp4", chapterNamePath, name),
		}
		cmd := exec.Command(ffmpeg, args...)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		color.Blue.Println("视频处理完成")
	}
}

func Mp4Flv(videoUrl, unitName, chapterNamePath string, Isflv2mp4 bool) {
	var filename string
	if Isflv2mp4 {
		filename = unitName + ".flv"
	} else {
		filename = unitName + ".mp4"
	}
	concurrencyN := viper.GetInt("download.concurrentn")
	resume := viper.GetBool("download.resume")
	NewDownloader(concurrencyN, resume).Download(videoUrl, chapterNamePath, filename)
	if Isflv2mp4 {
		ffmpeg, lookErr := exec.LookPath("ffmpeg")
		if lookErr != nil {
			color.Red.Println("\n您还没安装ffmpeg,无法下载，请您安装后下载!\n")
			return
		}
		mp4Path := chapterNamePath + "\\" + unitName + ".mp4"
		args := []string{
			"-i",
			chapterNamePath + "\\" + filename,
			"-vcodec",
			"copy",
			"-acodec",
			"copy",
			mp4Path,
		}
		cmd := exec.Command(ffmpeg, args...)
		_, err := cmd.Output()
		if err != nil {
			color.Red.Println("！！！！！！！！！！！！警告，flv视频转mp4失败！！！！！！！！！！！")
			return
			// panic(err)
		}
		os.Remove(mp4Path)
	}
	fmt.Println(filename + " done\n")
}
