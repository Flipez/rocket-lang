package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}

	return true
}

func TestFloatObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`2.1.to_s()`, "2.1"},
		{`2.1.to_f()`, 2.1},
		{`2.1.to_i()`, 2},
		{`10.0.type()`, "FLOAT"},
		{`2.2.nope()`, "test:1:4: undefined method `.nope()` for FLOAT"},
		{"1.1.to_json()", "1.1"},
		{"3.123456.to_json()", "3.123456"},
		// round, ceil, floor and abs preserve the receiver's type, so a
		// FLOAT stays a FLOAT rather than collapsing to an INTEGER.
		{`3.14.round()`, 3.0},
		{`3.14.ceil()`, 4.0},
		{`3.14.floor()`, 3.0},
		{`3.14.abs()`, 3.14},
		{`(0.0 - 3.14).abs()`, 3.14},
		{`0.0.abs()`, 0.0},
		// math.Round rounds halves away from zero.
		{`2.5.round()`, 3.0},
		{`(0.0 - 2.5).round()`, -3.0},
		{`(0.0 - 3.7).ceil()`, -3.0},
		{`(0.0 - 3.2).floor()`, -4.0},
		// The method form agrees with the Math module form.
		{`3.14.round().to_s() == Math.round(3.14).to_s()`, true},
	}

	testInput(t, tests)
}

func TestFloatHashKey(t *testing.T) {
	float1_1 := object.NewFloat(1.0)
	float1_2 := object.NewFloat(1.0)
	float2 := object.NewFloat(2.0)

	if float1_1.HashKey() != float1_2.HashKey() {
		t.Errorf("float with same content have different hash keys")
	}

	if float1_1.HashKey() == float2.HashKey() {
		t.Errorf("float with different content have same hash keys")
	}
}

func TestFloatInspect(t *testing.T) {
	float1 := object.NewFloat(1.0)

	if float1.Inspect() != "1.0" {
		t.Errorf("float inspect does not match value")
	}
}
