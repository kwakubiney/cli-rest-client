package parser

import (
	"flag"
	"fmt"

	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
)

func ParseCommands(args []string, opts *cli.Options) {
	switch args[1] {
	case "create":
		createFlag := flag.NewFlagSet("create", flag.ExitOnError)
		opts.Flag = *createFlag
		opts.Method = args[1]

		createFlag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
			"Example: -type user")

		createFlag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
			"Example: -fields username=kwame,age=9,email=kwakubiney@gmail.com")

		createFlag.Parse(args[2:])
		createFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			createFlag.PrintDefaults()
		}

	case "update":
		updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
		opts.Flag = *updateFlag
		opts.Method = args[1]
		updateFlag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
			"Example: -type user")

		updateFlag.StringVar(&opts.Fields, "fields", "", "Specify fields to retrieve, create or update user(s) data, available repositories:\n"+
			"Example: -fields username=kwame,age=9,email=kwakubiney@gmail.com")

		updateFlag.StringVar(&opts.Where, "where", "", "Specify where to retrieve or update user data:\n"+
			"Example(where is limited to object's name): -where kb")

		updateFlag.Parse(args[2:])
		updateFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			updateFlag.PrintDefaults()
		}

	case "filter":
		filterFlag := flag.NewFlagSet("filter", flag.ExitOnError)
		opts.Flag = *filterFlag
		opts.Method = args[1]
		filterFlag.StringVar(&opts.TypeOfObject, "type", "", "Specify type of repository:\n"+
			"Example: -type user")
		filterFlag.StringVar(&opts.Where, "where", "", "Specify specific 'by' value")
		filterFlag.StringVar(&opts.By, "by", "", "Specify field to filter on")
		filterFlag.Parse(args[2:])
		filterFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			filterFlag.PrintDefaults()
		}
	default:
		flag.Usage = func() {
			fmt.Printf("Usage:\n")
			//TODO: Replace default Go usage.
			flag.PrintDefaults()
		}
	}

}
