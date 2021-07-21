/**
 * 把翻译的数据写回到多语言文件中
 */
package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io/ioutil"

	"inAction/trans/def"
)

// Encode()方法会在文件末尾写入空行，下面去掉空行
func delEmptyLine(filePath string) {
	data, err:= ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("read file error: %v\n", err)
	}
	// 去掉最后一个字节 "\n"
	lastByte := data[len(data)-1]
	if (lastByte == 10) {
		ioutil.WriteFile(filePath, data[:len(data)-1], 0666)
	}
}

func WriteFile(filePath string, langMap def.LangType) {
	// 打开目标json文件
	currentFile, err1 := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_TRUNC,
		0666,
	)
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer currentFile.Close()

	enc := json.NewEncoder(currentFile)
	enc.SetIndent("", "    ") // 用四个空格缩进
	enc.SetEscapeHTML(false) // 序列化时不转义特殊字符（如&）
	err2 := enc.Encode(langMap)

	delEmptyLine(filePath)

	if err2 != nil {
		log.Fatalln(err2)
	} else {
		fmt.Printf("%s\n\n", "结果：处理成功")
	} 
}
