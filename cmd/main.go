package main

import (
	"log"
	"os"

	"github.com/kwakubiney/canonical-take-home/internal/config"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
	"github.com/kwakubiney/canonical-take-home/internal/postgres"
	"github.com/kwakubiney/canonical-take-home/internal/server"
	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/kwakubiney/canonical-take-home/parser"
)

func main() {
	opts := &cli.Options{}
	parser.ParseCommands(os.Args, opts)
	config.LoadNormalConfig()
	db, err := postgres.Init()
	if err != nil {
		log.Fatal("database failed to start.")
	}
	userRepo := repository.NewUserRepository(db)
	gameRepo := repository.NewGameRepository(db)
	
	newCliHandler := cli.NewCliHandler(opts)
	err = newCliHandler.Dispatch()
	if err != nil {
		opts.Flag.Usage()
		log.Fatal()
	}
	userService := services.NewUserService(userRepo, newCliHandler)
	gameService := services.NewGameService(gameRepo, newCliHandler)
	server := server.New(userService, gameService)
	server.Start()

	err = cli.ApiRequestDispatcher(newCliHandler)
	if err != nil {
		opts.Flag.Usage()
		log.Fatal()
	}
}
