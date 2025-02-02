### Required tools:

- goimports
- gofumpt

### TODO:

- decide imports (time, pgtypes, etc), if an entity is used composite types, add imports to dtos/services/repos/etc. (
  OR: just use goimports)
- composite PK, and also: do not rely on hardcoded PK field name (find/update/delete: ID,record_id,etc), decide the key from DB


