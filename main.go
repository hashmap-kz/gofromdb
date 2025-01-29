package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"genpg-v5/internal/app"
)

func main() {
	structs := app.GenStructs()
	outputPath := "tmpgen"

	// Paths example:
	// internal/api/client/entity/postgres/client_entity_pg.go
	// internal/api/client/repository/client_repository.go
	// internal/api/client/repository/impl/client_repository_pg.go

	for _, s := range structs {
		repository := app.GenRepository(s)

		// entity
		entityPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/entity/postgres/%s_entity_pg.go",
			s.DbTableName,
			s.DbTableName),
		)
		writeFile(entityPath, repository.RepoEntity)

		// interface
		interfacePath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository/%s_repository.go",
			s.DbTableName,
			s.DbTableName),
		)
		writeFile(interfacePath, repository.RepoInterface)

		// interface impl
		interfaceImplPath := path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository/impl/%s_repository_pg.go",
			s.DbTableName,
			s.DbTableName),
		)
		writeFile(interfaceImplPath, repository.RepoImpl)
	}
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
