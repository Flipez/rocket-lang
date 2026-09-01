package stdlib

import (
	"strings"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// captureExit redirects the process-ending path for one test and returns the
// codes exitProcess was called with.
//
// The real os.Exit does not return, so the code after it never runs. This one
// does, which is the one way the test differs from production.
func captureExit(t *testing.T) *[]int {
	t.Helper()

	previous := exitProcess
	t.Cleanup(func() {
		exitProcess = previous
	})

	recorded := []int{}
	exitProcess = func(code int) { recorded = append(recorded, code) }

	return &recorded
}

func call(name string, args ...object.Object) object.Object {
	return osFunctions[name].Call(args, *object.NewEnvironment())
}

func TestOSExitUsesTheGivenCode(t *testing.T) {
	for _, code := range []int{0, 1, 3, 127} {
		t.Run(object.NewInteger(code).Inspect(), func(t *testing.T) {
			codes := captureExit(t)

			call("exit", object.NewInteger(code))

			if len(*codes) != 1 || (*codes)[0] != code {
				t.Errorf("expected exit with %d, got %v", code, *codes)
			}
		})
	}
}

// OS.raise let the caller pick the process's exit code; the global raise
// always exits 1 when uncaught, so the two were not equivalent. It goes
// anyway because the capability is reconstructible as print(msg); OS.exit(n).
func TestOSRaiseIsGone(t *testing.T) {
	if _, exists := Modules["OS"].Functions["raise"]; exists {
		t.Error("OS.raise should be gone; use print(msg); OS.exit(n)")
	}
}

// TestOSFunctionsValidateTheirArguments covers the layout doing the checking, so
// exit does not reach the process-ending path with the wrong arguments.
func TestOSFunctionsValidateTheirArguments(t *testing.T) {
	tests := []struct {
		name          string
		function      string
		args          []object.Object
		expectedError string
	}{
		{"exit without a code", "exit", nil, "too few arguments"},
		{"exit with a string", "exit", []object.Object{object.NewString("1")}, "wrong argument type"},
		{"exit with too many", "exit", []object.Object{object.NewInteger(1), object.NewInteger(2)}, "too many arguments"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes := captureExit(t)

			result := call(tt.function, tt.args...)

			if !object.IsError(result) {
				t.Fatalf("expected an error, got %v", result)
			}

			if !strings.Contains(result.Inspect(), tt.expectedError) {
				t.Errorf("expected an error containing %q, got %q", tt.expectedError, result.Inspect())
			}

			if len(*codes) != 0 {
				t.Errorf("the process should not have been ended, got exits %v", *codes)
			}
		})
	}
}
