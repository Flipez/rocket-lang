package object_test

import (
	"errors"
	"io"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

type inputTestCase struct {
	input    string
	expected interface{}
}

func testEval(input string) object.Object {
	l := lexer.New(input, "test")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testInput(t *testing.T, tests []inputTestCase) {
	t.Helper()

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, expected)
		case float64:
			testFloatObject(t, evaluated, float64(expected))
		case string:
			arrObj, ok := evaluated.(*object.Array)
			if ok {
				testStringObject(t, object.NewString(arrObj.Inspect()), expected)
				continue
			}
			matObj, ok := evaluated.(*object.Matrix)
			if ok {
				testStringObject(t, object.NewString(matObj.Inspect()), expected)
				continue
			}
			strObj, ok := evaluated.(*object.String)
			if ok {
				testStringObject(t, strObj, expected)
				continue
			}
			// A NIL used to be accepted here, which meant every string and
			// every error-message expectation in this file silently passed
			// when the method under test returned nil instead.
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		case bool:
			testBooleanObject(t, evaluated, expected)
		case nil:
			// A nil expectation means the conversion could not be performed.
			// Without this case the type switch matched nothing and the
			// assertion silently passed whatever it was given.
			if _, ok := evaluated.(*object.Nil); !ok {
				t.Errorf("input %q: object is not Nil. got=%T (%+v)", tt.input, evaluated, evaluated)
			}
		}
	}
}

func TestIsError(t *testing.T) {
	trueErrors := []object.Object{
		object.NewError(errors.New("test error")),
		object.NewError("test error"),
		object.NewErrorFormat("test %s", "error"),
	}

	for _, err := range trueErrors {
		if !object.IsError(err) {
			t.Errorf("'%s' should be an error", err.Inspect())
		}
	}

	falseErrors := []object.Object{
		nil,
		object.NewString("a"),
		object.NIL,
	}
	for _, err := range falseErrors {
		if object.IsError(nil) {
			t.Errorf("'%#v' is not an error", err)
		}
	}
}

func TestIsNumber(t *testing.T) {
	if !object.IsNumber(object.NewInteger(1)) {
		t.Error("INTEGER_OBJ should be a number")
	}
	if !object.IsNumber(object.NewFloat(1.1)) {
		t.Error("FLOAT_OBJ should be a number")
	}
	if object.IsNumber(object.NIL) {
		t.Error("NIL_OBJ is not a number")
	}
}

func TestIsTruthy(t *testing.T) {
	if !object.IsTruthy(object.TRUE) {
		t.Error("BOOLEAN_OBJ=true should be truthy")
	}
	if !object.IsTruthy(object.NewString("")) {
		t.Error("STRING_OBJ should be truthy")
	}
	if object.IsTruthy(object.NIL) {
		t.Error("NIL_OBJ should not be truthy")
	}
	if object.IsTruthy(object.FALSE) {
		t.Errorf("BOOLEAN_OBJ=false, should not be truthy")
	}
}

func TestIsFalsy(t *testing.T) {
	if object.IsFalsy(object.TRUE) {
		t.Error("BOOLEAN_OBJ=true should not be falsy")
	}
	if object.IsFalsy(object.NewString("")) {
		t.Error("STRING_OBJ should not be falsy")
	}
	if !object.IsFalsy(object.NIL) {
		t.Error("NIL_OBJ should be falsy")
	}
	if !object.IsFalsy(object.FALSE) {
		t.Errorf("BOOLEAN_OBJ=false, should be falsy")
	}
}

func TestAnyToObject(t *testing.T) {
	testcases := map[any]object.Object{
		"a":        object.NewString("a"),
		1:          object.NewInteger(1),
		1.2:        object.NewFloat(1.2),
		true:       object.TRUE,
		struct{}{}: object.NIL,
	}

	for input, expected := range testcases {
		obj := object.AnyToObject(input)
		if obj.Type() != expected.Type() {
			t.Errorf("wrong object type, got=%s want=%s", obj.Type(), expected.Type())
		}
	}
}

