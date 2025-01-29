-- +goose Up
-- +goose StatementBegin

drop function if exists public.before_update_row;
create or replace function public.before_update_row()
    returns trigger as
$$
begin
    -- forbid to modify internal fields values
    if new.created_at is distinct from old.created_at then
        raise exception 'updates to the "created_at" column are not allowed';
    end if;
    if new.guid is distinct from old.guid then
        raise exception 'updates to the "guid" column are not allowed';
    end if;
    if new.record_id is distinct from old.record_id then
        raise exception 'updates to the "record_id" column are not allowed';
    end if;

    -- update "updated_at" values
    new.updated_at = current_timestamp;

    return new;
end;
$$ language plpgsql;

-- helper functions, used to attach triggers for all tables in a given schema
drop function if exists public.attach_before_update_triggers;
create function public.attach_before_update_triggers(p_schema text) returns void as
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
        EXECUTE format('DROP TRIGGER IF EXISTS before_update_row ON %I.%I', p_schema, tbl.tablename);

        EXECUTE format('
            CREATE TRIGGER before_update_row
            BEFORE UPDATE ON %I.%I
            FOR EACH ROW EXECUTE FUNCTION public.before_update_row();', p_schema, tbl.tablename);
    end loop;
end
$fn$ language plpgsql;

-- +goose StatementEnd
