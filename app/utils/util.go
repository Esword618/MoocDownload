package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func Iv(index int) []byte {
	indexStr := strconv.Itoa(index)
	length := len(indexStr)
	iv := make([]byte, 0, 16)
	for i := 0; i < 16-length; i++ {
		iv = append(iv, byte(0))
	}
	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(string(indexStr[i]))
		iv = append(iv, byte(num))
	}
	return iv
}

// PathExistsAndMK 判断文件夹是否存在并创建
func PathExistsAndMK(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}

func ReadCookie() string {
	basePath, _ := os.Getwd()
	cookiePath := fmt.Sprintf("%s\\cookie.txt", basePath)
	_, err := os.Stat(cookiePath)
	if err != nil {
		os.Create(cookiePath)
	}
	data, err := ioutil.ReadFile(cookiePath)
	if err != nil {
		os.Create(cookiePath)
		//color.Red.Println("\n请把你的cookie复制到cookie.txt文件中并重新运行本程序!")
		time.Sleep(5 * time.Second)
		syscall.Exit(0)

	}
	if len(data) == 0 {
		//color.Green.Println("\n检测到cookie.txt文件中没有cookie,请您粘贴cookie并重新运行本程序!")
		time.Sleep(5 * time.Second)
		syscall.Exit(0)
	}
	return string(data)
}

func SaveCookie() string {
	//color.Green.Printf("请把cookie复制到此处:")
	basePath, _ := os.Getwd()
	var cookieStr string
	fmt.Scanln(&cookieStr)
	for len(cookieStr) == 0 {
		//color.Red.Printf("请把cookie复制到此处:")
		fmt.Scanln(&cookieStr)
	}
	cookiePath := fmt.Sprintf("%s\\cookie.txt", basePath)
	f, err := os.OpenFile(cookiePath, os.O_RDWR|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	} else {
		_, err = f.Write([]byte(cookieStr))
	}
	return cookieStr
}

func CookieToMap(CookieStr string) map[string]string {
	CookieMap := make(map[string]string)
	CookieList := strings.Split(CookieStr, "; ")
	for _, i := range CookieList {
		iList := strings.Split(i, "=")
		CookieMap[iList[0]] = iList[1]
	}
	// fmt.Println(CookieMap)
	return CookieMap
	// cmp := regexp.MustCompile("; ")
	// fmt.Println(cmp.FindAllString(CookieStr, -1))
}

func RemoveInvalidChar(oldStr string) string {
	cmp := regexp.MustCompile("]|\\[|：|:|\\*|\\?|\t|\\x0b|/|>|<|")
	newStr := cmp.ReplaceAllString(oldStr, "")
	newStr = strings.Replace(newStr, "/", "", -1)
	re := regexp.MustCompile(`[\s+.!/_,$%^*("')]+|[+—()?【】“”！，。？、~@#￥%…&*（）]+|{+|}+|:`)
	newStr = re.ReplaceAllString(newStr, "")
	return newStr
}

func HttpCookieToMap(httpCookie []*http.Cookie) map[string]string {
	CookieMap := make(map[string]string)
	for _, cookie := range httpCookie {
		CookieMap[cookie.Name] = cookie.Value
	}
	return CookieMap
}

func CookieMapTOStr(CookieMap map[string]string) string {
	var CookieStr string
	for key, value := range CookieMap {
		CookieStr += key + "=" + value + "; "
	}
	fmt.Println(CookieStr)
	return CookieStr[:len(CookieStr)-2]
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 判断一个字符串是否在字符串数组中
func StringInArray(a string, b []string) bool {
	sort.Strings(b)
	index := sort.SearchStrings(b, a)
	if index < len(b) && b[index] == a {
		return true
	}
	return false
}

//// []interface {} 转 []string
//func Interface2Map(data interface{}) map[string]string {
//	// check if data is a slice of interface{} values
//	if slice, ok := data.([]interface{}); ok {
//		// create a new slice of strings and convert each value
//		Strings := make([]string, len(slice))
//		for i, v := range slice {
//			if str, ok := v.(string); ok {
//				Strings[i] = str
//			} else {
//				// handle error, value is not a string
//			}
//		}
//		// do something with the new slice of strings
//		return Strings
//	}
//	return []string{}
//}

// []interface {} 转 []string
func Interface2Strings(data interface{}) []string {
	// check if data is a slice of interface{} values
	if slice, ok := data.([]interface{}); ok {
		// create a new slice of strings and convert each value
		Strings := make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				Strings[i] = str
			} else {
				// handle error, value is not a string
			}
		}
		// do something with the new slice of strings
		return Strings
	}
	return []string{}
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
