# Running the backend
1. Ensure vendor is up to date (only needs to be run once)
```bash
go mod vendor
go mod tidy
```
2. Start docker
```bash
sudo service docker start
```
3. Start backend
```bash
make fresh-start
```

# Installation

### Required Dependencies
- [**Go** to build and run the application](https://golang.org/)
- [**Migrate** to perform database schema migrations](
https://github.com/golang-migrate/migrate)
- [**swag** to generate the API docs](https://github.com/swaggo/swag)
- [**docker**](https://www.docker.com/)
- Third party libraries

### Installing Dependencies
On Linux, most of these can be installed using the [yay](https://github.com/Jguer/yay) package manager
```
yay -S go migrate docker
```

The `swag` binary is included in the repo under `./scripts/swag`

Setup the database to be used when running the project locally
```
make db-start
```

You should also run `go mod` commands to ensure `./vendor` is up-to-date
```
go mod vendor
go mod tidy
```

Don't forget to `make db-stop` once you're done! :))

## Running The Project

You need to manually create the database by logging into Postgres via To login to Postgres: `psql -U postgres -h localhost`

Then run `create database example;` on Postgres

First setup the database (if you haven't already) `make db-start`

To kill the Postgres instance:  `pg_ctl -D /usr/local/var/postgres restart`

`make core-run` will run the core service

*note: you'll need to restart the server to pickup any file changes*

## DB Migrations

`make migrate-up` will migrate the database to the latest version of the schema, the API server will also do this automatically

If you run into any issues here, 
```
make migrate-reset
make migrate-up
```
should fix them, if not hmu :))

## Code gen stuff
`make fmt` will keep our code clean

`make core-docs` will update the swagger api docs from codegen
These docs can then be viewed at http://localhost:5000/core/docs while the server is running

## Project Structure
```
jamar
├── config
|   └── local.yml
├── pkg/ (common self-created packages for all services)
|   ├── errors/
|   └── log/
├── service.core/
|   ├── cmd/core (entrypoint)
|   |   ├── main.go
|   |   └── main_test.go
|   ├── internal/ (logic split by feature)
|   |   ├── users 
|   |   |   ├── api.go
|   |   |   ├── api_test.go
|   |   |   ├── service.go
|   |   |   ├── service_test.go
|   |   |   ├── repository.go
|   |   |   └── repository_test.go
├── vendor/ (common vendor-packages for all services)
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### Benefits of the above structure
1. Shared `pkg` between all services
2. Mono-repo and included benefits, eg version controlled changes between services
3. I really really like working with the internal/{feature} split
    - gives great encapsulation of relevant logic, and also allows exposing really clean & minimal interfaces

## Inspiration
- [Go Project Layout Standards](https://github.com/golang-standards/project-layout)
- [/internal Services Layout](https://github.com/qiangxue/go-rest-api)
