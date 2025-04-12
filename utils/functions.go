package utils

import "reflect"

func StructHasField(structure interface{}, fieldName string) bool {
	// 获取结构体的反射值
	val := reflect.ValueOf(structure)

	// 确保传入的是一个结构体
	if val.Kind() != reflect.Struct {
		return false
	}

	// 获取结构体类型
	typ := val.Type()

	// 遍历字段，查找指定字段
	for i := 0; i < val.NumField(); i++ {
		if typ.Field(i).Name == fieldName {
			return true
		}
	}

	return false
}
