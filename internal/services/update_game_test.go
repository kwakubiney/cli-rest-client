package services_test

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestUpdateGameEndpoint(t *testing.T) {
	game := services.CreateTestGame(t)
	services.SeedDB(game)
	f := faker.New()
	updateUserRequest := services.MakeTestRequest(t, fmt.Sprintf("/Game?title=%s", game.Title), map[string]interface{}{
		"age_rating":  "12+",
		"url":         f.Internet().URL(),
		"description": f.Lorem().Sentence(200),
		"publisher":   f.Company().Name(),
	}, "PUT")

	getUserResponse := services.BootstrapServer(updateUserRequest, engine)
	responseBody := services.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "game successfully updated", responseBody["message"])
}
