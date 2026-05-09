.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: docs
docs:
	swag init --parseDependency --parseInternal

.PHONY: run
run:
	go run ./main.go

.PHONY: build
build:
	go build ./main.go

.PHONY: test
test:
	go test -cover -v -race ./...
