/* 定义一些数据 */
package def

// 自定义命令行的帮助信息
const HelpInfo = `
1. 删除 @deprecated@ 字段
    trans -d test.messages.json
    trans -d dir/

2. 用英语替换其他语言 -- 只替换未翻译的字段(中文简体、繁体除外)
    trans -r test.messages.json

3. 用英语替换其他语言 -- 不管是否翻译，直接替换
    trans -ra test.messages.json

4. 用指定文案翻译指定文件
    trans -f test.messages.json -id 104407325.1157
`

const Api = "http://trans.viabtc.com/api/trans/card/"