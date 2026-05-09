-- +goose Up
-- +goose StatementBegin

drop function if exists public.add_columns_to_all_tables;
create or replace function public.add_columns_to_all_tables(p_schema_name text)
    returns void as
$$
declare
    v_table_name  text;
    v_column_name record;
    v_sql         text;
begin
    drop table if exists _tmp_default_columns;
    create temporary table _tmp_default_columns
    (
        column_name    text not null,
        column_def     text not null,
        column_comment text not null
    ) on commit drop;
    insert into _tmp_default_columns(column_name, column_def, column_comment)
    values ('created_at', 'timestamptz not null default current_timestamp', 'Internal field, creation TS'),
           ('updated_at', 'timestamptz not null default current_timestamp', 'Internal field, last updated TS'),
           ('guid', 'uuid not null default gen_random_uuid()', 'Internal field, UUID of the row');

    for v_table_name in
        select it.table_name
        from information_schema.tables it
        where it.table_schema = p_schema_name
          and it.table_type = 'BASE TABLE'
          and it.table_name not like '%log%'
          and it.table_name not like '%migrate%'
          and it.table_name not like '%flyway%'
    loop
        for v_column_name in (select * from _tmp_default_columns)
        loop
            if not exists (select 1
                           from information_schema.columns
                           where table_schema = p_schema_name
                             and table_name = v_table_name
                             and column_name = v_column_name.column_name) then
                -- add column
                v_sql = format(
                        'alter table %I.%I add column %I %s',
                        p_schema_name, v_table_name, v_column_name.column_name, v_column_name.column_def
                        );
                raise notice '%', v_sql;
                execute v_sql;

                -- add comment on column
                v_sql = format(
                        'comment on column %I.%I.%I is ''%s''',
                        p_schema_name, v_table_name, v_column_name.column_name, v_column_name.column_comment
                        );
                raise notice '%', v_sql;
                execute v_sql;
            end if;
        end loop;
    end loop;
end;
$$ language plpgsql;

-- +goose StatementEnd
