package gormmanagers

import (
	"fmt"
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/logging"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormQueryManager struct {
	DB     *gorm.DB
	Logger logging.Logger
}

func NewGormQueryManager(db *gorm.DB) *GormQueryManager {
	return &GormQueryManager{DB: db}
}

func (m *GormQueryManager) ApplyPaginationAndFilters(c *gin.Context, model any) *dto.PaginationDTO {
	page := c.DefaultQuery(consts.PageQueryParam, "1")
	limit := c.DefaultQuery(consts.LimitQueryParam, "10")
	sortBy := c.DefaultQuery(consts.SortByQueryParam, "id")
	sortOrder := c.DefaultQuery("order", consts.SortOperatorAsc)

	var pageInt, limitInt int
	fmt.Sscanf(page, "%d", &pageInt)
	fmt.Sscanf(limit, "%d", &limitInt)

	var totalRecords int64
	m.DB = applyFilters(m.DB, c, model)
	m.DB.Model(model).Count(&totalRecords)

	totalPages := (totalRecords + int64(limitInt) - 1) / int64(limitInt)
	if pageInt < 1 {
		pageInt = 1
	}
	offset := (pageInt - 1) * limitInt
	prevPage := int64(pageInt - 1)
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := int64(pageInt + 1)
	if nextPage > totalPages {
		nextPage = totalPages
	}

	// Crear una slice del mismo tipo que model
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem() // Obtener el tipo base si es un puntero
	}
	dataSliceType := reflect.SliceOf(modelType)
	dataPtr := reflect.New(dataSliceType) // Creamos un puntero a la slice
	data := dataPtr.Interface()           // Convertimos el puntero a interfaz vacía

	// Ejecutar la consulta con la slice correcta
	result := m.DB.Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
		Limit(limitInt).
		Offset(offset).
		Find(data)

	if result.Error != nil {
		m.Logger.LogError("Error fetching data: ", result.Error)
	}

	return &dto.PaginationDTO{
		TotalRecord: totalRecords,
		TotalPage:   totalPages,
		Offset:      int64(offset),
		Limit:       int64(limitInt),
		Page:        int64(pageInt),
		PrevPage:    prevPage,
		NextPage:    nextPage,
		Data:        dataPtr.Elem().Interface(),
	}
}

func applyDynamicJoin(relatedModelField string, model interface{}, db *gorm.DB, joinedTables map[string]bool) (*gorm.DB, string, error) {
	relationParts := strings.SplitN(relatedModelField, consts.RelatedFieldFilterSeparator, 2)
	if len(relationParts) != 2 {
		return db, "", fmt.Errorf("relación inválida: %s", relatedModelField)
	}

	relation := relationParts[0]
	field := relationParts[1]

	if _, exists := joinedTables[relation]; !exists {
		db = db.Joins(relation)
		joinedTables[relation] = true
	}

	relationField := fmt.Sprintf("%s.%s", relation, field)
	return db, relationField, nil
}

func applyFilters(db *gorm.DB, c *gin.Context, model interface{}) *gorm.DB {
	joinedTables := make(map[string]bool)

	for key, values := range c.Request.URL.Query() {
		for _, value := range values {
			db = processFilter(db, key, value, model, joinedTables)
		}
	}
	return db
}

func processFilter(db *gorm.DB, key, value string, model interface{}, joinedTables map[string]bool) *gorm.DB {
	parts := strings.SplitN(value, consts.OperatorFilterSeparator, 2)
	if len(parts) != 2 {
		return db
	}

	field := key
	operator, filterValue := parts[0], parts[1]

	if strings.Contains(field, consts.RelatedFieldFilterSeparator) {
		var err error
		db, field, err = applyDynamicJoin(field, model, db, joinedTables)
		if err != nil {
			return db
		}
	}

	return applyFilterCondition(db, field, operator, filterValue)
}

func applyFilterCondition(db *gorm.DB, field, operator, filterValue string) *gorm.DB {
	switch operator {
	case consts.FilterOperatorEq:
		return db.Where(fmt.Sprintf("%s = ?", field), filterValue)
	case consts.FilterOperatorNeq:
		return db.Where(fmt.Sprintf("%s != ?", field), filterValue)
	case consts.FilterOperatorLt:
		return db.Where(fmt.Sprintf("%s < ?", field), filterValue)
	case consts.FilterOperatorGt:
		return db.Where(fmt.Sprintf("%s > ?", field), filterValue)
	case consts.FilterOperatorLeq:
		return db.Where(fmt.Sprintf("%s <= ?", field), filterValue)
	case consts.FilterOperatorGeq:
		return db.Where(fmt.Sprintf("%s >= ?", field), filterValue)
	case consts.FilterOperatorContains:
		return db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+filterValue+"%")
	case consts.FilterOperatorIn:
		values := strings.Split(filterValue, consts.InOptionsSeparator)
		return db.Where(fmt.Sprintf("%s IN (?)", field), values)
	case consts.FilterOperatorNin:
		values := strings.Split(filterValue, consts.InOptionsSeparator)
		return db.Where(fmt.Sprintf("%s NOT IN (?)", field), values)
	default:
		return db
	}
}
