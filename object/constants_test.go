package object_test

import (
	"math"
	"testing"
)

func TestConstants(t *testing.T) {
	tests := []inputTestCase{
		{`E`, math.E},
		{`Ln10`, math.Ln10},
		{`Ln2`, math.Ln2},
		{`Log10E`, math.Log10E},
		{`Log2E`, math.Log2E},
		{`Phi`, math.Phi},
		{`Pi`, math.Pi},
		{`Sqrt2`, math.Sqrt2},
		{`SqrtE`, math.SqrtE},
		{`SqrtPhi`, math.SqrtPhi},
		{`SqrtPi`, math.SqrtPi},
	}

	testInput(t, tests)
}
