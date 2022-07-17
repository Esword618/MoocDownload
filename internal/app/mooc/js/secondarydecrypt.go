package js

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dop251/goja"
	"github.com/duke-git/lancet/v2/convertor"

	"MoocDownload/crypt"
	"MoocDownload/internal/consts"
)

type MoocSecondaryDecrypt interface {
	D(e string) string
	B(e string) []byte
	G(e string) []byte
	E(e []byte) string
	M(e []byte) []byte
	P(e []byte) string
}

type MoocSecondaryDecryptStruct struct {
}

var (
	VmSecondaryDecrypt         *goja.Runtime
	moocSecondaryDecryptStruct = MoocSecondaryDecryptStruct{}
)

func SecondaryDecryptFunc() MoocSecondaryDecrypt {
	return MoocSecondaryDecrypt(&moocSecondaryDecryptStruct)
}

func init() {
	VmSecondaryDecrypt = goja.New()
	_, err := VmSecondaryDecrypt.RunString(consts.SecondaryDecryptScript)
	if err != nil {
		panic(err)
	}
}

func (C *MoocSecondaryDecryptStruct) D(e string) string {
	var d func(string) string
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("d"), &d)
	if err != nil {
		panic(err)
	}
	return d(e)
}

func (C *MoocSecondaryDecryptStruct) B(e string) []byte {
	var b func(string) []byte
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("b"), &b)
	if err != nil {
		panic(err)
	}
	return b(e)
}

func (C *MoocSecondaryDecryptStruct) G(e string) []byte {
	var g func(string) []byte
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("g"), &g)
	if err != nil {
		panic(err)
	}
	return g(e)
}

func (C *MoocSecondaryDecryptStruct) E(e []byte) string {
	var E func([]byte) string
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("E"), &E)
	if err != nil {
		panic(err)
	}
	return E(e)
}

func (C *MoocSecondaryDecryptStruct) M(e []byte) []byte {
	var m func([]byte) []byte
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("m"), &m)
	if err != nil {
		panic(err)
	}
	return m(e)
}

func (C *MoocSecondaryDecryptStruct) P(e []byte) string {
	var p func([]byte) string
	err := VmSecondaryDecrypt.ExportTo(VmSecondaryDecrypt.Get("p"), &p)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	return p(e)
}

func SecondaryDecrypt(e, t string, i bool) string {
	e = SecondaryDecryptFunc().D(e)
	var o string
	if i {
		o = convertor.ToString(t)[len(t)-7:] + "66" + convertor.ToString(time.Now().UnixMilli())[:4] + convertor.ToString(time.Now().Year())[1:]
	} else {
		o = convertor.ToString(t)[len(t)-6:] + "66" + convertor.ToString(time.Now().UnixMilli())[:4] + convertor.ToString(time.Now().Year())
	}
	s := SecondaryDecryptFunc().B(e)
	r := s[:16]
	oB := crypt.Base64StdDecode(SecondaryDecryptFunc().E(SecondaryDecryptFunc().G(o)))
	// fmt.Println(oB)
	rB := crypt.Base64StdDecode(SecondaryDecryptFunc().E(r))
	eB := s[16:]
	eB = crypt.Base64StdDecode(SecondaryDecryptFunc().E(eB))
	l := crypt.CBCDecrypter(eB, oB, rB)
	_k := SecondaryDecryptFunc().M(l)
	h := SecondaryDecryptFunc().P(_k)
	h = string(bytes.Replace([]byte(h), []byte{0, 0, 0}, []byte{}, -1))
	return h
}
