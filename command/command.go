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

func ValidateCreateandUpdateUserFields(method string, fieldKeys []string) bool {
	createUserFields := []string{"username", "age", "email", "gamehours"}
	if method == "create" {
		if len(createUserFields) != len(fieldKeys) {
			return false
		}
	}
	exists := make(map[string]bool)
	for _, value := range createUserFields {
		exists[value] = true
	}

	for _, value := range fieldKeys {
		if !exists[value] {
			return false
		}
	}
	return true
}
