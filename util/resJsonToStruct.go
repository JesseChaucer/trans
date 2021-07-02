/* 将接口返回的JSON字符串转成struct */
package util

import (
	"fmt"
	"encoding/json"

	"inAction/trans/def"
)

// json str 转 struct
func ResJsonToStruct(jsonByte []byte) *def.ResDataStruct {
	var resStruct def.ResDataStruct
		err := json.Unmarshal(jsonByte, &resStruct)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return &resStruct
}