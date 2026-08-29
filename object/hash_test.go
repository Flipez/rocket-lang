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
		{`{}.type()`, "HASH"},
		{"a = {\"a\": \"b\", \"b\":\"a\"};b = []; foreach key, value in a \n b.push(key) \nend; b.size()", 2},
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
		// to_s used to be "" because Hash was not Stringable. A single entry,
		// because the order of a bigger hash is not defined.
		{`{"a": 1}.to_s()`, `{"a": 1}`},
		{`{"a": 1, "b": 2}.to_i()`, nil},
		{`{"a": 1, "b": 2}.to_f()`, nil},
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

// TestHashRubyMethods covers the methods added to close the gap with Ruby's
// Hash. Only the order-independent ones: keys(), values() and anything built on
// them still come back in map order, so to_a is left out until a hash keeps
// insertion order.
func TestHashRubyMethods(t *testing.T) {
	tests := []inputTestCase{
		{`{"a": 1}.size()`, 1},
		{`{}.size()`, 0},
		{`{}.empty?()`, true},
		{`{"a": 1}.empty?()`, false},

		// fetch errors on a missing key unless given a fallback. That is the
		// difference from get(), which always requires one.
		{`{"a": 1}.fetch("a")`, 1},
		{`{"a": 1}.fetch("z", 0)`, 0},
		{`{"a": 1}.fetch("z")`, `key not found: "z"`},
		{`{"a": 1}.fetch(nil)`, "unusable as hash key: NIL"},

		// delete reports the value that went, or nil when nothing did.
		{`h = {"a": 1}; h.delete("a")`, 1},
		{`h = {"a": 1}; h.delete("a"); h.size()`, 0},
		{`h = {"a": 1}; h.delete("z")`, nil},
		{`h = {"a": 1}; h.delete("z"); h.size()`, 1},

		{`h = {"a": 1}; h.clear().size()`, 0},
		{`h = {"a": 1}; h.clear(); h.size()`, 0},

		{`{"a": 1}.merge({"b": 2}).size()`, 2},
		// The argument wins a clash.
		{`{"a": 1}.merge({"a": 9}).get("a", 0)`, 9},
		{`{"a": 1}.invert().get(1, "missing")`, "a"},
		{`{"a": nil}.invert()`, "unusable as hash key: NIL"},

		{`{"a": 1, "b": nil}.compact().size()`, 1},
		{`{"a": nil}.compact().get("a", "gone")`, "gone"},

		// to_json renders a nil value as null now that Nil is serializable.
		{`{"a": nil}.to_json()`, `{"a":null}`},
	}
	testInput(t, tests)
}

// TestHashBangPairsAreComplete checks the convention on Hash's two pairs.
func TestHashBangPairsAreComplete(t *testing.T) {
	tests := []inputTestCase{
		// Pure.
		{`h = {"a": 1}; h.merge({"b": 2}); h.size()`, 1},
		{`h = {"a": nil}; h.compact(); h.size()`, 1},

		// In place.
		{`h = {"a": 1}; h.merge!({"b": 2}); h.size()`, 2},
		{`h = {"a": nil}; h.compact!(); h.size()`, 0},

		// ...and they chain, because each returns the hash.
		{`{"a": 1}.merge!({"b": nil}).compact!().size()`, 1},
		{`{"a": 1}.merge!({"b": 2}).type()`, "HASH"},
	}
	testInput(t, tests)
}
