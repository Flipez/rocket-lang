package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNumberObjects(t *testing.T) {
	tests := []inputTestCase{
		// integer | integer
		{"1 == 1", true},
		{"1 == 2", false},

		{"1 != 1", false},
		{"1 != 2", true},

		{"1 < 2", true},
		{"1 < 1", false},

		{"2 > 1", true},
		{"1 > 2", false},

		{"0+0", 0},
		{"0+1", 1},
		{"1+2", 3},

		{"0-0", 0},
		{"0-1", -1},
		{"1-0", 1},
		{"1-2", -1},
		{"2-1", 1},

		{"1*1", 1},
		{"1*2", 2},
		{"1*0", 0},
		{"2*1", 2},
		{"0*1", 0},

		{"1/1", 1},
		// Dividing two integers yields an integer, as in Ruby. The result is
		// truncated toward zero, which keeps it consistent with %, whose sign
		// follows the dividend.
		{"1/2", 0},
		{"2/1", 2},
		{"5/2", 2},
		{"4/2", 2},
		{"0-5", -5},
		{"(0-5)/2", -2},
		{"(0-5)%2", -1},
		// Mixing the two promotes to float.
		{"4.0/2", 2.0},
		{"4/2.0", 2.0},
		{"5/2.0", 2.5},

		// float | float
		{"1.0 == 1.0", true},
		{"1.0 == 2.0", false},

		{"1.0 != 1.0", false},
		{"1.0 != 2.0", true},

		{"1.0 < 2.0", true},
		{"1.0 < 1.0", false},

		{"2.0 > 1.0", true},
		{"1.0 > 2.0", false},

		{"0.0+0.0", 0.0},
		{"0.0+1.0", 1.0},
		{"1.0+2.0", 3.0},
		{"1.2+3.4", 4.6},

		{"0.0-0.0", 0.0},
		{"0.0-1.0", -1.0},
		{"1.0-0.0", 1.0},
		{"1.0-2.0", -1.0},
		{"2.0-1.0", 1.0},
		{"3.4-1.2", 2.2},

		{"1.0*1.0", 1.0},
		{"1.0*2.0", 2.0},
		{"1.0*0.0", 0.0},
		{"2.0*1.0", 2.0},
		{"0.0*1.0", 0.0},
		{"1.2*3.4", 4.08},

		{"1.0/1.0", 1.0},
		{"1.0/2.0", 0.5},
		{"2.0/1.0", 2.0},
		{"6.8/3.4", 2.0},

		// integer | float
		{"1 == 1.0", false},
		{"1 == 2.0", false},

		{"1 != 1.0", true},
		{"1 != 2.0", true},

		{"1 < 2.0", true},
		{"1 < 1.0", false},

		{"2 > 1.0", true},
		{"1 > 2.0", false},

		{"0+0.0", 0.0},
		{"0+1.0", 1.0},
		{"1+2.0", 3.0},
		{"1+3.4", 4.4},

		{"0-0.0", 0.0},
		{"0-1.0", -1.0},
		{"1-0.0", 1.0},
		{"1-2.0", -1.0},
		{"2-1.0", 1.0},
		{"3-1.2", 1.8},

		{"1*1.0", 1.0},
		{"1*2.0", 2.0},
		{"1*0.0", 0.0},
		{"2*1.0", 2.0},
		{"0*1.0", 0.0},
		{"3*1.2", 3.5999999999999996},

		{"1/1.0", 1.0},
		{"1/2.0", 0.5},
		{"2/1.0", 2.0},

		// float | integer
		{"1.0 == 1", false},
		{"1.0 == 2", false},

		{"1.0 != 1", true},
		{"1.0 != 2", true},

		{"1.0 < 2", true},
		{"1.0 < 1", false},

		{"2.0 > 1", true},
		{"1.0 > 2", false},

		{"0.0+0", 0.0},
		{"0.0+1", 1.0},
		{"1.0+2", 3.0},
		{"1.2+3", 4.2},

		{"0.0-0", 0.0},
		{"0.0-1", -1.0},
		{"1.0-0", 1.0},
		{"1.0-2", -1.0},
		{"2.0-1", 1.0},
		{"3.4-1", 2.4},

		{"1.0*1", 1.0},
		{"1.0*2", 2.0},
		{"1.0*0", 0.0},
		{"2.0*1", 2.0},
		{"0.0*1", 0.0},
		{"1.2*3", 3.5999999999999996},

		{"1.0/1", 1.0},
		{"1.0/2", 0.5},
		{"2.0/1", 2.0},

		// float var | integer
		{"a = 1.0; a == 1", false},
		{"a = 1.0; a == 2", false},

		{"a = 1.0; a != 1", true},
		{"a = 1.0; a != 2", true},

		{"a = 1.0; a < 2", true},
		{"a = 1.0; a < 1", false},

		{"a = 2.0; a > 1", true},
		{"a = 1.0; a > 2", false},

		{"a = 0.0; a = a+0; a", 0.0},
		{"a = 0.0; a = a+1; a", 1.0},
		{"a = 1.0; a = a+2; a", 3.0},
		{"a = 1.2; a = a+3; a", 4.2},

		{"a = 0.0; a = a-0; a", 0.0},
		{"a = 0.0; a = a-1; a", -1.0},
		{"a = 1.0; a = a-0; a", 1.0},
		{"a = 1.0; a = a-2; a", -1.0},
		{"a = 2.0; a = a-1; a", 1.0},
		{"a = 3.4; a = a-1; a", 2.4},

		{"a = 1.0; a = a*1; a", 1.0},
		{"a = 1.0; a = a*2; a", 2.0},
		{"a = 1.0; a = a*0; a", 0.0},
		{"a = 2.0; a = a*1; a", 2.0},
		{"a = 0.0; a = a*1; a", 0.0},
		{"a = 1.2; a = a*3; a", 3.5999999999999996},

		{"a = 1.0; a = a/1; a", 1.0},
		{"a = 1.0; a = a/2; a", 0.5},
		{"a = 2.0; a = a/1; a", 2.0},

		// integer var | float
		{"a = 1; a == 1.0", false},
		{"a = 1; a == 2.0", false},

		{"a = 1; a != 1.0", true},
		{"a = 1; a != 2.0", true},

		{"a = 1; a < 2.0", true},
		{"a = 1; a < 1.0", false},

		{"a = 2; a > 1.0", true},
		{"a = 1; a > 2.0", false},

		{"a = 0; a = a+0.0; a", 0.0},
		{"a = 0; a = a+1.0; a", 1.0},
		{"a = 1; a = a+2.0; a", 3.0},
		{"a = 1; a = a+3.4; a", 4.4},

		{"a = 0; a = a-0.0; a", 0.0},
		{"a = 0; a = a-1.0; a", -1.0},
		{"a = 1; a = a-0.0; a", 1.0},
		{"a = 1; a = a-2.0; a", -1.0},
		{"a = 2; a = a-1.0; a", 1.0},
		{"a = 3; a = a-1.2; a", 1.8},

		{"a = 1; a = a*1.0; a", 1.0},
		{"a = 1; a = a*2.0; a", 2.0},
		{"a = 1; a = a*0.0; a", 0.0},
		{"a = 2; a = a*1.0; a", 2.0},
		{"a = 0; a = a*1.0; a", 0.0},
		{"a = 3; a = a*1.2; a", 3.5999999999999996},

		{"a = 1; a = a/1.0; a", 1.0},
		{"a = 1; a = a/2.0; a", 0.5},
		{"a = 2; a = a/1.0; a", 2.0},

		// division by zero
		{"1.0 / 0", "division by zero not allowed"},
		{"1.0 / 0.0", "division by zero not allowed"},
		{"1 / 0", "division by zero not allowed"},
		{"1 / 0.0", "division by zero not allowed"},
	}

	testInput(t, tests)
}

