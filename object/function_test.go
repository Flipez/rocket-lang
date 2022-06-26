package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestFunctionObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{"def()\n\nend.nope()", "undefined method `.nope()` for FUNCTION"},
	}

	testInput(t, tests)
}
func TestFunctionType(t *testing.T) {
	tests := []inputTestCase{
		{"def()\n\nend", "def () \n\nend"},
		{"def(a)\nputs(a)\nend", "def (a) \nputs(a)\nend"},
	}

	for _, tt := range tests {
		def := testEval(tt.input).(*object.Function)
		defInspect := def.Inspect()

		if defInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, defInspect)
		}
	}
}
