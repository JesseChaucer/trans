package main

import (
	"fmt"
	"flag"

	"inAction/fileUtil/translate"
)

func main() {
	// 从命令行获取参数
	file := "";  // 文件路径
	operation := "";  // 操作
	flag.StringVar(&file, "f", "default", "文件路径")
	flag.StringVar(&operation, "o", "d", "删除操作")
	flag.Parse()

	switch operation {
	case "d":
		translate.DeleteDeprecated(file)
	case "r":
		translate.ReplaceWithEn(file)
	default: 
		fmt.Println("没有对应操作")
	}
}