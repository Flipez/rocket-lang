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
		{`"test".count()`, "too few arguments: got=0, want=1"},
		{`"test".index_of("e")`, 1},
		{`"test".index_of()`, "too few arguments: got=0, want=1"},
		{`"test".size()`, 4},
		{`"test".to_integer()`, nil},
		{`"125".to_integer()`, 125},
		{`"test125".to_integer()`, nil},
		{`"0125".to_integer()`, 85},
		{`"1010".to_integer()`, 1010},
		{`"-1010".to_integer()`, -1010},
		{`"0x1022".to_integer()`, 4130},
		// A bare "0" is decimal zero, not a leading-zero octal prefix with
		// nothing after it. Getting this wrong gave it base 8, so adding it to
		// any ordinary integer failed with "unequal base".
		{`"0".to_integer()`, 0},
		// A method's base is now only observable through the prefix its
		// to_string() prints, since Integer#base was dropped in favour of
		// to_base(n).
		{`"0".to_integer().to_string()`, "0"},
		{`"0".to_integer() + 1`, 1},
		{`"-0".to_integer() + 1`, 1},
		// A leading zero followed by a non-octal digit is a zero-padded
		// decimal, not a malformed octal literal.
		{`"08".to_integer()`, 8},
		{`"09".to_integer().to_string()`, "9"},
		// Base prefixes are case insensitive; uppercase ones used to fall
		// through to the legacy-octal branch and be tagged base 8.
		{`"0X1022".to_integer()`, 4130},
		{`"0X1022".to_integer().to_string()`, "0x1022"},
		{`"0B101".to_integer()`, 5},
		{`"0B101".to_integer().to_string()`, "0b101"},
		{`"0o17".to_integer()`, 15},
		{`"0O17".to_integer().to_string()`, "0o17"},
		// Legacy leading-zero octal keeps working.
		{`"0125".to_integer().to_string()`, "0o125"},
		{`"1022".to_float()`, 1022.0},
		{`"1022".to_string()`, "1022"},
		{`"test".replace("e", "s")`, "tsst"},
		{`"test".replace()`, "too few arguments: got=0, want=2"},
		{`"test".replace("e")`, "too few arguments: got=1, want=2"},
		{`"test".reverse()`, "tset"},
		{`"test test1".split()`, `["test", "test1"]`},
		{`"test test1".split(",")`, `["test test1"]`},
		{`"test test1".split(",", "x")`, `too many arguments: got=2, want=1`},
		{`"test".split(1)`, `wrong argument type on position 1: got=INTEGER, want=STRING`},
		{`"test ".trim()`, "test"},
		{`" test ".trim()`, "test"},
		{`"test".trim()`, "test"},
		{`"test".uppercase()`, "TEST"},
		{`"tESt".lowercase()`, "test"},
		{`"test".type()`, "STRING"},
		{`"test".nope()`, "test:1:7: undefined method `.nope()` for STRING"},
		{`"test".methods().type()`, "ARRAY"},
		{`("test".methods().size() > 0).to_string()`, "true"},
		{`"string".index_of("s")`, 0},
		{`"string".index_of("string")`, 0},
		{`"string".index_of("g")`, 5},
		{`"string".index_of("tr")`, 1},
		{`"string".index_of("ng")`, 4},
		{`"string".index_of("x")`, -1},
		{`"ab".reverse()`, "ba"},
		{`"abc".uppercase()`, "ABC"},
		{`"a b c".uppercase()`, "A B C"},
		{`"a%b!c".uppercase()`, "A%B!C"},
		{`"ABC".lowercase()`, "abc"},
		{`"A B C".lowercase()`, "a b c"},
		{`"A%B!C".lowercase()`, "a%b!c"},
		{`"     ".trim()`, ""},
		{`"
                       string".trim()`, "string"},
		{`"abc".replace("a", "A")`, "Abc"},
		{`"These are the days of summer".count("e")`, 5},
		{`a = "test"; a.uppercase!(); a`, "TEST"},
		{`a = "tESt"; a.lowercase!(); a`, "test"},
		{`a = "test"; a.reverse!(); a`, "tset"},
		{`a = " test "; a.trim!(); a`, "test"},
		{"a = \"test\"; b = []; foreach char in a \n b.append!(char) \nend; b.size()", 4},
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
		// codepoints always returns an ARRAY. It used to return a bare INTEGER
		// for a one-character string and -1 for an empty one, so callers had
		// to branch on the length of their own input.
		{`"".codepoints()`, `[]`},
		{`"a".codepoints()`, `[97]`},
		{`"abc".codepoints()`, `[97, 98, 99]`},
		{`"abc".codepoints().size() == "abc".size()`, true},
		// Sizing the result by byte length left the trailing slots of a
		// multi-character string nil, which segfaulted on Inspect.
		{`"\xc3\xa9".codepoints().size() == "\xc3\xa9".size()`, true},
	}

	testInput(t, tests)
}

