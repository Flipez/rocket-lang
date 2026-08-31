package object_test

import (
	"math"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestMathModule(t *testing.T) {
	tests := []inputTestCase{
		{`Math.E`, math.E},
		{`Math.Ln10`, math.Ln10},
		{`Math.Ln2`, math.Ln2},
		{`Math.Log10E`, math.Log10E},
		{`Math.Log2E`, math.Log2E},
		{`Math.Phi`, math.Phi},
		{`Math.Pi`, math.Pi},
		{`Math.Sqrt2`, math.Sqrt2},
		{`Math.SqrtE`, math.SqrtE},
		{`Math.SqrtPhi`, math.SqrtPhi},
		{`Math.SqrtPi`, math.SqrtPi},
	}

	testInput(t, tests)
}

// TestMathRandom covers what is left of Math's functions once every
// operation on a number moved onto the number itself: randomness, which has
// no receiver to be a method of.
func TestMathRandom(t *testing.T) {
	evaluated := testEval(`Math.random()`)

	result, ok := evaluated.(*object.Float)
	if !ok {
		t.Fatalf("object is not Float. got=%T (%+v)", evaluated, evaluated)
	}

	if result.Value < 0.0 || result.Value >= 1.0 {
		t.Errorf("Math.random() out of range [0.0, 1.0): got=%f", result.Value)
	}
}
