package utils

import "reflect"

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
