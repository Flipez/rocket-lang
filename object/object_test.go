package object_test

import (
	"errors"
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
	imports := make(map[string]struct{})
	p := parser.New(l, imports)
	program, _ := p.ParseProgram()
	object.AddEvaluator(evaluator.Eval)
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testInput(t *testing.T, tests []inputTestCase) {
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
			_, ok = evaluated.(*object.Nil)
			if ok {
				continue
			}

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
