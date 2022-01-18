package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestErrorType(t *testing.T) {
	err := object.NewError("test")

	if err.Type() != object.ERROR_OBJ {
		t.Errorf("error.Type() returns wrong type")
	}
	if err.Inspect() != "ERROR: test" {
		t.Errorf("error.Inspect() returns wrong type")
	}
}
func TestErrorObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope().nope()`, "undefined method `.nope()` for ERROR"},
	}

	testInput(t, tests)
}
