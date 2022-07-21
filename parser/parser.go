package parser

import (
	"flag"
	"fmt"
	"os"

	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
)

const genericUsageString = 
`
These are common subcommands used in various situations:
	create     Create a new record in the database
	filter     Filter a record in the database
	update     Update a record in the database`


const createUsageString = 
`NAME
       create - Create a new record in the database

SYNOPSIS
        create [--type] [--fields]

DESCRIPTION
       This command creates a new record in the repository.


FLAG VALUES
		<type>...
			Specifies the repository in which the record is to be created in. 
			Available repositories include <game> and <user> .

		<fields>...
			Specifies ALL the fields needed to be populated to create the record.
			Fields used depend on the repository selected.	
			Values are specified strictly in this format: 
			--fields username=john,age=44,email=k@mail.com
			Allowed fields for <user> repository are: [username, age, email]
			Allowed fields for <game> repository are: [title, description, url, publisher, age_rating]`



const updateUsageString = 
`NAME
       update - Update a record in the database

SYNOPSIS
        update [--type] [--fields] [--where]

DESCRIPTION
       This command updates a record in the repository.

	
FLAG VALUES
		<type>...
			Specifies the repository in which the record is to be updated in. 
			Available repositories include <game> and <user> .

		<fields>...
			Specifies ALL the fields needed to be populated to create the record.
			Fields used depend on the repository selected.
			Values are specified strictly in this format: --fields username=john,age=44,email=k@mail.com
			Allowed fields for <user> repository are: [username, age, email]
			Allowed fields for <game> repository are: [title, description, url, publisher, age_rating]
			
		<where>...
			Specifies the where clause value.
			Records to be updated are filtered only by the username (when using the user repo) or title 
			(when using the game repository)`


const filterUsageString = 
`NAME
       filter - Filter a record in the database

SYNOPSIS
        filter [--type] [--where] [--by]

DESCRIPTION
       This command filters a record in the repository by a field value.

	   
FLAG VALUES
		<type>...
			Specifies the repository in which the record is to be filtered in. 
			Available repositories include <game> and <user> .

		<by>...
			Specifies the field for which the record will be filtered by.
			Field used depend on the repository selected.
			Allowed fields for <user> repository are: [username, age, email]
			Allowed fields for <game> repository are: [title, description, url, publisher, age_rating]
			
		<where>...
			Specifies the specific value to be filtered by and must be a value of the <by> field specified.
			Records to be updated are filtered only by the username (when using the user repo) or title 
			(when using the game repository).`

func ParseCommands(args []string, opts *cli.Options) {

	if len(args) < 2{
		flag.Usage()
		os.Exit(1)
	}
	switch args[1] {
	case "create":
		createFlag := flag.NewFlagSet("create", flag.ExitOnError)
		opts.Flag = createFlag
		opts.Method = args[1]

		createFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			fmt.Println(createUsageString)
		}

		createFlag.StringVar(&opts.TypeOfObject, "type", "", "")
		createFlag.StringVar(&opts.Fields, "fields", "", "")
		createFlag.Parse(args[2:])



	case "update":
		updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
		opts.Flag = updateFlag
		opts.Method = args[1]

		updateFlag.StringVar(&opts.TypeOfObject, "type", "", "")
		updateFlag.StringVar(&opts.Fields, "fields", "", "")
		updateFlag.StringVar(&opts.Where, "where", "", "")

		updateFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			fmt.Println(updateUsageString)
		}

		updateFlag.Parse(args[2:])

	case "filter":
		filterFlag := flag.NewFlagSet("filter", flag.ExitOnError)
		opts.Flag = filterFlag
		opts.Method = args[1]

		filterFlag.StringVar(&opts.TypeOfObject, "type", "", "")
		filterFlag.StringVar(&opts.Where, "where", "", "")
		filterFlag.StringVar(&opts.By, "by", "", "")
		
		filterFlag.Usage = func() {
			fmt.Printf("Usage:\n")
			fmt.Println(filterUsageString)
		}

		filterFlag.Parse(args[2:])
	default:
		flag.Usage = func() {
			fmt.Printf("Usage:\n")
			fmt.Println(genericUsageString)		
		}
	}
}


