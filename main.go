package main

import (
	"flag"
	"fmt"
	"github.com/hashmap-kz/gofromdb/internal/core"
	"github.com/hashmap-kz/gofromdb/internal/logger"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

const scaffoldDir = "examples/go-project-template-v7"

func main() {
	start := time.Now()

	logger.Init(&logger.Opts{
		Level:     "debug",
		Format:    "text",
		AddSource: true,
	})

	outputFlag := flag.String("output", "", "output directory")
	connString := flag.String(
		"conn",
		"postgres://postgres:postgres@localhost:15432/bookstore",
		"postgresql connection string",
	)
	workers := flag.Int("workers", 2, "concurrency")
	flag.Parse()

	outputPath := scaffoldDir
	if *outputFlag != "" {
		outputPath = *outputFlag
		if err := prepareOutputDir(outputPath); err != nil {
			slog.Error(
				"failed to prepare output dir",
				slog.String("path", outputPath),
				slog.Any("err", err),
			)
			os.Exit(1)
		}
	}

	slog.Debug("resolved config",
		slog.String("output", outputPath),
		slog.String("conn", *connString),
	)

	structs, err := core.GenStructs(*connString)
	if err != nil {
		slog.Error("failed to introspect database", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Debug("introspected tables", slog.Int("count", len(structs)))

	if err := writeInterfaces(structs, outputPath); err != nil {
		slog.Error("failed to write interfaces", slog.Any("err", err))
		os.Exit(1)
	}

	// process tables
	w := *workers
	if w <= 0 {
		w = runtime.NumCPU()
	}
	if w > len(structs) {
		w = len(structs)
	}

	g := new(errgroup.Group)
	g.SetLimit(w)
	for _, s := range structs {
		g.Go(func() error {
			slog.Debug("generating layers", slog.String("table", s.DbTableName))
			if err := writeRepoFiles(s, outputPath); err != nil {
				return fmt.Errorf("%s.%s repo: %w", s.DbSchemaName, s.DbTableName, err)
			}
			if err := writeServiceFiles(s, outputPath); err != nil {
				return fmt.Errorf("%s.%s service: %w", s.DbSchemaName, s.DbTableName, err)
			}
			if err := writeHandlerFiles(s, outputPath); err != nil {
				return fmt.Errorf("%s.%s handler: %w", s.DbSchemaName, s.DbTableName, err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		slog.Error("failed to generate files", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("done",
		slog.String("output", outputPath),
		slog.Duration("duration", time.Since(start)),
	)
}

func prepareOutputDir(dst string) error {
	slog.Debug(
		"preparing output dir",
		slog.String("dst", dst),
		slog.String("scaffold", scaffoldDir),
	)
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	if err := os.CopyFS(dst, os.DirFS(scaffoldDir)); err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(dst, "internal/api"))
}

func writeInterfaces(s []core.TableToStructInfo, outputPath string) error {
	layer, err := core.GenInterfaces(s)
	if err != nil {
		return err
	}
	if err := writeFile(
		path.Join(outputPath, "internal/api/repository.go"),
		layer.RepoInterface,
	); err != nil {
		return err
	}
	if err := writeFile(
		path.Join(outputPath, "internal/api/service.go"),
		layer.ServiceInterface,
	); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath, "internal/api/handler.go"), layer.HandlerInterface)
}

func writeRepoFiles(s core.TableToStructInfo, outputPath string) error {
	layer, err := core.GenRepository(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(
		outputPath,
		fmt.Sprintf(
			"internal/api/%s/%s/entity.go",
			s.DbSchemaName,
			s.DbTableName,
		),
	), layer.Entity); err != nil {
		return err
	}
	return writeFile(path.Join(
		outputPath,
		fmt.Sprintf(
			"internal/api/%s/%s/repository.go",
			s.DbSchemaName,
			s.DbTableName,
		),
	), layer.Repository)
}

func writeServiceFiles(s core.TableToStructInfo, outputPath string) error {
	layer, err := core.GenService(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(
		outputPath,
		fmt.Sprintf(
			"internal/api/%s/%s/dto.go",
			s.DbSchemaName,
			s.DbTableName,
		),
	), layer.Dto); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath,
		fmt.Sprintf("internal/api/%s/%s/service.go", s.DbSchemaName, s.DbTableName)), layer.Service)
}

func writeHandlerFiles(s core.TableToStructInfo, outputPath string) error {
	layer, err := core.GenHandler(s)
	if err != nil {
		return err
	}
	if err := writeFile(path.Join(
		outputPath,
		fmt.Sprintf(
			"internal/api/%s/%s/payload.go",
			s.DbSchemaName,
			s.DbTableName,
		),
	), layer.Payload); err != nil {
		return err
	}
	return writeFile(path.Join(outputPath,
		fmt.Sprintf("internal/api/%s/%s/handler.go", s.DbSchemaName, s.DbTableName)), layer.Handler)
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
