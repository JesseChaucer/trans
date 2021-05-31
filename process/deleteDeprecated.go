/**
 * 删除 @deprecated@ 字段
 */
package process

import (
	"fmt"
	"strings"

	"inAction/trans/util"
)

func delFunc(filePath string) {
	fmt.Println("---- 删除 @deprecated@ 字段 ----")

	langMap := util.JsonToMap(filePath)

	fmt.Printf("%v", langMap)
	fmt.Println("del func ----")


	if langMap == nil {
		return
	}

	// 遍历嵌套的map
	for _, singleLangMap := range langMap {
		for key, _ := range singleLangMap {
			// 删除包含 @deprecated@ 的键值对
			if strings.Contains(key, "@DEPRECATED@") {
				delete(singleLangMap, key)
			}
		}
	}

	fmt.Println("--------")
	fmt.Printf("%#v\n", langMap)

	// 把翻译的数据写回到多语言文件中
	util.WriteFile(filePath, langMap)
}

func DeleteDeprecated(filePath string) {
	util.ProcessAllFile(filePath, delFunc)
}
