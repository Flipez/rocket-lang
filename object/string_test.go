package object_test

import (
	"testing"
)

func TestObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`"test".count("e")`, 1},
		{`"test".count()`, "Missing argument to count()!"},
		{`"test".find("e")`, 1},
		{`"test".find()`, "Missing argument to find()!"},
		{`"test".size()`, 4},
		{`"test".plz_i()`, 0},
		{`"125".plz_i()`, 125},
		{`"test125".plz_i()`, 0},
		{`"0125".plz_i()`, 125},
		{`"test".replace("e", "s")`, "tsst"},
		{`"test".replace()`, "Missing arguments to replace()!"},
		{`"test".replace("e")`, "Missing arguments to replace()!"},
		{`"test".reverse()`, "tset"},
		{`"test test1".split()`, `[test, test1]`},
		{`"test test1".split(",")`, `[test test1]`},
		{`"test test1".split(",", "x")`, `[test test1]`},
		{`"test ".strip()`, "test"},
		{`" test ".strip()`, "test"},
		{`"test".strip()`, "test"},
		{`"test".toupper()`, "TEST"},
		{`"tESt".tolower()`, "test"},
		{`"test".type()`, "string"},
		{`"test".nope()`, "Failed to invoke method: nope"},
		{`"test".methods()`, `[count, find, size, methods, replace, reverse, split, toupper, tolower, type]`},
	}

	testInput(t, tests)
}
