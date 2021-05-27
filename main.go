package main

import (
	"fmt"
	"flag"

	"inAction/trans/translate"
)

func main1() {
	// translate.FromTransSystem("104407325.1157")
	// translate.DeleteDeprecated("./asset/index.json")
}

func main() {
	// 从命令行获取参数
	var operation string // 操作
	var filePath string  // 文件路径
	var apiPath string  // api路径

	flag.StringVar(&operation, "o", "", "操作")
	flag.StringVar(&filePath, "f", "", "文件路径")
	flag.StringVar(&apiPath, "api", "", "文案系统的api路径")
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
		if (len(apiPath) <= 0) {
			fmt.Println("请指定api路径")
			return;
		}
		translate.FromTransSystem(filePath, apiPath)  // eg: apiPath = "104407325.1157"
	default: 
		fmt.Println("没有对应操作")
	}
}