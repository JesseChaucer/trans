/**
 * 用英语替换其他语言
 */
package process

import (
	"fmt"

	"inAction/trans/util"
)

func ReplaceAllWithEn(filePath string) {
	fmt.Println("操作：用英语替换其他语言 -- 不管是否翻译，直接替换")
	
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
