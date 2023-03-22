package js

import (
	"bytes"
	"encoding/base64"
	"strconv"
)

type MoocTokenCrypt interface {
	l_(e int, t int) string
	l(e int) []string
	u(e int) string
	b(e []byte) string
	p(e []byte) string
	m(e []byte) []byte
	E(e string) []byte
	g(e string) []byte
}

type MoocTokenCryptStruct struct{}

var (
	moocTokenCryptStruct = MoocTokenCryptStruct{}
)

func TokenFunc() MoocTokenCrypt {
	return MoocTokenCrypt(&moocTokenCryptStruct)
}

var charListMap = map[string][]map[string]int{
	"v_0": {
		{
			"t": 1,
			"i": 2,
		},
		{
			"t": 3,
			"i": 9,
		},
		{
			"t": 3,
			"i": 7,
		},
		{
			"t": 3,
			"i": 2,
		},
		{
			"t": 1,
			"i": 0,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 1,
			"i": 2,
		},
		{
			"t": 3,
			"i": 1,
		},
		{
			"t": 3,
			"i": 1,
		},
		{
			"t": 3,
			"i": 8,
		},
		{
			"t": 1,
			"i": 5,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 3,
			"i": 6,
		},
		{
			"t": 3,
			"i": 9,
		},
		{
			"t": 1,
			"i": 5,
		},
	},
	"v_1": {
		{
			"t": 3,
			"i": 3,
		},
		{
			"t": 1,
			"i": 5,
		},
		{
			"t": 1,
			"i": 15,
		},
		{
			"t": 3,
			"i": 4,
		},
		{
			"t": 1,
			"i": 23,
		},
		{
			"t": 1,
			"i": 18,
		},
		{
			"t": 3,
			"i": 9,
		},
		{
			"t": 3,
			"i": 2,
		},
		{
			"t": 3,
			"i": 2,
		},
		{
			"t": 1,
			"i": 14,
		},
		{
			"t": 1,
			"i": 20,
		},
		{
			"t": 1,
			"i": 22,
		},
		{
			"t": 3,
			"i": 5,
		},
		{
			"t": 1,
			"i": 16,
		},
		{
			"t": 3,
			"i": 7,
		},
		{
			"t": 3,
			"i": 2,
		},
	},
	"v_2": {
		{
			"t": 2,
			"i": 16,
		},
		{
			"t": 2,
			"i": 7,
		},
		{
			"t": 1,
			"i": 7,
		},
		{
			"t": 2,
			"i": 24,
		},
		{
			"t": 1,
			"i": 17,
		},
		{
			"t": 2,
			"i": 4,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 2,
			"i": 18,
		},
		{
			"t": 2,
			"i": 12,
		},
		{
			"t": 2,
			"i": 5,
		},
		{
			"t": 2,
			"i": 18,
		},
		{
			"t": 2,
			"i": 4,
		},
		{
			"t": 1,
			"i": 0,
		},
		{
			"t": 2,
			"i": 22,
		},
		{
			"t": 1,
			"i": 11,
		},
		{
			"t": 2,
			"i": 6,
		},
	},
	"v_3": {
		{
			"t": 2,
			"i": 18,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 1,
			"i": 7,
		},
		{
			"t": 2,
			"i": 24,
		},
		{
			"t": 1,
			"i": 17,
		},
		{
			"t": 2,
			"i": 15,
		},
		{
			"t": 1,
			"i": 4,
		},
		{
			"t": 2,
			"i": 18,
		},
		{
			"t": 1,
			"i": 11,
		},
		{
			"t": 2,
			"i": 5,
		},
		{
			"t": 2,
			"i": 18,
		},
		{
			"t": 1,
			"i": 14,
		},
		{
			"t": 1,
			"i": 0,
		},
		{
			"t": 2,
			"i": 22,
		},
		{
			"t": 1,
			"i": 11,
		},
		{
			"t": 3,
			"i": 5,
		},
	},
}

var charEnlist1 = []string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
}
var charEnlist2 = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}
var charEnlist3 = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

func (C *MoocTokenCryptStruct) l_(e int, t int) string {
	return C.l(t)[e]
}

func (C *MoocTokenCryptStruct) l(e int) []string {
	switch e {
	case 1:
		return charEnlist1
	case 2:
		return charEnlist2
	case 3:
		return charEnlist3
	default:
		return nil
	}
}

func (C *MoocTokenCryptStruct) u(e int) string {
	t := charListMap["v_"+strconv.Itoa(e)]
	var n bytes.Buffer
	for _, c := range t {
		n.WriteString(C.l_(c["i"], c["t"]))
	}
	return n.String()
}

func (C *MoocTokenCryptStruct) b(e []byte) string {
	return base64.StdEncoding.EncodeToString(e)
}

func (C *MoocTokenCryptStruct) p(e []byte) string {
	var n bytes.Buffer
	var i int
	for i = 0; i < len(e)/8192; i++ {
		n.Write(e[i*8192 : (i+1)*8192])
	}
	n.Write(e[i*8192:])
	return n.String()
}

func (C *MoocTokenCryptStruct) m(e []byte) []byte {
	t := make([]byte, 0, len(e)*4)
	for _, n := range e {
		for a := 3; a >= 0; a-- {
			t = append(t, byte(n>>(8*a)&255))
		}
	}
	return t
}

func (C *MoocTokenCryptStruct) E(e string) []byte {
	return []byte(e)
}

func (C *MoocTokenCryptStruct) g(e string) []byte {
	t, _ := base64.StdEncoding.DecodeString(e)
	return []byte(t)
}
