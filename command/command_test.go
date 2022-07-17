package command_test

import (
	"testing"

	"github.com/kwakubiney/canonical-take-home/command"
	"github.com/stretchr/testify/assert"
)

func TestParseFields(t *testing.T) {
	arrayOfParsedFields, mapOfData := command.ParseFields("username=kwame, age=9, email=kwakubiney@gmail.com")
	assert.Equal(t, []string{"username", "age", "email"}, arrayOfParsedFields)
	assert.Equal(t, map[string]string{"username": "kwame", "age": "9", "email": "kwakubiney@gmail.com"}, mapOfData)
	arrayOfParsedFields, mapOfData = command.ParseFields("username=kwame,age=9,email=kwakubiney@gmail.com")
	assert.Equal(t, []string{"username", "age", "email"}, arrayOfParsedFields)
	assert.Equal(t, map[string]string{"username": "kwame", "age": "9", "email": "kwakubiney@gmail.com"}, mapOfData)
}

func TestValidateFields(t *testing.T) {
	ok := command.ValidateFields("create", []string{"username", "age", "email"}, "user")
	assert.Equal(t, true, ok)

	ok = command.ValidateFields("create", []string{"use", "age"}, "game")
	assert.Equal(t, false, ok)

	ok = command.ValidateFields("update", []string{"user", "age"}, "user")
	assert.Equal(t, false, ok)

	ok = command.ValidateFields("update", []string{"username", "age", "email"}, "user")
	assert.Equal(t, true, ok)
}
