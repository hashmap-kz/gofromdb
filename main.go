package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"genpg-v5/internal/app"
)

func main() {
	connString := "postgres://postgres:postgres@localhost:5432/bookstore"
	structs := app.GenStructs(connString)
	outputPath := path.Join("examples", "go-project-template-v7")

	writeInterfaces(structs, outputPath)
	for _, s := range structs {
		writeRepoFiles(s, outputPath)
		writeServiceFiles(s, outputPath)
		writeHandlerFiles(s, outputPath)
	}

	// cleanup imports, format output
	execCleanupCmd("goimports", "-w", ".")
	execCleanupCmd("gofumpt", "-w", ".")
}

func execCleanupCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
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
	err := os.MkdirAll(path.Dir(entityPath), 0o755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(entityPath, []byte(content), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
