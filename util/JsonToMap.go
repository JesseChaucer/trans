/* 
 * 读取 多语言的 json 文件，返回 map
*/
package util

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

func JsonToMap(file string) LangType {
	// 读取 json 文件
	jsonByte, err1 := ioutil.ReadFile(file)
	if err1 != nil {
		log.Fatal(err1)
		return nil
	}

	var langMap LangType
	err2 := json.Unmarshal(jsonByte, &langMap)
	if err2 != nil {
		log.Fatal(err2)
		return nil
	}

	return langMap
}