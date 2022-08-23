package utilities

import (
	"fmt"
	"os"
	"path/filepath"
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
