package stdlib

import (
	"bytes"
	"strings"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// captureExit redirects the process-ending path for one test and returns what
// was recorded: the codes exitProcess was called with, and whatever was written
// alongside them.
func captureExit(t *testing.T) (codes *[]int, output *bytes.Buffer) {
	t.Helper()

	previousExit, previousOutput := exitProcess, exitOutput
	t.Cleanup(func() {
		exitProcess, exitOutput = previousExit, previousOutput
	})

	recorded := []int{}
	buffer := &bytes.Buffer{}

	// The real os.Exit does not return, so the code after it never runs. This
	// one does, which is the one way the test differs from production.
	exitProcess = func(code int) { recorded = append(recorded, code) }
	exitOutput = buffer

	return &recorded, buffer
}

func call(name string, args ...object.Object) object.Object {
	return osFunctions[name].Call(args, *object.NewEnvironment())
}

func TestOSExitUsesTheGivenCode(t *testing.T) {
	for _, code := range []int{0, 1, 3, 127} {
		t.Run(object.NewInteger(code).Inspect(), func(t *testing.T) {
			codes, output := captureExit(t)

			call("exit", object.NewInteger(code))

			if len(*codes) != 1 || (*codes)[0] != code {
				t.Errorf("expected exit with %d, got %v", code, *codes)
			}

			if output.Len() != 0 {
				t.Errorf("expected exit to print nothing, got %q", output.String())
			}
		})
	}
}

func TestOSRaisePrintsThenExits(t *testing.T) {
	codes, output := captureExit(t)

	call("raise", object.NewInteger(2), object.NewString("broken"))

	if len(*codes) != 1 || (*codes)[0] != 2 {
		t.Errorf("expected exit with 2, got %v", *codes)
	}

	// Pinned as documented in docs/builtins/os.yml, quotes included: the
	// message goes through Inspect.
	expected := "🔥 RocketLang raised an error: \"broken\"\n"
	if output.String() != expected {
		t.Errorf("expected %q, got %q", expected, output.String())
	}
}

// TestOSFunctionsValidateTheirArguments covers the layout doing the checking, so
// neither function reaches the process-ending path with the wrong arguments.
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
		{"raise without a message", "raise", []object.Object{object.NewInteger(1)}, "too few arguments"},
		{"raise with a non-string message", "raise", []object.Object{object.NewInteger(1), object.NewInteger(2)}, "wrong argument type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes, _ := captureExit(t)

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
