// 从文案系统获取翻译数据
package translate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/md5"
	"encoding/hex"

	"inAction/fileUtil/util"
)

// 获取api数据，返回对应的struct类型的地址
func GetTransData(apiPath string) *util.ResDataStruct {
	var url = "http://trans.coinex.cc/api/trans/card/" + apiPath
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
	// 去除字符串中的空格
	for key, val := range runeSlice {
		if (string(val) == " ") {
			runeSlice = append(runeSlice[:key], runeSlice[key+1:]...)
		}
	}
	if len(runeSlice) > 8 {
		h := md5.New()
		h.Write([]byte(string(runeSlice)))
		md5Str := hex.EncodeToString(h.Sum(nil))
		text = string(runeSlice[:8]) + "_" + md5Str[:4]
	}
	return text
} 

func FromTransSystem(apiPath string) {
	// 接口返回的翻译数据，转成结构体
	tranStruct := GetTransData(apiPath)

	// 多语言json文件转成map
	langMap := util.JsonToMap("./asset/index.json")
	langMapEn := langMap["en_US"]

	// 遍历翻译数据slice
	var tranSlice = tranStruct.Data.Trans
	for _, val := range tranSlice {
		var cn = val.Text  // 中文
		var en = val.EnUS  // 英文
		var md5Text = GetMD5Text(cn)
		if _, ok := langMapEn[md5Text]; ok {
			fmt.Printf("%v\n", md5Text)
			langMapEn[md5Text] = en;
		}
	}
	fmt.Printf("%v\n", langMapEn)
}
