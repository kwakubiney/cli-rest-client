package services_test

import (
	"testing"

	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestFilterUser200(t *testing.T) {
	user := services.CreateTestUser(t)
	services.SeedDB(user)
	updateUserRequest := services.MakeTestRequest(t, fmt.Sprintf("/User?username=%s", user.Username), map[string]interface{}{
		"by": "username",
	} , "GET")

	getUserResponse := services.BootstrapServer(updateUserRequest, engine)
	responseBody := services.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully filtered", responseBody["message"])
}