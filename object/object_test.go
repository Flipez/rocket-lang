package object_test

import (
	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"

	"testing"
)

type inputTestCase struct {
	input    string
	expected interface{}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	imports := make(map[string]struct{})
	p := parser.New(l, imports)
	program, _ := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testInput(t *testing.T, tests []inputTestCase) {
	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case float64:
			testFloatObject(t, evaluated, float64(expected))
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
