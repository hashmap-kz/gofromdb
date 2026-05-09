# go-gen-project-structure

Point it at a PostgreSQL database. Get a running Go REST API.

The generator connects to your database, reads the schema, and writes a complete Go project - one package per table, with entity structs, repositories, services, HTTP handlers, DTOs, and Swagger annotations. The result compiles and runs immediately.

---

## How it works

1. Connect to your PostgreSQL database
2. Introspect every table: columns, types, primary keys, foreign keys, comments
3. Write a full Go project to the output directory, based on a scaffold template
4. Format all generated code with `goimports` and `gofumpt`

```
database schema  -->  go-gen-project-structure  -->  running REST API
```

---

## Quick start

```bash
# build
make build

# generate a project from your database
./bin/go-gen-project-structure \
  --conn "postgres://postgres:postgres@localhost:5432/mydb" \
  --output ./my-api

# run the generated project
cd my-api
go run main.go
```

---

## Generated project structure

For each database table, the generator produces a self-contained package under `internal/api/<table_name>/`:

```
internal/api/
  repository.go          # top-level repository interfaces (all tables)
  service.go             # top-level service interfaces (all tables)
  handler.go             # router: mounts all routes

  products/
    entity.go            # DB struct mapped from table schema
    repository.go        # pgx queries: Save, Update, Delete, Find, FindAll, paginated
    service.go           # business logic layer, delegates to repository
    dto.go               # CreateDto, UpdateDto, Dto (internal transfer types)
    payload.go           # HTTP request/response types with JSON tags
    handler.go           # net/http handlers with Swagger annotations
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

## CLI flags

| Flag       | Default                                              | Description                  |
|------------|------------------------------------------------------|------------------------------|
| `--conn`   | `postgres://postgres:postgres@localhost:5432/bookstore` | PostgreSQL connection string |
| `--output` | _(in-place, updates the bundled example)_            | Output directory             |

If `--output` is omitted, the generator overwrites the bundled example under `examples/go-project-template-v7/`.

---

## Build and install

```bash
make build      # builds to bin/go-gen-project-structure
make install    # installs to /usr/local/bin
make test       # runs tests
```

**Prerequisites:** `goimports`, `gofumpt` must be available on PATH (used to format generated files).

---

## Example database

A ready-to-use example schema with `users`, `categories`, `products`, `orders`, and `order_items` is in `examples/database/`. Start it with Docker Compose:

```bash
cd examples/database
docker compose up -d
```

The generated project under `examples/go-project-template-v7/` was produced from this schema.
