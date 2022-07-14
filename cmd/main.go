package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kwakubiney/canonical-take-home/internal/config"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
	"github.com/kwakubiney/canonical-take-home/internal/postgres"
	"github.com/kwakubiney/canonical-take-home/internal/server"
	"github.com/kwakubiney/canonical-take-home/internal/services"
)

func main() {
	opts := &cli.Options{}
	flag.StringVar(&opts.Method, "m", "", "Specify method to retrieve user data with flag values:\n"+
		"Example: -m create")

	flag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
		"Example: -type user")

	flag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
		"Example: -fields username=kwame,age=9,email=kwakubiney@gmail.com")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	opts.Help = flag.Bool("help", false, "Show usage")
	flag.Parse()
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



