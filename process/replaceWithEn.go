/**
 * 用英语替换其他语言--只替换未翻译字段
 */
package process

import (
	"fmt"

	"inAction/trans/util"
)

func ReplaceWithEn(filePath string) {
	fmt.Println("---- 用英语替换其他语言 -- 只替换未翻译的字段(中文简体、繁体除外) ----")
	
	langMap := util.JsonToMap(filePath)
	// 中文json
	cn := langMap["zh_Hans_CN"]
	// 英文json
	us := langMap["en_US"]
	for lang, _ := range langMap {
		// 不是英语/简体汉语/繁体汉语，用英语替换
		if (lang != "en_US" && lang != "zh_Hans_CN" && lang != "zh_Hant_HK") {
			for field, val := range langMap[lang] {
				// 如果未翻译(是中文)，则替换成英文
				if (val == cn[field]) {
					langMap[lang][field] = us[field]
				}
			}
		}
	}

	/* 把翻译的数据写回到多语言文件中 */
	util.WriteFile(filePath, langMap)
}
