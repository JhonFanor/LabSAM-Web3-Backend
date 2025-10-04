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

		if name == "is_approved" {
			result[name] = field.Interface()
			continue
		}

		if isZeroValue(field) {
			continue
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

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		if z, ok := v.Interface().(interface{ IsZero() bool }); ok {
			return z.IsZero()
		}
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
