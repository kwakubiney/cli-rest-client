package services_test

import (
	"fmt"
	"testing"

	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserEndpoint(t *testing.T) {
	user := services.CreateTestUser(t)
	services.SeedDB(user)
	updateUserRequest := services.MakeTestRequest(t, fmt.Sprintf("/User?username=%s", user.Username), map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	} , "PUT")

	getUserResponse := services.BootstrapServer(updateUserRequest, engine)
	responseBody := services.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully updated", responseBody["message"])
}