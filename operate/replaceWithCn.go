/**
 * 用中文替换暂时用英文翻译的文案
 */
package operate

import (
	"fmt"

	"inAction/trans/util"
)

func replaceCnFunc(filePath string) {
	var flag bool = false

	langMap := util.JsonToMap(filePath)
	// 中文json
	cn := langMap["zh_Hans_CN"]
	// 英文json
	// us := langMap["en_US"]
	for lang, _ := range langMap {
		// 用中文替换指定语言
		// var excludeLangSlice = []string{"en_US", "zh_Hans_CN", "zh_Hant_HK"}
		var includeLangSlice = []string{"ja_JP", "id_ID", "vi_VN", "tr_TR", "ar_AE"}
		for _, includeLang := range includeLangSlice {
			if lang == includeLang {
				for field, val := range langMap[lang] {
					// 如果未翻译(是英文)，则替换成中文
					/* if (val == us[field]) {
						langMap[lang][field] = cn[field]
						flag = true
					} */
					// 直接替换成中文
					if len(val) > 0 {
						langMap[lang][field] = cn[field]
						flag = true
					}
				}
			}
		}
	}

	if flag {
		// 把翻译的数据写回到多语言文件中
		util.WriteFile(filePath, langMap)
	} else {
		fmt.Printf("结果：无处理\n\n")
	}
}

func ReplaceWithCn(filePath string) {
	fmt.Printf("操作：用中文替换暂时用英文翻译的文案\n\n")

	util.ProcessAllFile(filePath, replaceCnFunc)
}
