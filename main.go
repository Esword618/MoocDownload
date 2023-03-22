package main

import (
	"bytes"
	"embed"
	"github.com/Esword618/MoocDownload/app"
	"github.com/Esword618/MoocDownload/app/utils"
	"github.com/Esword618/MoocDownload/config"
	"github.com/spf13/viper"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	application := app.NewApplication()
	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:  "MoocDownload",
		Width:  1000,
		Height: 790,

		//MinWidth:          1000,
		//MinHeight:         780,
		//MaxWidth:          1200,
		//MaxHeight:         780,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         application.Startup,
		OnDomReady:        application.DomReady,
		OnBeforeClose:     application.BeforeClose,
		OnShutdown:        application.Shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			application,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			WebviewGpuIsDisabled:              false,
		},
		// Mac platform specific options
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
				Title:   "Wails Template Vue",
				Message: "A Wails template based on Vue and Vue-Router",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// 配置文件
	viper.SetConfigFile(config.ConfigPath)
	b, _ := utils.PathExist(config.ConfigPath)
	if !b {
		viper.ReadConfig(bytes.NewBuffer([]byte(config.DefaultYaml)))
		viper.WriteConfig()
	}
	viper.ReadInConfig()
}

//https://www.icourse163.org/learn/kaopei-1003292002?tid=1003524008#/learn/announce
