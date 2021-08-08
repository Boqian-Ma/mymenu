
test:
	- gofmt -s -w ./pkg ./service.core
	- go test -mod vendor -coverprofile=cov.out -p 1 ./... | { grep -v 'no test'; true; }

fmt:
	- gofmt -s -w ./pkg ./service.core

core-docs:
	- ./scripts/swag init -g ./service.core/cmd/core/main.go -o ./service.core/docs/

core-build: 
	- docker build -t service-core . -f service.core/Dockerfile

core-run:
	- go run ./service.core/cmd/core/main.go
cloc:
	- find . -type f \( -iname \*.go -o -iname \*.tf \) -not -path "*/vendor/*" -not -path "*/.terraform/*" -not -path "*/docs/*" | xargs wc -l

# Database setup and management for localhost & testing
POSTGRESQL_URL := 'postgres://postgres:password@localhost:5432/example?sslmode=disable'
db-create:
	- migrate -database ${POSTGRESQL_URL} -path ./migrations drop

db-start:
	- docker-compose up -d

db-stop:
	- docker-compose down

migrate-up:
	- migrate -database ${POSTGRESQL_URL} -path ./migrations up

migrate-down:
	- migrate -database ${POSTGRESQL_URL} -path ./migrations down 1

migrate-set5:
	- migrate -database ${POSTGRESQL_URL} -path ./migrations force 5

# Need to run this once on the DB to setup migrations before testing
migrate-reset:
	- migrate -database ${POSTGRESQL_URL} -path ./migrations drop -f

fresh-start:
	- docker-compose down
	- docker-compose up -d
	- migrate -database ${POSTGRESQL_URL} -path ./migrations drop -f
	- migrate -database ${POSTGRESQL_URL} -path ./migrations down 1
	- migrate -database ${POSTGRESQL_URL} -path ./migrations up
	- go run ./service.core/cmd/core/main.go

load-data:
	- bash load_data.sh