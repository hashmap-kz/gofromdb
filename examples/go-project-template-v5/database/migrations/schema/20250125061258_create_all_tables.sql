-- +goose Up

create table client
(
    record_id serial primary key,
    email     varchar(250) unique not null
);

comment on table client is 'Stores client information, identified by a unique email.';
comment on column client.record_id is 'Primary key for the client table.';
comment on column client.email is 'Unique email address of the client.';

create table category
(
    record_id    serial primary key,
    name         varchar(250) not null,
    parent_id    int          references category (record_id) on delete set null,
    valid_period daterange    not null default '[1970-01-01,)'::daterange,
    is_current   bool         not null generated always as (valid_period @> '9999-12-30'::date) stored,
    exclude using gist(name with =, valid_period with &&)
);

comment on table category is 'Represents product categories, supporting hierarchical relationships.';
comment on column category.record_id is 'Primary key for the category table.';
comment on column category.name is 'Name of the category.';
comment on column category.parent_id is 'Reference to the parent category. NULL if it is a root category.';
comment on column category.valid_period is 'Validity period of this category naming.';
comment on column category.is_current is 'Whether this category is the last actual.';

create table product
(
    record_id   serial primary key,
    category_id int          not null references category (record_id),
    name        varchar(250) not null,
    description text
);
create unique index ix_product_unq on product (category_id, name);

comment on table product is 'Stores products with a reference to their category.';
comment on column product.record_id is 'Primary key for the product table.';
comment on column product.category_id is 'Foreign key referencing the category to which the product belongs.';
comment on column product.name is 'Name of the product.';
comment on column product.description is 'Detailed description of the product.';

create table buy
(
    record_id   serial primary key,
    client_id   int not null references client (record_id),
    description text
);

comment on table buy is 'Represents purchases made by clients.';
comment on column buy.record_id is 'Primary key for the buy table.';
comment on column buy.client_id is 'Foreign key referencing the client who made the purchase.';
comment on column buy.description is 'Optional description or additional details of the purchase.';

create table buy_item
(
    record_id  serial primary key,
    buy_id     int            not null references buy (record_id),
    product_id int            not null references product (record_id),
    quantity   int            not null,
    price      numeric(15, 2) not null
);

comment on table buy_item is 'Represents items in a purchase, including quantity and price.';
comment on column buy_item.record_id is 'Primary key for the buy_item table.';
comment on column buy_item.buy_id is 'Foreign key referencing the associated purchase.';
comment on column buy_item.product_id is 'Foreign key referencing the purchased product.';
comment on column buy_item.quantity is 'Number of units of the product in the purchase.';
comment on column buy_item.price is 'Price per unit of the product at the time of purchase.';

create table purchase_step
(
    record_id serial primary key,
    step_name varchar(32) not null
);
create unique index ix_purchase_step_step_name_unq on purchase_step (step_name);

comment on table purchase_step is 'Purchase steps enum: ordered, delivered, etc...';
comment on column purchase_step.step_name is 'Step name, unique';

create table buy_step
(
    record_id        serial primary key,
    buy_step_period  daterange not null,
    buy_id           int       not null references buy,
    purchase_step_id int       not null references purchase_step,
    exclude using gist(buy_id with =, purchase_step_id with =, buy_step_period with &&)
);

comment on table buy_step is 'Buy steps tracking';
comment on column buy_step.record_id is 'PK';
comment on column buy_step.buy_step_period is 'Period, open means that the step is in progress';
comment on column buy_step.buy_id is 'Buy-order ID';
comment on column buy_step.purchase_step_id is 'Step ID';

create table product_price
(
    record_id            serial primary key,
    product_price_period daterange      not null default '[1970-01-01,)'::daterange,
    product_id           int            not null references product (record_id),
    product_price        numeric(15, 2) not null,
    exclude using gist (product_id with =, product_price_period with &&)
);

comment on table product_price is 'Prices of goods and services';
comment on column product_price.record_id is 'PK';
comment on column product_price.product_price_period is 'Effective date range for the price';
comment on column product_price.product_id is 'References to products';
comment on column product_price.product_price is 'Actual price';



