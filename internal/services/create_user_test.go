package services_test

import (
	"fmt"
	"log"
	"os"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
	"github.com/kwakubiney/canonical-take-home/internal/config"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
	"github.com/kwakubiney/canonical-take-home/internal/postgres"
	"github.com/kwakubiney/canonical-take-home/internal/server"
	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/stretchr/testify/assert"
)

var engine *gin.Engine
func TestMain(m *testing.M) {
	opts := &cli.Options{}
	err := config.LoadTestConfig()
	assert.NoError(&testing.T{}, err)
	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	gameRepo := repository.NewGameRepository(db)
	newCliHandler := cli.NewCliHandler(opts)
	userService := services.NewUserService(userRepo, newCliHandler)
	gameService := services.NewGameService(gameRepo, newCliHandler)
	server := server.New(userService, gameService)
	engine = server.SetupRoutes()
	os.Exit(m.Run())
}

func TestCreateUserEndpoint(t *testing.T) {
	f := faker.New()
	req := services.MakeTestRequest(t, "/User", map[string]interface{}{
		"username": "kwamz",
		"email":    f.Person().Contact().Email,
		"age":      fmt.Sprint(f.RandomNumber(2)),
	}, "POST")

	response := services.BootstrapServer(req, engine)
	responseBody := services.DecodeResponse(t, response)
	assert.Equal(t, "user successfully created", responseBody["message"])
}

