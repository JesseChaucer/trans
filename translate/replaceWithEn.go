/**
 * 用英语替换其他语言--只替换未翻译字段
 */
package translate

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"bytes"

	"inAction/trans/util"
)

func ReplaceWithEn(filePath string) {
	fmt.Println("replace 未翻译字段")
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

	resByte, err := json.Marshal(langMap)
	if err != nil {
		log.Fatalln(err)
	}

	// 格式化JSON
	var formattedBytesBuffer bytes.Buffer
	json.Indent(&formattedBytesBuffer, resByte, "", "    ")

	err = ioutil.WriteFile(filePath, []byte(formattedBytesBuffer.String()), 0644)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("replace 未翻译字段 success...")
	}
}
