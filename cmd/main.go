package main

import (
	"flag"
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
	parser.ParseCommands(os.Args[1], opts)
	config.LoadNormalConfig()
	db, err := postgres.Init()
	if err != nil {
		flag.Usage()
		log.Fatal()
	}
	repo := repository.NewUserRepository(db)
	newCliHandler := cli.NewCliHandler(opts)
	err = newCliHandler.Dispatch()
	if err != nil {
		flag.Usage()
		log.Fatal()
	}
	userService := services.NewUserService(repo, newCliHandler)
	server := server.New(userService)
	server.Start()

	err = cli.ApiRequestDispatcher(newCliHandler)
	if err != nil {
		flag.Usage()
		log.Fatal()
	}
}



