//go:build wasm

package repl

import (
	"fmt"
	"io"
)

var buildVersion = "v0.10.0"
var buildDate = "2021-12-27T21:13:44Z"

func Start(in io.Reader, out io.Writer) {}

func SplashVersion() string {
	return fmt.Sprintf("rocket-lang version %s (%s)\n", buildVersion, buildDate)
}
