package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestHashObject(t *testing.T) {
	tests := []inputTestCase{
		{`{"a": 1} == {"a": 1}`, true},
		{`{"a": 1} == {"a": 1, "b": 2}`, false},
		{`{"a": 1} == {"b": 1}`, false},
		{`{"a": 1} == {"a": "c"}`, false},
		{`{{1: true}: "a"}.keys()`, `[{1: true}]`},
	}

	testInput(t, tests)
}

func TestHashObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`{"a": 2}.keys()`, `["a"]`},
		{`{}.nope()`, "test:1:3: undefined method `.nope()` for HASH"},
		{`({}.wat().lines().size() == {}.methods().size() + 1).plz_s()`, "true"},
		{`{}.type()`, "HASH"},
		{"a = {\"a\": \"b\", \"b\":\"a\"};b = []; foreach key, value in a \n b.yoink(key) \nend; b.size()", 2},
		{`{"a": 1, "b": 2}["a"]`, 1},
		{`{"a": 1, "b": 2}.keys().size()`, 2},
		{`{"a": 1, "b": 2}.values().size()`, 2},
		{`{"a": "b"}.to_json()`, `{"a":"b"}`},
		{`{1: "b"}.to_json()`, `{"1":"b"}`},
		{`{true: "b"}.to_json()`, `{"true":"b"}`},
		{`{"a": 1, 1: "b"}.include?("a")`, true},
		{`{"a": 1, 1: "b"}.include?(1)`, true},
		{`{"a": 1, 1: "b"}.include?("c")`, false},
		{`{"a": 1, 1: "b"}.include?(nil)`, `wrong argument type on position 1: got=NIL, want=BOOLEAN|STRING|INTEGER|FLOAT|ARRAY|HASH`},
		{`{"a": 1, 1: "b"}.include?()`, `to few arguments: got=0, want=1`},
		{`{"a": 1, "b": 2}.get("a", 10)`, 1},
		{`{"a": 1, "b": 2}.get("c", 10)`, 10},
	}

	testInput(t, tests)
}

func TestHashInspect(t *testing.T) {
	tests := []inputTestCase{
		{"{}", "{}"},
		{`{"a": 1}`, `{"a": 1}`},
		{`{true: "a"}`, `{true: "a"}`},
	}

	for _, tt := range tests {
		hash := testEval(tt.input).(*object.Hash)
		hashInspect := hash.Inspect()

		if hash.Inspect() != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, hashInspect)
		}
	}
}

func TestHashType(t *testing.T) {
	hash1 := object.NewHash(nil)

	if hash1.Type() != object.HASH_OBJ {
		t.Errorf("hash.Type() returns wrong type")
	}
}

func TestHashSet(t *testing.T) {
	hash := object.NewHash(nil)

	hash.Set("a", 1)
	obj, ok := hash.Get("a")
	if !ok || obj == nil {
		t.Errorf("expected to get value")
	} else {
		if obj.Type() != object.INTEGER_OBJ {
			t.Errorf("unexpected type of value, got=%s want=%s", obj.Type(), object.INTEGER_OBJ)
		}
	}

	hash.Set(object.NewString("a"), 2)
	obj, ok = hash.Get("a")
	if !ok || obj == nil {
		t.Errorf("expected to get value")
	} else {
		if obj.Type() != object.INTEGER_OBJ {
			t.Errorf("unexpected type of value, got=%s want=%s", obj.Type(), object.INTEGER_OBJ)
		}
	}

	hash.Set(object.NewString("b"), object.NewInteger(3))
	obj, ok = hash.Get("b")
	if !ok || obj == nil {
		t.Errorf("expected to get value")
	} else {
		if obj.Type() != object.INTEGER_OBJ {
			t.Errorf("unexpected type of value, got=%s want=%s", obj.Type(), object.INTEGER_OBJ)
		}
	}
}
