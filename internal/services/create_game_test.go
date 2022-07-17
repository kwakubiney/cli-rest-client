package services_test

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateGameEndpoint(t *testing.T) {
	f := faker.New()
	req := services.MakeTestRequest(t, "/Game", map[string]interface{}{
		"age_rating":  "12+",
		"url":         f.Internet().URL(),
		"description": f.Lorem().Sentence(200),
		"publisher":   f.Company().Name(),
		"title":       f.App().Name(),
	}, "POST")

	response := services.BootstrapServer(req, engine)
	responseBody := services.DecodeResponse(t, response)
	assert.Equal(t, "game successfully created", responseBody["message"])
}