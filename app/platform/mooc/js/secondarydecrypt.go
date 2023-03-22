package js

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"unicode/utf8"
)

type MoocSecondaryDecryptStruct struct {
}

type MoocSecondaryDecrypt interface {
	d(e string) string
	b(e string) []byte
	g(e string) []byte
	E(e []byte) string
	p(e []byte) string
	m(e []byte) []byte
}

var (
	moocSecondaryDecryptStruct = MoocSecondaryDecryptStruct{}
)

func SecondaryDecryptFunc() MoocSecondaryDecrypt {
	return MoocSecondaryDecrypt(&moocSecondaryDecryptStruct)
}

func (C *MoocSecondaryDecryptStruct) d(e string) string {
	t := strings.TrimSpace(e)
	if strings.HasPrefix(strings.ToLower(t), "0x") {
		t = t[2:]
	}
	if len(t)%2 != 0 {
		panic("Illegal Format ASCII Code!")
	}
	data := make([]byte, hex.DecodedLen(len(t)))
	hex.Decode(data, []byte(t))
	return string(data)
}

func (C *MoocSecondaryDecryptStruct) b(e string) []byte {
	t, _ := base64.StdEncoding.DecodeString(e)
	n := make([]byte, len(t))
	for i := 0; i < len(t); i++ {
		n[i] = byte(t[i])
	}
	return n
}

func (C *MoocSecondaryDecryptStruct) g(e string) []byte {
	t := make([]byte, len(e))
	for i := 0; i < len(e); i++ {
		t[i] = e[i]
	}
	return t
}

func (C *MoocSecondaryDecryptStruct) E(e []byte) string {
	s := make([]byte, len(e))
	for i, b := range e {
		s[i] = byte(b)
	}
	t := string(s)
	return base64.StdEncoding.EncodeToString([]byte(t))
}

func (C *MoocSecondaryDecryptStruct) p(e []byte) string {
	var t strings.Builder
	for i := 0; i < len(e); i++ {
		r, _ := utf8.DecodeRune([]byte{e[i]})
		t.WriteRune(r)
	}
	return t.String()
}
func (C *MoocSecondaryDecryptStruct) m(e []byte) []byte {
	t := make([]byte, 0, len(e)*4)
	for _, n := range e {
		for a := 3; a >= 0; a-- {
			t = append(t, byte(n>>(8*a)&255))
		}
	}
	return t
}

//func m(e []byte) []byte {
//	t := make([]byte, len(e)*4)
//	for i := 0; i < len(e); i++ {
//		n := e[i]
//		for a := 3; a >= 0; a-- {
//			t[i*4+(3-a)] = n >> uint8(8*a) & 255
//		}
//	}
//	return t
//}
