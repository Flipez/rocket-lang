package main

import (
	"os"

	"github.com/flipez/rocket-lang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
