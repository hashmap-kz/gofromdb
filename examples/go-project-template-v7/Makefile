.PHONY: lint
lint:
	golangci-lint run --output.tab.path=stdout

.PHONY: docs
docs:
	swag init --parseDependency --parseInternal

.PHONY: run
run:
	go run ./main.go

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o go-project-template-v7 ./main.go

.PHONY: test
test:
	go test -v -race -cover ./...
