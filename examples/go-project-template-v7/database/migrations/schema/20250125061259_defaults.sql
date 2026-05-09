-- +goose Up

-- create default columns, attach triggers
select public.add_columns_to_all_tables('public');
select audit_logs.attach_audit_triggers('public');
select public.attach_before_update_triggers('public');
select public.create_indexes_to_fkeys('public');
select public.create_indexes_to_internal_fields('public');
