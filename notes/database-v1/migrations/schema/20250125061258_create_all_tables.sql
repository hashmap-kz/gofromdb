-- +goose Up

create table departments
(
    record_id       serial primary key,
    department_name varchar(250) not null unique
);

create table job_titles
(
    record_id  serial primary key,
    title_name varchar(250) not null unique
);

create table users
(
    record_id serial primary key,
    email     varchar(250) unique not null
);

comment on table users is 'Stores users information, identified by a unique email.';
comment on column users.record_id is 'Primary key for the users table.';
comment on column users.email is 'Unique email address of the user.';

create table customers
(
    record_id serial primary key,
    email     varchar(250) unique not null
);

comment on table customers is 'Stores customers information, identified by a unique email.';
comment on column customers.record_id is 'Primary key for the customers table.';
comment on column customers.email is 'Unique email address of the customer.';

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

create table purchases
(
    record_id   serial primary key,
    customer_id int not null references customers (record_id),
    description text
);

comment on table purchases is 'Represents purchases made by clients.';
comment on column purchases.record_id is 'Primary key for the buy table.';
comment on column purchases.customer_id is 'Foreign key referencing the customer who made the purchase.';
comment on column purchases.description is 'Optional description or additional details of the purchase.';

create table purchase_items
(
    record_id   serial primary key,
    purchase_id int            not null references purchases (record_id),
    product_id  int            not null references products (record_id),
    quantity    int            not null,
    price       numeric(15, 2) not null
);

comment on table purchase_items is 'Represents items in a purchase, including quantity and price.';
comment on column purchase_items.record_id is 'Primary key for the buy_item table.';
comment on column purchase_items.purchase_id is 'Foreign key referencing the associated purchase.';
comment on column purchase_items.product_id is 'Foreign key referencing the purchased product.';
comment on column purchase_items.quantity is 'Number of units of the product in the purchase.';
comment on column purchase_items.price is 'Price per unit of the product at the time of purchase.';

create table workflow_steps
(
    record_id serial primary key,
    step_name varchar(32) not null
);
create unique index ix_purchase_step_step_name_unq on workflow_steps (step_name);

comment on table workflow_steps is 'Purchase steps enum: ordered, delivered, etc...';
comment on column workflow_steps.step_name is 'Step name, unique';

create table purchase_statuses
(
    record_id        serial primary key,
    valid_period     daterange not null,
    purchase_id      int       not null references purchases,
    workflow_step_id int       not null references workflow_steps,
    exclude using gist(purchase_id with =, workflow_step_id with =, valid_period with &&)
);

comment on table purchase_statuses is 'Buy steps tracking';
comment on column purchase_statuses.record_id is 'PK';
comment on column purchase_statuses.valid_period is 'Period, open means that the step is in progress';
comment on column purchase_statuses.purchase_id is 'Buy-order ID';
comment on column purchase_statuses.workflow_step_id is 'Step ID';

create table product_pricing
(
    record_id     serial primary key,
    valid_period  daterange      not null default '[1970-01-01,)'::daterange,
    product_id    int            not null references products (record_id),
    product_price numeric(15, 2) not null,
    exclude using gist (product_id with =, valid_period with &&)
);

comment on table product_pricing is 'Prices of goods and services';
comment on column product_pricing.record_id is 'PK';
comment on column product_pricing.valid_period is 'Effective date range for the price';
comment on column product_pricing.product_id is 'References to products';
comment on column product_pricing.product_price is 'Actual price';

create table currencies
(
    record_id      serial primary key,
    currency_code  varchar(16)  not null,
    currency_value varchar(250) not null
);

