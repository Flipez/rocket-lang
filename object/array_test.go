package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestArrayObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3].size()`, 3},
		{`[1,2,3].yeet()`, 3},
		{`let a = []; a.yoink(1); a`, "[1]"},
		{`[].nope()`, "Failed to invoke method: nope"},
		{`([].wat().lines().size() == [].methods().size() + 1).plz_s()`, "true"},
	}

	testInput(t, tests)
}

func TestArrayInspect(t *testing.T) {
	arr1 := &object.Array{Elements: []object.Object{}}

	if arr1.Type() != object.ARRAY_OBJ {
		t.Errorf("array.Type() returns wrong type")
	}
}
