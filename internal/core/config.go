package core

// Config controls code-generation behaviour.
type Config struct {
	// SkipColumns lists DB column names that are excluded from insert/update
	// DTOs but still appear in the full entity struct (e.g. created_at, guid).
	SkipColumns []string `json:"skip_columns"`
}

func (c Config) skipSet() map[string]struct{} {
	m := make(map[string]struct{}, len(c.SkipColumns))
	for _, col := range c.SkipColumns {
		m[col] = struct{}{}
	}
	return m
}
