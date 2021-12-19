package object_test

import (
	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"

	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

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

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%T, want=%t", result.Value, expected)
		return false
	}

	return true
}

type inputTestCase struct {
	input    string
	expected interface{}
}

func testInput(t *testing.T, tests []inputTestCase) {
	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			arrObj, ok := evaluated.(*object.Array)
			if ok {
				testStringObject(t, &object.String{Value: arrObj.Inspect()}, expected)
				continue
			}
			strObj, ok := evaluated.(*object.String)
			if ok {
				testStringObject(t, strObj, expected)
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

func TestReturnValue(t *testing.T) {
	rv := &object.ReturnValue{Value: &object.String{Value: "a"}}

	if rv.Type() != object.RETURN_VALUE_OBJ {
		t.Errorf("returnValue.Type() returns wrong type")
	}
	if rv.Inspect() != "a" {
		t.Errorf("returnValue.Inspect() returns wrong type")
	}
}

func TestNullType(t *testing.T) {
	n := &object.Null{}

	if n.Type() != object.NULL_OBJ {
		t.Errorf("null.Type() returns wrong type")
	}
	if n.Inspect() != "null" {
		t.Errorf("null.Inspect() returns wrong type")
	}
}
func TestNullObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}

func TestErrorType(t *testing.T) {
	err := &object.Error{Message: "test"}

	if err.Type() != object.ERROR_OBJ {
		t.Errorf("error.Type() returns wrong type")
	}
	if err.Inspect() != "ERROR: test" {
		t.Errorf("error.Inspect() returns wrong type")
	}
}
func TestErrorObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope().nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}

func TestFunctionObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`fn(){}.nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}
func TestFunctionType(t *testing.T) {
	tests := []inputTestCase{
		{"fn(){}", "fn() {\n\n}"},
		{"fn(a){puts(a)}", "fn(a) {\nputs(a)\n}"},
	}

	for _, tt := range tests {
		fn := testEval(tt.input).(*object.Function)
		fnInspect := fn.Inspect()

		if fnInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, fnInspect)
		}
	}
}
