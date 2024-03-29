package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func testIntegerObject(t *testing.T, obj object.Object, expected int) bool {
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
		{`2.to_s()`, "2"},
		{`10.to_s(2)`, "1010"},
		{`2.to_f()`, 2.0},
		{`2.to_i()`, 2},
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
