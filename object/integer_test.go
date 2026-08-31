package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func testIntegerObject(t *testing.T, obj object.Object, expected int) bool {
	t.Helper()
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func TestIntegerObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`3.abs()`, 3},
		{`(0 - 5).abs()`, 5},
		{`0.abs()`, 0},
		// abs keeps the integer's base.
		{`"-0x10".to_integer().abs().base()`, 16},
		{`"0x10".to_integer().abs().to_string()`, "0x10"},
		{`2.to_string()`, "2"},
		{`2.to_float()`, 2.0},
		{`2.to_integer()`, 2},
		{`10.type()`, "INTEGER"},
		{`2.nope()`, "test:1:2: undefined method `.nope()` for INTEGER"},
		{"1.to_json()", "1"},
	}

	testInput(t, tests)
}

func TestIntegerHashKey(t *testing.T) {
	int1_1 := object.NewInteger(1)
	int1_2 := object.NewInteger(1)
	int2 := object.NewInteger(2)

	if int1_1.HashKey() != int1_2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if int1_1.HashKey() == int2.HashKey() {
		t.Errorf("integer with different content have same hash keys")
	}
}

func TestIntegerInspect(t *testing.T) {
	int1 := object.NewInteger(1)

	if int1.Inspect() != "1" {
		t.Errorf("integer inspect does not match value")
	}
}

func TestIntegerIteratable(t *testing.T) {
	int1 := object.NewInteger(3)
	int1Iterator := int1.GetIterator(0, 1, false)

	for expected := 0; expected < 3; expected++ {
		_, value, ok := int1Iterator.Next()
		actual := value.(*object.Integer)

		if !ok {
			t.Errorf("integer iteration finished too early")
		}

		if actual.Value != expected {
			t.Errorf(
				"integer next %d does not match value %d",
				actual.Value,
				expected,
			)
		}
	}

	_, _, ok := int1Iterator.Next()
	if ok {
		t.Errorf("integer iteration didn't finish")
	}

	int1Iterator = int1.GetIterator(0, 1, false)
	_, _, ok = int1Iterator.Next()
	if !ok {
		t.Errorf("new integer iteration shouldn't finish after first next")
	}
}

// TestIntegerRubyMethods covers the methods added to close the gap with Ruby's
// Integer. The expectations come from the example lines in
// https://ruby-doc.org/3.4.1/Integer.html, except where noted.
func TestIntegerRubyMethods(t *testing.T) {
	tests := []inputTestCase{
		{`12345.digits()`, "[5, 4, 3, 2, 1]"},
		{`12345.digits(7)`, "[4, 6, 6, 0, 5]"},
		{`0.digits()`, "[0]"},
		{`(0 - 1).digits()`, "digits of a negative number are not defined"},
		{`5.digits(1)`, "base has to be 2 or greater, got 1"},

		{`36.gcd(60)`, 12},
		{`3.gcd(0 - 7)`, 1},
		{`0.gcd(0)`, 0},
		{`36.lcm(60)`, 180},
		{`3.lcm(0 - 7)`, 21},
		{`0.lcm(2)`, 0},

		{`2.pow(3)`, 8},
		{`2.pow(3, 5)`, 3},
		{`2.pow(0)`, 1},
		{`2.pow(0 - 1)`, "negative exponent is not supported"},
		{`2.pow(3, 0)`, "modulus of zero not allowed"},

		{`0.bit_length()`, 0},
		{`255.bit_length()`, 8},
		{`256.bit_length()`, 9},
		// A negative number needs as many bits as its complement.
		{`(0 - 1).bit_length()`, 0},
		{`(0 - 256).bit_length()`, 8},

		{`1.succ()`, 2},
		{`1.pred()`, 0},
		{`(0 - 1).pred()`, -2},

		{`4.even?()`, true},
		{`5.even?()`, false},
		{`5.odd?()`, true},
		{`(0 - 3).odd?()`, true},
		{`(0 - 3).even?()`, false},
		{`0.zero?()`, true},
		{`1.positive?()`, true},
		{`(0 - 1).negative?()`, true},
		{`0.positive?()`, false},

		{`65.chr()`, "A"},
		{`(0 - 1).chr()`, "-1 is out of the range of a character"},

		// Rounding an integer only does something with a negative digit count.
		{`555.ceil(0 - 1)`, 560},
		{`555.floor(0 - 1)`, 550},
		{`555.round(0 - 1)`, 560},
		{`555.truncate(0 - 1)`, 550},
		{`555.truncate(0 - 2)`, 500},
		{`555.ceil()`, 555},
		{`555.round(2)`, 555},
		{`(0 - 555).floor(0 - 1)`, -560},
		{`(0 - 555).ceil(0 - 1)`, -550},
		{`(0 - 555).round(0 - 1)`, -560},
		{`(0 - 555).truncate(0 - 1)`, -550},
		// A value already on the boundary is left alone rather than pushed to
		// the next multiple.
		{`550.ceil(0 - 1)`, 550},
		{`550.floor(0 - 1)`, 550},

		// divmod follows this language's / and %, which truncate toward zero.
		// Ruby floors, so it answers [-3, -1] for the third case.
		{`11.divmod(4)`, "[2, 3]"},
		{`11.divmod(4).first()`, 2},
		{`11.divmod(0 - 4)`, "[-2, 3]"},
		{`11.divmod(0)`, "division by zero not allowed"},

		// An arithmetic method rejects a mixed base for the same reason 0x10 + 4
		// does.
		{`"0x10".to_integer().gcd(4)`, "infix operation with unequal base not allowed"},
		{`"0x10".to_integer().succ().base()`, 16},
		{`"0b101".to_integer().pred().base()`, 2},
	}
	testInput(t, tests)
}

