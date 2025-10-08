package utils

import (
	"reflect"
	"strings"
)

func StructToMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}

	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})

	ignoreFields := map[string]bool{
		"localitation": true,
		"user":         true,
		"subtopics":    true,
		"created_at":   true,
		"updated_at":   true,
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
		name := strings.Split(jsonTag, ",")[0]
		if name == "" {
			continue
		}

		if ignoreFields[name] {
			continue
		}

		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			if field.IsNil() {
				result[name] = nil
				continue
			}
		}

		if field.Kind() == reflect.Struct && field.Type().String() != "time.Time" {
			result[name] = StructToMap(field.Interface())
			continue
		}

		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			slice := []interface{}{}
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Struct {
					slice = append(slice, StructToMap(elem.Interface()))
				} else {
					slice = append(slice, elem.Interface())
				}
			}
			result[name] = slice
			continue
		}

		result[name] = field.Interface()
	}

	return result
}
