package js

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Esword618/MoocDownload/crypt"
	"github.com/duke-git/lancet/v2/convertor"
	"strings"
	"time"
)

//type MoocJs interface {
//	SecondaryDecrypt(e, t string, i bool) string
//	TokenCrypt(e string) string
//}

type MoocJs struct {
}

func NewMoocJs() *MoocJs {
	return &MoocJs{}
}

func (mj *MoocJs) SecondaryDecrypt(e string, t string, i bool) string {
	e = SecondaryDecryptFunc().d(e)
	//now := time.Now()
	var o string
	if i {
		o = convertor.ToString(t)[len(t)-7:] + "66" + convertor.ToString(time.Now().UnixMilli())[:4] + convertor.ToString(time.Now().Year())[1:]
	} else {
		o = convertor.ToString(t)[len(t)-6:] + "66" + convertor.ToString(time.Now().UnixMilli())[:4] + convertor.ToString(time.Now().Year())
	}

	s := SecondaryDecryptFunc().b(e)

	r := s[:16] // 创建一个长度为16的切片r，其底层数据指向s的前16个字节

	//key
	oBytes, _ := base64.StdEncoding.DecodeString(SecondaryDecryptFunc().E(SecondaryDecryptFunc().g(o)))

	//iv
	rBytes, _ := base64.StdEncoding.DecodeString(SecondaryDecryptFunc().E(r))

	// ciphertext
	eBytes, _ := base64.StdEncoding.DecodeString(SecondaryDecryptFunc().E(s[16 : len(s)-16]))

	ciphertext := crypt.CBCDecrypter(eBytes, oBytes, rBytes)

	cc := SecondaryDecryptFunc().m(ciphertext)

	h := SecondaryDecryptFunc().p(cc)
	h = string(bytes.Replace([]byte(h), []byte{0, 0, 0}, []byte{}, -1))
	if i {
		return h[:32]
	} else {
		return h
	}
}

func (mj *MoocJs) TokenCrypt(e string, t int) (link string, error error) {
	i := TokenFunc().u(t)

	a := TokenFunc().g(e)

	o := a[:16]
	// key
	iBytes, _ := base64.StdEncoding.DecodeString(TokenFunc().b(TokenFunc().E(i)))
	// iv
	oBytes, _ := base64.StdEncoding.DecodeString(TokenFunc().b(o))
	// ciphertext
	eBytes, _ := base64.StdEncoding.DecodeString(TokenFunc().b(a[16:]))

	s := crypt.CBCDecrypter(eBytes, iBytes, oBytes)

	r := TokenFunc().m(s)
	r = bytes.Replace(r, []byte{0, 0, 0}, []byte{}, -1)

	h := TokenFunc().p(r)

	d := strings.Index(h, "}")

	if d != -1 {
		var data map[string]string
		err := json.Unmarshal([]byte(h), &data)
		if err != nil {
			return link, err
		}
		link = "https:" + data["k"]
		return link, nil
	} else {
		return link, errors.New("失败")
	}
}
