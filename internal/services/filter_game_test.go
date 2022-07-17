package services_test

import (
	"testing"

	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestFilterGameEndpoint(t *testing.T) {
	game := services.CreateTestGame(t)
	services.SeedDB(game)
	updateUserRequest := services.MakeTestRequest(t, fmt.Sprintf("/Game?title=%s", game.Title), map[string]interface{}{
		"by": "title",
	} , "GET")

	getUserResponse := services.BootstrapServer(updateUserRequest, engine)
	responseBody := services.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "game successfully filtered", responseBody["message"])
}