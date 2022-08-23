//go:build !wasm

package utilities

import (
	"log"
	"os"
	"strings"
)

func initSearchPaths() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Printf("error getting cwd: %s", err)
	}

	if e := os.Getenv("ROCKETLANGPATH"); e != "" {
		tokens := strings.Split(e, ":")

		for _, token := range tokens {
			if err := AddPath(token); err != nil {
				log.Fatalf("error adding token: %s", err)
			}
		}
	} else {
		SearchPaths = append(SearchPaths, cwd)
	}
}
