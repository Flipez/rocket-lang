package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNewArrayWithObjects(t *testing.T) {
	arr := object.NewArrayWithObjects(object.NewString("a"))
	if v := arr.Type(); v != object.ARRAY_OBJ {
		t.Errorf("array.Type() return wrong type: %s", v)
	}

	if v := arr.Elements[0].Type(); v != object.STRING_OBJ {
		t.Errorf("first array element should be a string object")
	}
}

func TestArrayObject(t *testing.T) {
	tests := []inputTestCase{
		{"[1] == [1]", true},
		{"[1] == [true]", false},
		{"[1] == [true, 1]", false},
	}

	testInput(t, tests)
}

func TestArrayObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3][0]`, 1},
		{`[1,2,3].size()`, 3},
		{`[1,2,3].pop()`, 3},
		{`[1,2,3].type()`, "ARRAY"},
		{`a = []; a.push(1); a`, "[1]"},
		{`[].nope()`, "test:1:3: undefined method `.nope()` for ARRAY"},
		{`([].wat().lines().size() == [].methods().size() + 1).to_s()`, "true"},
		{"a = [\"a\", \"b\"]; b = []; foreach i, item in a \n b.push(item) \nend; b.size()", 2},
		{`[1,2,3].index(4)`, -1},
		{`[1,2,3].index(3)`, 2},
		{`[1,2,3].index(true)`, -1},
		{`[1,2,3].index()`, "to few arguments: got=0, want=1"},
		{"a = []; b = []; foreach i in a \n b.push(a[i]) \nend; a.size()==b.size()", true},
		{`[1,1,2].uniq().size()`, 2},
		{`[true,true,2].uniq().size()`, 2},
		{`["test","test",2].uniq().size()`, 2},
		{`["12".reverse!()].uniq()`, "failed because element NIL is not hashable"},
		{"[].first()", "NIL"},
		{"[1,2,3].first()", 1},
		{"[].last()", "NIL"},
		{"[1,2,3].last()", 3},
		{"[1,2,3].to_json()", "[1,2,3]"},
		{`["test",true,3].to_json()`, `["test",true,3]`},
		{`[3.4, 3.1, 2.0].sort()`, `[2.0, 3.1, 3.4]`},
		{`[3, 1, 4].sort()`, `[1, 3, 4]`},
		{`["Gopher", "Go", "Alpha"].sort()`, `["Alpha", "Go", "Gopher"]`},
		{`["Gopher", 1, "Alpha"].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[1, "Go", 1].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[2.0, "Go", 2.0].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[true, "Go", true].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[].sort()`, `[]`},
		{`["a", "b", 1, 2].reverse()`, `[2, 1, "b", "a"]`},
		{`[1,2,3].include?(4)`, false},
		{`[1,2,3].include?(3)`, true},
		{`[1,2,3].include?(true)`, false},
		{`[1,2,3].include?()`, "to few arguments: got=0, want=1"},
		{`[1,2,3,4,5,6,7,8,9].slices(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8, 9]]`},
		{`[1,2,3,4,5,6,7,8].slices(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8]]`},
		{`[1,2].slices(3)`, `[[1, 2]]`},
		{`[1,2].slices(0)`, `invalid slice size, needs to be > 0`},
		{"[1,2,3,{}].join()", "Found non stringable element HASH on index 3"},
		{"[1,2,3].join()", "123"},
		{"[1,2,3].join('-')", "1-2-3"},
		{"['1',2, 2.5,{}].sum()", "Found non number element HASH on index 3"},
		{"['1', 2, 2.5].sum()", 5},
	}

	testInput(t, tests)
}

func TestArrayInspect(t *testing.T) {
	arr1 := object.NewArray(nil)

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

func TestArrayAdd(t *testing.T) {
	array := object.NewArray(nil)
	array.Add("a")
	if len(array.Elements) != 1 || array.Elements[0].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 0")
	}

	array.Add(object.NewString("b"))
	if len(array.Elements) != 2 || array.Elements[1].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 1")
	}

	array.Add(object.NewString("c"), 1)
	if len(array.Elements) != 4 || array.Elements[2].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 2")
	}
	if len(array.Elements) != 4 || array.Elements[3].Type() != object.INTEGER_OBJ {
		t.Errorf("expected array to have an integer value on index 3")
	}
}
