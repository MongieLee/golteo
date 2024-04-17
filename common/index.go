package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// JsonEncode 结构体转成JSON数据
func JsonEncode(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("转换Json字符串失败，err：%v!\n", err)
		return ""
	}
	return string(jsonByte)
}

// JsonDecode JSON数据解析成结构体
func JsonDecode(data string, val interface{}) error {
	return json.Unmarshal([]byte(data), val)
}

// StructToMap 结构体转成Map
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		result[typeOfT.Field(i).Name] = field.Interface()
	}
	return result
}
