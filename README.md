<picture>
  <img alt="gofromdb - Generate ready-to-run Go REST microservices from PostgreSQL" src="docs/assets/gofromdb-logo.svg">
</picture>

# gofromdb

Point it at a PostgreSQL database. Get a running Go REST API.

[![License](https://img.shields.io/github/license/hashmap-kz/gofromdb)](https://github.com/hashmap-kz/gofromdb/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hashmap-kz/gofromdb)](https://goreportcard.com/report/github.com/hashmap-kz/gofromdb)
[![Go Reference](https://pkg.go.dev/badge/github.com/hashmap-kz/gofromdb.svg)](https://pkg.go.dev/github.com/hashmap-kz/gofromdb)
[![Workflow Status](https://img.shields.io/github/actions/workflow/status/hashmap-kz/gofromdb/ci.yml?branch=master)](https://github.com/hashmap-kz/gofromdb/actions/workflows/ci.yml?query=branch:master)
[![GitHub Issues](https://img.shields.io/github/issues/hashmap-kz/gofromdb)](https://github.com/hashmap-kz/gofromdb/issues)
[![Go Version](https://img.shields.io/github/go-mod/go-version/hashmap-kz/gofromdb)](https://github.com/hashmap-kz/gofromdb/blob/master/go.mod#L3)
[![Latest Release](https://img.shields.io/github/v/release/hashmap-kz/gofromdb)](https://github.com/hashmap-kz/gofromdb/releases/latest)
[![Start contributing](https://img.shields.io/github/issues/hashmap-kz/gofromdb/good%20first%20issue?color=7057ff&label=Contribute)](https://github.com/hashmap-kz/gofromdb/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22)

`gofromdb` connects to your PostgreSQL database, reads the schema, and writes a **complete Go project** -
entity structs, repositories, services, HTTP handlers, DTOs, and Swagger annotations.

**The result compiles and runs immediately.**

---

## Purpose

This tool is for **database-first** Go projects where the schema is the source of truth.

It helps when:

- You're tired of writing the same CRUD boilerplate for every new table.
- You want plain SQL with predictable performance over hidden magic.
- You need a backend MVP fast, or a consistent starting point to customize from.

---

## Install

**Homebrew**

```bash
brew tap hashmap-kz/homebrew-tap
brew install gofromdb
```

**Go**

```bash
go install github.com/hashmap-kz/gofromdb@latest
```

---

## Usage

```bash
gofromdb \
  -conn="postgres://user:pass@localhost:5432/mydb" \
  -output=myapp \
  -workers=8
```

A `-config` flag accepts a JSON file for filtering schemas and tables:

```bash
gofromdb -conn="..." -config=gofromdb.json -output=myapp
```

---

## Generated project structure

For each table, the generator produces a self-contained package under `internal/api/<schema>/<table>/`:

```
internal/api/
  repository.go          # top-level repository interfaces (all tables)
  service.go             # top-level service interfaces (all tables)
  handler.go             # router: mounts all routes

  public/
    products/
      entity.go          # DB struct mapped from table schema
      repository.go      # pgx queries: Save, Update, Delete, Find, FindAll, paginated
      service.go         # business logic layer, delegates to repository
      dto.go             # CreateDto, UpdateDto, Dto (internal transfer types)
      payload.go         # HTTP request/response types with JSON tags
      handler.go         # net/http handlers with Swagger annotations
```

---

## Example: from schema to code

Given this PostgreSQL table:

```sql
create table products (
    record_id   serial primary key,
    category_id int          not null references categories (record_id),
    name        varchar(250) not null,
    description text
);

comment on table products is 'Stores products with a reference to their category.';
```

**`entity.go`** - struct with proper Go types and JSON tags:

```go
type Products struct {
    RecordID    int       `json:"record_id"    db:"record_id"`
    CategoryID  int       `json:"category_id"  db:"category_id"`
    Name        string    `json:"name"         db:"name"`
    Description *string   `json:"description"  db:"description"`
}
```

**`repository.go`** - typed SQL queries via pgx/v5, including paginated list:

```go
func (r *repo) Save(ctx context.Context, inputEntity *Products) (*Products, error) {
    query := `
        insert into public.products (category_id, name, description)
        values ($1, $2, $3)
        returning record_id, category_id, name, description
    `
    row := r.db.Pool.QueryRow(ctx, query,
        inputEntity.CategoryID,
        inputEntity.Name,
        inputEntity.Description,
    )
    return scanFullRow(row)
}
```

**`handler.go`** - stdlib `net/http` handlers with Swagger annotations:

```go
// @Summary Create new item
// @Tags products
// @Accept json
// @Produce json
// @Param request body productsCreateRequest true "Create input"
// @Success 201 {object} productsResponse
// @Router /api/v1/products [post]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
    req := &productsCreateRequest{}
    if err := httputils.ReadJSON(r, &req); err != nil {
        httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
        return
    }
    if err := validator.ValidateStruct(req); err != nil {
        httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
        return
    }
    resp, err := h.svc.Save(r.Context(), mapCreateRequestToCreateInputDto(req))
    if err != nil {
        httputils.WriteJSON(w, http.StatusInternalServerError, httputils.ErrorResponse{Message: err.Error()})
        return
    }
    httputils.WriteJSON(w, http.StatusCreated, mapDtoToPayload(resp))
}
```

**Routes mounted automatically** for every table:

```
POST   /api/v1/products
PUT    /api/v1/products/{record_id}
DELETE /api/v1/products/{record_id}
GET    /api/v1/products/{record_id}
GET    /api/v1/products
GET    /api/v1/products/pageable
```

**Wiring** - the entire DI chain, no framework:

```go
router := http.NewServeMux()

apiHandler := api.NewHandler(api.NewServices(ctx, api.Deps{
    Repos: api.NewRepositories(ctx, pg),
}))
apiHandler.Mount(router)

...
httpserver.NewServer(router)
```

---

## Stack

| Concern      | Library / approach                |
|--------------|-----------------------------------|
| HTTP server  | stdlib `net/http`                 |
| PostgreSQL   | `pgx/v5` with connection pool     |
| Validation   | `go-playground/validator`         |
| API docs     | Swagger via `swaggo/swag`         |
| Pagination   | built-in `pageable` package       |
| Config       | YAML per environment              |
| Middleware   | CORS, structured logging (`slog`) |
| Containerize | Dockerfile included               |

---

## Example database

A ready-to-use bookstore database is in `examples/database/`, spanning multiple PostgreSQL schemas:

| Schema                | Tables                                                              |
|-----------------------|---------------------------------------------------------------------|
| `public`              | `users`, `categories`, `products`, `orders`, `order_items`          |
| `bookstore_catalog`   | `books`, `authors`, `book_authors`, `book_translations`, `publishers` |
| `bookstore_sales`     | `orders`, `order_lines`, `customers`, `discount_codes`              |
| `bookstore_inventory` | `warehouses`, `stock_levels`, `stock_events`                        |
| `bookstore_import`    | `import_batches`, `import_errors`                                   |

Start it with Docker Compose:

```bash
cd examples/database
docker compose up -d
```

The generated project under `examples/go-project-template-v7/` was produced from this schema.

---

## License

MIT. See [LICENSE](./LICENSE) for details.
