package utils

import (
	"encoding/json"
)

func StructToMap(input interface{}) map[string]interface{} {
	var result map[string]interface{}

	data, _ := json.Marshal(input)
	json.Unmarshal(data, &result)

	cleanedResult := make(map[string]interface{})
	for key, value := range result {
		if value != nil {
			cleanedResult[key] = value
		}
	}

	return cleanedResult
}
