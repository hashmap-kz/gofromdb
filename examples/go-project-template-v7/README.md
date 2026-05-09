### Notes (pgx):

- https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
- https://github.com/terfo1/linkedout/blob/main/cmd/database.go
- https://pieces.app/blog/how-to-build-and-document-a-go-rest-api-with-gin-and-go-swagger
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

### Notes (logging):

- https://betterstack.com/community/guides/logging/logging-in-go/
- https://github.com/timtoronto634/pgx-slog/blob/main/adapter.go

### Notes (api design):

- https://medium.com/@shershnev/layered-architecture-implementation-in-golang-6318a72c1e10
- https://github.com/evrone/go-clean-template
- https://github.com/Creatly/creatly-backend
- https://github.com/reshimahendra/lbw-go/tree/master

### Projects:
- https://github.com/moby/moby (a lot of work around HTTP utilities, error-handling, etc...)
- https://github.com/go-gitea/gitea (project structure, etc...)

### Notes (network):

- https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b
- https://github.com/ante-neh/Harmony-Hotel-Reservation/blob/main/util/util.go
- https://github.com/pkarakal/aws-skg-meetup-otel-demo/blob/master/src/cart/router/router.go
- https://github.com/otaleghani/swms/blob/main/internal/server/init.go

### Notes (pagination):

- https://github.com/knadh/paginator/tree/master

### Notes (auth):

- https://developer.auth0.com/resources/guides/api/standard-library/basic-authorization
- https://github.com/mihaiflorentin88/go-keycloak-guard/blob/master/infrastructure/client/keycloak/client.go
- https://github.com/lapitskyss/chat-service/blob/main/internal/clients/keycloak/api_introspect.go

### Prepare:

```
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH
```

### Docs

- http://localhost:5000/swagger-ui/index.html

### POST http://localhost:5000/api/v1/files

```
{
  "type": "mp3"
}
```

### GET

- http://localhost:5000/api/v1/files
- http://localhost:5000/api/v1/files/1

### Migration:

#### golang-migrate

```
# install tool:
# https://github.com/golang-migrate/migrate/releases

# create migration file
migrate create -ext sql -dir database/migration/ -seq files_table

# apply migrations
migrate -path database/migration/ -database "postgresql://postgres:postgres@localhost:5432/bookstore?sslmode=disable" -verbose up
```

#### goose:

```
# install tool: 
# https://github.com/pressly/goose/releases

# create migration file
goose create -dir=database/migrations/schema     add_files_table sql
goose create -dir=database/migrations/data       seed_files_table sql
goose create -dir=database/migrations/repeatable utils_funcs sql

# apply migrations
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable"
goose up -table migrate_schema -dir=database/migrations/schema
goose up -table=migrate_data -dir=database/migrations/data
```

### TODO:

- slog logging
- pgx pool (remove sqlx)
- context (svc, repo, handlers, main, etc...)
- sql queries logging
- tracing
- initDb, initCfg, initLogger -> move to config PKG
- pagination queries (swms, template-v2)
- redis
- responses -> (success, error: swms)
- kafka (with microservices)
- auth (keycloak)
- goose migrations
- resty for HTTP-requests (https://github.com/lapitskyss/chat-service/blob/main/internal/clients/keycloak/api_introspect.go)
- config: github.com/caarlos0/env
- httputil from https://github.com/moby







