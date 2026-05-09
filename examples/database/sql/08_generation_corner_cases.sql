-- +goose Up

-- Extra bookstore schemas used to test generator corner cases.
-- These tables are intentionally meaningful, but they cover:
--   * non-public schemas
--   * natural primary keys
--   * composite primary keys, including natural composite keys
--   * uuid and bigserial primary keys
--   * tables without primary keys
--   * tables with defaults / generated columns / nullable fields
--   * partitioned tables and partitions
--   * duplicate table names across schemas, e.g. public.orders and bookstore_sales.orders

create schema if not exists bookstore_catalog;
create schema if not exists bookstore_sales;
create schema if not exists bookstore_inventory;
create schema if not exists bookstore_import;
create schema if not exists bookstore_events;

-- --------------------------------------------------------------------
-- Catalog: natural keys, uuid keys, generated columns, arrays, jsonb.
-- --------------------------------------------------------------------

create table bookstore_catalog.publishers
(
    code         text primary key,
    name         text not null,
    country_code char(2) not null,
    website      text,
    founded_on   date,
    active       boolean not null default true
);

comment on table bookstore_catalog.publishers is 'Book publishers. Uses a natural text primary key.';
comment on column bookstore_catalog.publishers.code is 'Natural publisher code, for example no_starch or manning.';

create table bookstore_catalog.authors
(
    author_id   uuid primary key default gen_random_uuid(),
    display_name text not null,
    legal_name   text,
    biography    text,
    metadata     jsonb not null default '{}'::jsonb,
    active       boolean not null default true,
    born_on      date
);

comment on table bookstore_catalog.authors is 'Authors. Uses a UUID primary key with a database default.';

create table bookstore_catalog.books
(
    book_id       bigserial primary key,
    publisher_code text not null references bookstore_catalog.publishers (code),
    isbn13        varchar(13) not null unique,
    title         text not null,
    subtitle      text,
    description   text,
    price         numeric(12, 2) not null,
    weight_grams  integer,
    rating        numeric(3, 2),
    published_on  date,
    tags          text[] not null default '{}'::text[],
    attrs         jsonb not null default '{}'::jsonb,
    cover_image   bytea,
    archived_at   timestamptz,
    title_search  text generated always as (lower(title)) stored
);

comment on table bookstore_catalog.books is 'Books available in the store. Covers serial PK, arrays, jsonb, bytea, nullable fields, and generated columns.';

create table bookstore_catalog.book_authors
(
    book_id            bigint not null references bookstore_catalog.books (book_id) on delete cascade,
    author_id          uuid   not null references bookstore_catalog.authors (author_id) on delete cascade,
    contribution_order smallint not null default 1,
    role               text not null default 'author',
    notes              text,
    primary key (book_id, author_id),
    unique (book_id, contribution_order)
);

comment on table bookstore_catalog.book_authors is 'Many-to-many link. Tests composite primary key order: book_id, author_id.';

create table bookstore_catalog.book_translations
(
    book_id       bigint not null references bookstore_catalog.books (book_id) on delete cascade,
    language_code text   not null,
    translated_title text not null,
    translated_by   text,
    released_on     date,
    primary key (book_id, language_code)
);

comment on table bookstore_catalog.book_translations is 'Natural composite key: one translation per book and language.';

-- --------------------------------------------------------------------
-- Sales: duplicate table name with public.orders, composite order lines.
-- --------------------------------------------------------------------

create table bookstore_sales.customers
(
    customer_id uuid primary key default gen_random_uuid(),
    email       varchar(250) not null unique,
    full_name   text not null,
    phone       text,
    marketing_opt_in boolean not null default false,
    registered_at timestamptz not null default current_timestamp
);

comment on table bookstore_sales.customers is 'Store customers. UUID primary key and simple unique field.';

create table bookstore_sales.orders
(
    order_id    bigserial primary key,
    customer_id uuid not null references bookstore_sales.customers (customer_id),
    status      text not null default 'draft',
    placed_at   timestamptz,
    paid_at     timestamptz,
    cancelled_at timestamptz,
    comment     text,
    check (status in ('draft', 'placed', 'paid', 'cancelled'))
);

comment on table bookstore_sales.orders is 'Sales orders. Duplicate table name with public.orders to test schema-aware generation.';

