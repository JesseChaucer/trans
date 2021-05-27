package main

import (
	"fmt"
	"flag"

	"inAction/trans/translate"
)

func main1() {
	fmt.Printf("%v\n", translate.GetMD5Text("2. 被邀请人累计交易量包括现货交易、杠杆交易和合约交易。"))
	// translate.DeleteDeprecated("./asset/index.json")
}

func main() {
	// 从命令行获取参数
	var operation string
	var filePath string
	var tranId string

	flag.StringVar(&operation, "o", "", "操作")
	flag.StringVar(&filePath, "f", "", "文件路径")
	flag.StringVar(&tranId, "id", "", "文案系统的任务id")
	flag.Parse()
	

	// 如果未提供文件路径
	if (len(filePath) <= 0) {
		fmt.Println("请指定文件路径")
		return;
	}

	switch operation {
	// 删除 @deprecated@ 字段
	case "d":
		translate.DeleteDeprecated(filePath)
	// 用英语替换其他语言--只替换未翻译的字段(中文简体、繁体除外)
	case "r":
		translate.ReplaceWithEn(filePath)
	// 用英语替换其他语言--不管是否翻译，直接替换
	case "ra":
		translate.ReplaceAllWithEn(filePath)
	// 从文案系统获取翻译数据
	case "api":
		// 如果未提供api路径
		if (len(tranId) <= 0) {
			fmt.Println("请指定api路径")
			return;
		}
		translate.FromTransSystem(filePath, tranId)  // eg: tranId = "104407325.1157"
	default: 
		fmt.Println("没有对应操作")
	}
}