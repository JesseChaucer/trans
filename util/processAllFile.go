package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"log"
)

// 定义全局变量
var processFunc func(string)

// 处理文件
func processFile(filePath string, fileName string) {
	// 如果文件名后缀为 messages.json，则处理
	if strings.HasSuffix(fileName, "messages.json") {
		fmt.Printf("文件：%s\n", filePath)
		processFunc(filePath)
	}
}

// 处理目录
func processDir(filePath string) {
	fmt.Printf("---- 当前目录：%s ----\n", filePath)

	// 遍历文件夹下的所有文件，并处理
	rd, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Printf("read dir error: %v\n", err)
		return
	}

	for _, fi := range rd {
		/**
		* 1. 当前文件为目录
		* 2. 不以"."号开头（即不是隐藏文件）
		 */
		var fileName = fi.Name()
		var newPath = filePath + "/" + fileName
		if fi.IsDir() && !strings.HasPrefix(fileName, ".") {
			ProcessAllFile(newPath, processFunc)
		} else {
			processFile(newPath, fileName)
		}
	}
}

// 处理文件 & 目录
func ProcessAllFile(filePath string, paramProcessFunc func(string)) {
	processFunc = paramProcessFunc

	myFileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	if myFileInfo.IsDir() {
		processDir(filePath) // 处理目录
	} else {
		processFile(filePath, myFileInfo.Name()) // 处理文件
	}
}
