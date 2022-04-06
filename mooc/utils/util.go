package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gookit/color"
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

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
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
		color.Red.Println("\n请把你的cookie复制到cookie.txt文件中并重新运行本程序!")
		time.Sleep(5 * time.Second)
		syscall.Exit(0)

	}
	if len(data) == 0 {
		color.Green.Println("\n检测到cookie.txt文件中没有cookie,请您粘贴cookie并重新运行本程序!")
		time.Sleep(5 * time.Second)
		syscall.Exit(0)
	}
	return string(data)
}

func SaveCookie() string {
	color.Green.Printf("请把cookie复制到此处:")
	basePath, _ := os.Getwd()
	var cookieStr string
	fmt.Scanln(&cookieStr)
	for len(cookieStr) == 0 {
		color.Red.Printf("请把cookie复制到此处:")
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
