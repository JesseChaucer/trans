// 从文案系统获取翻译数据
package operate

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"inAction/trans/def"
	"inAction/trans/util"
)

// 定义全局变量
var tranSlice def.TransType

// 获取api数据，返回对应的struct类型的地址
func getTransData(tranId string) *def.ResDataStruct {
	var url = def.Api + tranId
	fmt.Printf("Get data from: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Printf("状态：请求翻译系统文案成功\n\n")
		var jsonStr = string(body)
		resStructPointer := util.ResJsonToStruct(jsonStr)
		return resStructPointer
	} else {
		fmt.Println("状态：请求翻译系统文案失败")
		return nil
	}
}

/**
 * 翻译指定语言
 * 返回值表示文件是否有修改
 */
func transSpecifiedLang(lang string, langMap def.LangType) bool {
	var flag bool = false

	currentLangMap := langMap[lang]
	for _, val := range tranSlice {
		// fmt.Printf("%#v\n", val)
		var cn = val.Text       // 中文
		var translatedText = "" // 对应语言的翻译
		switch lang {
		case "en_US":
			translatedText = val.EnUS
			/* case "en_US":
				translatedText = val.EnUS
			case "es_ES":
				translatedText = val.EsES
			case "ja_JP":
				translatedText = val.JaJP
			case "ko_KP":
				translatedText = val.KoKP
			case "ru_KZ":
				translatedText = val.RuKZ
			case "fa_IR": // 波斯语
				translatedText = val.FaIR
			case "id_ID": // 印度尼西亚语
				translatedText = val.IdID
			case "tr_TR": // 土耳其语
				translatedText = val.TrTR
			case "vi_VN": // 越南语
				translatedText = val.ViVN
			case "ar_AE": // 阿拉伯
				translatedText = val.ArAE */
		}

		var md5Text = util.GetMD5Text(cn)
		if len(translatedText) > 0 {
			if _, ok := currentLangMap[md5Text]; ok {
				// fmt.Printf("%v = %v\n", md5Text, translatedText)
				currentLangMap[md5Text] = translatedText
				flag = true
			}
		}
	}
	return flag
}

/* 定义处理文件的函数 */
func processFunc(filePath string) {
	var flag bool = false
	// 多语言json文件转成map
	langMap := util.JsonToMap(filePath)

	for key, _ := range langMap {
		// fmt.Printf("%#v\n", key)
		// 不是简体和繁体，则翻译
		if key != "zh_Hans_CN" && key != "zh_Hant_HK" {
			var res = transSpecifiedLang(key, langMap)
			if res {
				flag = true
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

func FromTransSystem(filePath string, tranId string) {
	fmt.Printf("操作：用指定文案翻译指定文件\n\n")

	// 接口返回的翻译数据，转成结构体
	tranStruct := getTransData(tranId)
	tranSlice = tranStruct.Data.Trans

	util.ProcessAllFile(filePath, processFunc)
}
