### Required tools:

- goimports
- gofumpt

### TODO:

- decide imports (time, pgtypes, etc), if an entity is used composite types, add imports to dtos/services/repos/etc. (
  OR: just use goimports)
- when a table has default values: decide which of them are included in create-update-requests (exclude system generated
  fields like: guid, created_at, etc..., include others)

### Filters:
```
// decide: if a field is required for 'INSERT'
func isFieldRequiredForInsert() {
  if isPkey(field) {
    // serial, or has 'default' (gen_random_uuid(), etc...)
    if isGenerated(field) {
      return false
    }
  }
  
  // internal fields like: created_at/updated_at/guid,
  // that were generated for all tables, 
  // that have default values, controlled by triggers, not updatable.
  if isInternal(field) {
    return false
  }
  
  // even if a field has 'default', it may be insertable
  return true
}

// decide: if a field is required for 'UPDATE'

```



