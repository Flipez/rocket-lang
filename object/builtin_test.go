package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestBuiltinObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`puts.nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}

func TestBuiltinType(t *testing.T) {
	tests := []inputTestCase{
		{"puts", "builtin function"},
	}

	for _, tt := range tests {
		def := testEval(tt.input).(*object.Builtin)
		defInspect := def.Inspect()

		if defInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, defInspect)
		}
	}
}
