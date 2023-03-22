package download

import (
	"fmt"
	"github.com/Esword618/MoocDownload/app/global"
	"github.com/Esword618/MoocDownload/app/platform/mooc/model"
	"github.com/Esword618/MoocDownload/app/utils"
	"github.com/imroc/req/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (md *MoocDownload) Text(textUrl string, Unit model.UnitStruct, chapterNamePath string) {
	utils.PathExistsAndMK(chapterNamePath)
	path := fmt.Sprintf("%s//%s.pdf", chapterNamePath, Unit.UnitName)
	client := req.C()
	callback := func(info req.DownloadInfo) {
		if info.Response.Response != nil {
			//fmt.Printf("downloaded %.2f%%\n", float64(info.DownloadedSize)/float64(info.Response.ContentLength)*100.0)
			percentage := float64(info.DownloadedSize) / float64(info.Response.ContentLength) * 100.0
			Unit.Progress.Percentage = int(percentage)
			Unit.Progress.Name = Unit.UnitName + ".pdf"
			Unit.Progress.Uuid = Unit.Uuid
			Unit.Progress.Status = "warning"
			runtime.EventsEmit(global.GlobalCtx, Unit.Uuid, Unit.Progress)
		}
	}
	client.R().
		SetOutputFile(path).
		SetDownloadCallback(callback).
		Get(textUrl)
}
