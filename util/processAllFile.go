package util

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

// 处理文件 & 目录
func ProcessAllFile(filePath string, processFunc func(string)) {
	myFile, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	if err != nil {
		fmt.Printf("open error: %v\n", err)
		return
	}

	var myFileInfo, _ = myFile.Stat()
	// 如果是目录
	if myFileInfo.IsDir() {
		// 遍历文件夹下的所有文件，并处理
		rd, err := ioutil.ReadDir(filePath)
		if err != nil {
			fmt.Printf("read dir error: %v\n", err)
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
				fmt.Printf("---- 当前目录：%s ----\n", filePath)
				ProcessAllFile(newPath, processFunc)
			} else if strings.HasSuffix(fileName, "message.json") {
				// 如果文件名后缀为 message.json，则处理
				fmt.Printf("    处理文件：%s\n", newPath)
				processFunc(newPath)
			}
		}
	} else if strings.HasSuffix(myFileInfo.Name(), "message.json") {
		// 如果文件名后缀为 message.json，则处理
		fmt.Printf("处理文件：%s\n", filePath)
		processFunc(filePath)
	}
}
