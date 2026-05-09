package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"genpg-v5/internal/app"
	"genpg-v5/internal/scaffold"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

func main() {
	conn := flag.String("conn", "postgres://postgres:postgres@localhost:5432/bookstore", "postgres connection string")
	out := flag.String("out", path.Join("examples", "go-project-template-v7"), "output directory")
	module := flag.String("module", "go-project-template-v5", "go module name for the generated project")
	flag.Parse()

	if err := scaffold.CopyTo(*out, *module); err != nil {
		log.Fatal(err)
	}

	structs := app.GenStructs(*conn)

	writeInterfaces(structs, *out)
	for _, s := range structs {
		writeRepoFiles(s, *out)
		writeServiceFiles(s, *out)
		writeHandlerFiles(s, *out)
	}
}

func writeInterfaces(s []app.TableToStructInfo, outputPath string) {
	layer := app.GenInterfaces(s)
	writeFile(path.Join(outputPath, "internal/api/repository.go"), layer.RepoInterface)
	writeFile(path.Join(outputPath, "internal/api/service.go"), layer.ServiceInterface)
	writeFile(path.Join(outputPath, "internal/api/handler.go"), layer.HandlerInterface)
}

func writeRepoFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenRepository(s)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/entity.go", s.DbTableName)), layer.Entity)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository.go", s.DbTableName)), layer.Repository)
}

func writeServiceFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenService(s)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/dto.go", s.DbTableName)), layer.Dto)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/service.go", s.DbTableName)), layer.Service)
}

func writeHandlerFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenHandler(s)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/payload.go", s.DbTableName)), layer.Payload)
	writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/handler.go", s.DbTableName)), layer.Handler)
}

func writeFile(entityPath, content string) {
	if err := os.MkdirAll(path.Dir(entityPath), 0o755); err != nil {
		log.Fatal(err)
	}
	src, err := imports.Process(entityPath, []byte(content), nil)
	if err != nil {
		log.Fatal(err)
	}
	src, err = format.Source(src, format.Options{})
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile(entityPath, src, 0o644); err != nil {
		log.Fatal(err)
	}
}
