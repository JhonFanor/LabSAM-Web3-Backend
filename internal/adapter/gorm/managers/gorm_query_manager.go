package gormmanagers

import (
	"context"
	"fmt"
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/logging"
	"lamsam-web3-backend/internal/queryparams"
	"lamsam-web3-backend/internal/reflection"
	"lamsam-web3-backend/internal/utils"
	"reflect"
	"strings"

	"github.com/jinzhu/inflection"
	"gorm.io/gorm"
)

type GormQueryManager struct {
	DB     *gorm.DB
	Logger logging.Logger
}

func NewGormQueryManager(db *gorm.DB) *GormQueryManager {
	return &GormQueryManager{DB: db}
}

func (m *GormQueryManager) ApplyPaginationAndFilters(ctx context.Context, tx *gorm.DB, model any) (*gorm.DB, *dto.PaginationDTO, error) {
	paginationInfo := &dto.PaginationDTO{}
	joinedTables := make(map[string]bool)

	if sort, ok := ctx.Value(consts.ContextKeySort).(*queryparams.SortQueryParams); ok {
		valid, err := reflection.IsValidField(model, sort.Field)

		if err != nil {
			return nil, nil, err
		}

		if !valid {
			return nil, nil, err
		}

		fields, err := reflection.ParseField(sort.Field)

		if err != nil {
			return nil, nil, err
		}

		if len(fields) > 1 {
			relatedModel := fields[0]
			fieldToSort := fields[1]

			dbJoin, relatedTable, err := applyDynamicJoin(relatedModel, model, tx, joinedTables)

			if err != nil {
				return nil, nil, err
			}

			if utils.IsCamelCase(fieldToSort) {
				fieldToSort = utils.CamelToSnake(fieldToSort)
			} else {
				fieldToSort = utils.PascalToSnake(fieldToSort)
			}

			tx = dbJoin.Order(fmt.Sprintf("%s %s", fmt.Sprintf("%s.%s", relatedTable, fieldToSort), sort.Order))
		} else {
			if utils.IsCamelCase(fields[0]) {
				fields[0] = utils.CamelToSnake(fields[0])
			} else {
				fields[0] = utils.PascalToSnake(fields[0])
			}

			tx = tx.Order(fmt.Sprintf("%s %s", fields[0], sort.Order))
		}
	}

	// Apply filters if they're present in the context
	if filters, ok := ctx.Value(consts.ContextKeyFilters).(queryparams.Filters); ok {
		for i, filter := range filters {
			// Check if is a valid field
			valid, err := reflection.IsValidField(model, filter.Field)

			if err != nil {
				return nil, nil, err
			}

			if !valid {
				return nil, nil, err
			}

			fields, err := reflection.ParseField(filter.Field)

			if err != nil {
				return nil, nil, err
			}

			// Check if the filter field is a relation
			if len(fields) > 1 {
				// Preload the necessary relations
				relatedModelField := fields[0]
				fieldToFilter := fields[1]

				// Join with the related table and apply filter
				dbJoin, relatedTable, err := applyDynamicJoin(relatedModelField, model, tx, joinedTables)

				if err != nil {
					return nil, nil, err
				}

				tx = dbJoin

				if utils.IsCamelCase(fieldToFilter) {
					fieldToFilter = utils.CamelToSnake(fieldToFilter)
				} else {
					fieldToFilter = utils.PascalToSnake(fieldToFilter)
				}

				filters[i].Field = fmt.Sprintf("%s.%s", relatedTable, fieldToFilter)
			} else {
				var fieldToFilter string
				if utils.IsCamelCase(fields[0]) {
					fieldToFilter = utils.CamelToSnake(fields[0])
				} else {
					fieldToFilter = utils.PascalToSnake(fields[0])
				}
				filters[i].Field = fieldToFilter
			}
		}

		tx = applyFilters(tx, filters)
	}

	if pagination, ok := ctx.Value(consts.ContextKeyPagination).(*queryparams.PaginationQueryParams); ok {
		var totalRecords int64

		if err := tx.Model(model).Count(&totalRecords).Error; err != nil {
			return nil, nil, err
		}

		paginationInfo.TotalRecords = int(totalRecords)
		paginationInfo.TotalPages = int((totalRecords + int64(pagination.Limit) - 1) / int64(pagination.Limit))
		paginationInfo.CurrentPage = pagination.Page
		paginationInfo.PageSize = pagination.Limit

		tx = tx.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit)
	}

	return tx, paginationInfo, nil
}

