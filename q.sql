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
                            case
                                when a.attnotnull then
                                    true
                                else false
                                end
                       end)                                            as attnotnull,
                   a.attgenerated                                      as attgenerated,
                   coalesce(obj_description(c.oid, 'pg_class'), '')    as tab_desc,
                   coalesce((select d.description
                             from pg_description d
                             where d.objoid = c.oid
                               and d.objsubid = a.attnum), '')         as col_desc,
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
            from pg_class c,
                 pg_attribute a,
                 pg_type t
            where a.attnum > 0
              and a.attrelid = c.oid
              and a.atttypid = t.oid
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
                where colinfo.table_schema = 'public'
                order by relpath, ordinal_position),

     def as (select (select c.relnamespace::regnamespace::text || '.' || c.relname
                     from pg_class c
                     where c.oid = pa.attrelid) as relpath,
                    pa.attname                  as attname,
                    pa.attnum                   as attnum,
                    (case
                         when pa.attgenerated = 's' then
                             'generated always as (' ||
                             pg_catalog.pg_get_expr(d.adbin, d.adrelid) ||
                             ') stored'
                         else
                             'default ' || pg_catalog.pg_get_expr(d.adbin, d.adrelid)
                        end)                    as def,
                    pa.attrelid                 as attrelid
             from pg_catalog.pg_attrdef d,
                  pg_catalog.pg_attribute pa
             where d.adrelid = pa.attrelid
               and d.adnum = pa.attnum
               and pa.atthasdef
             order by relpath, pa.attnum),

     type_map AS (SELECT '{
       "_bool": "[]bool",
       "_date": "[]time.Time",
       "_float4": "[]float32",
       "_float8": "[]float64",
       "_int2": "[]int16",
       "_int4": "[]int",
       "_int8": "[]int64",
       "_numeric": "[]string",
       "_text": "[]string",
       "_uuid": "[]string",
       "_varchar": "[]string",
       "bytea": "[]byte",
       "bigserial": "int64",
       "bool": "bool",
       "bpchar": "string",
       "char": "string",
       "cid": "uint32",
       "date": "time.Time",
       "decimal": "string",
       "float4": "float32",
       "float8": "float64",
       "int2": "int16",
       "int4": "int",
       "int8": "int64",
       "interval": "string",
       "json": "string",
       "jsonb": "string",
       "name": "string",
       "numeric": "string",
       "oid": "uint32",
       "pg_lsn": "string",
       "regclass": "string",
       "regproc": "string",
       "regtype": "string",
       "serial": "int",
       "smallserial": "int16",
       "text": "string",
       "time": "time.Time",
       "timestamp": "time.Time",
       "timestamptz": "time.Time",
       "timetz": "time.Time",
       "tsquery": "string",
       "tsvector": "string",
       "uuid": "string",
       "varchar": "string",
       "xid": "uint32",
       "xml": "string"
     }'::jsonb AS mapping),

     typemap_tab as (SELECT key   AS pg_type,
                            value AS go_type
                     FROM jsonb_each_text((SELECT mapping FROM type_map)))

select ti.relpath,
       ti.attname,
       ti.atttype,
       ti.atttype2,
       coalesce(fk.con_frelpath, '') refto,
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
           when coalesce(tmt.go_type, '') = ''
               then ''
           else
               case
                   when not ti.attnotnull then
                       case
                           when tmt.go_type ilike '%[]%' then
                               replace(tmt.go_type, '[]', '[]*')
                           else
                               '*' || tmt.go_type
                           end
                   else
                       tmt.go_type
                   end
           end as                    go_type
from ti
         left join fk on fk.con_relpath = ti.relpath and fk.conkey_first = ti.attnum
         left join limits l on l.relpath = ti.relpath and l.attnum = ti.attnum
         left join def on def.relpath = ti.relpath and def.attnum = ti.attnum
         left join typemap_tab tmt on tmt.pg_type = ti.atttype2
where ti.relpath ~* 'public'
  and ti.relpath !~* 'migrate_'
order by ti.relpath, ti.attnum
;
