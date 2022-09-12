package object_test

import (
	"os"
	"testing"
)

func TestFileObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`open("../fixtures/module.rl", "r", "0644").close()`, true},
		{`a = open("../fixtures/module.rl", "r", "0644"); a.close(); a.position()`, -1},
		{`open("../fixtures/module.rl", "r", "0644").content().size()`, 51},
		{`open("../fixtures/module.rl", "r", "0644").content(1)`, "to many arguments: got=1, want=0"},
		{`open("../fixtures/module.rl", "r", "0644").read()`, "to few arguments: got=0, want=1"},
		{`open("../fixtures/module.rl", "r", "0644").read(1)`, "a"},
		{`open("../fixtures/module.rl", "r", "0644").position()`, 0},
		{`a = open("../fixtures/module.rl", "r", "0644"); a.read(1); a.content(); a.position()`, 0},
		{`a = open("../fixtures/module.rl", "r", "0644"); a.read(1); a.position()`, 1},
		{`a = open("../fixtures/module.rl", "r", "0644"); a.read(1); a.content().size()`, 51},
		{`open("../fixtures/module.rl", "r", "0644").lines().size()`, 7},
		{`a = open("../fixtures/module.rl", "r", "0644"); a.read(25); a.lines().size()`, 7},
		{`open("../fixtures/nope", "r", "0644")`, "open ../fixtures/nope: no such file or directory"},
		{`open("../fixtures/nope", "r", "0644").content()`, "undefined method `.content()` for ERROR"},
		{`a = open("../fixtures/nope", "rw", "0644"); a.content()`, ""},
		{`f = open("../fixtures/module.rl", "r", "0644"); (f.wat().lines().size() == f.methods().size() + 1).plz_s()`, "true"},
		{`open("").type()`, "ERROR"},
		{`open("../fixtures/module.rl", "r", "0644").type()`, "FILE"},
	}

	testInput(t, tests)

	os.Remove("../fixtures/nope")
}
