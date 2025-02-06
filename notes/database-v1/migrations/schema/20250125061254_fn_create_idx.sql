-- +goose Up
-- +goose StatementBegin

-- helper functions, used to create indexes for all foreign-keys
drop function if exists public.create_indexes_to_fkeys;
create function public.create_indexes_to_fkeys(p_schema text) returns void as
$fn$
declare
    v_fk_record record;
    v_sql       text;
begin
    for v_fk_record in
        select conrelid::regclass as table_name,
               conname            as constraint_name,
               a.attname          as column_name
        from pg_constraint c
                 join pg_attribute a on a.attrelid = c.conrelid and a.attnum = any (c.conkey)
        where c.contype = 'f'
          and c.connamespace = p_schema::regnamespace
    loop
        v_sql = format(
                'CREATE INDEX IF NOT EXISTS idx_%I_%I_fk ON %I (%I);',
                v_fk_record.table_name, v_fk_record.column_name,
                v_fk_record.table_name, v_fk_record.column_name
                );
        raise notice '%', v_sql;
        execute v_sql;
    end loop;
end
$fn$ language plpgsql;

-- helper functions, used to create indexes for internal-fields
drop function if exists public.create_indexes_to_internal_fields;
create function public.create_indexes_to_internal_fields(p_schema text) returns void as
$fn$
declare
    v_tbl_record record;
    v_sql        text;
begin
    for v_tbl_record in
        select t.tablename
        from pg_tables t
        where t.schemaname = p_schema
          and not t.tablename ilike '%migrate%'
    loop
        -- guid column
        v_sql = format(
                'CREATE UNIQUE INDEX IF NOT EXISTS idx_%I_guid_unq ON %I (guid);',
                v_tbl_record.tablename,
                v_tbl_record.tablename
                );
        raise notice '%', v_sql;
        execute v_sql;

    end loop;
end
$fn$ language plpgsql;

-- +goose StatementEnd
