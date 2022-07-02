package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestBreakValue(t *testing.T) {
	bv := object.NewBreakValue(object.NewString("a"))

	if bv.Type() != object.BREAK_VALUE_OBJ {
		t.Errorf("breakValue.Type() returns wrong type")
	}
	if bv.Inspect() != `"a"` {
		t.Errorf("breakValue.Inspect() returns wrong string")
	}
}
