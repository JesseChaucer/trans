package main

import (
	"flag"
	"fmt"
	"os"

	"inAction/trans/def"
	"inAction/trans/operate"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("请使用 -h 选项参考帮助文档。\n")
		return
	}

	// 从命令行获取参数
	var filePath string

	flag.StringVar(&filePath, "d", "", "delete")
	flag.StringVar(&filePath, "r", "", "replace")
	flag.StringVar(&filePath, "ra", "", "replace all")
	flag.StringVar(&filePath, "rcn", "", "replace with chinese")
	flag.StringVar(&filePath, "g", "", "get untranslated text")

	flag.StringVar(&filePath, "f", "", "文件路径")
	flag.StringVar(&def.TranId, "id", "", "文案系统的任务id")  // eg: TranId = "104407325.1157"

	flag.StringVar(&def.Lang, "lang", "", "指定要处理的语言")

	flag.Usage = func() {
		fmt.Print(def.HelpInfo)
	}
	flag.Parse()

	// 如果提供了文件路径和文案id
	if len(filePath) > 0 && len(def.TranId) > 0 {
		operate.FromTransSystem(filePath)
		return
	}

	// 如果未提供文案id
	if len(def.TranId) <= 0 {
		var operation = os.Args[1]
		switch operation {
		// 删除 @deprecated@ 字段
		case "-d":
			operate.DeleteDeprecated(filePath)
		// 用英语替换其他语言--只替换未翻译的字段(中文简体、繁体除外)
		case "-r":
			operate.ReplaceWithEn(filePath)
		// 用英语替换其他语言--不管是否翻译，直接替换
		case "-ra":
			operate.ReplaceAllWithEn(filePath)
		case "-rcn":
			operate.ReplaceWithCn(filePath)
		case "-g":
			operate.GetUntransText(filePath)
		default:
			fmt.Println("没有对应操作")
		}
		return
	}
}
