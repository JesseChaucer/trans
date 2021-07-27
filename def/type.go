/* 定义一些公共类型 */
package def

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
	PtPT string `json:"pt_PT"`  // 葡萄牙语
	RuKZ string `json:"ru_KZ"`  // 俄语
	FrFR string `json:"fr_FR"`  // 法语
	FaIR string `json:"fa_IR"`  // 波斯语
	IdID string `json:"id_ID"`  // 印度尼西亚语
	TrTR string `json:"tr_TR"`  // 土耳其语
	ViVN string `json:"vi_VN"`  // 越南语
	ArAE string `json:"ar_AE"`  // 阿拉伯
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