/**
 * 提取未来翻译的文案
 */
package operate

import (
	"fmt"
	"os"

	"inAction/trans/util"
)

// 新建一个json文件，写入未翻译的文案
func writeUntransText(fileName string) {
	var filePath = "./asset/" + fileName
	// 打开文件，不存在时则创建
	myFile, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Printf("open file error...\n")
	}
	// 关闭文件
	defer myFile.Close()

	// 向文件中写入内容
	byteSlice := []byte("abc\n")
	myFile.Write(byteSlice)
}

func getUntransTextFunc(filePath string) {
	langMap := util.JsonToMap(filePath)
	// 简体汉语/繁体汉语/英语
	// var excludeLang = []string{"zh_Hans_CN", "zh_Hant_HK", "en_US"}
	// 日、韩、俄、西
	var targetLang = map[string]bool{
		"ja_JP": true,
		// "ko_KP": true,
		// "ru_KZ": true,
		// "es_ES": true,
	}

	// 中文json
	cn := langMap["zh_Hans_CN"]
	// 英文json
	en := langMap["en_US"]
	for lang, _ := range langMap {
		// 不是英语/简体汉语/繁体汉语，用英语替换
		if targetLang[lang] {
			fmt.Printf("---- %v ----\n", lang)

			// 打开文件，不存在时则创建
			myFile, err := os.OpenFile(
				"./asset/" + lang + ".json",
				os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
				0666,
			)
			if err != nil {
				fmt.Printf("open file error...\n")
			}
			// 关闭文件
			defer myFile.Close()

			// 写入左花括号
			myFile.Write([]byte("{\n"))

			for field, val := range langMap[lang] {
				if val == cn[field] || val == en[field] {
					fmt.Printf("%v\n", cn[field])
					// 向文件中写入内容
					byteSlice := []byte(fmt.Sprintf("    \"%v\": \"\",\n", cn[field]))
					myFile.Write(byteSlice)
				}
			}

			// 写入右花括号
			myFile.Write([]byte("}\n"))

			fmt.Print("\n\n")
		}
	}
}

// 如果文案对应的翻译为中文或英文，则表示其未被翻译，提取出来
func GetUntransText(filePath string) {
	fmt.Printf("操作：提取未翻译的文案\n\n")

	util.ProcessAllFile(filePath, getUntransTextFunc)
}
