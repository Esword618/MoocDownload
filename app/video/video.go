/*
 * @Author: Esword
 * @Description:
 * @FileName:  mergeTs
 * @Version: 1.0.0
 * @Date: 2022-08-06 22:41
 */

package video

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/yapingcat/gomedia/go-codec"
	"github.com/yapingcat/gomedia/go-flv"
	"github.com/yapingcat/gomedia/go-mp4"
	"github.com/yapingcat/gomedia/go-mpeg2"
)

// ts视频列表合成mp4
func MergeTsFileListToSingleMp4(count int, name string, chapterNamePath string) {
	var tsFileList []string
	for i := 0; i < count; i++ {
		tsFileList = append(tsFileList, fmt.Sprintf("%s\\tem\\%d.ts", chapterNamePath, i))
	}
	OutPutMp4 := fmt.Sprintf("%s\\%s.mp4", chapterNamePath, name)
	mp4file, err := os.OpenFile(OutPutMp4, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer mp4file.Close()

	muxer, err := mp4.CreateMp4Muxer(mp4file)
	if err != nil {
		fmt.Println(err)
	}
	vtid := muxer.AddVideoTrack(mp4.MP4_CODEC_H264)
	atid := muxer.AddAudioTrack(mp4.MP4_CODEC_AAC)

	demuxer := mpeg2.NewTSDemuxer()
	var OnFrameErr error
	var audioTimestamp uint64 = 0
	aacSampleRate := -1
	demuxer.OnFrame = func(cid mpeg2.TS_STREAM_TYPE, frame []byte, pts uint64, dts uint64) {
		if OnFrameErr != nil {
			return
		}
		if cid == mpeg2.TS_STREAM_AAC {
			audioTimestamp = pts
			codec.SplitAACFrame(frame, func(aac []byte) {
				if aacSampleRate == -1 {
					adts := codec.NewAdtsFrameHeader()
					adts.Decode(aac)
					aacSampleRate = codec.AACSampleIdxToSample(int(adts.Fix_Header.Sampling_frequency_index))
				}
				err = muxer.Write(atid, aac, audioTimestamp, audioTimestamp)
				audioTimestamp += uint64(1024 * 1000 / aacSampleRate) // 每帧aac采样固定为1024。aac_sampleRate 为采样率
				if err != nil {
					OnFrameErr = err
					return
				}
			})
		} else if cid == mpeg2.TS_STREAM_H264 {
			err = muxer.Write(vtid, frame, pts, dts)
			if err != nil {
				OnFrameErr = err
				return
			}
		} else {
			OnFrameErr = errors.New("unknown cid " + strconv.Itoa(int(cid)))
			return
		}
	}
	// simulating some work
	for _, tsFile := range tsFileList {
		var buf []byte
		buf, err = ioutil.ReadFile(tsFile)
		if err != nil {
			fmt.Println(err)
		}
		err = demuxer.Input(bytes.NewReader(buf))
		if err != nil {

		}
		//bar.Increment()
	}
	// wait for our bar to complete and flush
	//p.Wait()

	err = muxer.WriteTrailer()
	if err != nil {
		fmt.Println(err)
	}
	err = mp4file.Sync()
	if err != nil {
		fmt.Println(err)
	}
	temPath := fmt.Sprintf("%s\\tem", chapterNamePath)
	err = os.RemoveAll(temPath)
	if err != nil {
		panic(err)
	}
}

// flv转mp4
func Flv2Mp4(Mp4FilePath, FlvFilePath string) {
	mp4file, err := os.OpenFile(Mp4FilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mp4file.Close()

	muxer, err := mp4.CreateMp4Muxer(mp4file)
	if err != nil {
		fmt.Println(err)
		return
	}
	vtid := muxer.AddVideoTrack(mp4.MP4_CODEC_H264)
	atid := muxer.AddAudioTrack(mp4.MP4_CODEC_AAC)

	flvfilereader, _ := os.Open(FlvFilePath)
	defer flvfilereader.Close()
	fr := flv.CreateFlvReader()

	fr.OnFrame = func(ci codec.CodecID, b []byte, pts, dts uint32) {
		if ci == codec.CODECID_AUDIO_AAC {
			err := muxer.Write(atid, b, uint64(pts), uint64(dts))
			if err != nil {
				fmt.Println(err)
			}

		} else if ci == codec.CODECID_VIDEO_H264 {
			err := muxer.Write(vtid, b, uint64(pts), uint64(dts))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	cache := make([]byte, 4096)
	for {
		n, err := flvfilereader.Read(cache)
		if err != nil {
			fmt.Println(err)
			break
		}
		fr.Input(cache[0:n])
	}
	muxer.WriteTrailer()
	fmt.Println("flv转mp4完成！")
}
