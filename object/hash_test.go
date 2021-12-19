package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestHashObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`{1: 1, "a": 2}.keys()`, "[1, a]"},
		{`{}.nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}

func TestHashInspect(t *testing.T) {
	hash1 := testEval(`{}`).(*object.Hash)
	hash2 := testEval(`{"a": 1, 2: 3}`).(*object.Hash)

	if hash1.Inspect() != "{}" {
		t.Errorf("wrong string. expected=%q, got=%q", "{}", hash1.Inspect())
	}

	if hash2.Inspect() != "{a: 1, 2: 3}" {
		t.Errorf("wrong string. expected=%q, got=%q", "{a: 1, 2: 3}", hash2.Inspect())
	}
}

func TestHashType(t *testing.T) {
	hash1 := &object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}

	if hash1.Type() != object.HASH_OBJ {
		t.Errorf("hash.Type() returns wrong type")
	}
}
