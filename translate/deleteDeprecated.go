/**
 * 删除 @deprecated@ 字段
 */
package translate

import (
	"fmt"
	"strings"

	"inAction/trans/util"
)

func DeleteDeprecated(filePath string) {
	fmt.Println("---- delete deprecated ----")
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

	/* 把翻译的数据写回到多语言文件中 */
	util.WriteFile(filePath, langMap)
}
