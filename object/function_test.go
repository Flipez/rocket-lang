package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestFunctionObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`def(){}.nope()`, "undefined method `.nope()` for FUNCTION"},
	}

	testInput(t, tests)
}
func TestFunctionType(t *testing.T) {
	tests := []inputTestCase{
		{"def(){}", "def () {\n\n}"},
		{"def(a){puts(a)}", "def (a) {\nputs(a)\n}"},
	}

	for _, tt := range tests {
		def := testEval(tt.input).(*object.Function)
		defInspect := def.Inspect()

		if defInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, defInspect)
		}
	}
}
