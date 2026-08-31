package stdlib

import "testing"

// print replaces puts. Checked here rather than through the evaluator for the
// same reason: this package registers the name, so this package can see it.
func TestPrintIsRegisteredAndPutsIsNot(t *testing.T) {
	if _, exists := Functions["print"]; !exists {
		t.Error("print should be registered")
	}
	if _, exists := Functions["puts"]; exists {
		t.Error("puts should be gone")
	}
}
