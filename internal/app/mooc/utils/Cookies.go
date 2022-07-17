package utils

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func NewCookies() *cookiejar.Jar {
	cookies, _ := cookiejar.New(nil)
	return cookies
}

func ParseCookies(rawurl, cookies string) *cookiejar.Jar {
	c := NewCookies()
	cookieList := strings.Split(cookies, ";")
	urls, _ := url.Parse(rawurl)
	for _, cookie := range cookieList {
		cookie = strings.TrimSpace(cookie)
		if cookie == "" {
			continue
		}
		keyValue := strings.SplitN(cookie, "=", 2)
		if len(keyValue) != 2 {
			panic(errors.New("该字符串不符合Cookies标准"))
		}
		key := keyValue[0]
		value := keyValue[1]
		if strings.Contains(value, `"`) {
			value = strings.Replace(value, `"`, ``, -1)
		}
		c.SetCookies(urls, []*http.Cookie{{
			Name:  key,
			Value: value,
		}})
	}
	return c
}
