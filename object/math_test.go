package object_test

import (
	"math"
	"testing"
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
