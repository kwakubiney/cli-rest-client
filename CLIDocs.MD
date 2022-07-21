This client uses subcommands and flags to make API requests.

Assuming we built binary and named it `gamedev`

* All subcommands can be seen by passing the `-- help` flag
    - Example => `./gamedev --help`

* Specific commands for subcommands can be found by passing `--help` flag to subcommands.
    - Example => `./gamedev create --help`


**Example**

* Sample flow for creating a user

After building the go project and spinning DockerCompose:

Assuming we built binary and named it `gamedev`

1) `./gamedev create -type user --fields username=john,age=44,email=k@mail.com`


```
These are common subcommands used in various situations:
	create     Create a new record in the database
	filter     Filter a record in the database
	update     Update a record in the database`


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


NAME
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
			(when using the game repository)

NAME
       filter - Filter a record in the database

SYNOPSIS
        update [--type] [--where] [--by]

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

```