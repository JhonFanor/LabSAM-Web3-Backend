package gormmanagers

import (
	"fmt"
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/logging"
	"lamsam-web3-backend/internal/utils"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormQueryManager struct {
	Logger logging.Logger
}

func NewGormQueryManager(logger logging.Logger) *GormQueryManager {
	return &GormQueryManager{
		Logger: logger,
	}
}

func (m *GormQueryManager) ApplyPaginationAndFilters(c *gin.Context, db *gorm.DB, model any) *dto.PaginationDTO {
	page := c.DefaultQuery(consts.PageQueryParam, "1")
	limit := c.DefaultQuery(consts.LimitQueryParam, "10")
	sortBy := c.DefaultQuery(consts.SortByQueryParam, "id")
	sortOrder := c.DefaultQuery("order", consts.SortOperatorAsc)

	var pageInt, limitInt int
	fmt.Sscanf(page, "%d", &pageInt)
	fmt.Sscanf(limit, "%d", &limitInt)

	var totalRecords int64
	db = applyFilters(db, c, model)
	db.Model(model).Count(&totalRecords)

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

	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem() 
	}
	dataSliceType := reflect.SliceOf(modelType)
	dataPtr := reflect.New(dataSliceType)
	data := dataPtr.Interface()         

	result := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
		Distinct().
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
	parts := strings.Split(relatedModelField, ".")
	if len(parts) < 2 {
		return db, "", fmt.Errorf("relación inválida: %s", relatedModelField)
	}

	currentTable := utils.GetTableNameFromModel(model)

	for i := 0; i < len(parts)-1; i++ {
		parent := currentTable
		child := parts[i]

		if isManyToMany(db, parent, child) {
			joinTable := getJoinTableForManyToMany(db, parent, child)
			childTable := utils.ResolveTableName(child)

			parentForeignKey := getForeignKey(parent, child)
			childForeignKey := getForeignKey(child, parent)
			// Ajustar la lógica para el JOIN many-to-many de manera genérica
			joinClause := fmt.Sprintf(
				"JOIN %s ON %s.%s = %s.id JOIN %s ON %s.%s = %s.id",
				joinTable, joinTable, childForeignKey, parent, childTable, joinTable, parentForeignKey, childTable,
			)

			joinKey := fmt.Sprintf("%s.%s", parent, childTable)
			if !joinedTables[joinKey] {
				db = db.Joins(joinClause)
				joinedTables[joinKey] = true
			}
			currentTable = childTable
		} else {
			// Si no es una relación many-to-many, hacer un join normal
			childTable := utils.ResolveTableName(child)
			joinKey := fmt.Sprintf("%s.%s", parent, childTable)

			if !joinedTables[joinKey] {
				// Obtener el nombre de la clave foránea genérica
				foreignKey := getForeignKey(parent, child)

				joinClause := fmt.Sprintf("JOIN %s ON %s.id = %s.%s", childTable, childTable, parent, foreignKey)
				db = db.Joins(joinClause)
				joinedTables[joinKey] = true
			}
			currentTable = childTable
		}
	}

	finalField := fmt.Sprintf("%s.%s", currentTable, parts[len(parts)-1])
	return db, finalField, nil
}

func getForeignKey(parentTable, childTable string) string {
	return fmt.Sprintf("%s_id", childTable)
}

func isManyToMany(db *gorm.DB, parentTable string, childTable string) bool {
	intermediateTable := fmt.Sprintf("%s_%s", parentTable, childTable)

	var count int64
	db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = ?", intermediateTable).Scan(&count)

	return count > 0
}

func getJoinTableForManyToMany(db *gorm.DB, parentTable string, childTable string) string {
	if parentTable < childTable {
		return fmt.Sprintf("%s_%s", parentTable, childTable)
	}
	return fmt.Sprintf("%s_%s", childTable, parentTable)
}

func applyFilters(db *gorm.DB, c *gin.Context, model interface{}) *gorm.DB {
	joinedTables := make(map[string]bool)

	for key, values := range c.Request.URL.Query() {
		for _, value := range values {
			log.Println("Filtro recibido:", key, value)
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
		log.Print("entro en el join")
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
