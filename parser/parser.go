package parser

import (
	"flag"
	"fmt"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
)

func ParseCommands(args string, opts *cli.Options){

	switch args{
		case "create":
		createFlag := flag.NewFlagSet("create", flag.ExitOnError)

		createFlag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
		"Example: -type user")
		
		createFlag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
		"Example: -fields username=kwame,age=9,email=kwakubiney@gmail.com")

		createFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			flag.PrintDefaults()
			}
	
		case "update":
			updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
			updateFlag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
			"Example: -type user")
			
			updateFlag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
			"Example: -fields username=kwame,age=9,email=kwakubiney@gmail.com")

				
			updateFlag.StringVar(&opts.Where, "where", "", "Specify where to retrieve or update user data:\n"+
			"Example(where is limited to object's name): -where kb")

			updateFlag.Usage = func() {
				fmt.Printf("Usage:\n")
				flag.PrintDefaults()
				}

		case "filter":
			//
	} 
}
