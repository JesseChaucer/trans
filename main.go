package main

import (
	"fmt"
	"flag"
	"os"

	"inAction/fileUtil/translate"
)

func main1() {
	translate.FromTransSystem("73966177.704")
}

func main() {
	// 从命令行获取参数
	var filePath string  // 文件路径
	var apiPath string  // api路径

	flag.StringVar(&filePath, "d", "", "传入文件路径, 删除 @deprecated@ 字段")
	flag.StringVar(&filePath, "r", "", "传入文件路径, 用英语替换其他语言--只替换未翻译的字段")
	flag.StringVar(&filePath, "ra", "", "传入文件路径, 用英语替换其他语言--不管是否翻译，直接替换")
	flag.StringVar(&apiPath, "api", "", "传入api路径，从文案系统获取翻译数据")
	flag.Parse()
	
	operation := os.Args[1]
	switch operation {
	case "-d":
		translate.DeleteDeprecated(filePath)
	case "-r":
		translate.ReplaceWithEn(filePath)
	case "-ra":
		translate.ReplaceAllWithEn(filePath)
	// case "-api":
	// 	translate.FromTransSystem(apiPath)
	default: 
		fmt.Println("没有对应操作")
	}
}