//go:build !wasm

package utilities

import (
	"log"
	"os"
)

func initSearchPaths() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Printf("error getting cwd: %s", err)
	}

	for _, path := range searchPathList(os.Getenv("ROCKETLANGPATH"), cwd) {
		if err := AddPath(path); err != nil {
			// A bad entry costs us one search path, not the whole process.
			log.Printf("error adding search path %q: %s", path, err)
		}
	}
}
