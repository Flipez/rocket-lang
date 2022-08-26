package object_test

import (
	"math"
	"testing"
)

func TestMathObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`Math.abs(-2.0)`, math.Abs(-2.0)},
		{`Math.acos(1.0)`, math.Acos(1.0)},
		{`Math.asin(0.0)`, math.Asin(0.0)},
		{`Math.atan(0.0)`, math.Atan(0.0)},
		{`Math.ceil(1.49)`, math.Ceil(1.49)},
		{`Math.copysign(3.2, -1.0)`, math.Copysign(3.2, -1.0)},
		{`Math.cos(Pi/2)`, math.Cos(math.Pi / 2)},
		{`Math.exp(1.0)`, math.Exp(1.0)},
		{`Math.floor(1.51)`, math.Floor(1.51)},
		{`Math.log(2.7183)`, math.Log(2.7183)},
		{`Math.log10(100.0)`, math.Log10(100.0)},
		{`Math.log2(256.0)`, math.Log2(256.0)},
		{`Math.pow(2.0, 3.0)`, math.Pow(2.0, 3.0)},
		{`Math.remainder(100.0, 30.0)`, math.Remainder(100.0, 30.0)},
		{`Math.sin(Pi)`, math.Sin(math.Pi)},
		{`Math.sqrt(3.0 * 3.0 + 4.0 * 4.0)`, math.Sqrt(3.0*3.0 + 4.0*4.0)},
		{`Math.tan(0.0)`, math.Tan(0.0)},
	}

	testInput(t, tests)
}
