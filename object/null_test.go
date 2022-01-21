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
		{`[1][1].nope()`, "Failed to invoke method: nope"},
		{`[1][1].plz_s()`, ""},
		{`[1][1].plz_i()`, 0},
		{`[1][1].plz_f()`, 0.0},
	}

	testInput(t, tests)
}
