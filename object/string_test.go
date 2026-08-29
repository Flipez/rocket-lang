package object_test

import (
	"strings"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	t.Helper()
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
		{`"test".to_i()`, nil},
		{`"125".to_i()`, 125},
		{`"test125".to_i()`, nil},
		{`"0125".to_i()`, 85},
		{`"1010".to_i()`, 1010},
		{`"-1010".to_i()`, -1010},
		{`"0x1022".to_i()`, 4130},
		// A bare "0" is decimal zero, not a leading-zero octal prefix with
		// nothing after it. Getting this wrong gave it base 8, so adding it to
		// any ordinary integer failed with "unequal base".
		{`"0".to_i()`, 0},
		{`"0".to_i().base()`, 10},
		{`"0".to_i() + 1`, 1},
		{`"-0".to_i() + 1`, 1},
		// A leading zero followed by a non-octal digit is a zero-padded
		// decimal, not a malformed octal literal.
		{`"08".to_i()`, 8},
		{`"09".to_i().base()`, 10},
		// Base prefixes are case insensitive; uppercase ones used to fall
		// through to the legacy-octal branch and be tagged base 8.
		{`"0X1022".to_i()`, 4130},
		{`"0X1022".to_i().base()`, 16},
		{`"0B101".to_i()`, 5},
		{`"0B101".to_i().base()`, 2},
		{`"0o17".to_i()`, 15},
		{`"0O17".to_i().base()`, 8},
		// Legacy leading-zero octal keeps working.
		{`"0125".to_i().base()`, 8},
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
		{"a = \"test\"; b = []; foreach char in a \n b.push(char) \nend; b.size()", 4},
		{`"test" * 2`, "testtest"},
		{`2 * "test"`, "testtest"},
		{`"test".to_json()`, `"test"`},
		{`{"test": HTTP.new()}.to_json()`, `Error while marshal value: json: error calling MarshalJSON for type *object.Hash: unable to serialize value: "test"`},
		{`"te\nst".size()`, 5},
		{`"te\"st".size()`, 5},
		{`'te\"st'.size()`, 6},
		{`"te\"st" == 'te"st'`, true},
		{`"te\tst".size()`, 5},
		{`"te\rst".size()`, 5},
		{`"te\\st".size()`, 5},
		{`"a\tb\tc".split("\t")`, `["a", "b", "c"]`},
		{`"line1\r\nline2".split("\r\n")`, `["line1", "line2"]`},
		{`"path\\to\\file".count("\\")`, 2},
		{`"test\xabc".size()`, 9},
		{`"test\xabc"`, "test\\xabc"},
		{`"test\x"`, "test\\x"},
		{`"test%d".format(1)`, "test1"},
		{`"%dtest%d".format(1,2)`, "1test2"},
		{`"test%5d".format(1)`, "test    1"},
		{`"test%f".format(1.3)`, "test1.300000"},
		{`"test%1.1f".format(1.3)`, "test1.3"},
		{`"test%s".format("test")`, "testtest"},
		{`"test%t".format(true)`, "testtrue"},
		// ascii always returns an ARRAY. It used to return a bare INTEGER for
		// a one-character string and -1 for an empty one, so callers had to
		// branch on the length of their own input.
		{`"".ascii()`, `[]`},
		{`"a".ascii()`, `[97]`},
		{`"abc".ascii()`, `[97, 98, 99]`},
		{`"abc".ascii().size() == "abc".size()`, true},
		// Sizing the result by byte length left the trailing slots of a
		// multi-character string nil, which segfaulted on Inspect.
		{`"\xc3\xa9".ascii().size() == "\xc3\xa9".size()`, true},
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

func TestStringInspect(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", `"test"`},
		{"test\nline2", `"test\nline2"`},
		{"tab\there", `"tab\there"`},
		{"carriage\rreturn", `"carriage\rreturn"`},
		{"back\\slash", `"back\\slash"`},
		{"quote\"test", `"quote\"test"`},
		{"multi\n\t\r\\\"test", `"multi\n\t\r\\\"test"`},
	}

	for _, tt := range tests {
		str := object.NewString(tt.input)
		result := str.Inspect()
		if result != tt.expected {
			t.Errorf("Inspect() for %q: got=%q, want=%q", tt.input, result, tt.expected)
		}
	}
}

