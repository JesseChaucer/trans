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
	langMap := util.JsonToMap(filePath)

	if langMap == nil {
		return
	}

	var flag bool = false;
	// 遍历嵌套的map
	for _, singleLangMap := range langMap {
		for key, _ := range singleLangMap {
			// 删除包含 @deprecated@ 的键值对
			if strings.Contains(key, "@DEPRECATED@") {
				flag = true;
				delete(singleLangMap, key)
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

func DeleteDeprecated(filePath string) {
	fmt.Printf("操作：删除 @deprecated@ 字段\n\n")

	util.ProcessAllFile(filePath, delFunc)
}
