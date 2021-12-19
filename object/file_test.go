package object_test

import (
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`let a = open("../main.go"); a.close()`, true},
		{`let a = open("../main.go"); a.read()`, "package main\n"},
		{`let a = open("../main.go"); a.lines().size()`, 45},
	}

	testInput(t, tests)
}
