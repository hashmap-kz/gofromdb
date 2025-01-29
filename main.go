package main

import (
	"fmt"

	"genpg-v5/internal/app"
)

func main() {
	structs := app.GenStructs()

	for _, s := range structs {
		repository := app.GenRepository(s)
		fmt.Printf("// @@@@@ REPO: %s\n", s.StructName)
		fmt.Println(repository.RepoEntity)
		fmt.Println(repository.RepoImpl)
	}
}