create table bookstore_sales.order_lines
(
    order_id   bigint not null references bookstore_sales.orders (order_id) on delete cascade,
    line_no    integer not null,
    book_id    bigint not null references bookstore_catalog.books (book_id),
    quantity   smallint not null,
    unit_price numeric(12, 2) not null,
    discount_amount numeric(12, 2) not null default 0,
    note       text,
    primary key (order_id, line_no)
);

comment on table bookstore_sales.order_lines is 'Order lines. Tests composite primary key order: order_id, line_no.';

create table bookstore_sales.discount_codes
(
    code          text primary key,
    description   text,
    percent_off   numeric(5, 2) not null,
    valid_period  daterange not null default '[1970-01-01,)'::daterange,
    max_uses      integer,
    active        boolean not null default true
);

comment on table bookstore_sales.discount_codes is 'Discount codes. Natural text primary key plus daterange.';

-- --------------------------------------------------------------------
-- Inventory: natural composite keys and a PK-less event table.
-- --------------------------------------------------------------------

create table bookstore_inventory.warehouses
(
    code        text primary key,
    name        text not null,
    address     jsonb not null default '{}'::jsonb,
    timezone    text not null default 'UTC',
    active      boolean not null default true
);

comment on table bookstore_inventory.warehouses is 'Warehouses. Natural text primary key.';

create table bookstore_inventory.stock_levels
(
    warehouse_code    text   not null references bookstore_inventory.warehouses (code),
    book_id           bigint not null references bookstore_catalog.books (book_id),
    available_qty     integer not null default 0,
    reserved_qty      integer not null default 0,
    reorder_threshold integer not null default 0,
    last_counted_at   timestamp,
    primary key (warehouse_code, book_id)
);

comment on table bookstore_inventory.stock_levels is 'Current stock per warehouse and book. Natural + surrogate composite foreign key primary key.';

create table bookstore_inventory.stock_events
(
    happened_at    timestamptz not null default current_timestamp,
    warehouse_code text not null references bookstore_inventory.warehouses (code),
    book_id        bigint not null references bookstore_catalog.books (book_id),
    delta_qty      integer not null,
    reason         text not null,
    payload        jsonb not null default '{}'::jsonb
);

comment on table bookstore_inventory.stock_events is 'Append-only stock event stream. Intentionally has no primary key.';

-- --------------------------------------------------------------------
-- Import: composite natural key and PK-less error rows.
-- --------------------------------------------------------------------

create table bookstore_import.import_batches
(
    source_name text not null,
    batch_no    integer not null,
    started_at  timestamptz not null default current_timestamp,
    finished_at timestamptz,
    file_name   text not null,
    row_count   integer not null default 0,
    metadata    jsonb not null default '{}'::jsonb,
    primary key (source_name, batch_no)
);

comment on table bookstore_import.import_batches is 'Import batch metadata. Composite natural primary key.';

create table bookstore_import.import_errors
(
    source_name text not null,
    batch_no    integer not null,
    row_no      integer not null,
    column_name text,
    message     text not null,
    raw_payload jsonb not null default '{}'::jsonb,
    created_at  timestamptz not null default current_timestamp,
    foreign key (source_name, batch_no)
        references bookstore_import.import_batches (source_name, batch_no)
        on delete cascade
);

comment on table bookstore_import.import_errors is 'Import validation errors. Intentionally has no primary key.';

-- --------------------------------------------------------------------
-- Events: partitioned parent and partitions.
-- A generator that excludes partitioned tables and child partitions should
-- not generate CRUD for any of these tables.
-- --------------------------------------------------------------------

create table bookstore_events.book_events
(
    happened_at timestamptz not null,
    book_id     bigint not null,
    event_type  text not null,
    payload     jsonb not null default '{}'::jsonb
) partition by range (happened_at);

comment on table bookstore_events.book_events is 'Partitioned event stream parent. Should normally be excluded from CRUD generation.';

create table bookstore_events.book_events_2026_01
    partition of bookstore_events.book_events
    for values from ('2026-01-01') to ('2026-02-01');

create table bookstore_events.book_events_2026_02
    partition of bookstore_events.book_events
    for values from ('2026-02-01') to ('2026-03-01');
