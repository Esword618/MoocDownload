package main

import (
	"fmt"
	"net/http"
)

// url = "http://mooc1vod.stu.126.net/nos/mp4/2015/05/18/1571003_sd.mp4?ak=7909bff134372bffca53cdc2c17adc27a4c38c6336120510aea1ae1790819de82a5e95edf50731dafe6574167c941a01734a0b0389ae407faf13dc1ff1a078113059f726dc7bb86b92adbc3d5b34b132e6866222d16d6728d622da1f3663d3cb"

func main() {
	videoUrl := "http://mooc1vod.stu.126.net/nos/mp4/2015/05/18/1571003_sd.mp4?ak=7909bff134372bffca53cdc2c17adc27a4c38c6336120510aea1ae1790819de82a5e95edf50731dafe6574167c941a01734a0b0389ae407faf13dc1ff1a078113059f726dc7bb86b92adbc3d5b34b132e6866222d16d6728d622da1f3663d3cb"
	r, _ := http.Get(videoUrl)
	fmt.Println(r.Header.Get("Accept-Ranges"))
	fmt.Println(r.Header.Get("Content-Length"))
}
