package object_test

import (
	"github.com/flipez/rocket-lang/object"
	"testing"
)

func TestObjectMethods(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"test".count("e")`, 1},
		{`"test".count()`, "Missing argument to count()!"},
		{`"test".find("e")`, 1},
		{`"test".find()`, "Missing argument to find()!"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}
