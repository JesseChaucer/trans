// 从文案系统获取翻译数据
package translate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/md5"
	"encoding/hex"

	"inAction/trans/util"
)

// 获取api数据，返回对应的struct类型的地址
func GetTransData(tranId string) *util.ResDataStruct {
	var url = "http://trans.viabtc.com/api/trans/card/" + tranId
	fmt.Printf("Get data from: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Printf("====请求成功====\n\n")
		var jsonStr = string(body)
		resStructPointer := util.ResJsonToStruct(jsonStr)
		return resStructPointer
	} else {
		fmt.Println("请求api数据失败")
		return nil
	}
}

/** 
 * 获取字符串的md5加密后的字符串
 * 如："7×24小时 专家支持" --> "7×24小时专家_0ecc"
 * 含中文字符，要先转成[]rune
*/
func GetMD5Text(text string) string {
	var runeSlice = []rune(text)
	if len(runeSlice) > 8 {
		h := md5.New()
		h.Write([]byte(string(runeSlice)))
		md5Str := hex.EncodeToString(h.Sum(nil))

		// 去除字符串中的空格
		for key, val := range runeSlice {
			if (string(val) == " ") {
				runeSlice = append(runeSlice[:key], runeSlice[key+1:]...)
			}
		}

		text = string(runeSlice[:8]) + "_" + md5Str[:4]
	}
	return text
} 


// 翻译指定语言
func trans(lang string, langMap util.LangType, tranSlice util.TransType) {
	currentLangMap := langMap[lang]
	for _, val := range tranSlice {
		// fmt.Printf("%#v\n", val)
		var cn = val.Text  // 中文
		var translatedText = "";  // 对应语言的翻译
		switch lang {
		case "en_US":
			translatedText = val.EnUS
		case "es_ES":
			translatedText = val.EsES
		case "ja_JP":
			translatedText = val.JaJP
		case "ko_KP":
			translatedText = val.KoKP
		case "ru_KZ":
			translatedText = val.RuKZ
		}
		var md5Text = GetMD5Text(cn)
		if (len(translatedText) > 0) {
			if _, ok := currentLangMap[md5Text]; ok {
				// fmt.Printf("%v\n", md5Text)
				currentLangMap[md5Text] = translatedText;
			}
		}
	}
}

func FromTransSystem(filePath string, tranId string) {
	fmt.Println("---- 用指定文案翻译指定文件 ----")
	// 接口返回的翻译数据，转成结构体
	tranStruct := GetTransData(tranId)
	var tranSlice = tranStruct.Data.Trans

	// 多语言json文件转成map
	// filePath = "./asset/demo.json"
	langMap := util.JsonToMap(filePath)
	
	for key, _ := range langMap {
		// fmt.Printf("%#v\n", key)
		// 不是简体和繁体，则翻译
		if (key != "zh_Hans_CN" && key != "zh_Hant_HK") {
			trans(key, langMap, tranSlice);
		}
	}
	

	/* 把翻译的数据写回到多语言文件中 */
	util.WriteFile(filePath, langMap)
}