func applyDynamicJoin(relatedModelField string, model any, db *gorm.DB, joinedTables map[string]bool) (*gorm.DB, string, error) {
	// Get the real related model name
	if utils.IsCamelCase(relatedModelField) {
		relatedModelField = utils.CamelToPascal(relatedModelField)
	}

	relatedModelType, err := reflection.GetFieldType(model, relatedModelField)

	if err != nil {
		return nil, "", err
	}

	// parse the related model name from the type
	relatedModel := relatedModelType.Name()

	relatedTable := utils.PascalToSnake(relatedModel)

	if !utils.IsPlural(relatedTable) {
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

	modelTable := utils.PascalToSnake(modelType.Name())

	if !utils.IsPlural(modelTable) {
		modelTable = inflection.Plural(modelTable)
	}

	modelTable = strings.ToLower(modelTable)
	// Join with the related table and apply filtering
	joinKey := utils.PascalToSnake(relatedModelField) + "_id"

	db = db.Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.%s", relatedTable, relatedTable, modelTable, joinKey))

	return db, relatedTable, nil
}

func applyFilters(db *gorm.DB, filters []queryparams.FilterQueryParams) *gorm.DB {
	for _, filter := range filters {
		if isDateField(filter.Field) {
			switch filter.Operator {
			case consts.FilterOperatorEq:
				db = db.Where(fmt.Sprintf("%s = ?", filter.Field), filter.Value)
			case consts.FilterOperatorNeq:
				db = db.Where(fmt.Sprintf("%s != ?", filter.Field), filter.Value)
			case consts.FilterOperatorLt:
				db = db.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value)
			case consts.FilterOperatorGt:
				db = db.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value)
			case consts.FilterOperatorLeq:
				db = db.Where(fmt.Sprintf("%s <= ?", filter.Field), filter.Value)
			case consts.FilterOperatorGeq:
				db = db.Where(fmt.Sprintf("%s >= ?", filter.Field), filter.Value)
			case consts.FilterOperatorContains:
				// Transform the field to text for filtering
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", filter.Field), fmt.Sprintf("%%%v%%", filter.Value))
			case consts.FilterOperatorNContains:
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) NOT ILIKE ?", filter.Field), fmt.Sprintf("%%%v%%", filter.Value))
			case consts.FilterOperatorStartsWith:
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", filter.Field), fmt.Sprintf("%v%%", filter.Value))
			case consts.FilterOperatorEndsWith:
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", filter.Field), fmt.Sprintf("%%%v", filter.Value))
			}
		} else {
			switch filter.Operator {
			case consts.FilterOperatorEq:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) = ?", filter.Field), filter.Value)
			case consts.FilterOperatorNeq:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) != ?", filter.Field), filter.Value)
			case consts.FilterOperatorLt:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) < ?", filter.Field), filter.Value)
			case consts.FilterOperatorGt:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) > ?", filter.Field), filter.Value)
			case consts.FilterOperatorLeq:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) <= ?", filter.Field), filter.Value)
			case consts.FilterOperatorGeq:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) >= ?", filter.Field), filter.Value)
			case consts.FilterOperatorContains:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) ILIKE ?", filter.Field), fmt.Sprintf("%%%v%%", filter.Value))
			case consts.FilterOperatorNContains:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) NOT ILIKE ?", filter.Field), fmt.Sprintf("%%%v%%", filter.Value))
			case consts.FilterOperatorStartsWith:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) ILIKE ?", filter.Field), fmt.Sprintf("%v%%", filter.Value))
			case consts.FilterOperatorEndsWith:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) ILIKE ?", filter.Field), fmt.Sprintf("%%%v", filter.Value))
			case consts.FilterOperatorStartsWithCaseSensitive:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) LIKE ?", filter.Field), fmt.Sprintf("%v%%", filter.Value))
			case consts.FilterOperatorEndsWithCaseSensitive:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) LIKE ?", filter.Field), fmt.Sprintf("%%%v", filter.Value))
			case consts.FilterOperatorNStartsWith:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) NOT ILIKE ?", filter.Field), fmt.Sprintf("%v%%", filter.Value))
			case consts.FilterOperatorNEndsWith:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) NOT ILIKE ?", filter.Field), fmt.Sprintf("%%%v", filter.Value))
			case consts.FilterOperatorIn:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) IN (?)", filter.Field), filter.Value)
			case consts.FilterOperatorNin:
				db = db.Where(fmt.Sprintf("public.unaccent(%s) NOT IN (?)", filter.Field), filter.Value)
			}
		}
	}
	return db
}

func isDateField(field string) bool {
	dateFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"issued_at":  true,
		"expires_at": true,
	}
	return dateFields[field]
}
