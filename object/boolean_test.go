package object_test

import (
	"github.com/flipez/rocket-lang/object"
	"testing"
)

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%T, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestBooleanObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`true.plz_s()`, "true"},
		{`false.plz_s()`, "false"},
		{`false.type()`, "BOOLEAN"},
		{`false.nope()`, "Failed to invoke method: nope"},
		{`(true.wat().lines().size() == true.methods().size() + 1).plz_s()`, "true"},
	}

	testInput(t, tests)
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &object.Boolean{Value: true}
	true2 := &object.Boolean{Value: true}
	false1 := &object.Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestBooleanInspect(t *testing.T) {
	true1 := &object.Boolean{Value: true}
	false1 := &object.Boolean{Value: false}

	if true1.Inspect() != "true" {
		t.Errorf("boolean inspect does not match value")
	}

	if false1.Inspect() != "false" {
		t.Errorf("boolean inspect does not match value")
	}
}
