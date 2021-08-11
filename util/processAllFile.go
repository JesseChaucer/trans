package util

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

// 处理文件
func processFile(filePath string, processFunc func(string)) {
	// 如果文件名后缀为 messages.json，则处理
	var fileName string = filepath.Base(filePath)
	if strings.HasSuffix(fileName, "messages.json") {
		fmt.Printf("文件：%s\n", filePath)
		processFunc(filePath)
	}
}

// 使用 Walk 函数遍历指定目录下所有的文件和目录
func ProcessAllFile(filePath string, processFunc func(string)) {
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		// walk the tree to count files and folders
		if info.IsDir() {
			fmt.Printf("---- 当前目录：%s ----\n\n", path)
		} else {
			processFile(path, processFunc) // 处理文件
		}
		return nil
	})
}
