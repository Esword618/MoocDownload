package main

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gookit/color"

	"MoocDownload/mooc/utils"
)

func main() {
	ffmpeg, lookErr := exec.LookPath("ffmpeg")
	if lookErr != nil {
		panic(lookErr)
	}
	name := "存储系统大题操作系统计组结合章节"
	count := 1018
	chapterNamePath := "C:\\Users\\jyj34\\Desktop\\test\\download\\22考研计算机组成原理暑期强化直播\\强化直播第一场存储系统大题操作系统计组结合章节"
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
	utils.PathExists(fmt.Sprintf("%s\\mp4", chapterNamePath))
	for i, j := range mp4List {
		concatStr := strings.Join(j, "|")
		// ffmpeg, lookErr := exec.LookPath("ffmpeg")
		// if lookErr != nil {
		// 	panic(lookErr)
		// }
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
		fmt.Println(args)
		cmd = exec.Command(ffmpeg, args...)
		_, err = cmd.Output()
		if err != nil {
			// panic(err)
			fmt.Println(err)
		}
	}
	concatStr := strings.Join(mp4PathList, "|")
	// ffmpeg, lookErr := exec.LookPath("ffmpeg")
	// if lookErr != nil {
	// 	panic(lookErr)
	// }
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
