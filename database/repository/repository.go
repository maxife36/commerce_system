package repository

import (
	"commerce-system/database"
	"commerce-system/helpers"
	"commerce-system/models"
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

type Repository[M Models] interface {
	Create() (*M, error)
	GetAll() ([]M, error)
	GetById() (*M, error)
	GetByFilter() ([]M, error)
	Update() (*M, error)
	Delete() (*int64, error)
}

type Models interface {
	models.Category | models.Sale | models.User | models.Product
}

type FilterConditions struct {
	Operator string //(like, eq, neq, lt, lte, gt, gte )
	value    any
}

type baseModelDao[M Models] struct {
	Model  M
	Filter map[string][]FilterConditions
}

func NewDao[M Models](model M) *baseModelDao[M] {
	return &baseModelDao[M]{
		Model: model,
	}
}

func (md *baseModelDao[M]) Create() (*M, error) {
	result := db.Create(&md.Model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &md.Model, nil
}

func (md *baseModelDao[M]) GetAll() ([]M, error) {
	var records []M

	result := db.Find(&records)

	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (md *baseModelDao[M]) GetById() (*M, error) {

	if idChecker(md.Model) {
		return nil, errors.New("not valid ID")
	}

	result := db.First(&md.Model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &md.Model, nil
}

func (md *baseModelDao[M]) GetByFilter() ([]M, error) {
	// Verifico si se suministro la configuracion del filtro

	if md.Filter == nil {
		return nil, errors.New("unspecified filter")
	}

	allAttr, err := helpers.GetKeysStruct(md.Model)

	if err != nil {
		return nil, err
	}

	var records []M

	query := db.Model(md.Model)

	for field, conditions := range md.Filter {
		if fieldType, exists := allAttr[field]; exists {
			query, err = queryConfig(query, field, fieldType, conditions)

			if err != nil {
				return nil, err
			}
		}
	}

	result := query.Find(&records)

	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (md *baseModelDao[M]) Update() (*M, error) {
	if idChecker(md.Model) {
		return nil, errors.New("not valid ID")
	}

	result := db.Save(&md.Model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &md.Model, nil
}

func (md *baseModelDao[M]) Delete() (*int64, error) {
	if idChecker(md.Model) {
		return nil, errors.New("not valid ID")
	}

	result := db.Delete(&md.Model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &result.RowsAffected, nil
}

func queryConfig(db *gorm.DB, field string, fieldType reflect.Type, conditions []FilterConditions) (*gorm.DB, error) {
	for _, cond := range conditions {

		var queryStr string

		allowedOperators := map[string]string{
			"like": "LIKE",
			"eq":   "=",
			"neq":  "!=",
			"lt":   "<",
			"lte":  "<=",
			"gt":   ">",
			"gte":  ">=",
		}

		if operatorStr, exists := allowedOperators[cond.Operator]; exists {
			queryStr = fmt.Sprintf("%s %s ?", field, operatorStr)
		}

		if queryStr != "" {
			valueType := reflect.TypeOf(cond.value)

			if valueType != fieldType {
				return nil, fmt.Errorf("on field %s: expected %s, obteined %s", field, fieldType, valueType)
			}

			db = db.Where(queryStr, cond.value)
		}
	}

	return db, nil
}

func idChecker[M Models](model M) bool {
	v := reflect.ValueOf(model)

	field := v.FieldByName("ID")

	if !field.IsValid() || !field.CanInterface() {
		return false
	}

	if field.Kind() == reflect.Uint && field.Uint() != 0 {
		return true
	}

	return false
}