func TestObjectToAny(t *testing.T) {
	stringObj := object.NewString("a")
	intObj := object.NewInteger(1)
	floatObj := object.NewFloat(1.2)
	testcases := map[object.Object]any{
		stringObj:   "a",
		intObj:      1,
		floatObj:    float64(1.2),
		object.TRUE: true,
		object.NewArrayWithObjects(stringObj, intObj, floatObj): []any{"a", 1, float64(1.2)},
		object.NIL: nil,
	}

	hash := object.NewHash(nil)
	hash.Set("a", 1)
	testcases[hash] = map[any]any{"a": 1}

	for input, expected := range testcases {
		require.Equal(t, expected, object.ObjectToAny(input))
	}
}

// TestMethodListingIsSorted covers methods() and help(), which used to read
// the method names straight out of a map. That made the same program print a
// different order on every run, so their output could not be relied on or
// documented.
func TestMethodListingIsSorted(t *testing.T) {
	for _, subject := range []string{`[1,2,3]`, `"a"`, `{"a": 1}`, `1`, `1.0`, `true`} {
		names := object.NewArray(nil)

		// Repeat, because a map-order bug reproduces by chance rather than
		// every time.
		for range 20 {
			listed, ok := testEval(`(` + subject + `).methods()`).(*object.Array)
			require.True(t, ok, "methods() should return an array for %s", subject)

			rendered := listed.Inspect()
			require.Equal(t, sortedInspect(listed), rendered,
				"methods() of %s should be sorted", subject)

			if len(names.Elements) == 0 {
				names = listed
				continue
			}
			require.Equal(t, names.Inspect(), rendered,
				"methods() of %s should not vary between calls", subject)
		}

		// help() prints its listing and returns nil, so comparing return
		// values would compare "nil" to "nil" and assert nothing. Captured so
		// the listing does not end up in the test output.
		var helpResult object.Object
		captureStdout(t, func() { helpResult = testEval(`(` + subject + `).help()`) })
		require.Equal(t, object.ObjectType(object.NIL_OBJ), helpResult.Type(),
			"help() of %s should return nil", subject)

		first := captureStdout(t, func() { testEval(`(` + subject + `).help()`) })
		require.Equal(t, sortedHelpListing(first), first,
			"help() of %s should list its methods sorted", subject)

		// One header line plus one line per method, so nothing is dropped from
		// the listing or counted twice.
		listed, ok := testEval(`(` + subject + `).methods()`).(*object.Array)
		require.True(t, ok)
		require.Len(t, strings.Split(strings.TrimSuffix(first, "\n"), "\n"), len(listed.Elements)+1,
			"help() of %s should print one line per method plus the header", subject)

		for range 20 {
			printed := captureStdout(t, func() { testEval(`(` + subject + `).help()`) })
			require.Equal(t, first, printed,
				"help() of %s should not vary between calls", subject)
		}

		// wat is the only alias in the language, kept as an easter egg; it must
		// print the exact same listing help() does rather than a copy that could
		// drift.
		aliased := captureStdout(t, func() { testEval(`(` + subject + `).wat()`) })
		require.Equal(t, first, aliased, "wat() of %s should match help()", subject)
	}
}

// sortedInspect renders the array with its elements in sorted order, giving the
// expectation to compare the real rendering against. It sorts the bare names
// rather than their quoted renderings, because a quote sorts after "!" and
// would put "reverse!" ahead of "reverse".
func sortedInspect(a *object.Array) string {
	names := make([]string, len(a.Elements))
	for i, element := range a.Elements {
		name, ok := element.(*object.String)
		if !ok {
			return "element " + string(element.Type()) + " is not a string"
		}
		names[i] = name.Value
	}
	sort.Strings(names)

	for i, name := range names {
		names[i] = `"` + name + `"`
	}

	return "[" + strings.Join(names, ", ") + "]"
}

