package utils

import (
	"reflect"
	"strings"
)

func GetTableNameFromModel(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strings.ToLower(t.Name())
}
