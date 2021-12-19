package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestHashObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`{"a": 2}.keys()`, "[a]"},
		{`{}.nope()`, "Failed to invoke method: nope"},
		{`({}.wat().lines().size() == {}.methods().size() + 1).plz_s()`, "true"},
		{`{}.type()`, "HASH"},
	}

	testInput(t, tests)
}

func TestHashInspect(t *testing.T) {
	tests := []inputTestCase{
		{"{}", "{}"},
		{`{"a": 1}`, "{a: 1}"},
		{`{true: "a"}`, "{true: a}"},
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
	hash1 := &object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}

	if hash1.Type() != object.HASH_OBJ {
		t.Errorf("hash.Type() returns wrong type")
	}
}
