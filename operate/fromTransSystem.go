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
func getTransData() *def.ResDataStruct {
	var url = def.Api + def.TranId
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
		resStructPointer := util.ResJsonToStruct(body)
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
		var mapLang = map[string]string{
			"en_US": val.EnUS, // 英语
			"es_ES": val.EsES, // 西班牙
			"ja_JP": val.JaJP, // 日语
			"ko_KP": val.KoKP, // 韩语
			"ru_KZ": val.RuKZ, // 俄语
			"fr_FR": val.FrFR, // 法语
			"fa_IR": val.FaIR, // 波斯语
			"ar_AE": val.ArAE, // 阿拉伯语
			"id_ID": val.IdID, // 印度尼西亚语
			"tr_TR": val.TrTR, // 土耳其语
			"vi_VN": val.ViVN, // 越南语
		}
		translatedText = mapLang[lang]

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
		// 如果命令处参数提供了lang，则只处理指定的语言
		if len(def.Lang) > 0 {
			flag = transSpecifiedLang(def.Lang, langMap)
		} else if key != "zh_Hans_CN" && key != "zh_Hant_HK" {  // 不是简体和繁体，则翻译
			flag = transSpecifiedLang(key, langMap)
		}
	}

	if flag {
		// 把翻译的数据写回到多语言文件中
		util.WriteFile(filePath, langMap)
	} else {
		fmt.Printf("结果：无处理\n\n")
	}
}

func FromTransSystem(filePath string) {
	fmt.Printf("操作：用指定文案翻译指定文件\n\n")

	// 接口返回的翻译数据，转成结构体
	tranStruct := getTransData()
	tranSlice = tranStruct.Data.Trans

	util.ProcessAllFile(filePath, processFunc)
}
