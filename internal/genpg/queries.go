package genpg

var GetInfoQuery = `
with ti as (select (select cls.relnamespace::regnamespace::text || '.' || cls.relname::text
                    from pg_class cls
                    where cls.oid = c.oid)                             as relpath,
                   c.oid                                               as reloid,
                   a.attname                                           as attname,
                   pg_catalog.format_type(a.atttypid, a.atttypmod)     as atttype,
                   t.typname                                           as atttype2,
                   a.attnum                                            as attnum,
                   (select max(ma.attnum)
                    from pg_attribute ma
                    where ma.attrelid = a.attrelid)                    as max_attnum,
                   (case
                        when a.attgenerated = 's' then
                            false
                        else
                            a.attnotnull
                       end)                                            as attnotnull,
                   a.attgenerated                                      as attgenerated,
                   a.attidentity                                       as attidentity,
                   coalesce(obj_description(c.oid, 'pg_class'), '')    as tab_desc,
                   coalesce(col_description(c.oid, a.attnum), '')      as col_desc,
                   (select exists (select *
                                   from pg_constraint cnested
                                            join pg_class tnested on cnested.conrelid = tnested.oid
                                            join pg_attribute anested on anested.attnum = any (cnested.conkey) and
                                                                         anested.attrelid = tnested.oid
                                   where cnested.contype = 'p'
                                     and tnested.relnamespace = c.relnamespace
                                     and tnested.relname = c.relname
                                     and anested.attnum = a.attnum
                                     and anested.attname = a.attname)) as is_pk
            from pg_attribute a
                     left join pg_class c on a.attrelid = c.oid
                     left join pg_type t on a.atttypid = t.oid
            where a.attnum > 0
              and c.relnamespace in
                  (select n.oid
                   from pg_namespace n
                   where n.nspname !~* '^pg_'
                     and n.nspname <> 'information_schema')
              and c.relkind = 'r'),

     fk as (select cn.oid                                      as conoid,
                   (select cls.relnamespace::regnamespace::text || '.' || cls.relname::text
                    from pg_class cls
                    where cls.oid = cn.conrelid)               as con_relpath,
                   coalesce(
                           (select cls.relnamespace::regnamespace::text || '.' || cls.relname::text
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

     limits as (select colinfo.table_schema || '.' || colinfo.table_name as relpath,
                       colinfo.column_name                               as attname,
                       colinfo.ordinal_position                          as attnum,
                       colinfo.character_maximum_length                  as char_max_len,
                       colinfo.numeric_precision                         as n_prec,
                       colinfo.numeric_scale                             as n_scal
                from information_schema.columns colinfo
                order by relpath, ordinal_position),

     def as (select (select c.relnamespace::regnamespace::text || '.' || c.relname
                     from pg_class c
                     where c.oid = pa.attrelid)                as relpath,
                    pa.attname                                 as attname,
                    pa.attnum                                  as attnum,
                    pg_catalog.pg_get_expr(d.adbin, d.adrelid) as attdefexpr,
                    (case
                         when pa.attgenerated = 's' then
                             'generated always as (' ||
                             pg_catalog.pg_get_expr(d.adbin, d.adrelid) ||
                             ') stored'
                         else
                             'default ' || pg_catalog.pg_get_expr(d.adbin, d.adrelid)
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
              ('_uuid', '[]string', $t$'{}'::uuid[]$t$),
              ('_varchar', '[]string', $t$'{}'::varchar[]$t$),
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
              ('xml', 'string', (select str from zeroes)))

select ti.relpath,
       ti.attname,
       ti.atttype,
       ti.atttype2,
       coalesce(fk.con_frelpath, '')                                    refto,
       ti.col_desc,
       ti.attnotnull,
       ti.tab_desc,
       ti.attnum,
       l.char_max_len,
       l.n_prec,
       l.n_scal,
       def.def,
       ti.is_pk,
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
       -- a column is insertable if:
       -- is not an 'IDENTITY', is not a 'GENERATED', is not a one of a 'SERIAL'
       (ti.attidentity = ''
           and ti.attgenerated = ''
           and (not coalesce(def.attdefexpr, '') ilike 'nextval(%')) as is_insertable
from ti
         left join fk on fk.con_relpath = ti.relpath and fk.conkey_first = ti.attnum
         left join limits l on l.relpath = ti.relpath and l.attnum = ti.attnum
         left join def on def.relpath = ti.relpath and def.attnum = ti.attnum
         left join typemap_tab tmt on tmt.column1 = ti.atttype2
where ti.relpath ~* 'public'
  and ti.relpath !~* 'migrate_'
order by ti.relpath, ti.attnum
;
`
