### Required tools:

- goimports
- gofumpt

### TODO:

- decide imports (time, pgtypes, etc), if an entity is used composite types, add imports to dtos/services/repos/etc. (
  OR: just use goimports)
- composite PK, and also: do not rely on hardcoded PK field name (find/update/delete: ID,record_id,etc), decide the key from DB
- gen tests

1. PK-less tables must not generate ID-based methods.
2. Composite PK order must be fixed.
3. Natural PKs must not be updated accidentally.
4. Path parsing must match actual PK Go types.
5. Type mapping needs a stricter tested subset or pgtype-heavy mapping.
6. Schema/table/column identifiers should be quoted.
7. Update semantics should avoid nullif-zero-value unless you intentionally want PATCH-like behavior.
