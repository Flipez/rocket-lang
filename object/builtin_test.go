package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestBuiltinObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`puts.nope()`, "undefined method `.nope()` for BUILTIN_FUNCTION"},
	}

	testInput(t, tests)
}

func TestBuiltinType(t *testing.T) {
	tests := []inputTestCase{
		{"puts", "puts"},
	}

	for _, tt := range tests {
		def := testEval(tt.input).(*object.BuiltinFunction)
		defInspect := def.Inspect()

		if defInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, defInspect)
		}
	}
}
