package command_test

import (
	"testing"

	"github.com/kwakubiney/canonical-take-home/command"
	"github.com/stretchr/testify/assert"
)

func TestParseFields(t *testing.T) {
	//spaces in string
	arrayOfParsedFields, mapOfData := command.ParseFields("username=kwame, age=9, email=kwakubiney@gmail.com")
	assert.Equal(t, []string{"username", "age", "email"}, arrayOfParsedFields)
	assert.Equal(t, map[string]string{"username": "kwame", "age": "9", "email": "kwakubiney@gmail.com"}, mapOfData)
	arrayOfParsedFields, mapOfData = command.ParseFields("username=kwame,age=9,email=kwakubiney@gmail.com")
	assert.Equal(t, []string{"username", "age", "email"}, arrayOfParsedFields)
	assert.Equal(t, map[string]string{"username": "kwame", "age": "9", "email": "kwakubiney@gmail.com"}, mapOfData)
}

func TestValidateCreateAndUpdateUserFields(t *testing.T) {
	//all arguments for create
	ok := command.ValidateCreateandUpdateUserFields("create", []string{"username", "age", "email"})
	assert.Equal(t, true, ok)

	//missing arguments for create
	ok = command.ValidateCreateandUpdateUserFields("create", []string{"user", "age"})
	assert.Equal(t, false, ok)

	//wrong arguments for update
	ok = command.ValidateCreateandUpdateUserFields("update", []string{"user", "age"})
	assert.Equal(t, false, ok)

	//correct arguments for update
	ok = command.ValidateCreateandUpdateUserFields("update", []string{"username", "age", "email"})
	assert.Equal(t, true, ok)
}
