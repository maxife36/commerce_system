package helpers

import (
	"errors"
	"reflect"
)

func FilterStruct(obj map[string]any, notAllowedAttr ...string) *map[string]any {

	filtered := make(map[string]any)

	for _, attr := range notAllowedAttr {
		if value, exists := obj[attr]; exists {
			filtered[attr] = value
		}
	}

	return &filtered
}

func GetKeysStruct(obj any) (map[string]reflect.Type, error) {

	t := reflect.TypeOf(obj)

	if t.Kind() != reflect.Struct {
		return nil, errors.New("the parameter must be type structure")
	}

	save := make(map[string]reflect.Type, 0)

	extractFields(t, save)

	return save, nil
}

func extractFields(t reflect.Type, save map[string]reflect.Type) {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous {
			extractFields(field.Type, save)
			continue
		}

		save[field.Name] = field.Type
	}
}
