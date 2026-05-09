drop table if exists all_types_test_arrays;
create table all_types_test_arrays
(
    record_id       serial primary key,
    bool_arr_col    bool[]    not null,
    date_arr_col    date[]    not null,
    float4_arr_col  float4[]  not null,
    float8_arr_col  float8[]  not null,
    int2_arr_col    int2[]    not null,
    int4_arr_col    int4[]    not null,
    int8_arr_col    int8[]    not null,
    numeric_arr_col numeric[] not null,
    text_arr_col    text[]    not null,
    uuid_arr_col    uuid[]    not null,
    varchar_arr_col varchar[] not null
);

INSERT INTO all_types_test_arrays
(bool_arr_col,
 date_arr_col,
 float4_arr_col,
 float8_arr_col,
 int2_arr_col,
 int4_arr_col,
 int8_arr_col,
 numeric_arr_col,
 text_arr_col,
 uuid_arr_col,
 varchar_arr_col)
VALUES (ARRAY [true, false, true],
        ARRAY ['2023-01-01'::date, '2024-01-01'::date],
        ARRAY [1.1, 2.2, 3.3],
        ARRAY [1.11, 2.22, 3.33],
        ARRAY [1, 2, 3],
        ARRAY [10, 20, 30],
        ARRAY [100, 200, 300],
        ARRAY [1.123, 2.234, 3.345],
        ARRAY ['hello', 'world'],
        ARRAY ['550e8400-e29b-41d4-a716-446655440000'::uuid, '123e4567-e89b-12d3-a456-426614174000'::uuid],
        ARRAY ['test1', 'test2']);

drop table if exists all_types_test_integers;
create table all_types_test_integers
(
    record_id       serial primary key,
    smallserial_col smallserial not null,
    serial_col      serial      not null,
    bigserial_col   bigserial   not null,
    int2_col        int2        not null default 0,
    int4_col        int4        not null default 0,
    int8_col        int8        not null default 0
);
insert into all_types_test_integers(int2_col, int4_col, int8_col)
VALUES (42,
        100,
        1000);
