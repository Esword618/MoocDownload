package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// 使用选项创建应用
	TesTChan <- "hello world"
	err := wails.Run(&options.App{
		Title:             "MoocDownload",
		Width:             1200,
		Height:            768,
		DisableResize:     false, // 不可调节窗口大小
		Fullscreen:        false, // 是否全屏
		Frameless:         false, // 无边框
		StartHidden:       false, // 启动隐藏
		HideWindowOnClose: false, // 退出隐藏窗口
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "MoocDownload",
				Message: "MoocDownload By Esword",
				Icon:    icon,
			},
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
