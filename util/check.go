package util

import (
	"encoding/json"
	"os/exec"

	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/wangluozhe/requests"

	"MoocDownload/model"
)

func CheckVersion() {
	r, _ := requests.Get("http://1.15.249.230:6000/checkVersion1", nil)
	var checkStruct model.CheckStruct
	json.Unmarshal([]byte(r.Text), &checkStruct)
	if checkStruct.Data.Version != viper.GetString("info.version") {
		color.Red.Println(checkStruct.Data.Info)
	}
	if checkStruct.Msg != nil {
		color.Red.Println(checkStruct.Msg)
	}

}

func CheckFfmpeg() bool {
	_, lookErr := exec.LookPath("ffmpeg")
	if lookErr != nil {
		return false
	}
	return true
}
