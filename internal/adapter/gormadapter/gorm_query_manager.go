package gormadapter

import (
	"LABSAM-WEB3-BACKEND/config"
	"LABSAM-WEB3-BACKEND/internal/reflection"
	"LABSAM-WEB3-BACKEND/internal/tools"
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/inflection"
	"gorm.io/gorm"
)

type GormQueryManager struct {
	config *config.AppConfig
}

func NewGormQueryManager(config *config.AppConfig) *GormQueryManager {
	return &GormQueryManager{
		config: config,
	}
}

func applyDynamicJoin(relatedModelField string, model any, db *gorm.DB, joinedTables map[string]bool) (*gorm.DB, string, error) {
	// Get the real related model name
	if tools.IsCamelCase(relatedModelField) {
		relatedModelField = tools.CamelToPascal(relatedModelField)
	}

	relatedModelType, err := reflection.GetFieldType(model, relatedModelField)

	if err != nil {
		return nil, "", err
	}

	// parse the related model name from the type
	relatedModel := relatedModelType.Name()

	relatedTable := tools.PascalToSnake(relatedModel)

	if !tools.IsPlural(relatedTable) {
		relatedTable = inflection.Plural(relatedTable)
	}

	// Check if the related table has already been joined
	if _, exists := joinedTables[relatedTable]; exists {
		// If it has, don't join it again
		return db, relatedTable, nil
	}

	// Add the related table to the joinedTables map
	joinedTables[relatedTable] = true

	// Get main model table name
	modelType := reflect.TypeOf(model)
	modelType = modelType.Elem()

	modelTable := tools.PascalToSnake(modelType.Name())

	if !tools.IsPlural(modelTable) {
		modelTable = inflection.Plural(modelTable)
	}

	modelTable = strings.ToLower(modelTable)
	// Join with the related table and apply filtering
	joinKey := tools.PascalToSnake(relatedModelField) + "_id"

	db = db.Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.%s", relatedTable, relatedTable, modelTable, joinKey))

	return db, relatedTable, nil
}

// isDateField returns true if the field is a date field
func isDateField(field string) bool {
	dateFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"issued_at":  true,
		"expires_at": true,
	}
	return dateFields[field]
}
