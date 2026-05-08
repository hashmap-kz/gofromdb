package app

import (
	"bytes"
	"log"
	"text/template"

	"genpg-v5/internal/tmplts"
)

func ExecTemplate(name string, data any, funcMap map[string]any) string {
	content, err := tmplts.FS.ReadFile(name + ".tmpl")
	if err != nil {
		log.Fatalf("ExecTemplate: read %s: %v", name, err)
	}
	var result bytes.Buffer
	tmpl, err := template.New(name).Funcs(funcMap).Parse(string(content))
	if err != nil {
		log.Fatalf("ExecTemplate: parse %s: %v", name, err)
	}
	if err = tmpl.Execute(&result, data); err != nil {
		log.Fatalf("ExecTemplate: execute %s: %v", name, err)
	}
	return result.String()
}
