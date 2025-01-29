package main

import (
	"fmt"
	"os"

	"go-project-template-v5/internal/app"
)

func main() {
	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "localdev"
	}
	app.Run(fmt.Sprintf("./config/config-%s.yml", env))
}
