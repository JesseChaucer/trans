package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 遍历文件夹下的所有文件，并处理
func ProcessAllFile(filePath string, processFunc func(string)) error {
	rd, err := ioutil.ReadDir(filePath)
	fmt.Printf("%#v\n", rd)
	for _, fi := range rd {
		/**
		 * 1. 当前文件为目录
		 * 2. 不以"."号开头（即不是隐藏文件）
		 */
		var fileName = fi.Name()
		var currentFile = filePath + "/" + fileName
		if fi.IsDir() && !strings.HasPrefix(fileName, ".") {
			fmt.Printf("---- 当前目录：%s ----\n", filePath)
			ProcessAllFile(currentFile, processFunc)
		} else if strings.HasSuffix(fileName, "message.json") {
			// 如果文件名后缀为 message.json，则处理
			fmt.Printf("    处理文件：%s\n", currentFile)
			processFunc(currentFile)
		}
	}
	fmt.Println()
	return err
}
