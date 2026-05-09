package core

import (
	"bytes"
	"fmt"
	"io/fs"
	"strings"
	"sync"
	"text/template"

	"genpg-v5/internal/tmplts"
)

var (
	tmplCache    map[string]*template.Template
	tmplCacheErr error
	tmplOnce     sync.Once
)

// compileTemplates parses every *.tmpl file from the embedded FS exactly once.
// text/template.Execute is safe for concurrent use with separate writers.
func compileTemplates(funcMap template.FuncMap) (map[string]*template.Template, error) {
	tmplOnce.Do(func() {
		cache := make(map[string]*template.Template)
		entries, err := fs.ReadDir(tmplts.FS, ".")
		if err != nil {
			tmplCacheErr = fmt.Errorf("read template dir: %w", err)
			return
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".tmpl") {
				continue
			}
			name := strings.TrimSuffix(e.Name(), ".tmpl")
			content, err := tmplts.FS.ReadFile(e.Name())
			if err != nil {
				tmplCacheErr = fmt.Errorf("read template %s: %w", e.Name(), err)
				return
			}
			tmpl, err := template.New(name).Funcs(funcMap).Parse(string(content))
			if err != nil {
				tmplCacheErr = fmt.Errorf("parse template %s: %w", name, err)
				return
			}
			cache[name] = tmpl
		}
		tmplCache = cache
	})
	return tmplCache, tmplCacheErr
}

func ExecTemplate(name string, data any, funcMap map[string]any) (string, error) {
	cache, err := compileTemplates(funcMap)
	if err != nil {
		return "", fmt.Errorf("init templates: %w", err)
	}
	tmpl, ok := cache[name]
	if !ok {
		return "", fmt.Errorf("template %q not found", name)
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute template %s: %w", name, err)
	}
	return buf.String(), nil
}
