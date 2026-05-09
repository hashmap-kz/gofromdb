package app

import (
	"bytes"
	"fmt"
	"log/slog"
	"text/template"

	"genpg-v5/internal/tmplts"
)

func ExecTemplate(name string, data any, funcMap map[string]any) (string, error) {
	slog.Debug("exec template", slog.String("name", name))
	content, err := tmplts.FS.ReadFile(name + ".tmpl")
	if err != nil {
		return "", fmt.Errorf("read template %s: %w", name, err)
	}
	tmpl, err := template.New(name).Funcs(funcMap).Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("parse template %s: %w", name, err)
	}
	var result bytes.Buffer
	if err = tmpl.Execute(&result, data); err != nil {
		return "", fmt.Errorf("execute template %s: %w", name, err)
	}
	return result.String(), nil
}
