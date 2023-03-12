package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetVaildMsg(err error, obj interface{}) string {
	// 反射获取结构体类型
	getObj := reflect.TypeOf(obj)
	// err.(Type)是为了断言err是否为Type类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 迭代errs中的内容(errs是一个切片，所以需要迭代)
		for _, e := range errs {
			// 根据结构体报错的字段名，获取结构体标签中的字段信息
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}
