package tools

import "reflect"

func IsNilPointer(i interface{}) bool {
	iv := reflect.ValueOf(i)
	return iv.Kind() != reflect.Ptr || iv.IsNil()
}
