package command

import (
	"strings"
)

func ParseFields(arguments string) ([]string, map[string]string) {
	var fieldKeys []string
	mapOfFields := make(map[string]string)
	arrayOfParsedFields := strings.Split(arguments, ",")
	for _, parsedString := range arrayOfParsedFields {
		field, value, _ := strings.Cut(strings.TrimSpace(parsedString), "=")
		mapOfFields[field] = value
		fieldKeys = append(fieldKeys, field)
	}
	return fieldKeys, mapOfFields
}

func ValidateFields(method string, fieldKeys []string, typeOfObject string) bool {
	var validFields []string

	if typeOfObject == "user"{
		validFields = []string{"username", "age", "email"}
	}else {
		validFields = []string{"description", "age_rating", "title", "publisher", "url"}
	}

	if method == "create" {
		if len(validFields) != len(fieldKeys) {
			return false
		}
	}
	exists := make(map[string]bool)
	for _, value := range validFields {
		exists[value] = true
	}

	for _, value := range fieldKeys {
		if !exists[value] {
			return false
		}
	}
	return true
}
