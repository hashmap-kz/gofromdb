package genpg

var GetInfoQuery = `
with ti as (select (select format('%I.%I', cls.relnamespace::regnamespace, cls.relname)
                    from pg_class cls
                    where cls.oid = c.oid)                          as relpath,
                   c.oid                                            as reloid,
                   a.attname                                        as attname,
                   pg_catalog.format_type(a.atttypid, a.atttypmod)  as atttype,
                   t.typname                                        as atttype2,
                   a.attnum                                         as attnum,
                   (select max(ma.attnum)
                    from pg_attribute ma
                    where ma.attrelid = a.attrelid)                 as max_attnum,
                   (case
                        when a.attgenerated = 's' then
                            false
                        else
                            a.attnotnull
                       end)                                         as attnotnull,
                   a.attgenerated                                   as attgenerated,
                   a.attidentity                                    as attidentity,
                   coalesce(obj_description(c.oid, 'pg_class'), '') as tab_desc,
                   coalesce(col_description(c.oid, a.attnum), '')   as col_desc
            from pg_attribute a
                     left join pg_class c on a.attrelid = c.oid
                     left join pg_type t on a.atttypid = t.oid
            where a.attnum > 0
              and (not a.attisdropped)
              and c.relnamespace in
                  (select n.oid
                   from pg_namespace n
                   where n.nspname !~* '^pg_'
                     and n.nspname <> 'information_schema')
              and c.relkind = 'r'
              -- exclude partitioned tables
              and not exists (select 1
                              from pg_inherits inh
                              where inh.inhrelid = c.oid)),

     pk as (select format('%I.%I', n.nspname, c.relname) as relpath,
				   -- preserve pkeys order	
                   array_agg(a.attname order by k.ord)   as primary_keys
            from pg_constraint pkcn
                     join pg_class c on pkcn.conrelid = c.oid
                     join pg_namespace n on n.oid = c.relnamespace
                     join unnest(pkcn.conkey) with ordinality as k(attnum, ord) on true
                     join pg_attribute a on a.attnum = k.attnum and a.attrelid = c.oid
            where n.nspname !~* '^pg_'
              and n.nspname <> 'information_schema'
              and pkcn.contype = 'p'
            group by n.nspname, c.relname),

     fk as (select cn.oid                                      as conoid,
                   (select format('%I.%I', cls.relnamespace::regnamespace, cls.relname)
                    from pg_class cls
                    where cls.oid = cn.conrelid)               as con_relpath,
                   coalesce(
                           (select format('%I.%I', cls.relnamespace::regnamespace, cls.relname)
                            from pg_class cls
                            where cls.oid = cn.confrelid), '') as con_frelpath,
                   cn.conname                                  as conname,
                   cn.contype                                  as contype,
                   cn.conrelid                                 as conrelid,
                   cn.confrelid                                as confrelid,
                   cn.conkey                                   as conkey,
                   cn.conkey[1]                                as conkey_first,
                   cn.confkey                                  as confkey,
                   pg_get_constraintdef(cn.oid)                   condef
            from pg_constraint cn
            where cn.connamespace in
                  (select n.oid
                   from pg_namespace n
                   where n.nspname !~* '^pg_'
                     and n.nspname <> 'information_schema')
              and cn.contype = 'f'
            order by con_relpath, conname),

     limits as (select format('%I.%I', isc.table_schema, isc.table_name)
                                                    as relpath,
                       isc.column_name              as attname,
                       isc.ordinal_position         as attnum,
                       isc.character_maximum_length as char_max_len,
                       isc.numeric_precision        as n_prec,
                       isc.numeric_scale            as n_scal
                from information_schema.columns isc
                where isc.table_schema in
                      (select n.nspname::text
                       from pg_namespace n
                       where n.nspname !~* '^pg_'
                         and n.nspname <> 'information_schema')
                order by relpath, ordinal_position),

     def as (select (select format('%I.%I', c.relnamespace::regnamespace, c.relname)
                     from pg_class c
                     where c.oid = pa.attrelid)                as relpath,
                    pa.attname                                 as attname,
                    pa.attnum                                  as attnum,
                    pg_catalog.pg_get_expr(d.adbin, d.adrelid) as attdefexpr,
                    (case
                         when pa.attgenerated = 's' then
                             format('generated always as (%s) stored',
                                    pg_catalog.pg_get_expr(d.adbin, d.adrelid))
                         else
                             format('default %s',
                                    pg_catalog.pg_get_expr(d.adbin, d.adrelid))
                        end)                                   as def,
                    pa.attrelid                                as attrelid
             from pg_catalog.pg_attribute pa
                      left join pg_catalog.pg_attrdef d
                                on d.adrelid = pa.attrelid
                                    and d.adnum = pa.attnum
             where pa.atthasdef
             order by relpath, pa.attnum),

     zeroes as (select $t$''$t$                                           as str,
                       $t$''::bytea$t$                                    as byt,
                       $t$0::numeric$t$                                   as num,
                       $t$'{}'::json$t$                                   as jsn,
                       $t$'{}'::jsonb$t$                                  as jsb,
                       $t$'0001-01-01 00:00:00'$t$                        as dtz,
                       $t$0::int2$t$                                      as i16,
                       $t$0::int4$t$                                      as i32,
                       $t$0::int8$t$                                      as i64,
                       $t$0::float4$t$                                    as f32,
                       $t$0::float4$t$                                    as f64,
                       $t$'false'::bool$t$                                as bol,
                       $t$'empty'::int4range$t$                           as ri4,
                       $t$'empty'::int8range$t$                           as ri8,
                       $t$'empty'::numrange$t$                            as rnm,
                       $t$'empty'::daterange$t$                           as rdt,
                       $t$'empty'::tsrange$t$                             as rts,
                       $t$'empty'::tstzrange$t$                           as rtz,
                       $t$'00000000-0000-0000-0000-000000000000'::uuid$t$ as gid,
                       $t$'0'::interval$t$                                as itv,
                       $t$'0'::cid$t$                                     as ecd,
                       $t$'0'::oid$t$                                     as eod,
                       $t$'0'::xid$t$                                     as exd),
     typemap_tab as
         (values
              -- Arrays
              ('_bool', '[]bool', $t$'{}'::bool[]$t$),
              ('_date', '[]time.Time', $t$'{}'::date[]$t$),
              ('_float4', '[]float32', $t$'{}'::float4[]$t$),
              ('_float8', '[]float64', $t$'{}'::float8[]$t$),
              ('_int2', '[]int16', $t$'{}'::int2[]$t$),
              ('_int4', '[]int', $t$'{}'::int4[]$t$),
              ('_int8', '[]int64', $t$'{}'::int8[]$t$),
              ('_numeric', '[]string', $t$'{}'::numeric[]$t$),
              ('_text', '[]string', $t$'{}'::text[]$t$),
              ('_bpchar', '[]string', $t$'{}'::bpchar[]$t$),
              ('_uuid', '[]string', $t$'{}'::uuid[]$t$),
              ('_varchar', '[]string', $t$'{}'::varchar[]$t$),
              ('_timestamp', '[]time.Time', $t$'{}'::timestamp[]$t$),
              ('_timestamptz', '[]time.Time', $t$'{}'::timestamptz[]$t$),
              ('_time', '[]string', $t$'{}'::time[]$t$),
              ('_timetz', '[]string', $t$'{}'::timetz[]$t$),
              ('_json', '[]string', $t$'{}'::json[]$t$),
              ('_jsonb', '[]string', $t$'{}'::jsonb[]$t$),
              ('_bytea', '[][]byte', $t$'{}'::bytea[]$t$),
              -- Bytea
              ('bytea', '[]byte', (select byt from zeroes)),
              -- Serial Types
              ('smallserial', 'int16', (select i16 from zeroes)),
              ('serial', 'int', (select i32 from zeroes)),
              ('bigserial', 'int64', (select i64 from zeroes)),
              -- Integer Types
              ('int2', 'int16', (select i16 from zeroes)),
              ('int4', 'int', (select i32 from zeroes)),
              ('int8', 'int64', (select i64 from zeroes)),
              -- Arbitrary Precision Numbers
              ('decimal', 'string', (select num from zeroes)),
              ('numeric', 'string', (select num from zeroes)),
              -- Floating-Point Types
              ('float4', 'float32', (select f32 from zeroes)),
              ('float8', 'float64', (select f64 from zeroes)),
              -- Date/Time Types
              ('date', 'time.Time', (select dtz from zeroes) || '::date'),
              ('timestamp', 'time.Time', (select dtz from zeroes) || '::timestamp'),
              ('timestamptz', 'time.Time', (select dtz from zeroes) || '::timestamptz'),
              ('time', 'time.Time', (select dtz from zeroes) || '::time'),
              -- NOTE: should be parsed from string during scan - time.Parse("15:04:05-07", timetzStr)
              ('timetz', 'string', (select dtz from zeroes) || '::timetz'),
              -- Character Types
              ('bpchar', 'string', (select str from zeroes)),
              ('char', 'string', (select str from zeroes)),
              ('text', 'string', (select str from zeroes)),
              ('varchar', 'string', (select str from zeroes)),
              -- Boolean
              ('bool', 'bool', (select bol from zeroes)),

              -- pgtype handled types ++
              -- Range types
              ('int4range', 'pgtype.Range[int]', (select ri4 from zeroes)),
              ('int8range', 'pgtype.Range[int64]', (select ri8 from zeroes)),
              ('numrange', 'pgtype.Range[string]', (select rnm from zeroes)),
              ('daterange', 'pgtype.Range[time.Time]', (select rdt from zeroes)),
              ('tsrange', 'pgtype.Range[time.Time]', (select rts from zeroes)),
              ('tstzrange', 'pgtype.Range[time.Time]', (select rtz from zeroes)),
              -- Date/Time Types
              ('interval', 'pgtype.Interval', (select itv from zeroes)),
              -- pgtype handled types --

              -- JSON
              ('json', 'string', (select jsn from zeroes)),
              ('jsonb', 'string', (select jsb from zeroes)),
              -- ETC
              ('name', 'string', (select str from zeroes)),
              ('uuid', 'string', (select gid from zeroes)),
              ('cid', 'uint32', (select ecd from zeroes)),
              ('oid', 'uint32', (select eod from zeroes)),
              ('xid', 'uint32', (select exd from zeroes)),
              ('pg_lsn', 'string', (select str from zeroes)),
              ('regclass', 'string', (select str from zeroes)),
              ('regproc', 'string', (select str from zeroes)),
              ('regtype', 'string', (select str from zeroes)),
              ('tsquery', 'string', (select str from zeroes)),
              ('tsvector', 'string', (select str from zeroes)),
              ('xml', 'string', (select str from zeroes)),
              ('inet', 'string', (select str from zeroes)),
              ('cidr', 'string', (select str from zeroes)),
              ('macaddr', 'string', (select str from zeroes)),
              ('macaddr8', 'string', (select str from zeroes)),
              ('money', 'string', (select str from zeroes)),
              ('bit', 'string', (select str from zeroes)),
              ('varbit', 'string', (select str from zeroes)),
              ('jsonpath', 'string', (select str from zeroes)))

select ti.relpath,
       ti.attname,
       ti.atttype2,
       ti.col_desc,
       ti.attnotnull,
       ti.tab_desc,
       ti.attnum,
       pk.primary_keys,
       case
           when coalesce(tmt.column2, '') = ''
               then ''
           else
               case
                   -- is nullable, and not an array, make a pointer to a type
                   -- slices are already pointers
                   when (not ti.attnotnull) and (not tmt.column2 ilike '[]%') then
                       '*' || tmt.column2
                   else
                       tmt.column2
                   end
           end                                                       as go_type,
       tmt.column3                                                   as nullif_rhs,
       -- A column is insertable if it is neither an 'IDENTITY' nor 'GENERATED', nor one of the 'SERIAL' types.
       -- NOTE:
       -- Some columns are technically insertable, but have default values and should not be
       -- user-managed.
       -- These columns require a dedicated config file so they can be excluded from
       -- insert/update methods while still being preserved in get/list methods.
       (ti.attidentity = ''
           and ti.attgenerated = ''
           and (not coalesce(def.attdefexpr, '') ilike 'nextval(%')) as is_insertable
from ti
         left join fk on fk.con_relpath = ti.relpath and fk.conkey_first = ti.attnum
         left join pk on pk.relpath = ti.relpath
         left join limits l on l.relpath = ti.relpath and l.attnum = ti.attnum
         left join def on def.relpath = ti.relpath and def.attnum = ti.attnum
         left join typemap_tab tmt on tmt.column1 = ti.atttype2
where ti.relpath !~* 'migrate_'
order by ti.relpath, ti.attnum
;
`
