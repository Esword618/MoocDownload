package js

import (
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
	"github.com/wangluozhe/requests/utils"
)

func M3u8(e string, t string) string {
	i := "False"
	data := url.NewData()
	data.Set("e", e)
	data.Set("t", t)
	data.Set("i", i)
	res, _ := requests.Post("http://1.15.249.230:6000/secondarydecrypt", &url.Request{Data: data})

	m3u8 := res.Text
	m3u8 = utils.Base64Decode(m3u8)
	return m3u8
}

func Key(e string, t string) string {
	i := "True"
	data := url.NewData()
	data.Set("e", e)
	data.Set("t", t)
	data.Set("i", i)
	res, _ := requests.Post("http://1.15.249.230:6000/secondarydecrypt", &url.Request{Data: data})

	key := res.Text
	key = utils.Base64Decode(key)
	return key
}