func TestStringIndexOf(t *testing.T) {
	tests := []inputTestCase{
		{`"hello".index_of("l")`, 2},
		{`"hello".last_index_of("l")`, 3},
		{`"hello".index_of("z")`, -1},
		{`"hello".last_index_of("z")`, -1},
		// index_of/last_index_of must count runes like every other String
		// method (size, s[i], slicing, reverse), not bytes. "語" starts at
		// byte 6 in "日本語" but rune (character) index 2.
		{`"日本語".index_of("語")`, 2},
		{`"日本語".index_of("本")`, 1},
		{`"日本語".index_of("missing")`, -1},
		{`"日本語本".last_index_of("本")`, 3},
		// last_index_of's not-found case, on a multi-byte string.
		{`"日本語".last_index_of("missing")`, -1},
		// The property this fix exists to restore: find where a substring
		// is, then look there, and get it back -- on a multi-byte string.
		{`s = "日本語"; s[s.index_of("語")]`, "語"},
		{`s = "日本語本"; s[s.last_index_of("本")]`, "本"},
	}

	testInput(t, tests)
}

// TestStringIndexOfInvalidUTF8Needle covers a needle that cannot come from a
// RocketLang string literal: raw, non-UTF-8 bytes, the shape of whatever
// IO.read_line hands back from stdin with no validation. strings.Index can
// still match such a needle at a byte offset that is not a rune boundary in
// the haystack's own (valid) decoding, which used to make index_of report a
// position that pointed at the wrong character instead of "not present".
// Searching in rune space closes that: an invalid needle decodes to the
// replacement rune U+FFFD, which cannot equal a validly-decoded rune in the
// haystack, so it is correctly -1.
func TestStringIndexOfInvalidUTF8Needle(t *testing.T) {
	env := *object.NewEnvironment()

	// "X" + U+0080 (valid 2-byte UTF-8: 0xC2 0x80) + "Y" -- 3 runes, 4 bytes.
	haystack := object.NewString("X" + string(rune(128)) + "Y")
	// A lone 0x80 byte: a continuation byte with no leading byte, invalid on
	// its own. strings.Index would find it inside haystack's 0xC2 0x80 pair,
	// at byte offset 2 -- not a rune boundary.
	invalidNeedle := object.NewString("\x80")

	result := haystack.InvokeMethod("index_of", env, invalidNeedle)
	testIntegerObject(t, result, -1)

	result = haystack.InvokeMethod("last_index_of", env, invalidNeedle)
	testIntegerObject(t, result, -1)
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
		{`"0".to_integer()`, 0},
		{`"abc".to_integer()`, nil},
		{`"".to_integer()`, nil},
		{`"12abc".to_integer()`, nil},
		{`"0.0".to_float()`, 0.0},
		{`"abc".to_float()`, nil},
		{`"".to_float()`, nil},
		// to_string never fails, so nil keeps rendering as the empty string.
		{`nil.to_string()`, ""},
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
		{`"hELLO".swap_case()`, "Hello"},
		{`"Hello World".swap_case()`, "hELLO wORLD"},

		{`"  a  ".trim_start()`, "a  "},
		{`"  a  ".trim_end()`, "  a"},
		{`"a".trim_start()`, "a"},
		{`"\t\n a".trim_start()`, "a"},

		// trim_line_end with no argument removes one CR, LF or CRLF...
		{`"abc\r".trim_line_end()`, "abc"},
		{`"abc\n".trim_line_end()`, "abc"},
		{`"abc\r\n".trim_line_end()`, "abc"},
		// ...but "\n\r" loses only the "\r", as in Ruby.
		{`"abc\n\r".trim_line_end()`, "abc\n"},
		{`"abc".trim_line_end()`, "abc"},
		// With a separator it removes one trailing occurrence.
		{`"abcd".trim_line_end("d")`, "abc"},
		{`"abcdd".trim_line_end("d")`, "abcd"},
		// The empty separator is Ruby's "drop every trailing blank line".
		{`"abc\n\n\n".trim_line_end("")`, "abc"},
		{`"abc\r\n\r\n\r\n".trim_line_end("")`, "abc"},
		// It leaves bare CRs alone, which is the part that is easy to get wrong.
		{`"abc\r\r\r".trim_line_end("")`, "abc\r\r\r"},

		{`"abcd".remove_last()`, "abc"},
		// A trailing CRLF goes as a unit, so remove_last never splits a line ending.
		{`"abc\r\n".remove_last()`, "abc"},
		{`"".remove_last()`, ""},
		{`"a".remove_last()`, ""},

		{`"".empty?()`, true},
		{`"a".empty?()`, false},
		{`"abc".contains?("b")`, true},
		{`"abc".contains?("z")`, false},
		{`"abc".contains?("")`, true},
		{`"abc".starts_with?("ab")`, true},
		{`"abc".starts_with?("z")`, false},
		{`"abc".ends_with?("bc")`, true},
		{`"abc".ends_with?("z")`, false},
		// starts_with? and ends_with? take more than one candidate and are true
		// if any of them matches, as Ruby's start_with? and end_with? do.
		{`"abc".starts_with?("z", "y", "a")`, true},
		{`"abc".starts_with?("z", "y")`, false},
		{`"abc".ends_with?("z", "c")`, true},
		{`"abc".ends_with?("z", "y")`, false},
		{`"abc".starts_with?()`, "too few arguments: got=0, want=1"},
		{`"abc".contains?()`, "too few arguments: got=0, want=1"},
		{`"abc".empty?("x")`, "too many arguments: got=1, want=0"},
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
		{`s = "hello World"; s.swap_case(); s`, "hello World"},
		{`s = "  a  "; s.trim_start(); s`, "  a  "},
		{`s = "  a  "; s.trim_end(); s`, "  a  "},
		{`s = "abc\n"; s.trim_line_end(); s`, "abc\n"},
		{`s = "abcd"; s.remove_last(); s`, "abcd"},
		{`s = "a-b"; s.replace("-", "+"); s`, "a-b"},

		// ! methods change the receiver.
		{`s = "hello World"; s.capitalize!(); s`, "Hello world"},
		{`s = "hello World"; s.swap_case!(); s`, "HELLO wORLD"},
		{`s = "  a  "; s.trim_start!(); s`, "a  "},
		{`s = "  a  "; s.trim_end!(); s`, "  a"},
		{`s = "abc\n"; s.trim_line_end!(); s`, "abc"},
		{`s = "abcd"; s.remove_last!(); s`, "abc"},
		{`s = "a-b"; s.replace!("-", "+"); s`, "a+b"},

		// ...and return it, so they chain even when a call changes nothing.
		{`"  hello World  ".trim!().capitalize!().swap_case!()`, "hELLO WORLD"},
		{`"ABC".uppercase!().reverse!()`, "CBA"},
		{`"a-b\n".trim_line_end!().replace!("-", "+").uppercase!()`, "A+B"},
		{`"abc".trim_line_end!().trim_line_end!()`, "abc"},

		// Every ! method returns a STRING rather than NIL.
		{`"a".uppercase!().type()`, "STRING"},
		{`"a".lowercase!().type()`, "STRING"},
		{`"a".capitalize!().type()`, "STRING"},
		{`"a".swap_case!().type()`, "STRING"},
		{`"a".trim!().type()`, "STRING"},
		{`"a".trim_start!().type()`, "STRING"},
		{`"a".trim_end!().type()`, "STRING"},
		{`"a".reverse!().type()`, "STRING"},
		{`"a".trim_line_end!().type()`, "STRING"},
		{`"a".remove_last!().type()`, "STRING"},
		{`"a".replace!("a", "b").type()`, "STRING"},

		// A ! method takes the same arguments as its plain counterpart.
		{`"abcd".trim_line_end!("d")`, "abc"},
		{`"a-b".replace!("-")`, "too few arguments: got=1, want=2"},
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
		"capitalize", "trim_line_end", "remove_last", "lowercase", "trim_start",
		"replace", "reverse", "trim_end", "trim", "swap_case", "uppercase",
	} {
		if !names[name] {
			t.Errorf("expected String to have %s()", name)
		}
		if !names[name+"!"] {
			t.Errorf("expected String to have %s!()", name)
		}
	}

	// Predicates have nothing to change, so they must not have a ! form.
	for _, name := range []string{"empty?", "contains?", "starts_with?", "ends_with?"} {
		if !names[name] {
			t.Errorf("expected String to have %s()", name)
		}
		if names[name+"!"] {
			t.Errorf("%s should not have a ! form", name)
		}
	}
}
