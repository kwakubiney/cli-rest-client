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
	"github.com/kwakubiney/canonical-take-home/internal/utils"
)
const (

)
func main() {
	opts := &cli.Options{}
	flag.StringVar(&opts.Method, "m", "", "Specify method to retrieve user(s) data with flag values:\n"+
		"create, read, update, delete\n"+
		"Example: -m create")

	flag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
		"user, games\n"+
		"Example: -type user")

	flag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
		"-fields username=kwame,age=9,email=kwakubiney@gmail.com,gamehours=9")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	opts.Help = flag.Bool("help", false, "Show usage")
	flag.Parse()
	config.LoadNormalConfig()
	db, err := postgres.Init()
	if err != nil {
		log.Println(err)
	}

	//set db up here
	repo := repository.NewUserRepository(db)
	newCliHandler := cli.NewCliHandler(opts)
	err = newCliHandler.Dispatch()
	if err != nil {
		flag.Usage()
	}
	userService := services.NewUserService(repo, newCliHandler)
	server := server.New(userService)
	server.Start()
	resp, err := utils.MakeRequest("http://127.0.0.1:8233/createUser/", "8000", nil, "GET")
	if err !=nil{
		log.Println(err)
	}
	log.Println(resp)
}



