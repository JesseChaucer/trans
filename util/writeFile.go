/**
 * 把翻译的数据写回到多语言文件中
*/
package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"bytes"
)

func WriteFile(filePath string, langMap LangType) {
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
		fmt.Println("write success...")
	}
}