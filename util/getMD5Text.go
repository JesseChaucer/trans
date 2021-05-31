package util

import (
	"crypto/md5"
	"encoding/hex"
)

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
			if string(val) == " " {
				runeSlice = append(runeSlice[:key], runeSlice[key+1:]...)
			}
		}

		text = string(runeSlice[:8]) + "_" + md5Str[:4]
	}
	return text
}