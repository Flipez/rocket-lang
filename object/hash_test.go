package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestHashObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`{1: 1, "a": 2}.keys()`, "[1, a]"},
		{`[].nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}

func TestHashInspect(t *testing.T) {
	hash1 := &object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}

	if hash1.Inspect() != "{}" {
		t.Errorf("hash.Inspect() returns wrong string")
	}
}

func TestHashType(t *testing.T) {
	hash1 := &object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}

	if hash1.Type() != object.HASH_OBJ {
		t.Errorf("hash.Type() returns wrong type")
	}
}
