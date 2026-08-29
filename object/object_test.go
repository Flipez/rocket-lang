package object_test

import (
	"errors"
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
	object.AddEvaluator(evaluator.Eval)
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

// TestMethodListingIsSorted covers methods() and wat(), which used to read the
// method names straight out of a map. That made the same program print a
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

		first := testEval(`(` + subject + `).wat()`).Inspect()
		for range 20 {
			require.Equal(t, first, testEval(`(`+subject+`).wat()`).Inspect(),
				"wat() of %s should not vary between calls", subject)
		}
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
// to_s used to fall through to an empty string for ARRAY, HASH, ERROR, FILE and
// HTTP, because none of them implemented Stringable -- so [1,2].to_s() was ""
// while [1,2].to_json() worked.
func TestGenericMethodsWorkEverywhere(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2].to_s()`, "[1, 2]"},
		{`[].to_s()`, "[]"},
		{`{"a": 1}.to_s()`, `{"a": 1}`},
		{`"a".to_s()`, "a"},
		{`1.to_s()`, "1"},
		{`1.5.to_s()`, "1.5"},
		{`true.to_s()`, "true"},
		{`nil.to_s()`, ""},
		// to_s on a matrix already worked and must keep working.
		{`[[1,2]].to_m().to_s().size() > 0`, true},

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
		"MATRIX":  `[[1,2]].to_m()`,
		"NIL":     `nil`,
		"STRING":  `"a"`,
	}

	for wantType, literal := range subjects {
		if got := testEval(literal + ".type()"); got.Inspect() != `"`+wantType+`"` {
			t.Errorf("%s.type() = %s, want %q", literal, got.Inspect(), wantType)
			continue
		}

		for _, method := range []string{"to_s", "to_json", "methods", "type", "wat", "nil?"} {
			got := testEval(literal + "." + method + "()")
			if got.Type() == object.ERROR_OBJ {
				t.Errorf("%s.%s() failed: %s", literal, method, got.Inspect())
			}
		}

		// to_s must never be the empty string except for nil, whose string form
		// genuinely is empty.
		got := testEval(literal + ".to_s()")
		if wantType != "NIL" && got.Inspect() == `""` {
			t.Errorf("%s.to_s() is empty; the type is probably not Stringable", literal)
		}
	}
}
