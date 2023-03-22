package app

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Esword618/MoocDownload/app/global"
	"github.com/Esword618/MoocDownload/app/platform/mooc"
	"github.com/Esword618/MoocDownload/app/utils"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"log"
)

type Info struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
	Status     string `json:"status"`
}

// App struct
type App struct {
	ctx          context.Context
	file         []byte
	progressInfo []Info
	mooc         *mooc.Mooc
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApplication() *App {
	Mooc := mooc.NewMooc()
	return &App{
		mooc: Mooc,
	}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
	// 赋 mooc ctx
	//a.mooc.Ctx = ctx
	// 全局变量
	global.GlobalCtx = ctx
	// 窗口居中
	runtime.WindowCenter(a.ctx)

	// 解析 link
	runtime.EventsOn(ctx, "parse_link", func(optionalData ...interface{}) {

		a.mooc.UpdateCookie()
		//a.mooc.CheckStatus()
		link := optionalData[0].(string)
		//fmt.Println(link)
		a.mooc.ParseLink(link)
	})

	// 下载
	runtime.EventsOn(ctx, "download", func(UuidList ...interface{}) {
		//fmt.Println(UuidList[0])
		uuidList := utils.Interface2Strings(UuidList[0])
		downloadInfo := a.mooc.HandledownloadInfo(uuidList)
		a.mooc.Download(downloadInfo)
	})

	//runtime.EventsEmit(ctx, "default_setting", viper.AllSettings())

	// 更新设置
	runtime.EventsOn(ctx, "update_setting", func(optionalData ...interface{}) {
		//moocCookie := optionalData[0].(string)
		//fmt.Println(moocCookie)
		setting := optionalData[0].(map[string]interface{})
		viper.Set("mooc.cookie", setting["moocCookie"].(string))
		//fmt.Println(viper.AllSettings())
		a.mooc.UpdateCookie()
	})
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) DomReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作

}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
	fmt.Println("Shutting down")
	viper.WriteConfig()
}

func (a *App) initClipboard() {
	clipboard.Init()
	ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	for img := range ch {
		// print out clipboard data whenever it is changed
		//println(string(data))
		//img := clipboard.Read(clipboard.FmtImage)
		a.file = img
		base64Img := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(img))
		//fmt.Println(base64Img)
		//fmt.Println("go")
		//EventsEmit发消息
		//EventsOn接收消息
		runtime.EventsEmit(a.ctx, "clipboard_image", base64Img)
		runtime.WindowShow(a.ctx)
	}
}

func (a *App) parse() {
	//fmt.Println("parse")
	client := req.C()
	res, err := client.R().
		SetHeaders(map[string]string{
			"cookie":             "__atuvc=1^%^7C8; __atuvs=63f992ec1481c1e1000",
			"origin":             "https://p2t.behye.com",
			"referer":            "https://p2t.behye.com/",
			"sec-ch-ua-mobile":   "?0",
			"sec-fetch-dest":     "empty",
			"sec-ch-ua-platform": "Windows",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		}).
		SetFileBytes("image", "image", a.file). // Set form param name and filename
		SetFormData(map[string]string{          // Set form data while uploading
			"session_id": "session-AiYJrY-bxFfzzCwOA9Kb4cyWfpdJnt6q",
		}).
		Post("https://p2t.behye.com/api/pix2text")
	if err != nil {
		log.Println(err)
	}
	result, _ := res.ToString()
	//fmt.Println(result)
	runtime.EventsEmit(a.ctx, "parse_result", result)
	//fmt.Println("发送完成")
}

// 图片压缩功能
func compressImage(data []byte) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return data
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 80})
	if err != nil {
		return data
	}
	if buf.Len() > len(data) {
		return data
	}
	return buf.Bytes()
}
