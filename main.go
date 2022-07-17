package main

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gookit/color"

	"MoocDownload/util"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"

	"MoocDownload/internal/app/mooc"
)

func main() {
	// util.CheckVersion()
	ffmpegB := util.CheckFfmpeg()
	if !ffmpegB {
		color.Red.Println("\n检查到您还没安装ffmpeg,如果你要下载视频，可能无法下载，建议您安装后下载!如果仅仅是下载文本资料将不受影响！\n")
		color.Blue.Println("ffmpeg配置教程:" + "https://blog.csdn.net/weixin_42132415/article/details/118294517")
		time.Sleep(3 * time.Second)
	}
	mooc.MoocMain()
}

func init() {
	Config()
	Table()
}

func Config() {
	_, err := os.Stat(".\\conf.ini")
	if err != nil {
		cfg := ini.Empty()
		DefaultSection := cfg.Section("Info")
		NameSection, _ := DefaultSection.NewKey("Name", "慕课下载器")
		NameSection.Comment = "# 名字"

		VersionSection, _ := DefaultSection.NewKey("Version", "3.0.4")
		VersionSection.Comment = "# 版本号"

		PathSection := cfg.Section("Path")
		PathSection.Comment = "# 路径"

		downloadSection, _ := PathSection.NewKey("SavePath", "download")
		downloadSection.Comment = "# 保存路径"

		judgeSection := cfg.Section("Download")
		judgeSection.Comment = "# 下载为1，不下载为0"
		videoSection, _ := judgeSection.NewKey("Video", "1")
		videoSection.Comment = "# 是否下载视频"

		coursewareSection, _ := judgeSection.NewKey("Courseware", "1")
		coursewareSection.Comment = "# 是否下载资料"

		paperSection, _ := judgeSection.NewKey("Paper", "1")
		paperSection.Comment = "# 是否下载试卷(暂时不支持)"

		concurrencyN := runtime.NumCPU()
		concurrencyN = concurrencyN * 3 / 4
		concurrentSection, _ := judgeSection.NewKey("ConcurrentN", strconv.Itoa(concurrencyN))
		concurrentSection.Comment = "# 并发下载数(默认为电脑的cpu数量)"

		resumeSection, _ := judgeSection.NewKey("Resume", "true")
		resumeSection.Comment = "# 是否支持断点下载(默认为true,目前仅仅部分视频支出)"

		cfg.SaveTo("conf.ini")
	}
	viper.AddConfigPath(".\\")
	viper.SetConfigName("conf")
	viper.SetConfigType("ini")
	viper.WatchConfig()
	viper.ReadInConfig()
	// fmt.Println(viper.AllSettings())
}

func Table() {
	data := [][]string{
		{"慕", "", "", ""},
		{"课", "", "", ""},
		{"下", "Esword", "Spiders and AI", "https://github.com/Esword618/MoocDownload"},
		{"载", "", "", ""},
		{"器", "", "", ""},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{viper.GetString("info.version"), "Author", "公众号", "Github", ""})
	table.SetFooter([]string{"", "", "♥ Cmf", "Super invincible little cute", ""}) // Add Footer
	table.SetBorder(false)                                                         // Set Border to false

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{},
	)

	table.SetFooterColor(
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiYellowColor},
		tablewriter.Colors{},
	)

	table.AppendBulk(data)
	table.Render()
}
