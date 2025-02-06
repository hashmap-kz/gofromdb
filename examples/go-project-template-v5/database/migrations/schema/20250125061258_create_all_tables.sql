-- +goose Up

create table clients
(
    record_id serial primary key,
    email     varchar(250) unique not null
);

comment on table clients is 'Stores users information, identified by a unique email.';
comment on column clients.record_id is 'Primary key for the users table.';
comment on column clients.email is 'Unique email address of the user.';

create table categories
(
    record_id    serial primary key,
    name         varchar(250) not null,
    parent_id    int          references categories (record_id) on delete set null,
    valid_period daterange    not null default '[1970-01-01,)'::daterange,
    is_current   bool         not null generated always as (valid_period @> '9999-12-30'::date) stored,
    exclude using gist(name with =, valid_period with &&)
);

comment on table categories is 'Represents product categories, supporting hierarchical relationships.';
comment on column categories.record_id is 'Primary key for the category table.';
comment on column categories.name is 'Name of the category.';
comment on column categories.parent_id is 'Reference to the parent category. NULL if it is a root category.';
comment on column categories.valid_period is 'Validity period of this category naming.';
comment on column categories.is_current is 'Whether this category is the last actual.';

create table products
(
    record_id   serial primary key,
    category_id int          not null references categories (record_id),
    name        varchar(250) not null,
    description text
);
create unique index ix_product_unq on products (category_id, name);

comment on table products is 'Stores products with a reference to their category.';
comment on column products.record_id is 'Primary key for the product table.';
comment on column products.category_id is 'Foreign key referencing the category to which the product belongs.';
comment on column products.name is 'Name of the product.';
comment on column products.description is 'Detailed description of the product.';

create table customer_orders
(
    record_id   serial primary key,
    client_id   int not null references clients (record_id),
    description text
);

comment on table customer_orders is 'Represents purchases made by clients.';
comment on column customer_orders.record_id is 'Primary key for the customer_orders table.';
comment on column customer_orders.client_id is 'Foreign key referencing the client who made the order.';
comment on column customer_orders.description is 'Optional description or additional details of the order.';

create table customer_order_items
(
    record_id         serial primary key,
    customer_order_id int            not null references customer_orders (record_id),
    product_id        int            not null references products (record_id),
    quantity          numeric(15, 3) not null,
    price             numeric(15, 2) not null
);

comment on table customer_order_items is 'Represents items in a sales, including quantity and price.';
comment on column customer_order_items.record_id is 'Primary key for the sales_items table.';
comment on column customer_order_items.customer_order_id is 'Foreign key referencing the associated customer-order.';
comment on column customer_order_items.product_id is 'Foreign key referencing the product.';
comment on column customer_order_items.quantity is 'Number of units of the product.';
comment on column customer_order_items.price is 'Price per unit of the product at the time of ordering.';
