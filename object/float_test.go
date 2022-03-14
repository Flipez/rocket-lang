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
		{`2.1.plz_s()`, "2.1"},
		{`2.1.plz_f()`, 2.1},
		{`2.1.plz_i()`, 2},
		{`10.0.type()`, "FLOAT"},
		{`2.2.nope()`, "undefined method `.nope()` for FLOAT"},
		{`(2.0.wat().lines().size() == 2.0.methods().size() + 1).plz_s()`, "true"},
		{"1.1.to_json()", "1.1"},
		{"3.123456.to_json()", "3.123456"},
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
