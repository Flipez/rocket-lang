package object_test

import (
	"os"
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`open("../fixtures/module.rl").close()`, true},
		{`a = open("../fixtures/module.rl"); a.close(); a.position()`, -1},
		{`open("../fixtures/module.rl").content().size()`, 51},
		{`open("../fixtures/module.rl").content(1)`, "To many arguments: want=0, got=1"},
		{`open("../fixtures/module.rl").read()`, "To few arguments: want=1, got=0"},
		{`open("../fixtures/module.rl").read(1)`, "a"},
		{`open("../fixtures/module.rl").position()`, 0},
		{`a = open("../fixtures/module.rl"); a.read(1); a.content(); a.position()`, 0},
		{`a = open("../fixtures/module.rl"); a.read(1); a.position()`, 1},
		{`a = open("../fixtures/module.rl"); a.read(1); a.content().size()`, 51},
		{`open("../fixtures/module.rl").lines().size()`, 7},
		{`a = open("../fixtures/module.rl"); a.read(25); a.lines().size()`, 7},
		{`open("../fixtures/nope")`, "open ../fixtures/nope: no such file or directory"},
		{`open("../fixtures/nope").content()`, "Failed to invoke method: content"},
		{`a = open("../fixtures/nope", "rw"); a.content()`, ""},
		{`(open("../fixtures/module.rl").wat().lines().size() == open("../fixtures/module.rl").methods().size() + 1).plz_s()`, "true"},
		{`open("").type()`, "ERROR"},
		{`open("../fixtures/module.rl").type()`, "FILE"},
	}

	testInput(t, tests)

	os.Remove("../fixtures/nope")
}
