package config

import (
	"log"
	"os"
)

// 断点下载
const Resume = true

const ConfigPath = "./config.yaml"

// 每次下载视频个数
const BatchSize = 3

// 协程数
const PoolSize = 5

const DefaultYaml = `
author: esword
mooc:
  cookie: ""
xtzx:
  cookie: ""
download:
  batchsize: 3
  poolsize: 5
`

// 判断文件是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// 创建文件夹
func CreateDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
