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
	outputPath := path.Join("examples", "go-project-template-v5")

	// Paths example:
	//
	// internal/api/client/entity/postgres/client_entity_pg.go
	// internal/api/client/repository/client_repository.go
	// internal/api/client/repository/impl/client_repository_pg.go
	//
	// internal/api/client/dto/client_dto.go
	// internal/api/client/service/client_service.go
	// internal/api/client/service/impl/client_service_impl.go
	//
	// internal/api/client/handler/v1/client_payload.go

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

func writeHandlerFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenHandler(s)

	// models
	modelsPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/handler/v1/%s_payload.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(modelsPath, layer.HandlerDtos)

	// impl
	implPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/handler/v1/%s_handler.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(implPath, layer.HandlerImpl)
}

func writeServiceFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenService(s)

	// models
	modelsPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/dto/%s_dto.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(modelsPath, layer.ServiceDtos)

	// interface
	interfacePath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/service/%s_service.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(interfacePath, layer.ServiceInterface)

	// interface impl
	implPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/service/impl/%s_service_impl.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(implPath, layer.ServiceImpl)
}

func writeRepoFiles(s app.TableToStructInfo, outputPath string) {
	layer := app.GenRepository(s)

	// models
	modelsPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/entity/postgres/%s_entity_pg.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(modelsPath, layer.RepoEntity)

	// interface
	interfacePath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository/%s_repository.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(interfacePath, layer.RepoInterface)

	// interface impl
	implPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository/impl/%s_repository_pg.go",
		s.DbTableName,
		s.DbTableName),
	)
	writeFile(implPath, layer.RepoImpl)
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