// TestNumberMathMethods covers all thirteen functions absorbed from Math onto
// both Integer and Float, registered from the numberMath table. Each is
// pinned on both receiver types -- the table exists because thirteen
// near-identical blocks is thirteen chances for one of them to call the wrong
// function, and a wiring mistake like "asin": math.Acos would only show up if
// every entry, not just six of them, is checked.
func TestNumberMathMethods(t *testing.T) {
	tests := []inputTestCase{
		{`9.sqrt()`, 3.0},
		{`9.0.sqrt()`, 3.0},
		{`1.exp()`, 2.718281828459045},
		{`1.0.exp()`, 2.718281828459045},
		{`4.log()`, 1.3862943611198906},
		{`4.0.log()`, 1.3862943611198906},
		{`8.log2()`, 3.0},
		{`8.0.log2()`, 3.0},
		{`100.log10()`, 2.0},
		{`100.0.log10()`, 2.0},
		{`1.sin()`, 0.8414709848078965},
		{`1.0.sin()`, 0.8414709848078965},
		{`1.cos()`, 0.5403023058681398},
		{`1.0.cos()`, 0.5403023058681398},
		{`1.tan()`, 1.557407724654902},
		{`1.0.tan()`, 1.557407724654902},
		{`1.asin()`, 1.5707963267948966},
		{`1.0.asin()`, 1.5707963267948966},
		{`0.acos()`, 1.5707963267948966},
		{`0.5.acos()`, 1.0471975511965976},
		{`1.atan()`, 0.7853981633974483},
		{`1.0.atan()`, 0.7853981633974483},

		// copysign(NUMERIC) and remainder(NUMERIC) are the two binary entries
		// in the table; each needs its own case because the unary cases above
		// cannot exercise the second argument at all.
		{`3.copysign(0 - 1)`, -3.0},
		{`3.0.copysign(0.0 - 1.0)`, -3.0},
		{`3.copysign(1)`, 3.0},
		{`100.remainder(30)`, 10.0},
		{`100.0.remainder(30.0)`, 10.0},

		// The ArgPattern restricts the second argument to NUMERIC, so a
		// non-number is rejected before either function runs.
		{`3.copysign(nil)`, "wrong argument type on position 1: got=NIL, want=NUMERIC"},
		{`3.0.copysign("x")`, "wrong argument type on position 1: got=STRING, want=NUMERIC"},
		{`100.remainder(nil)`, "wrong argument type on position 1: got=NIL, want=NUMERIC"},
		{`100.0.remainder("x")`, "wrong argument type on position 1: got=STRING, want=NUMERIC"},

		{`5.successor()`, 6},
		{`5.predecessor()`, 4},
		{`65.to_character()`, "A"},
	}

	testInput(t, tests)
}

// Math keeps constants and randomness only. Everything that operates on a
// number is a method on the number, so there is one place to look.
func TestMathKeepsOnlyConstantsAndRandom(t *testing.T) {
	for _, gone := range []string{"abs", "ceil", "floor", "round", "pow", "sqrt", "max", "min", "rand"} {
		evaluated := testEval("Math." + gone + "(1)")

		if !object.IsError(evaluated) {
			t.Errorf("Math.%s should be gone, got %s", gone, evaluated.Inspect())
		}
	}
}
