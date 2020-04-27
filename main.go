package main

import (
	"fmt"
	"flag"
	"os"

	"inAction/fileUtil/translate"
)

func main() {
	// 从命令行获取参数
	var filePath string  // 文件路径

	flag.StringVar(&filePath, "d", "", "传入文件路径, 删除 @deprecated@ 字段")
	flag.StringVar(&filePath, "r", "", "传入文件路径, 用英语替换其他语言--只替换未翻译的字段")
	flag.StringVar(&filePath, "ra", "", "传入文件路径, 用英语替换其他语言--不管是否翻译，直接替换")
	flag.Parse()

	// 获取命令行参数
	/* fmt.Println("命令行参数数量:",len(os.Args))
	for k,v:= range os.Args{
			fmt.Printf("args[%v]=[%v]\n",k,v)
	} */
	
	operation := os.Args[1]
	switch operation {
	case "-d":
		translate.DeleteDeprecated(filePath)
	case "-r":
		translate.ReplaceWithEn(filePath)
	case "-ra":
		translate.ReplaceAllWithEn(filePath)
	default: 
		fmt.Println("没有对应操作")
	}
}