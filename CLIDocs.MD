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