// TestStringConversionFailureIsNil covers the change from silent zeros to nil:
// a failed conversion has to be distinguishable from a successful one that
// happens to produce 0.
func TestStringConversionFailureIsNil(t *testing.T) {
	tests := []inputTestCase{
		{`"0".to_i()`, 0},
		{`"abc".to_i()`, nil},
		{`"".to_i()`, nil},
		{`"12abc".to_i()`, nil},
		{`"0.0".to_f()`, 0.0},
		{`"abc".to_f()`, nil},
		{`"".to_f()`, nil},
		// to_s never fails, so nil keeps rendering as the empty string.
		{`nil.to_s()`, ""},
	}

	testInput(t, tests)
}

// TestStringRubyMethods covers the methods added to close the gap with Ruby's
// String. Each expectation is taken from the example lines in
// https://ruby-doc.org/3.4.1/String.html so the behaviour matches Ruby rather
// than whatever Go's strings package happens to do.
func TestStringRubyMethods(t *testing.T) {
	tests := []inputTestCase{
		// capitalize downcases the rest, so a capital in the middle is lost.
		{`"hello World!".capitalize()`, "Hello world!"},
		{`"".capitalize()`, ""},
		{`"hELLO".swapcase()`, "Hello"},
		{`"Hello World".swapcase()`, "hELLO wORLD"},

		{`"  a  ".lstrip()`, "a  "},
		{`"  a  ".rstrip()`, "  a"},
		{`"a".lstrip()`, "a"},
		{`"\t\n a".lstrip()`, "a"},

		// chomp with no argument removes one CR, LF or CRLF...
		{`"abc\r".chomp()`, "abc"},
		{`"abc\n".chomp()`, "abc"},
		{`"abc\r\n".chomp()`, "abc"},
		// ...but "\n\r" loses only the "\r", as in Ruby.
		{`"abc\n\r".chomp()`, "abc\n"},
		{`"abc".chomp()`, "abc"},
		// With a separator it removes one trailing occurrence.
		{`"abcd".chomp("d")`, "abc"},
		{`"abcdd".chomp("d")`, "abcd"},
		// The empty separator is Ruby's "drop every trailing blank line".
		{`"abc\n\n\n".chomp("")`, "abc"},
		{`"abc\r\n\r\n\r\n".chomp("")`, "abc"},
		// It leaves bare CRs alone, which is the part that is easy to get wrong.
		{`"abc\r\r\r".chomp("")`, "abc\r\r\r"},

		{`"abcd".chop()`, "abc"},
		// A trailing CRLF goes as a unit, so chop never splits a line ending.
		{`"abc\r\n".chop()`, "abc"},
		{`"".chop()`, ""},
		{`"a".chop()`, ""},

		{`"".empty?()`, true},
		{`"a".empty?()`, false},
		{`"abc".include?("b")`, true},
		{`"abc".include?("z")`, false},
		{`"abc".include?("")`, true},
		{`"abc".start_with?("ab")`, true},
		{`"abc".start_with?("z")`, false},
		{`"abc".end_with?("bc")`, true},
		{`"abc".end_with?("z")`, false},
		// start_with? and end_with? take more than one candidate and are true
		// if any of them matches, as Ruby's do.
		{`"abc".start_with?("z", "y", "a")`, true},
		{`"abc".start_with?("z", "y")`, false},
		{`"abc".end_with?("z", "c")`, true},
		{`"abc".end_with?("z", "y")`, false},
		{`"abc".start_with?()`, "to few arguments: got=0, want=1"},
		{`"abc".include?()`, "to few arguments: got=0, want=1"},
		{`"abc".empty?("x")`, "to many arguments: got=1, want=0"},
	}
	testInput(t, tests)
}

