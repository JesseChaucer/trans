/**
 * 用英语替换其他语言
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

func ReplaceAllWithEn(filePath string) {
	fmt.Println("---- 用英语替换其他语言 -- 不管是否翻译，直接替换 ----")
	langMap := util.JsonToMap(filePath)
	for key, _ := range langMap {
		// 不是英语/简体汉语/繁体汉语，用英语替换
		if (key != "en_US" && key != "zh_Hans_CN" && key != "zh_Hant_HK") {
			langMap[key] = langMap["en_US"]
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
		fmt.Println("replace all success...")
	}
}