// TestIntegerCallbackMethods covers times, upto and downto. All three hand back
// the receiver so a walk chains, the way Array#each does.
func TestIntegerCallbackMethods(t *testing.T) {
	tests := []inputTestCase{
		// times counts 0 to n-1, so 3.times sees 0, 1, 2.
		{`out = []; 3.times(def(i) out.append(i) end); out.to_json()`, "[0,1,2]"},
		{`3.times(def(i) end)`, 3},
		{`3.times(def(i) end).type()`, "INTEGER"},
		// A count of zero or less calls nothing rather than erroring, which is
		// what makes it safe to hand a computed count.
		{`out = []; 0.times(def(i) out.append(i) end); out.to_json()`, "[]"},
		{`out = []; (0 - 3).times(def(i) out.append(i) end); out.to_json()`, "[]"},

		// upto and downto are inclusive at both ends.
		{`out = []; 1.upto(3, def(i) out.append(i) end); out.to_json()`, "[1,2,3]"},
		{`out = []; 3.downto(1, def(i) out.append(i) end); out.to_json()`, "[3,2,1]"},
		{`out = []; 2.upto(2, def(i) out.append(i) end); out.to_json()`, "[2]"},
		// A limit on the wrong side calls nothing instead of running away.
		{`out = []; 3.upto(1, def(i) out.append(i) end); out.to_json()`, "[]"},
		{`out = []; 1.downto(3, def(i) out.append(i) end); out.to_json()`, "[]"},
		{`1.upto(3, def(i) end)`, 1},

		{`out = []; 9.times(def(i) if i == 3 break end out.append(i) end); out.to_json()`, "[0,1,2]"},
		{`out = []; 1.upto(9, def(i) if i == 3 break end out.append(i) end); out.to_json()`, "[1,2]"},

		// The counter keeps the receiver's base, and upto refuses a limit of
		// another base for the same reason the infix operators do.
		{`"0x10".to_integer().times(def(i) end).base()`, 16},
		{`out = []; "0b10".to_integer().times(def(i) out.append(i.base()) end); out.to_json()`, "[2,2]"},
		{`"0x10".to_integer().upto(4, def(i) end)`, "infix operation with unequal base not allowed"},

		{`3.times(def(i) i.nope() end)`, "test:1:17: undefined method `.nope()` for INTEGER"},
		{`3.times(def(i, j) end)`, "too few arguments: got=1, want=2"},
		{`3.times(1)`, "wrong argument type on position 1: got=INTEGER, want=CALLABLE"},
		{`1.upto(3)`, "too few arguments: got=1, want=2"},
	}
	testInput(t, tests)
}
