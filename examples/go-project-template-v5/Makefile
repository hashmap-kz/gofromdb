.PHONY: lint docs run build test migrate-up

lint:
	golangci-lint run ./...

docs:
	swag init --parseDependency --parseInternal

run:
	go run ./main.go

build:
	go build ./main.go

test:
	go test -cover ./...

.ONESHELL:
migrate-up:
	export GOOSE_DRIVER=postgres
	export GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable"
	goose up -table migrate_schema -dir=database/migrations/schema
	goose up -table=migrate_routine -dir=database/migrations/functions
	goose up -table=migrate_data -dir=database/migrations/data
