/**
 * 用英语替换其他语言
 */
package operate

import (
	"fmt"

	"inAction/trans/util"
)

func replaceAllFunc(filePath string) {	
	langMap := util.JsonToMap(filePath)
	for key, _ := range langMap {
		// 不是英语/简体汉语/繁体汉语，用英语替换
		if (key != "en_US" && key != "zh_Hans_CN" && key != "zh_Hant_HK") {
			langMap[key] = langMap["en_US"]
		}
	}

	/* 把翻译的数据写回到多语言文件中 */
	util.WriteFile(filePath, langMap)
}

func ReplaceAllWithEn(filePath string) {
	fmt.Printf("操作：用英语替换其他语言 -- 不管是否翻译，直接替换\n\n")

	util.ProcessAllFile(filePath, replaceAllFunc)
}
