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

func TestValidateCreateAndUpdateUserFields(t *testing.T) {
	ok := command.ValidateCreateandUpdateUserFields("create", []string{"username", "age", "email"})
	assert.Equal(t, true, ok)

	ok = command.ValidateCreateandUpdateUserFields("create", []string{"user", "age"})
	assert.Equal(t, false, ok)

	ok = command.ValidateCreateandUpdateUserFields("update", []string{"user", "age"})
	assert.Equal(t, false, ok)

	ok = command.ValidateCreateandUpdateUserFields("update", []string{"username", "age", "email"})
	assert.Equal(t, true, ok)
}
