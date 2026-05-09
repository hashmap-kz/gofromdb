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

`gofromdb` connects to your database, reads the schemas, and writes a **complete Go project** - 
with entity structs, repositories, services, HTTP handlers, DTOs, and Swagger annotations. 

**The result compiles and runs immediately.**

---

## Purpose

This tool is for **database-first** Go projects where the database schema is the source of truth.

It helps when:

- You are tired of writing the same REST API boilerplate again and again.
- Every new table means another handler, service, repository, DTOs, and CRUD methods.
- You prefer _plain old SQL_ with predictable performance over hidden magic.
- You want generated code with a consistent structure, even if you later customize it by hand.
- You're designing a database and need a backend MVP quickly.
- You're migrating from another language to Go and want to generate a starter API from your existing database.

An imperfect standard is better than no standard.

---

## Install

#### Package

```bash
go install github.com/hashmap-kz/gofromdb@latest
```

#### Brew

```bash
brew tap hashmap-kz/homebrew-tap
brew install gofromdb
```

## Usage

```bash
gofromdb -conn="postgres://postgres:postgres@localhost:5432/bookstore" -workers=8 -output=myapp
```

---

## Generated project structure

For each database table, the generator produces a self-contained package under `internal/api/<schema>/<table>/`:

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

Given this table in PostgreSQL:

```sql
create table products (
    record_id   serial primary key,
    category_id int          not null references categories (record_id),
    name        varchar(250) not null,
    description text
);

comment on table products is 'Stores products with a reference to their category.';
comment on column products.name is 'Name of the product.';
```

The generator produces:

**entity.go** - struct with proper Go types, JSON and db tags, comments from the schema:

```go
type Products struct {
    RecordID    int      `json:"record_id"    db:"record_id"`
    CategoryID  int      `json:"category_id"  db:"category_id"`
    Name        string   `json:"name"         db:"name"`
    Description *string  `json:"description"  db:"description"`
    CreatedAt   time.Time `json:"created_at"  db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"  db:"updated_at"`
    GUID        string   `json:"guid"         db:"guid"`
}
```

**repository.go** - typed SQL queries using pgx/v5, including paginated list:

```go
func (r *repo) Save(ctx context.Context, inputEntity *Products) (*Products, error) {
    query := `
        insert into public.products (category_id, name, description)
        values ($1, $2, $3)
        returning record_id, category_id, name, description, created_at, updated_at, guid
    `
    row := r.db.Pool.QueryRow(ctx, query,
        inputEntity.CategoryID,
        inputEntity.Name,
        inputEntity.Description,
    )
    return scanFullRow(row)
}
```

**handler.go** - stdlib `net/http` handlers with Swagger annotations:

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
    if err := httputils.ReadJSON(r, req); err != nil {
        httputils.WriteJSON(w, http.StatusBadRequest, httputils.ErrorResponse{Message: err.Error()})
        return
    }
    // validate -> map to DTO -> call service -> map to response
    resp, err := h.svc.Save(r.Context(), mapCreateRequestToCreateInputDto(req))
    ...
    httputils.WriteJSON(w, http.StatusCreated, dtoToPayload)
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

---

## Generated project stack

The scaffold the generator builds on top of:

| Concern         | Library / approach                        |
|-----------------|-------------------------------------------|
| HTTP server     | stdlib `net/http`                         |
| PostgreSQL      | `pgx/v5` with connection pool             |
| Validation      | `go-playground/validator`                 |
| API docs        | Swagger via `swaggo/swag`                 |
| Pagination      | built-in `pageable` package               |
| Config          | YAML per environment                      |
| Middleware      | CORS, structured logging (`slog`)         |
| Containerize    | Dockerfile included                       |

---

## Example database

A ready-to-use example database is in `examples/database/`. It spans multiple PostgreSQL schemas:

| Schema               | Tables                                                   |
|----------------------|----------------------------------------------------------|
| `public`             | `users`, `categories`, `products`, `orders`, `order_items` |
| `bookstore_catalog`  | `books`, `authors`, `book_authors`, `book_translations`, `publishers` |
| `bookstore_sales`    | `orders`, `order_lines`, `customers`, `discount_codes`   |
| `bookstore_inventory`| `warehouses`, `stock_levels`, `stock_events`             |
| `bookstore_import`   | `import_batches`, `import_errors`                        |

Start it with Docker Compose:

```bash
cd examples/database
docker compose up -d
```

The generated project under `examples/go-project-template-v7/` was produced from this schema.

---

## License

MIT License. See [LICENSE](./LICENSE) for details.
