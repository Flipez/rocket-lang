package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNullType(t *testing.T) {
	if object.NULL.Type() != object.NULL_OBJ {
		t.Errorf("null.Type() returns wrong type")
	}
	if object.NULL.Inspect() != "null" {
		t.Errorf("null.Inspect() returns wrong type")
	}
}
func TestNullObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope()`, "undefined method `.nope()` for NULL"},
	}

	testInput(t, tests)
}
