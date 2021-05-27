# translation

## -o: 指定操作 

```shell
# d: delete
# 删除 @deprecated@ 字段
trans -o d

# r: replace, 用英语替换其他语言--只替换未翻译的字段(中文简体、繁体除外)
trans -o r

# ra: replace all, 用英语替换其他语言--不管是否翻译，直接替换
trans -o ra
```

## -f: 指定文件

```
trans -o r -f test.message.json
```

## -id: 指定文案id 

```shell
# 根据文案系统中指定的内容，翻译指定的多语言文件
trans -f test.message.json -id 104407325.1157
```

