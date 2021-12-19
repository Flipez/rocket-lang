package object_test

import (
	"github.com/flipez/rocket-lang/object"
	"testing"
)

func TestStringObjectMethods(t *testing.T) {
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
		{`"test".upcase()`, "TEST"},
		{`"tESt".downcase()`, "test"},
		{`"test".type()`, "string"},
		{`"test".nope()`, "Failed to invoke method: nope"},
		{`"test".methods()`, `[count, find, size, methods, replace, reverse, split, toupper, tolower, type]`},
		{`"string".find("s")`, 0},
		{`"string".find("string")`, 0},
		{`"string".find("g")`, 5},
		{`"string".find("tr")`, 1},
		{`"string".find("ng")`, 4},
		{`"string".find("x")`, -1},
		{`"ab".reverse()`, "ba"},
		{`"abc".upcase()`, "ABC"},
		{`"a b c".upcase()`, "A B C"},
		{`"a%b!c".upcase()`, "A%B!C"},
		{`"ABC".downcase()`, "abc"},
		{`"A B C".downcase()`, "a b c"},
		{`"A%B!C".downcase()`, "a%b!c"},
		{`"     ".strip()`, ""},
		{`"
                       string".strip()`, "string"},
		{`"abc".replace("a", "A")`, "Abc"},
		{`"These are the days of summer".count("e")`, 5},
	}

	testInput(t, tests)
}

func TestStringHashKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "My name is johnny"}
	diff2 := &object.String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with different content have different hash keys")
	}
}
