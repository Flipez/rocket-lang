package object_test

import (
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`let a = open("../main.go"); a.close()`, true},
		{`let a = open("../main.go"); a.read()`, "package main\n"},
		{`let a = open("../main.go"); a.lines().size()`, 45},
		{`(open("").wat().lines().size() == open("").methods().size() + 1).plz_s()`, "true"},
		{`open("").type()`, "FILE"},
	}

	testInput(t, tests)
}
