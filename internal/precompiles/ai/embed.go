// embed.go
package ai

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed python_files/*
var pythonFiles embed.FS

// ExtractPythonFiles extracts the embedded Python files to the target directory.
func ExtractPythonFiles(targetDir string) error {
	return fs.WalkDir(pythonFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := pythonFiles.ReadFile(path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(targetDir, path)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}
		return os.WriteFile(destPath, data, 0644) //nolint: gosec // Other users may need to read these files.
	})
}

func GetRequirementsFile() []byte {
	content, err := pythonFiles.ReadFile("python_files/requirements.txt")
	if err != nil {
		panic(err)
	}
	return content
}
