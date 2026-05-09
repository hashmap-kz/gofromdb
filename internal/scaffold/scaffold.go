package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// moduleNamePlaceholder is the sentinel value embedded in scaffold files that gets
// replaced with the user-supplied module name on copy.
const moduleNamePlaceholder = "go-project-template-v5"

// CopyTo writes all scaffold files into destDir, substituting moduleNamePlaceholder
// with moduleName in every file's content and in go.mod's module declaration.
func CopyTo(destDir, moduleName string) error {
	return fs.WalkDir(FS, "files", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel("files", path)
		// .tmpl suffix is used to prevent go:embed treating the file as a module
		// boundary (e.g. go.mod.tmpl → go.mod on copy).
		dest := filepath.Join(destDir, strings.TrimSuffix(rel, ".tmpl"))
		if d.IsDir() {
			return os.MkdirAll(dest, 0o755)
		}
		data, err := fs.ReadFile(FS, path)
		if err != nil {
			return err
		}
		content := strings.ReplaceAll(string(data), moduleNamePlaceholder, moduleName)
		return os.WriteFile(dest, []byte(content), 0o644)
	})
}
