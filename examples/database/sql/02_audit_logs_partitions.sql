-- +goose Up
-- +goose StatementBegin

-- Monthly partitions
CREATE TABLE audit_logs.logs_2025_01 PARTITION OF audit_logs.logs
    FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');

CREATE TABLE audit_logs.logs_2025_02 PARTITION OF audit_logs.logs
    FOR VALUES FROM ('2025-02-01') TO ('2025-03-01');

-- Simplify, prepare partitions for a year beforehand
do
$$
    declare
        v_start_date     date := '2025-01-01'; -- starting date for partitions
        v_end_date       date := '2027-01-01'; -- ending date (exclusive)
        v_current_date   date;
        v_partition_name text;
        v_sql_text       text;
    begin
        v_current_date := v_start_date;
        while v_current_date < v_end_date
        loop
            v_partition_name := format('logs_%s', to_char(v_current_date, 'YYYY_MM'));

            if not exists (select 1
                           from pg_tables pt
                           where pt.schemaname = 'audit_logs'
                             and pt.tablename = v_partition_name) then

                v_sql_text = format('
                    create table if not exists audit_logs.%I partition of audit_logs.logs
                    for values from (%L) to (%L);',
                                    v_partition_name,
                                    v_current_date,
                                    (v_current_date + interval '1 month')::date
                             );
                raise notice 'created partition: %', v_partition_name;
                raise notice '%', v_sql_text;
                execute v_sql_text;
            end if;

            -- move to the next month
            v_current_date := v_current_date + interval '1 month';
        end loop;
    end
$$;


-- +goose StatementEnd
