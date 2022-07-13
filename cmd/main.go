package main

import (
	"flag"
	"fmt"

	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
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
		"-type User or -type Game")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	opts.Help = flag.Bool("help", false, "Show usage")
	flag.Parse()

	baseUrl := ""
	repo := repository.NewUserRepository(baseUrl)

	userService := cli.NewUserService(*repo)

	handler := cli.NewCliHandler(*opts, *userService)

	handler.Dispatch()
}
