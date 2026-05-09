-- +goose Up
-- +goose StatementBegin

-- Links:
-- https://github.com/kvesteri/postgresql-audit
-- https://github.com/2ndQuadrant/audit-trigger
-- http://coussej.github.io/2016/05/24/A-Minus-Operator-For-PostgreSQLs-JSONB/

-- Docs:
-- https://postgrespro.ru/docs/postgresql/17/plpgsql-trigger
-- https://postgrespro.ru/docs/postgresql/17/functions-json

create schema audit_logs;

create table audit_logs.logs
(
    record_id             bigserial   not null,
    schema_name           name        not null,
    table_name            name        not null,
    table_oid             oid         not null,
    operation_type        varchar(32) not null,
    changed_at            timestamptz not null,
    changed_by            name        not null,
    native_transaction_id xid8        not null,
    old_data              jsonb       not null,
    new_data              jsonb       not null,
    row_guid              uuid        not null
) partition by range (changed_at);

create index ix_audit_logs_search on audit_logs.logs (table_name, operation_type, row_guid);

comment on table audit_logs.logs is $comment$
Centralized repository for tracking all data changes in the database.
$comment$;
comment on column audit_logs.logs.record_id is $comment$ Unique ID for each log entry $comment$;
comment on column audit_logs.logs.schema_name is $comment$ Name of the schema where the change occurred $comment$;
comment on column audit_logs.logs.table_name is $comment$ Name of the table where the change occurred $comment$;
comment on column audit_logs.logs.table_oid is $comment$ OID of the table where the change occurred $comment$;
comment on column audit_logs.logs.operation_type is $comment$ Type of operation ('INSERT', 'UPDATE', 'DELETE') $comment$;
comment on column audit_logs.logs.changed_at is $comment$ Timestamp of the change $comment$;
comment on column audit_logs.logs.changed_by is $comment$ Session user responsible for the change $comment$;
comment on column audit_logs.logs.native_transaction_id is $comment$ Current transaction ID pg_current_xact_id() $comment$;
comment on column audit_logs.logs.old_data is $comment$ Data before the change (for UPDATE/DELETE) $comment$;
comment on column audit_logs.logs.new_data is $comment$ Data after the change (for INSERT/UPDATE) $comment$;
comment on column audit_logs.logs.row_guid is $comment$ Each table should has a column 'guid' $comment$;

drop function if exists audit_logs.jsonb_subtract;
create function audit_logs.jsonb_subtract(arg1 jsonb, arg2 jsonb) returns jsonb as
$$
select coalesce(json_object_agg(key, value), '{}')::jsonb
from
    jsonb_each(arg1)
where (arg1 -> key) <> (arg2 -> key)
   or (arg2 -> key) is null
$$ language SQL;

drop operator if exists - (jsonb, jsonb);
create operator - ( leftarg = jsonb, rightarg = jsonb, procedure = audit_logs.jsonb_subtract );

-- helper functions, used to attach triggers for all tables in a given schema
drop function if exists audit_logs.attach_audit_triggers;
create function audit_logs.attach_audit_triggers(p_schema text) returns void as
$fn$
declare
    tbl record;
