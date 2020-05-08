/* 
 * 读取 多语言的 json 文件，返回 map
*/
package util

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

type LangType map[string]map[string]string

func JsonToMap(file string) LangType {
	// 读取 json 文件
	jsonByte, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var langMap LangType
	err = json.Unmarshal(jsonByte, &langMap)
	if err != nil {
		log.Fatal(err)
	}

	return langMap
}