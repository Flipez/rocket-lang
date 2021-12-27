package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

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
