FROM golang:1.26.0-alpine3.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/root/go-build go mod download -x \
  && go install github.com/swaggo/swag/cmd/swag@latest \
  && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN swag init --parseDependency --parseInternal \
  && golangci-lint run ./... \
  && go test -v ./... \
  && go build -ldflags="-s -w" -o ./service

# Deploy the application binary into a lean image
# FROM golang:1.26.0-alpine3.23 AS build-release-stage
FROM gcr.io/distroless/base-debian12 AS build-release-stage

ARG APP_ENVIRONMENT
ENV APP_ENVIRONMENT=${APP_ENVIRONMENT}

WORKDIR /app

COPY --from=build-stage /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app/service"]
