package object_test

import (
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`let a = open("../main.go"); a.close()`, true},
		{`let a = open("../fixtures/module.rl"); a.read()`, "a = 1\nA = 5\n\nSum = fn(a, b) {\n    return a + b\n}\n"},
		{`let a = open("../fixtures/module.rl"); a.lines().size()`, 7},
		{`(open("").wat().lines().size() == open("").methods().size() + 1).plz_s()`, "true"},
		{`open("").type()`, "FILE"},
	}

	testInput(t, tests)
}
