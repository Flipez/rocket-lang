//go:build wasm

package planet

import (
	"fmt"
	"io"
)

// Command is unavailable under wasm: there is no filesystem to install into and
// no process to run git in. Imports do not resolve there either.
func Command(_ []string, _ io.Writer, stderr io.Writer) int {
	fmt.Fprintln(stderr, "planet is not available in this build")

	return 2
}
