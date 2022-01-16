package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestArrayObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3][0]`, 1},
		{`[1,2,3].size()`, 3},
		{`[1,2,3].yeet()`, 3},
		{`[1,2,3].type()`, "ARRAY"},
		{`a = []; a.yoink(1); a`, "[1]"},
		{`[].nope()`, "Failed to invoke method: nope"},
		{`([].wat().lines().size() == [].methods().size() + 1).plz_s()`, "true"},
		{`a = ["a", "b"]; b = []; foreach i, item in a { b.yoink(item) }; b.size()`, 2},
		{`[1,2,3].index(4)`, -1},
		{`[1,2,3].index(3)`, 2},
		{`[1,2,3].index(true)`, -1},
		{`[1,2,3].index()`, "to few arguments: want=1, got=0"},
		{`a = []; b = []; foreach i in a { b.yoink(a[i]) }; a.size()==b.size()`, true},
		{`[1,1,2].uniq().size()`, 2},
		{`[true,true,2].uniq().size()`, 2},
		{`["test","test",2].uniq().size()`, 2},
	}

	testInput(t, tests)
}

func TestArrayInspect(t *testing.T) {
	arr1 := &object.Array{Elements: []object.Object{}}

	if arr1.Type() != object.ARRAY_OBJ {
		t.Errorf("array.Type() returns wrong type")
	}
}

func TestArrayHashKey(t *testing.T) {
	arr1 := &object.Array{Elements: []object.Object{}}
	arr2 := &object.Array{Elements: []object.Object{}}
	diff1 := &object.Array{Elements: []object.Object{&object.String{Value: "Hello World"}}}
	diff2 := &object.Array{Elements: []object.Object{&object.String{Value: "Hello Another World"}}}

	if arr1.HashKey() != arr2.HashKey() {
		t.Errorf("arrays with same content have different hash keys")
	}

	if diff1.HashKey() == diff2.HashKey() {
		t.Errorf("arrays with different content have same hash keys")
	}
}
