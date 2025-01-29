package app

import (
	"bytes"
	"log"
	"text/template"
)

func ExecTemplate(name, t string, data map[string]any, funcMap map[string]any) string {
	var result bytes.Buffer
	tmpl, err := template.New(name).Funcs(funcMap).Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(&result, data)
	if err != nil {
		log.Fatal(err)
	}
	return result.String()
}
