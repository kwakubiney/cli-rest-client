# Canonical Take Home REST Client

This client makes REST API requests using the CLI.

# Development Usage (Installing)

- Run `docker-compose up` to spin up Postgres service for the API service

- Create an `.env` file in root directory by cloning the `.env.test` file. 

- You can modify .env but it is compatible with postgres service being spun. If modified, `docker-compose.yaml` must be modified to correspond.

- `cd` into `internal/migrations` and run `goose postgres "user=postgres password=postgres dbname=gamedev sslmode=disable" up` to make migrations.

- Build the go project with `go build -o gamedev cmd/main.go` in root and run `--help` to see subcommands, you can check out [CLI documentation](CLIDocs.MD)

# Tests

- Tests can be run in root with `go test ./...` 

# API Documentation

- You can jump right into the [API Documentation](APIDocs.MD).

# Limitations

- Filters are done based on one field.

- Updates are done based only on username and game title.

- No authentication provided.
