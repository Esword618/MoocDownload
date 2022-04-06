package js

import (
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
	"github.com/wangluozhe/requests/utils"
)

func Token(e string) string {
	data := url.NewData()
	data.Set("e", e)
	data.Set("qualityValue", "1")
	res, _ := requests.Post("http://1.15.249.230:6000/token", &url.Request{Data: data})
	videoToke := res.Text
	// fmt.Println(videoToke)
	videoToke = utils.Base64Decode(videoToke)
	// fmt.Println(videoToke)
	return videoToke
}
