package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNullType(t *testing.T) {
	if object.NIL.Type() != object.NIL_OBJ {
		t.Errorf("nil.Type() returns wrong type")
	}
	if object.NIL.Inspect() != "nil" {
		t.Errorf("nil.Inspect() returns wrong type")
	}
}
func TestNilObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope()`, "test:1:7: undefined method `.nope()` for NIL"},
		{`[1][1].plz_s()`, ""},
		{`[1][1].plz_i()`, 0},
		{`[1][1].plz_f()`, 0.0},
	}

	testInput(t, tests)
}
