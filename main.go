package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"genpg-v5/internal/app"
	"genpg-v5/internal/logger"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

const scaffoldDir = "examples/go-project-template-v7"

func main() {
	logger.Init(&logger.Opts{
		Level:     "debug",
		Format:    "text",
		AddSource: true,
	})

	outputFlag := flag.String("output", "", "output directory")
	connString := flag.String("conn", "postgres://postgres:postgres@localhost:5432/bookstore", "postgresql connection string")
	flag.Parse()

	outputPath := scaffoldDir
	if *outputFlag != "" {
		outputPath = *outputFlag
		if err := prepareOutputDir(outputPath); err != nil {
			slog.Error("failed to prepare output dir", slog.String("path", outputPath), slog.Any("err", err))
			os.Exit(1)
		}
	}

	slog.Debug("resolved config",
		slog.String("output", outputPath),
		slog.String("conn", *connString),
	)

	structs, err := app.GenStructs(*connString)
	if err != nil {
		slog.Error("failed to introspect database", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Debug("introspected tables", slog.Int("count", len(structs)))

	if err := writeInterfaces(structs, outputPath); err != nil {
		slog.Error("failed to write interfaces", slog.Any("err", err))
		os.Exit(1)
	}
	for _, s := range structs {
		slog.Debug("generating layers", slog.String("table", s.DbTableName))
		if err := writeRepoFiles(s, outputPath); err != nil {
			slog.Error("failed to write repo files", slog.String("table", s.DbTableName), slog.Any("err", err))
			os.Exit(1)
		}
		if err := writeServiceFiles(s, outputPath); err != nil {
			slog.Error("failed to write service files", slog.String("table", s.DbTableName), slog.Any("err", err))
			os.Exit(1)
		}
		if err := writeHandlerFiles(s, outputPath); err != nil {
			slog.Error("failed to write handler files", slog.String("table", s.DbTableName), slog.Any("err", err))
			os.Exit(1)
		}
	}
	slog.Info("done", slog.String("output", outputPath))
}

func prepareOutputDir(dst string) error {
	slog.Debug("preparing output dir", slog.String("dst", dst), slog.String("scaffold", scaffoldDir))
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	if err := os.CopyFS(dst, os.DirFS(scaffoldDir)); err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(dst, "internal/api"))
}

func writeInterfaces(s []app.TableToStructInfo, outputPath string) error {
	layer, err := app.GenInterfaces(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(outputPath, "internal/api/repository.go"), layer.RepoInterface); err != nil {
		return err
	}
	if err := writeFile(path.Join(outputPath, "internal/api/service.go"), layer.ServiceInterface); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath, "internal/api/handler.go"), layer.HandlerInterface)
}

func writeRepoFiles(s app.TableToStructInfo, outputPath string) error {
	layer, err := app.GenRepository(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/entity.go", s.OutDirName)), layer.Entity); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/repository.go", s.OutDirName)), layer.Repository)
}

func writeServiceFiles(s app.TableToStructInfo, outputPath string) error {
	layer, err := app.GenService(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/dto.go", s.OutDirName)), layer.Dto); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/service.go", s.OutDirName)), layer.Service)
}

func writeHandlerFiles(s app.TableToStructInfo, outputPath string) error {
	layer, err := app.GenHandler(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/payload.go", s.OutDirName)), layer.Payload); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath, fmt.Sprintf("internal/api/%s/handler.go", s.OutDirName)), layer.Handler)
}

func writeFile(entityPath, content string) error {
	slog.Debug("writing file", slog.String("path", entityPath))
	if err := os.MkdirAll(path.Dir(entityPath), 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", path.Dir(entityPath), err)
	}
	src, err := imports.Process(entityPath, []byte(content), nil)
	if err != nil {
		return fmt.Errorf("goimports %s: %w", entityPath, err)
	}
	src, err = format.Source(src, format.Options{})
	if err != nil {
		return fmt.Errorf("gofumpt %s: %w", entityPath, err)
	}
	if err = os.WriteFile(entityPath, src, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", entityPath, err)
	}
	return nil
}
