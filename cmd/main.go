package main

import (
	"log"
	"os"
	"flag"
	"fmt"

	"github.com/kwakubiney/canonical-take-home/internal/config"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
	"github.com/kwakubiney/canonical-take-home/internal/postgres"
	"github.com/kwakubiney/canonical-take-home/internal/server"
	"github.com/kwakubiney/canonical-take-home/internal/services"
	"github.com/kwakubiney/canonical-take-home/parser"
)


const genericUsageString = 
`
These are common subcommands used in various situations:
	create     Create a new record in the database
	filter     Filter a record in the database
	update     Update a record in the database`



func main() {
	opts := &cli.Options{}
	flag.Usage = func() {
		fmt.Printf("Usage:[--help]\n")
		fmt.Println(genericUsageString)
	}
	flag.Parse()
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
		log.Println(opts.Flag)
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
