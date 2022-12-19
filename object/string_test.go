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
		{`"test".count()`, "to few arguments: got=0, want=1"},
		{`"test".find("e")`, 1},
		{`"test".find()`, "to few arguments: got=0, want=1"},
		{`"test".size()`, 4},
		{`"test".to_i()`, 0},
		{`"125".to_i()`, 125},
		{`"test125".to_i()`, 0},
		{`"0125".to_i()`, 125},
		{`"1010".to_i()`, 1010},
		{`"1010".to_i(2)`, 10},
		{`"0x1022".to_i()`, 530},
		{`"0x1022".to_i(8)`, 530},
		{`"1022".to_i(8)`, 530},
		{`"1022".to_f()`, 1022.0},
		{`"1022".to_s()`, "1022"},
		{`"test".replace("e", "s")`, "tsst"},
		{`"test".replace()`, "to few arguments: got=0, want=2"},
		{`"test".replace("e")`, "to few arguments: got=1, want=2"},
		{`"test".reverse()`, "tset"},
		{`"test test1".split()`, `["test", "test1"]`},
		{`"test test1".split(",")`, `["test test1"]`},
		{`"test test1".split(",", "x")`, `to many arguments: got=2, want=1`},
		{`"test".split(1)`, `wrong argument type on position 1: got=INTEGER, want=STRING`},
		{`"test ".strip()`, "test"},
		{`" test ".strip()`, "test"},
		{`"test".strip()`, "test"},
		{`"test".upcase()`, "TEST"},
		{`"tESt".downcase()`, "test"},
		{`"test".type()`, "STRING"},
		{`"test".nope()`, "test:1:7: undefined method `.nope()` for STRING"},
		{`"test".methods().type()`, "ARRAY"},
		{`("test".methods().size() > 0).to_s()`, "true"},
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
		{`("test".wat().lines().size() == "test".methods().size() + 1).to_s()`, "true"},
		{"a = \"test\"; b = []; foreach char in a \n b.push(char) \nend; b.size()", 4},
		{`"test" * 2`, "testtest"},
		{`2 * "test"`, "testtest"},
		{`"test".to_json()`, `"test"`},
		{`{"test": HTTP.new()}.to_json()`, `Error while marshal value: json: error calling MarshalJSON for type *object.Hash: unable to serialize value: "test"`},
		{`"te\nst".size()`, 6},
		{`"te\"st".size()`, 5},
		{`'te\"st'.size()`, 6},
		{`"te\"st" == 'te"st'`, true},
		{`"test%d".format(1)`, "test1"},
		{`"%dtest%d".format(1,2)`, "1test2"},
		{`"test%5d".format(1)`, "test    1"},
		{`"test%f".format(1.3)`, "test1.300000"},
		{`"test%1.1f".format(1.3)`, "test1.3"},
		{`"test%s".format("test")`, "testtest"},
		{`"test%t".format(true)`, "testtrue"},
		{`"".ascii()`, -1},
		{`"a".ascii()`, 97},
		{`"abc".ascii()`, `[97, 98, 99]`},
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
