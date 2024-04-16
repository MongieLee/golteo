package common

import (
	"encoding/json"
	"fmt"
)

// JsonEncode 对象转成JSON数据
func JsonEncode(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("转换Json字符串失败，err：%v!\n", err)
		return ""
	}
	return string(jsonByte)
}

// JsonDecode JSON数据解析成对象
func JsonDecode(data string, val interface{}) error {
	return json.Unmarshal([]byte(data), val)
}
