package util

// 自定义命令行的帮助信息
const HelpInfo = `
1. 删除 @deprecated@ 字段
    trans -d test.message.json

2. 用英语替换其他语言 -- 只替换未翻译的字段(中文简体、繁体除外)
    trans -r test.message.json

3. 用英语替换其他语言 -- 不管是否翻译，直接替换
    trans -ra test.message.json

4. 用指定文案翻译指定文件
    trans -f test.message.json -id 104407325.1157
`

/* 定义类型 */
// 多语言文件的数据结构
type LangType map[string]map[string]string

type TransType []struct {
	ID   string `json:"id"`
	Mark string `json:"mark"`
	Text string `json:"text"`
	EnUS string `json:"en_US"`  // 英语
	EsES string `json:"es_ES"`  // 西班牙
	JaJP string `json:"ja_JP"`  // 日语
	KoKP string `json:"ko_KP"`  // 韩语
	RuKZ string `json:"ru_KZ"`  // 俄语
	FaIR string `json:"fa_IR"`  // 波斯语
}

// 接口返回的数据结构
type ResDataStruct struct {
	Data struct {
		CreateTime int64    `json:"createTime"`
		Creator    string   `json:"creator"`
		ID         string   `json:"id"`
		OwnerTrans []string `json:"ownerTrans"`
		Title      string   `json:"title"`
		Trans      TransType `json:"trans"`
	} `json:"data"`
	Success bool `json:"success"`
}