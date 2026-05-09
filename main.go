package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"genpg-v5/internal/app"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

const scaffoldDir = "examples/go-project-template-v7"

func main() {
	outputFlag := flag.String("output", "", "output directory")
	connString := flag.String("conn", "postgres://postgres:postgres@localhost:5432/bookstore", "postgresql connection string")
	flag.Parse()

	outputPath := scaffoldDir
	if *outputFlag != "" {
		outputPath = *outputFlag
		if err := prepareOutputDir(outputPath); err != nil {
			log.Fatal(err)
		}
	}

	structs := app.GenStructs(*connString)

	writeInterfaces(structs, outputPath)
	for _, s := range structs {
		writeRepoFiles(s, outputPath)
		writeServiceFiles(s, outputPath)
		writeHandlerFiles(s, outputPath)
	}
}

func prepareOutputDir(dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	if err := os.CopyFS(dst, os.DirFS(scaffoldDir)); err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(dst, "internal/api"))
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
