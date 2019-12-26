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



	/* input, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	var resSlice []string
	for index, line := range lines {
		if !strings.Contains(line, "@DEPRECATED@") {
			// 如果下一行是要删除的行，并且是当前语言的最后一行(即最后没有逗号)，则当前行最后的逗号要去掉
			if index+1 < len(lines) {
				nextLine := lines[index+1]
				if strings.Contains(nextLine, "@DEPRECATED@") && nextLine[len(nextLine)-1:] != "," {
					line = strings.TrimRight(line, ",")
				}
			}
			resSlice = append(resSlice, line)
		}
	}

	// 覆盖原文件
	output := strings.Join(resSlice, "\n")
	err = ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("delete success...")
	} */
}
