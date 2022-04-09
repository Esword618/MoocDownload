package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

type Downloader struct {
	Resume          bool
	Concurrency     int
	ChapterNamePath string
	Filename        string
	VideoURL        string
}

func NewDownloader(chapterNamePath, filename, videoURL string) *Downloader {
	return &Downloader{Resume: true, Concurrency: 15, ChapterNamePath: chapterNamePath, Filename: filename, VideoURL: videoURL}
}

func (d *Downloader) Download() error {
	if d.Filename == "" {
		d.Filename = path.Base(d.VideoURL)
	}

	resp, err := http.Head(d.VideoURL)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.multiDownload(int(resp.ContentLength))
	}

	return d.singleDownload()
}

// 视频分片下载
func (d *Downloader) multiDownload(contentLen int) error {
	partSize := contentLen / 15
	var wg sync.WaitGroup
	wg.Add(d.Concurrency)
	rangeStart := 0
	barP := mpb.New(mpb.WithWidth(60), mpb.WithWaitGroup(&wg))
	bar := barP.New(int64(d.Concurrency),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(d.Filename, decor.WC{W: len(d.Filename), C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	for i := 0; i < d.Concurrency; i++ {
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.Concurrency-1 {
				rangeEnd = contentLen
			}

			downloaded := 0
			if d.Resume {
				partFileName := d.getPartFilename(d.ChapterNamePath, d.Filename, i)
				content, err := os.ReadFile(partFileName)
				if err == nil {
					downloaded = len(content)
				}
				bar.IncrBy(downloaded)
			}

			d.downloadPartial(rangeStart+downloaded, rangeEnd, i)

		}(i, rangeStart)
		bar.Increment()
		rangeStart += partSize + 1
	}

	wg.Wait()

	d.merge()

	return nil
}

// 单个视频下载
func (d *Downloader) singleDownload() error {
	r, err := http.Get(d.VideoURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	// 获取文件大小
	fileSize, _ := strconv.Atoi(r.Header["Content-Length"][0])
	// 创建文件

	target, _ := os.Create(d.ChapterNamePath + "\\" + d.Filename)
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
		return err
	}
	target.Close()
	p.Wait()
	fmt.Println(d.Filename, " done\n")
	return nil
	// resp, err := http.Get(strURL)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	//
	// d.setBar(int(resp.ContentLength))
	//
	// f, err := os.OpenFile(chapterPath+"//"+filename, os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	//
	// buf := make([]byte, 32*1024)
	// _, err = io.CopyBuffer(io.MultiWriter(f, d.bar), resp.Body, buf)
	// return err
}

// getPartFilename 构造部分文件的名字
func (d *Downloader) getPartFilename(chapterPath, filename string, partNum int) string {
	return fmt.Sprintf("%s/tem/%s-%d", chapterPath, filename, partNum)
}

func (d *Downloader) downloadPartial(rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", d.VideoURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	flags := os.O_CREATE | os.O_WRONLY
	if d.Resume {
		flags |= os.O_APPEND
	}
	partFile, err := os.OpenFile(d.getPartFilename(d.ChapterNamePath, d.Filename, i), flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer partFile.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(partFile), resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}

// 视频合成
func (d *Downloader) merge() error {
	destFile, err := os.OpenFile(d.ChapterNamePath+"\\"+d.Filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.Concurrency; i++ {
		partFileName := d.getPartFilename(d.ChapterNamePath, d.Filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}
	return nil
}
func main() {
	ChapterNamePath := "C:\\Users\\jyj34\\Desktop\\MoocDownload\\download"
	Filename := "1.mp4"
	videoUrl := "http://mooc1vod.stu.126.net/nos/mp4/2015/05/18/1571003_sd.mp4?ak=7909bff134372bffca53cdc2c17adc27a4c38c6336120510aea1ae1790819de82a5e95edf50731dafe6574167c941a01734a0b0389ae407faf13dc1ff1a078113059f726dc7bb86b92adbc3d5b34b132e6866222d16d6728d622da1f3663d3cb"
	NewDownloader(ChapterNamePath, Filename, videoUrl).Download()
}
