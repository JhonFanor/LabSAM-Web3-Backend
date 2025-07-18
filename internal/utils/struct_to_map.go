package utils

import (
	"reflect"
)

func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)

		if !field.CanInterface() {
			continue
		}

		jsonTag := typeField.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		if jsonTag == "is_approved" {
			result[jsonTag] = field.Interface()
			continue
		}

		if isZeroValue(field) {
			continue
		}

		result[jsonTag] = field.Interface()
	}

	return result
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
