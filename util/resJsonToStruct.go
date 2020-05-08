/* 将接口返回的JSON字符串转成struct */
package util

import (
	"fmt"
	"encoding/json"
)

// 接口返回的数据结构
type ResDataStruct struct {
	Data struct {
		CreateTime int64    `json:"createTime"`
		Creator    string   `json:"creator"`
		ID         string   `json:"id"`
		OwnerTrans []string `json:"ownerTrans"`
		Title      string   `json:"title"`
		Trans      []struct {
			EnUS string `json:"en_US"`
			ID   string `json:"id"`
			JaJP string `json:"ja_JP"`
			KoKP string `json:"ko_KP"`
			Mark string `json:"mark"`
			Text string `json:"text"`
		} `json:"trans"`
	} `json:"data"`
	Success bool `json:"success"`
}

// json str 转 struct
func ResJsonToStruct(jsonStr string) *ResDataStruct {
	var resStruct ResDataStruct
		err := json.Unmarshal([]byte(jsonStr), &resStruct)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return &resStruct
}