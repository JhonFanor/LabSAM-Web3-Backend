package reflection

import (
	"errors"
	"fmt"
	"lamsam-web3-backend/internal/consts"
	"reflect"
	"strings"
)

// ParseField parse a field's name into its parts (including relation).
func ParseField(field string) ([]string, error) {
	parts := strings.Split(field, consts.RelatedFieldFilterSeparator)
	if len(parts) == 0 {
		return nil, errors.New("void field")
	}
	return parts, nil
}

// IsValidField verifies if a field exists in a model (including relations).
func IsValidField(model interface{}, field string) (bool, error) {
	fieldParts, err := ParseField(field)

	if err != nil {
		return false, err
	}

	return isValidFieldRecursive(reflect.TypeOf(model), fieldParts, field), nil
}

// isValidFieldRecursive verifies if a field exists in a model recursively.
func isValidFieldRecursive(modelType reflect.Type, fieldParts []string, originalField string) bool {
	if len(fieldParts) == 0 {
		return false
	}

	currentField := fieldParts[0]

	// Verify if the model is a pointer
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	// The model must be a struct
	if modelType.Kind() != reflect.Struct {
		return false
	}

	// Search for the field in the model
	field, found := modelType.FieldByNameFunc(func(name string) bool {
		return strings.EqualFold(name, currentField) // Case-insensitive comparison
	})

	if !found {
		return false
	}

	// Base case: if it's the last field, return true
	if len(fieldParts) == 1 {
		return true
	}

	// If it's not the last field, continue recursively
	return isValidFieldRecursive(field.Type, fieldParts[1:], originalField)
}

func GetFieldType(model interface{}, field string) (reflect.Type, error) {
	value := reflect.ValueOf(model)
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	typ := value.Type()

	fieldValue, ok := typ.FieldByName(field)
	if !ok {
		return nil, fmt.Errorf("field '%s' not found in type '%s'", field, typ.Name())
	}
	// Get the type of the field
	typ = fieldValue.Type

	// If the type is a pointer or slice, get the underlying type
	for typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}

	return typ, nil
}
