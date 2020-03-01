/**
 * 删除 @deprecated@ 字段
 */
package translate

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"encoding/json"
	"bytes"

	"inAction/fileUtil/util"
)

func DeleteDeprecated(filePath string) {
	fmt.Println("delete")
	langMap := util.JsonToMap(filePath)
	// 遍历嵌套的map
	for _, singleLangMap := range langMap {
		for key, _ := range singleLangMap {
			// 删除包含 @deprecated@ 的键值对
			if strings.Contains(key, "@DEPRECATED@") {
				delete(singleLangMap, key)
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
		fmt.Println("delete success...")
	}
}
