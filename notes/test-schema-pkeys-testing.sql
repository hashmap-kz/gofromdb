-- pkeys combination

create table serial_pk
(
    id   bigserial primary key,
    name text not null
);

create table uuid_pk
(
    id   uuid primary key,
    name text not null
);

create table composite_pk
(
    tenant_id bigint not null,
    code      text   not null,
    name      text   not null,
    primary key (tenant_id, code)
);

create table natural_pk
(
    code text primary key,
    name text not null
);

create table no_pk
(
    event_time timestamptz not null,
    message    text        not null
);

create table nullable_types
(
    id         bigserial primary key,
    name       text,
    amount     numeric,
    payload    jsonb,
    tags       text[],
    active     boolean,
    created_at timestamptz
);