begin
    for tbl in
        select t.tablename
        from pg_tables t
        where t.schemaname = p_schema
          and not t.tablename ilike '%migrate%'
    loop
        EXECUTE format('DROP TRIGGER IF EXISTS audit_trigger_update ON %I.%I', p_schema, tbl.tablename);
        EXECUTE format('DROP TRIGGER IF EXISTS audit_trigger_insert ON %I.%I', p_schema, tbl.tablename);
        EXECUTE format('DROP TRIGGER IF EXISTS audit_trigger_delete ON %I.%I', p_schema, tbl.tablename);

        -- update
        EXECUTE format('
            CREATE TRIGGER audit_trigger_update
            AFTER UPDATE ON %I.%I
            REFERENCING NEW TABLE AS new_table OLD TABLE AS old_table
            FOR EACH STATEMENT EXECUTE FUNCTION audit_logs.log_changes();', p_schema, tbl.tablename);

        -- insert
        EXECUTE format('
            CREATE TRIGGER audit_trigger_insert
            AFTER INSERT ON %I.%I
            REFERENCING NEW TABLE AS new_table
            FOR EACH STATEMENT EXECUTE FUNCTION audit_logs.log_changes();', p_schema, tbl.tablename);

        -- delete
        EXECUTE format('
            CREATE TRIGGER audit_trigger_delete
            AFTER DELETE ON %I.%I
            REFERENCING OLD TABLE AS old_table
            FOR EACH STATEMENT EXECUTE FUNCTION audit_logs.log_changes();', p_schema, tbl.tablename);
    end loop;
end
$fn$ language plpgsql;

drop function if exists audit_logs.log_changes;
create function audit_logs.log_changes() returns trigger as
$fn$
declare
begin
    if TG_WHEN <> 'AFTER' then
        raise exception 'audit_logs.log_changes() may only run as an after trigger';
    end if;
    if TG_LEVEL <> 'STATEMENT' then
        raise exception 'audit_logs.log_changes() may only run as a statement-level trigger';
    end if;

    if (TG_OP = 'UPDATE') then
        insert into audit_logs.logs
        (schema_name,
         table_name,
         table_oid,
         operation_type,
         changed_at,
         changed_by,
         native_transaction_id,
         old_data,
         new_data,
         row_guid)
        select TG_TABLE_SCHEMA       as schema_name,
               TG_TABLE_NAME         as table_name,
               TG_RELID              as table_oid,
               LOWER(TG_OP)          as operation_type,
               statement_timestamp() as changed_at,
               session_user          as changed_by,
               pg_current_xact_id()  as native_transaction_id,
               sub.old_data          as old_data,
               sub.new_data          as new_data,
               sub.old_guid          as row_guid
        from (select o.guid        as old_guid,
                     to_jsonb(o.*) as old_data,
                     to_jsonb(n.*) as new_data
              from old_table o
                       join new_table n using (record_id)) as sub;

    elsif (TG_OP = 'INSERT') then
        insert into audit_logs.logs
        (schema_name,
         table_name,
         table_oid,
         operation_type,
         changed_at,
         changed_by,
         native_transaction_id,
         old_data,
         new_data,
         row_guid)
        select TG_TABLE_SCHEMA       AS schema_name,
               TG_TABLE_NAME         AS table_name,
               TG_RELID              AS table_oid,
               LOWER(TG_OP)          AS operation_type,
               statement_timestamp() AS changed_at,
               session_user          as changed_by,
               pg_current_xact_id()  AS native_transaction_id,
               '{}'::jsonb           AS old_data,
               to_jsonb(new_table.*) as new_data,
               new_table.guid        as row_guid
        from new_table;

    elseif TG_OP = 'DELETE' then
        insert into audit_logs.logs
        (schema_name,
         table_name,
         table_oid,
         operation_type,
         changed_at,
         changed_by,
         native_transaction_id,
         old_data,
         new_data,
         row_guid)
        select TG_TABLE_SCHEMA       AS schema_name,
               TG_TABLE_NAME         AS table_name,
               TG_RELID              AS table_oid,
               LOWER(TG_OP)          AS operation_type,
               statement_timestamp() AS changed_at,
               session_user          as changed_by,
               pg_current_xact_id()  AS native_transaction_id,
               to_jsonb(old_table.*) AS old_data,
               '{}'::jsonb           as new_data,
               old_table.guid        as row_guid
        from old_table;

    end if;

    return null;

    --
    -- Implementation notes:
    --

    -- UNOPTIMIZED JOIN VERSION-orig
    --
    -- FROM (
    --     SELECT *
    --     FROM (
    --         SELECT
    --             row_to_json(old_table.*)::jsonb AS old_data,
    --             row_number() OVER ()
    --         FROM old_table
    --     ) AS old_table
    --     JOIN (
    --         SELECT
    --             row_to_json(new_table.*)::jsonb AS new_data,
    --             row_number() OVER ()
    --         FROM new_table
    --     ) AS new_table
    --     USING(row_number)
    -- ) as sub
    -- WHERE new_data - old_data - excluded_cols != '{}'::jsonb;

    -- UNOPTIMIZED JOIN VERSION-1
    --
    -- from (select new_data, old_data, old_guid
    --       from (select old_table.record_id,
    --                    guid                  as old_guid,
    --                    to_jsonb(old_table.*) as old_data
    --             from old_table) as old_table
    --                join (select new_table.record_id,
    --                             to_jsonb(new_table.*) as new_data
    --                      from new_table) as new_table
    --                     using (record_id)) as sub;
    --
    -- EXAMPLE CHECK QUERY
    --
    -- with old_table as (select b.record_id, b.description, b.guid
    --                    from public.buy b
    --                    where b.record_id = 1),
    --      new_table as (select n.record_id, upper(n.description) as description, n.guid
    --                    from public.buy n
    --                    where record_id = 1)
    -- select *
    -- from (select new_data, old_data, old_guid, new_data - old_data as diff
    --       from (select old_table.record_id,
    --                    guid                  as old_guid,
    --                    to_jsonb(old_table.*) as old_data
    --             from old_table) as old_table
    --                join (select new_table.record_id,
    --                             to_jsonb(new_table.*) as new_data
    --                      from new_table) as new_table
    --                     using (record_id)) as sub
    -- ;

    -- OPTIMIZED JOIN V2
    --
    -- EXAMPLE CHECK QUERY
    --
    -- with old_table as (select b.record_id, b.description, b.guid
    --                    from public.buy b
    --                    where b.record_id between 1 and 5),
    --      new_table as (select n.record_id, upper(n.description) as description, n.guid
    --                    from public.buy n
    --                    where record_id between 1 and 5)
    -- select *
    -- from (select o.guid        as old_guid,
    --              to_jsonb(o.*) as old_data,
    --              to_jsonb(n.*) as new_data
    --       from old_table o
    --                join new_table n using (record_id)) as sub

    -- USAGE EXAMPLE:
    --
    -- select l.old_data,
    --        l.new_data,
    --        l.new_data - l.old_data as diff
    -- from audit_logs.logs l
    -- where l.operation_type = 'update'
    -- limit 10
    -- ;

end
$fn$ language plpgsql
    security definer
    set search_path to pg_catalog, public;

-- +goose StatementEnd
