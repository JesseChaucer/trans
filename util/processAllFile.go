package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 遍历文件夹下的所有文件，并处理
func ProcessAllFile(pathname string, processFunc func(string)) error {
	fmt.Printf("---- 当前目录：%s ----\n", pathname)
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		// 如果当前文件为目录，并且不以"."号开头（即不是隐藏文件）
		if fi.IsDir() && !strings.HasPrefix(fi.Name(), ".") {
			var subDir = pathname + "/" + fi.Name()
			ProcessAllFile(subDir)
		} else {
			// 处理文件
			// fmt.Println(fi.Name())
			processFunc(fi.Name())
		}
	}
	return err
}
