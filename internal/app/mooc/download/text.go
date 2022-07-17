package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/wangluozhe/requests/utils"
)

func Text(textUrl string, path string) {
	r, _ := http.Get(textUrl)
	// 获取文件名字
	text := r.Header["Content-Disposition"][0]
	re := regexp.MustCompile(`filename="(.*?)"`)
	filename := utils.DecodeURI(re.FindStringSubmatch(text)[1])
	// 获取文件大小
	fileSize, _ := strconv.Atoi(r.Header["Content-Length"][0])
	// 创建文件
	target, _ := os.Create(path)
	p := mpb.New(mpb.WithWidth(60))
	bar := p.New(int64(fileSize),
		mpb.BarStyle().Rbound("|"),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f"),
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), " done",
			),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
			decor.Name(" ] "),
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
		),
	)

	reader := bar.ProxyReader(r.Body)
	defer reader.Close()
	// 将下载的文件流拷贝到临时文件
	if _, err := io.Copy(target, reader); err != nil {
		target.Close()
	}
	target.Close()
	p.Wait()
	fmt.Println(filename, " done\n")
}