// TestStringBangConvention checks the convention across every String pair: the
// plain method leaves the receiver alone, the ! method changes it and returns
// it. Ruby's String bangs return nil when they changed nothing, which is why
// "ABC".upcase!.reverse! raises there; RocketLang returns the receiver instead
// so chains hold.
func TestStringBangConvention(t *testing.T) {
	tests := []inputTestCase{
		// Plain methods are pure.
		{`s = "hello World"; s.capitalize(); s`, "hello World"},
		{`s = "hello World"; s.swapcase(); s`, "hello World"},
		{`s = "  a  "; s.lstrip(); s`, "  a  "},
		{`s = "  a  "; s.rstrip(); s`, "  a  "},
		{`s = "abc\n"; s.chomp(); s`, "abc\n"},
		{`s = "abcd"; s.chop(); s`, "abcd"},
		{`s = "a-b"; s.replace("-", "+"); s`, "a-b"},

		// ! methods change the receiver.
		{`s = "hello World"; s.capitalize!(); s`, "Hello world"},
		{`s = "hello World"; s.swapcase!(); s`, "HELLO wORLD"},
		{`s = "  a  "; s.lstrip!(); s`, "a  "},
		{`s = "  a  "; s.rstrip!(); s`, "  a"},
		{`s = "abc\n"; s.chomp!(); s`, "abc"},
		{`s = "abcd"; s.chop!(); s`, "abc"},
		{`s = "a-b"; s.replace!("-", "+"); s`, "a+b"},

		// ...and return it, so they chain even when a call changes nothing.
		{`"  hello World  ".strip!().capitalize!().swapcase!()`, "hELLO WORLD"},
		{`"ABC".upcase!().reverse!()`, "CBA"},
		{`"a-b\n".chomp!().replace!("-", "+").upcase!()`, "A+B"},
		{`"abc".chomp!().chomp!()`, "abc"},

		// Every ! method returns a STRING rather than NIL.
		{`"a".upcase!().type()`, "STRING"},
		{`"a".downcase!().type()`, "STRING"},
		{`"a".capitalize!().type()`, "STRING"},
		{`"a".swapcase!().type()`, "STRING"},
		{`"a".strip!().type()`, "STRING"},
		{`"a".lstrip!().type()`, "STRING"},
		{`"a".rstrip!().type()`, "STRING"},
		{`"a".reverse!().type()`, "STRING"},
		{`"a".chomp!().type()`, "STRING"},
		{`"a".chop!().type()`, "STRING"},
		{`"a".replace!("a", "b").type()`, "STRING"},

		// A ! method takes the same arguments as its plain counterpart.
		{`"abcd".chomp!("d")`, "abc"},
		{`"a-b".replace!("-")`, "to few arguments: got=1, want=2"},
	}
	testInput(t, tests)
}

// TestStringPairsAreComplete guards the convention itself: every String method
// that changes the receiver has a pure counterpart of the same name and the
// other way round. Adding one half of a pair and forgetting the other is the
// gap this test exists to catch.
func TestStringPairsAreComplete(t *testing.T) {
	listed, ok := testEval(`"a".methods()`).(*object.Array)
	if !ok {
		t.Fatal(`"a".methods() should return an array`)
	}

	names := make(map[string]bool, len(listed.Elements))
	for _, element := range listed.Elements {
		name, ok := element.(*object.String)
		if !ok {
			t.Fatalf("method name is not a string, got %s", element.Type())
		}
		names[name.Value] = true
	}

	// Every ! method must have a pure counterpart. The other direction does not
	// hold: size() and split() return something other than a string, so there
	// is nothing for a size!() to mean.
	for name := range names {
		if !strings.HasSuffix(name, "!") {
			continue
		}
		if !names[strings.TrimSuffix(name, "!")] {
			t.Errorf("%s has no pure counterpart", name)
		}
	}

	// The pairs that must exist. Listed explicitly so that dropping one is a
	// failure rather than a silently smaller loop.
	for _, name := range []string{
		"capitalize", "chomp", "chop", "downcase", "lstrip",
		"replace", "reverse", "rstrip", "strip", "swapcase", "upcase",
	} {
		if !names[name] {
			t.Errorf("expected String to have %s()", name)
		}
		if !names[name+"!"] {
			t.Errorf("expected String to have %s!()", name)
		}
	}

	// Predicates have nothing to change, so they must not have a ! form.
	for _, name := range []string{"empty?", "include?", "start_with?", "end_with?"} {
		if !names[name] {
			t.Errorf("expected String to have %s()", name)
		}
		if names[name+"!"] {
			t.Errorf("%s should not have a ! form", name)
		}
	}
}