// TestGenericMethodsWorkEverywhere checks the methods every type answers to.
// to_string used to fall through to an empty string for ARRAY, HASH, ERROR,
// FILE and HTTP, because none of them implemented Stringable -- so
// [1,2].to_string() was "" while [1,2].to_json() worked.
func TestGenericMethodsWorkEverywhere(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2].to_string()`, "[1, 2]"},
		{`[].to_string()`, "[]"},
		{`{"a": 1}.to_string()`, `{"a": 1}`},
		{`"a".to_string()`, "a"},
		{`1.to_string()`, "1"},
		{`1.5.to_string()`, "1.5"},
		{`true.to_string()`, "true"},
		{`nil.to_string()`, ""},
		// to_string on a matrix already worked and must keep working.
		{`[[1,2]].to_matrix().to_string().size() > 0`, true},

		// nil? asks the question that comparing against nil already allowed,
		// but reads better in a chain.
		{`nil.nil?()`, true},
		{`1.nil?()`, false},
		{`"".nil?()`, false},
		{`[].nil?()`, false},
		{`{}.nil?()`, false},
		{`false.nil?()`, false},
		// A method that returns nil on a miss can be asked directly.
		{`[].first().nil?()`, true},
		{`[1].first().nil?()`, false},

		{`nil.to_json()`, "null"},
	}
	testInput(t, tests)
}

// TestEveryTypeAnswersTheGenericMethods walks every registered type and calls
// each generic method on a value of that type, so a type added later cannot
// quietly fail to answer one of them.
func TestEveryTypeAnswersTheGenericMethods(t *testing.T) {
	// One literal per type that can be written down.
	subjects := map[string]string{
		"ARRAY":   `[1,2]`,
		"BOOLEAN": `true`,
		"FLOAT":   `1.5`,
		"HASH":    `{"a": 1}`,
		"INTEGER": `1`,
		"MATRIX":  `[[1,2]].to_matrix()`,
		"NIL":     `nil`,
		"STRING":  `"a"`,
	}

	for wantType, literal := range subjects {
		if got := testEval(literal + ".type()"); got.Inspect() != `"`+wantType+`"` {
			t.Errorf("%s.type() = %s, want %q", literal, got.Inspect(), wantType)
			continue
		}

		for _, method := range []string{"to_string", "to_json", "methods", "type", "help", "nil?"} {
			got := testEval(literal + "." + method + "()")
			if got.Type() == object.ERROR_OBJ {
				t.Errorf("%s.%s() failed: %s", literal, method, got.Inspect())
			}
		}

		// to_string must never be the empty string except for nil, whose string
		// form genuinely is empty.
		got := testEval(literal + ".to_string()")
		if wantType != "NIL" && got.Inspect() == `""` {
			t.Errorf("%s.to_string() is empty; the type is probably not Stringable", literal)
		}
	}
}

// TestUsageMarksOptionalAndVariadicArgs covers what a rendered signature says
// about its arguments. Usage() used to print only the types, so three different
// contracts came out identically: count(STRING) takes exactly one, split(STRING)
// takes zero or one, and starts_with?(STRING) takes one or more.
func TestUsageMarksOptionalAndVariadicArgs(t *testing.T) {
	tests := []struct {
		name   string
		layout object.MethodLayout
		want   string
	}{
		{"size", object.MethodLayout{}, "size()"},
		{
			"count",
			object.MethodLayout{ArgPattern: object.Args(object.Arg(object.STRING_OBJ))},
			"count(STRING)",
		},
		{
			"split",
			object.MethodLayout{ArgPattern: object.Args(object.OptArg(object.STRING_OBJ))},
			"split([STRING])",
		},
		{
			"starts_with?",
			object.MethodLayout{ArgPattern: object.Args(object.OverloadArg(object.STRING_OBJ))},
			"starts_with?(STRING...)",
		},
		{
			// A required argument followed by an optional one, as pow has.
			"pow",
			object.MethodLayout{ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
				object.OptArg(object.INTEGER_OBJ),
			)},
			"pow(INTEGER, [INTEGER])",
		},
		{
			// A union keeps its alternatives inside the brackets.
			"first",
			object.MethodLayout{ArgPattern: object.Args(
				object.OptArg(object.INTEGER_OBJ, object.STRING_OBJ),
			)},
			"first([INTEGER|STRING])",
		},
		{
			// A variadic union is parenthesised, so the ellipsis does not read
			// as applying to STRING alone.
			"format",
			object.MethodLayout{ArgPattern: object.Args(
				object.OverloadArg(object.STRING_OBJ, object.INTEGER_OBJ),
			)},
			"format((STRING|INTEGER)...)",
		},
	}

	for _, tt := range tests {
		if got := tt.layout.Usage(tt.name); got != tt.want {
			t.Errorf("Usage(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

// TestArgumentStringStaysBare checks that the wording of the "wrong argument
// type" error did not pick up the signature brackets: there the reader is being
// told what one position accepts, and "want=[STRING]" would only be confusing.
func TestArgumentStringStaysBare(t *testing.T) {
	if got := object.OptArg(object.STRING_OBJ).String(); got != "STRING" {
		t.Errorf("OptArg(...).String() = %q, want %q", got, "STRING")
	}
	if got := object.OverloadArg(object.STRING_OBJ).String(); got != "STRING" {
		t.Errorf("OverloadArg(...).String() = %q, want %q", got, "STRING")
	}

	tests := []inputTestCase{
		{`"a".split(1)`, "wrong argument type on position 1: got=INTEGER, want=STRING"},
		{`"a".starts_with?(1)`, "wrong argument type on position 1: got=INTEGER, want=STRING"},
	}
	testInput(t, tests)
}

// captureStdout collects what fn writes to os.Stdout. help() prints its
// listing instead of returning it, so this is the only way to assert on it.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	read, write, err := os.Pipe()
	require.NoError(t, err)

	original := os.Stdout
	os.Stdout = write
	defer func() { os.Stdout = original }()

	fn()
	require.NoError(t, write.Close())

	printed, err := io.ReadAll(read)
	require.NoError(t, err)

	return string(printed)
}

// sortedHelpListing returns the listing with its method lines sorted, giving
// the expectation to compare the real output against. The header stays first.
//
// It sorts on the bare method name rather than the rendered line, because "("
// sorts after "!" and would otherwise put trim_line_end!([STRING]) ahead of
// trim_line_end([STRING]).
func sortedHelpListing(listing string) string {
	lines := strings.Split(strings.TrimSuffix(listing, "\n"), "\n")
	if len(lines) < 2 {
		return listing
	}

	sorted := make([]string, len(lines)-1)
	copy(sorted, lines[1:])
	sort.SliceStable(sorted, func(i, j int) bool {
		return helpMethodName(sorted[i]) < helpMethodName(sorted[j])
	})

	return lines[0] + "\n" + strings.Join(sorted, "\n") + "\n"
}

func helpMethodName(line string) string {
	name := strings.TrimPrefix(line, "\t")
	if open := strings.Index(name, "("); open >= 0 {
		return name[:open]
	}

	return name
}

// TestTypeGroups covers the argument groups introduced for #296. A group names
// what an object must be able to do instead of listing the types that can do it,
// which is how append came to accept FUNCTION but not FLOAT.
func TestTypeGroups(t *testing.T) {
	tests := []inputTestCase{
		// ANY takes everything, so these are no longer type errors. Each of
		// them was one, because the hand-written list behind it had forgotten
		// FLOAT and MATRIX.
		{`a = [1]; a.append(1.5); a.to_json()`, "[1,1.5]"},
		{`a = [1]; a.prepend(1.5); a.to_json()`, "[1.5,1]"},
		{`a = [1]; a.insert(0, 1.5); a.to_json()`, "[1.5,1]"},
		{`[1.5].contains?(1.5)`, true},
		{`[1.5].index_of(1.5)`, 0},
		{`[1.5].last_index_of(1.5)`, 0},
		{`[1.5].count(1.5)`, 1},
		{`a = [1.5]; a.remove(1.5); a.to_json()`, "[]"},
		{`a = [1]; a.append([[1,2]].to_matrix()); a.size()`, 2},
		{`[1].contains?([[1,2]].to_matrix())`, false},

		// HASHABLE takes what can be a key. get and has_key? used to assert
		// without checking, so {"a": 1}.get(nil, 0) panicked and took the
		// process with it.
		{`{"a": 1}.get(nil, 0)`, "wrong argument type on position 1: got=NIL, want=HASHABLE"},
		{`{"a": 1}.get([[1,2]].to_matrix(), 0)`, "wrong argument type on position 1: got=MATRIX, want=HASHABLE"},
		{`{"a": 1}.has_key?(nil)`, "wrong argument type on position 1: got=NIL, want=HASHABLE"},
		{`{"a": 1}.fetch(nil)`, "wrong argument type on position 1: got=NIL, want=HASHABLE"},
		{`{"a": 1}.remove(nil)`, "wrong argument type on position 1: got=NIL, want=HASHABLE"},
		// All four take the same keys, which was not true before: get accepted
		// a NIL and crashed, has_key? rejected a FLOAT, remove accepted one.
		{`{1.5: "a"}.get(1.5, "missing")`, "a"},
		{`{1.5: "a"}.has_key?(1.5)`, true},
		{`{1.5: "a"}.fetch(1.5)`, "a"},
		{`h = {1.5: "a"}; h.remove(1.5)`, "a"},
		// A hash and an array are hashable, so they are keys too.
		{`{[1]: "a"}.get([1], "missing")`, "a"},

		// The fallback value of get is ANY, not HASHABLE -- it is never used
		// as a key. fetch has no fallback argument at all; see hash_test.go.
		{`{"a": 1}.get("z", nil)`, nil},

		// NUMERIC takes INTEGER or FLOAT.
		{`m = [[1,2]].to_matrix(); m.set(0, 0, 9); m.to_array().to_json()`, "[[9,2]]"},
		{`m = [[1,2]].to_matrix(); m.set(0, 0, 9.5); m.to_array().to_json()`, "[[9.5,2]]"},
		{`[[1,2]].to_matrix().set(0, 0, "x")`, "wrong argument type on position 3: got=STRING, want=NUMERIC"},
		{`[[1,2]].to_matrix().set(0, 0, nil)`, "wrong argument type on position 3: got=NIL, want=NUMERIC"},
	}
	testInput(t, tests)
}

// TestTypeGroupsRenderInSignatures checks that a group prints as its own name
// rather than expanding. Hash#get used to render as a 118-character union of
// nine types repeated twice, which told the reader nothing.
func TestTypeGroupsRenderInSignatures(t *testing.T) {
	tests := []struct {
		name   string
		layout object.MethodLayout
		want   string
	}{
		{
			"append",
			object.MethodLayout{ArgPattern: object.Args(object.Arg(object.ANY))},
			"append(ANY)",
		},
		{
			"get",
			object.MethodLayout{ArgPattern: object.Args(
				object.Arg(object.HASHABLE),
				object.OptArg(object.ANY),
			)},
			"get(HASHABLE, [ANY])",
		},
		{
			"set",
			object.MethodLayout{ArgPattern: object.Args(object.Arg(object.NUMERIC))},
			"set(NUMERIC)",
		},
		{
			"format",
			object.MethodLayout{ArgPattern: object.Args(object.OverloadArg(object.ANY))},
			"format(ANY...)",
		},
		{
			// A group and a concrete type can sit in the same argument.
			"mixed",
			object.MethodLayout{ArgPattern: object.Args(
				object.Arg(object.NUMERIC, object.STRING_OBJ),
			)},
			"mixed(NUMERIC|STRING)",
		},
	}

	for _, tt := range tests {
		if got := tt.layout.Usage(tt.name); got != tt.want {
			t.Errorf("Usage(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

// TestTypeGroupNamesDoNotCollideWithTypes guards the one thing that would make
// groups silently wrong: a group shares its namespace with the object type
// names, so a group called STRING would shadow the type.
func TestTypeGroupNamesDoNotCollideWithTypes(t *testing.T) {
	groups := []string{object.ANY, object.HASHABLE, object.STRINGABLE, object.NUMERIC}

	types := []string{
		object.INTEGER_OBJ, object.STRING_OBJ, object.BOOLEAN_OBJ, object.ARRAY_OBJ,
		object.HASH_OBJ, object.MATRIX_OBJ, object.FLOAT_OBJ, object.ERROR_OBJ,
		object.NIL_OBJ, object.FILE_OBJ, object.FUNCTION_OBJ, object.HTTP_OBJ,
	}

	for _, group := range groups {
		for _, objType := range types {
			if group == objType {
				t.Errorf("type group %q collides with the object type of the same name", group)
			}
		}
	}
}

// TestIsA covers the predicate for both a type name and a group name, and the
// mistake it has to refuse: a name that is neither.
func TestIsA(t *testing.T) {
	tests := []inputTestCase{
		// A group.
		{`"a".is_a?("HASHABLE")`, true},
		{`nil.is_a?("HASHABLE")`, false},
		{`1.is_a?("COMPARABLE")`, true},
		{`true.is_a?("COMPARABLE")`, false},
		{`true.is_a?("INTEGERABLE")`, true},
		{`[1].is_a?("STRINGABLE")`, true},
		{`def() end.is_a?("STRINGABLE")`, false},
		{`1.is_a?("NUMERIC")`, true},
		{`1.5.is_a?("NUMERIC")`, true},
		{`"1".is_a?("NUMERIC")`, false},
		// ANY is true for everything, including a function and a nil.
		{`"a".is_a?("ANY")`, true},
		{`nil.is_a?("ANY")`, true},
		{`def() end.is_a?("ANY")`, true},

		// A concrete type, which collapses type() == "X" into the same question.
		{`"a".is_a?("STRING")`, true},
		{`"a".is_a?("INTEGER")`, false},
		{`1.5.is_a?("FLOAT")`, true},
		{`1.is_a?("FLOAT")`, false},
		{`nil.is_a?("NIL")`, true},
		{`[1].is_a?("ARRAY")`, true},
		{`{"a": 1}.is_a?("HASH")`, true},
		{`def() end.is_a?("FUNCTION")`, true},
		{`[[1,2]].to_matrix().is_a?("MATRIX")`, true},

		// A name that is neither is an error, not a false. A typo would
		// otherwise answer "no" and read like a real result.
		{`"a".is_a?("HASHBALE")`, "unknown type or type group: HASHBALE"},
		{`"a".is_a?("STRIGN")`, "unknown type or type group: STRIGN"},
		{`"a".is_a?("")`, "unknown type or type group: "},
		// Names are exact; type() answers in upper case, so is_a? asks in it.
		{`"a".is_a?("hashable")`, "unknown type or type group: hashable"},
		{`"a".is_a?("string")`, "unknown type or type group: string"},

		{`"a".is_a?(1)`, "wrong argument type on position 1: got=INTEGER, want=STRING"},
		{`"a".is_a?()`, "too few arguments: got=0, want=1"},
	}
	testInput(t, tests)
}

// TestTypeGroupsMethod covers the listing. The expectations are the rows of the
// membership table in docs/docs/language/types.md, so the two cannot
// drift apart unnoticed.
func TestTypeGroupsMethod(t *testing.T) {
	tests := []inputTestCase{
		{`"a".type_groups().to_json()`, `["COMPARABLE","HASHABLE","INTEGERABLE","STRINGABLE"]`},
		{`1.type_groups().to_json()`, `["COMPARABLE","HASHABLE","INTEGERABLE","NUMERIC","STRINGABLE"]`},
		{`1.5.type_groups().to_json()`, `["COMPARABLE","HASHABLE","INTEGERABLE","NUMERIC","STRINGABLE"]`},
		{`true.type_groups().to_json()`, `["HASHABLE","INTEGERABLE","STRINGABLE"]`},
		{`[1].type_groups().to_json()`, `["HASHABLE","STRINGABLE"]`},
		{`{"a": 1}.type_groups().to_json()`, `["HASHABLE","STRINGABLE"]`},
		{`nil.type_groups().to_json()`, `["STRINGABLE"]`},
		{`[[1,2]].to_matrix().type_groups().to_json()`, `["STRINGABLE"]`},
		// A function is CALLABLE and nothing else -- not STRINGABLE, not
		// HASHABLE, which is the whole reason the element checks in join, sum,
		// unique and sort exist.
		{`def() end.type_groups().to_json()`, `["CALLABLE"]`},
		{`print.type_groups().to_json()`, `["CALLABLE"]`},
		// Nothing else is callable.
		{`1.type_groups().contains?("CALLABLE")`, false},
		{`"a".type_groups().contains?("CALLABLE")`, false},

		{`"a".type_groups("x")`, "too many arguments: got=1, want=0"},
	}
	testInput(t, tests)
}

// TestIsAAgreesWithTypeGroups walks every type against every group and checks
// that the two methods answer the same. They share one predicate, and this is
// what keeps a future change from giving them separate copies of it.
func TestIsAAgreesWithTypeGroups(t *testing.T) {
	subjects := []string{
		`"a"`, `1`, `1.5`, `true`, `[1]`, `{"a": 1}`, `nil`,
		`[[1,2]].to_matrix()`, `def() end`,
	}

	for _, subject := range subjects {
		for _, group := range object.TypeGroupNames() {
			predicate := testEval(`(` + subject + `).is_a?("` + group + `")`)

			// ANY is deliberately absent from the listing: it says an argument
			// accepts anything, which describes a parameter rather than this
			// value, and it is true of everything. is_a? still answers it.
			if group == object.ANY {
				require.Equal(t, "true", predicate.Inspect(),
					"is_a?(ANY) should be true for %s", subject)
				require.Equal(t, "false",
					testEval(`(`+subject+`).type_groups().contains?("ANY")`).Inspect(),
					"type_groups() should not list ANY for %s", subject)

				continue
			}

			listed := testEval(`(` + subject + `).type_groups().contains?("` + group + `")`)
			require.Equal(t, listed.Inspect(), predicate.Inspect(),
				"is_a?(%q) and type_groups() disagree for %s", group, subject)
		}

		// A value is always its own type, whatever that is.
		own := testEval(`(` + subject + `).is_a?((` + subject + `).type())`)
		require.Equal(t, "true", own.Inspect(),
			"%s should be a %s", subject, testEval(`(`+subject+`).type()`).Inspect())

	}
}

// TestKnownObjectTypesAreComplete guards the list is_a? validates against. A
// type missing from it would make is_a? report a real type name as unknown.
func TestKnownObjectTypesAreComplete(t *testing.T) {
	known := make(map[string]bool)
	for _, name := range object.KnownObjectTypes() {
		known[name] = true
	}

	for objType := range object.ListObjectMethods() {
		name := string(objType)
		if name == "*" {
			continue
		}
		if !known[name] {
			t.Errorf("%s registers methods but is not in KnownObjectTypes()", name)
		}
	}

	// Every group name has to stay out of the type list, or is_a? would resolve
	// it twice and the two answers could differ.
	for _, group := range object.TypeGroupNames() {
		if known[group] {
			t.Errorf("type group %q is also listed as an object type", group)
		}
	}
}

// TestContains covers the split of the old include?, one name shared by three
// types that meant two different things. String and Array test membership of
// an element, so both answer to contains?. Hash tests membership of a key,
// which contains? never said -- {"a": 1}.include?(1) gave no clue whether 1
// was being looked for as a key or a value. Hash answers to has_key? instead.
func TestContains(t *testing.T) {
	tests := []inputTestCase{
		{`"hello".contains?("ell")`, true},
		{`"hello".contains?("z")`, false},
		{`[1,2,3].contains?(2)`, true},
		{`[1,2,3].contains?(9)`, false},
		{`{"a": 1}.has_key?("a")`, true},
		{`{"a": 1}.has_key?("b")`, false},
	}

	testInput(t, tests)
}
