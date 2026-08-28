package utilities

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var SearchPaths []string
var once sync.Once

func AddPath(path string) error {
	path = os.ExpandEnv(filepath.Clean(path))
	absolutePath, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	SearchPaths = append(SearchPaths, absolutePath)

	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func FindModule(name string) string {
	once.Do(initSearchPaths)

	basename := fmt.Sprintf("%s.rl", name)

	for _, p := range SearchPaths {
		filename := filepath.Join(p, basename)

		if Exists(filename) {
			return filename
		}
	}

	return ""
}

// FindModuleFrom resolves a module path. Paths beginning with "./" or "../"
// resolve against importerDir, the directory of the file doing the
// importing. Every other path goes through the search paths.
func FindModuleFrom(name string, importerDir string) string {
	if !strings.HasPrefix(name, "./") && !strings.HasPrefix(name, "../") {
		return FindModule(name)
	}

	if importerDir == "" {
		return ""
	}

	candidate := filepath.Join(importerDir, name+".rl")

	if !Exists(candidate) {
		return ""
	}

	absolutePath, err := filepath.Abs(candidate)
	if err != nil {
		return ""
	}

	return absolutePath
}
