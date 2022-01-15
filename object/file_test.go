package object_test

import (
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`a = open("../main.go"); a.close()`, true},
		{`a = open("../main.go"); a.close(); a.position()`, -1},
		{`a = open("../fixtures/module.rl"); a.content().size()`, 49},
		{`a = open("../fixtures/module.rl"); a.content(1)`, "To many arguments: want=0, got=1"},
		{`a = open("../fixtures/module.rl"); a.read()`, "To few arguments: want=1, got=0"},
		{`a = open("../fixtures/module.rl"); a.read(1)`, "a"},
		{`a = open("../fixtures/module.rl"); a.position()`, 0},
		{`a = open("../fixtures/module.rl"); a.read(1); a.content(); a.position()`, 0},
		{`a = open("../fixtures/module.rl"); a.read(1); a.position()`, 1},
		{`a = open("../fixtures/module.rl"); a.read(1); a.content().size()`, 49},
		{`a = open("../fixtures/module.rl"); a.lines().size()`, 7},
		{`a = open("../fixtures/module.rl"); a.read(25); a.lines().size()`, 7},
		{`(open("").wat().lines().size() == open("").methods().size() + 1).plz_s()`, "true"},
		{`open("").type()`, "FILE"},
	}

	testInput(t, tests)
}
