/**
 * 把翻译的数据写回到多语言文件中
 */
package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"inAction/trans/def"
)

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
	// 用四个空格缩进
	enc.SetIndent("", "    ")
	enc.SetEscapeHTML(false) // 序列化时不转义特殊字符（如&）
	err2 := enc.Encode(langMap)
	if err2 != nil {
		log.Fatalln(err2)
	} else {
		fmt.Printf("%s\n\n", "结果：处理成功")
	} 
}
