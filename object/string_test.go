package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("obj is not String. got=%T(%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s",
			result.Value, expected)
		return false
	}
	return true
}

func TestStringObject(t *testing.T) {
	tests := []inputTestCase{
		{`"a" == "a"`, true},
	}
	testInput(t, tests)
}

func TestStringObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`"test".count("e")`, 1},
		{`"test".count()`, "to few arguments: want=1, got=0"},
		{`"test".find("e")`, 1},
		{`"test".find()`, "to few arguments: want=1, got=0"},
		{`"test".size()`, 4},
		{`"test".plz_i()`, 0},
		{`"125".plz_i()`, 125},
		{`"test125".plz_i()`, 0},
		{`"0125".plz_i()`, 125},
		{`"1010".plz_i()`, 1010},
		{`"1010".plz_i(2)`, 10},
		{`"0x1022".plz_i()`, 530},
		{`"0x1022".plz_i(8)`, 530},
		{`"1022".plz_i(8)`, 530},
		{`"test".replace("e", "s")`, "tsst"},
		{`"test".replace()`, "to few arguments: want=2, got=0"},
		{`"test".replace("e")`, "to few arguments: want=2, got=1"},
		{`"test".reverse()`, "tset"},
		{`"test test1".split()`, `["test", "test1"]`},
		{`"test test1".split(",")`, `["test test1"]`},
		{`"test test1".split(",", "x")`, `to many arguments: want=1, got=2`},
		{`"test".split(1)`, `wrong argument type on position 0: got=INTEGER, want=STRING`},
		{`"test ".strip()`, "test"},
		{`" test ".strip()`, "test"},
		{`"test".strip()`, "test"},
		{`"test".upcase()`, "TEST"},
		{`"tESt".downcase()`, "test"},
		{`"test".type()`, "STRING"},
		{`"test".nope()`, "undefined method `.nope()` for STRING"},
		{`"test".methods().type()`, "ARRAY"},
		{`("test".methods().size() > 0).plz_s()`, "true"},
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
		{`a = "test"; a.upcase!(); a`, "TEST"},
		{`a = "tESt"; a.downcase!(); a`, "test"},
		{`a = "test"; a.reverse!(); a`, "tset"},
		{`a = " test "; a.strip!(); a`, "test"},
		{`("test".wat().lines().size() == "test".methods().size() + 1).plz_s()`, "true"},
		{"a = \"test\"; b = []; foreach char in a \n b.yoink(char) \nend; b.size()", 4},
		{`"test" * 2`, "testtest"},
		{`2 * "test"`, "testtest"},
		{`"test".to_json()`, `"test"`},
		{`{"test": HTTP.new()}.to_json()`, `Error while marshal value: json: error calling MarshalJSON for type *object.Hash: unable to serialize value: "test"`},
		{`"te\nst".size()`, 6},
		{`"te\"st".size()`, 5},
	}

	testInput(t, tests)
}

func TestStringHashKey(t *testing.T) {
	hello1 := object.NewString("Hello World")
	hello2 := object.NewString("Hello World")
	diff1 := object.NewString("My name is johnny")
	diff2 := object.NewString("My name is johnny")

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with different content have different hash keys")
	}
}
