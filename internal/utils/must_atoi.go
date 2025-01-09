package utils

import (
	"log"
	"strconv"
)

func MustAtoi(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Failed to parse integer from string: %v", err)
	}
	return value
}
