package js

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/dop251/goja"

	"MoocDownload/crypt"
	"MoocDownload/internal/consts"
)

type MoocTokenCrypt interface {
	U(e string) string
	G(e string) []byte
	E(e string) []byte
	B(e []byte) string
	M(e []byte) []byte
	P(e []byte) string
}

type MoocTokenCryptStruct struct{}

var (
	VmToken              *goja.Runtime
	moocTokenCryptStruct = MoocTokenCryptStruct{}
)

func TokenFunc() MoocTokenCrypt {
	return MoocTokenCrypt(&moocTokenCryptStruct)
}

func init() {
	VmToken = goja.New()
	_, err := VmToken.RunString(consts.TokenSCRIPT)
	if err != nil {
		panic(err)
	}

}

func (C *MoocTokenCryptStruct) U(e string) string {
	var u func(string) string
	err := VmToken.ExportTo(VmToken.Get("u"), &u)
	if err != nil {
		panic(err)
	}
	return u(e)
}

func (C *MoocTokenCryptStruct) G(e string) []byte {
	var g func(string) []byte
	err := VmToken.ExportTo(VmToken.Get("g"), &g)
	if err != nil {
		panic(err)
	}
	return g(e)
}

func (C *MoocTokenCryptStruct) E(e string) []byte {
	var E func(string) []byte
	err := VmToken.ExportTo(VmToken.Get("E"), &E)
	if err != nil {
		panic(err)
	}
	return E(e)
}

func (C *MoocTokenCryptStruct) B(e []byte) string {
	var b func([]uint8) string
	err := VmToken.ExportTo(VmToken.Get("b"), &b)
	if err != nil {
		panic(err)
	}
	return b(e)
}

func (C *MoocTokenCryptStruct) M(e []byte) []byte {
	var m func([]byte) []byte
	err := VmToken.ExportTo(VmToken.Get("m"), &m)
	if err != nil {
		panic(err)
	}
	return m(e)
}

func (C *MoocTokenCryptStruct) P(e []byte) string {
	var p func([]byte) string
	err := VmToken.ExportTo(VmToken.Get("p"), &p)
	if err != nil {
		panic(err)
	}
	return p(e)
}

func Base64StdDecode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func TokenCrypt(e string) string {
	t := "1"
	i := TokenFunc().U(t)
	a := TokenFunc().G(e)
	o := a[:16]
	iB := Base64StdDecode(TokenFunc().B(TokenFunc().E(i)))
	oB := Base64StdDecode(TokenFunc().B(o))
	eB := a[16:]
	eB = Base64StdDecode(TokenFunc().B(eB))
	s := crypt.CBCDecrypter(eB, iB, oB)
	r := TokenFunc().M(s)
	d := TokenFunc().P(r)
	if strings.Contains(d, "}") {
		re := regexp.MustCompile(`"(.*?)"`)
		url := "https:" + re.FindAllStringSubmatch(d, -1)[1][1]
		return url
	}
	return ""
}
