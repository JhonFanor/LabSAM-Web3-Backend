package utils

import "reflect"

func GetModifiedFields[T any](old, new *T) map[string]interface{} {
	updates := make(map[string]interface{})
	oldVal := reflect.ValueOf(*old)
	newVal := reflect.ValueOf(*new)
	typ := oldVal.Type()

	for i := 0; i < oldVal.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" || jsonTag == "-" || field.Name == "ID" {
			continue
		}

		if oldVal.Field(i).Interface() != newVal.Field(i).Interface() {
			updates[jsonTag] = newVal.Field(i).Interface()
		}
	}

	return updates
}
