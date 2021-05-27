package main

import (
	"flag"
	"fmt"

	"inAction/trans/util"
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
	flag.Usage = func() {
        fmt.Print(util.HelpInfo)
    }
	flag.Parse()

	// 如果未提供文件路径（文件路径为必选参数）
	if len(filePath) <= 0 {
		fmt.Println("请指定文件路径")
		return
	}

	// 如果提供了 "文案id" 参数，则根据文案系统中指定的内容，翻译指定的多语言文件
	if len(tranId) > 0 {
		translate.FromTransSystem(filePath, tranId) // eg: tranId = "104407325.1157"
		return
	}

	// 如果提供了“操作”参数
	if len(operation) > 0 {
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
		default:
			fmt.Println("没有对应操作")
		}
	}
}
