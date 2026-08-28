package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestEnvironment(t *testing.T) {
	parent := object.NewEnvironment()
	child := object.NewEnclosedEnvironment(parent)

	/*
		initial situation, both environments are empty
		parrent:
		child:
	*/
	v, ok := parent.Get("a")
	if v != nil {
		t.Errorf(`v, ok := parent.Get("a"): 'v' expected to be nil, but got %v`, v)
	}
	if ok {
		t.Errorf(`v, ok := parent.Get("a"): 'ok' expected to be false, but got %t`, ok)
	}

	v, ok = child.Get("a")
	if v != nil {
		t.Errorf(`v, ok := child.Get("a"): 'v' expected to be nil, but got %v`, v)
	}
	if ok {
		t.Errorf(`v, ok := child.Get("a"): 'ok' expected to be false, but got %t`, ok)
	}

	/*
		set new variable in parent env
		parent:
			a=1
		child:
	*/
	value := object.NewInteger(1)
	retObj := parent.Set("a", value)
	if retType := retObj.Type(); retType != object.INTEGER_OBJ {
		t.Errorf(`retObj := parent.Set("a", value): expected 'retObj' to be type INTEGER but got %s`, retType)
	}
	if retObj.Inspect() != value.Inspect() {
		t.Errorf(`retObj := parent.Set("a", value): expected 'retObj' and 'value' to be equal but got %s and %s`, retObj.Inspect(), value.Inspect())
	}

	v, ok = parent.Get("a")
	if vType := v.Type(); vType != object.INTEGER_OBJ {
		t.Errorf(`v, ok := parent.Get("a"): expected 'v' to be type INTEGER but got %s`, vType)
	}
	if v.Inspect() != value.Inspect() {
		t.Errorf(`v, ok := parent.Get("a"): expected 'v' and 'value' to be equal but got %s and %s`, v.Inspect(), value.Inspect())
	}
	if !ok {
		t.Errorf(`v, ok := parent.Get("a"): expected 'ok' to be true, but got %t`, ok)
	}

	v, ok = child.Get("a")
	if vType := v.Type(); vType != object.INTEGER_OBJ {
		t.Errorf(`v, ok := child.Get("a"): expected 'v' to be type INTEGER but got %s`, vType)
	}
	if v.Inspect() != value.Inspect() {
		t.Errorf(`v, ok := child.Get("a"): expected 'v' and 'value' to be equal but got %s and %s`, v.Inspect(), value.Inspect())
	}
	if !ok {
		t.Errorf(`v, ok := child.Get("a"): expected 'ok' to be true, but got %t`, ok)
	}

	/*
		update a variable that exists in parent env
		parent:
			a=2
		child:
	*/
	value = object.NewInteger(2)
	retObj = parent.Set("a", value)
	if retType := retObj.Type(); retType != object.INTEGER_OBJ {
		t.Errorf(`retObj := parent.Set("a", value): expected 'retObj' to be type INTEGER but got %s`, retType)
	}
	if retObj.Inspect() != value.Inspect() {
		t.Errorf(`retObj := parent.Set("a", value): expected 'retObj' and 'value' to be equal but got %s and %s`, retObj.Inspect(), value.Inspect())
	}

	v, ok = parent.Get("a")
	if vType := v.Type(); vType != object.INTEGER_OBJ {
		t.Errorf(`v, ok := parent.Get("a"): expected 'v' to be type INTEGER but got %s`, vType)
	}
	if v.Inspect() != value.Inspect() {
		t.Errorf(`v, ok := parent.Get("a"): expected 'v' and 'value' to be equal but got %s and %s`, v.Inspect(), value.Inspect())
	}
	if !ok {
		t.Errorf(`v, ok := parent.Get("a"): expected 'ok' to be true, but got %t`, ok)
	}

	v, ok = child.Get("a")
	if vType := v.Type(); vType != object.INTEGER_OBJ {
		t.Errorf(`v, ok := child.Get("a"): expected 'v' to be type INTEGER but got %s`, vType)
	}
	if v.Inspect() != value.Inspect() {
		t.Errorf(`v, ok := child.Get("a"): expected 'v' and 'value' to be equal but got %s and %s`, v.Inspect(), value.Inspect())
	}
	if !ok {
		t.Errorf(`v, ok := child.Get("a"): expected 'ok' to be true, but got %t`, ok)
	}

	/*
		set a new variable in child env
		parent:
			a=2
		child:
			b=3
	*/
	value = object.NewInteger(3)
	retObj = child.Set("b", value)
	if retType := retObj.Type(); retType != object.INTEGER_OBJ {
		t.Errorf(`retObj := child.Set("a", value): expected 'retObj' to be type INTEGER but got %s`, retType)
	}
	if retObj.Inspect() != value.Inspect() {
		t.Errorf(`retObj := child.Set("a", value): expected 'retObj' and 'value' to be equal but got %s and %s`, retObj.Inspect(), value.Inspect())
	}

	v, ok = parent.Get("b")
	if v != nil {
		t.Errorf(`v, ok := parent.Get("b"): 'v' expected to be nil, but got %v`, v)
	}
	if ok {
		t.Errorf(`v, ok := parent.Get("b"): 'ok' expected to be false, but got %t`, ok)
	}

	v, ok = child.Get("b")
	if vType := v.Type(); vType != object.INTEGER_OBJ {
		t.Errorf(`v, ok := child.Get("b"): expected 'v' to be type INTEGER but got %s`, vType)
	}
	if v.Inspect() != value.Inspect() {
		t.Errorf(`v, ok := child.Get("b"): expected 'v' and 'value' to be equal but got %s and %s`, v.Inspect(), value.Inspect())
	}
	if !ok {
		t.Errorf(`v, ok := child.Get("b"): expected 'ok' to be true, but got %t`, ok)
	}
}

// TestRebindAllowedDefaultsToFalse proves a fresh environment does not allow
// import rebinding until AllowRebind is explicitly called.
func TestRebindAllowedDefaultsToFalse(t *testing.T) {
	env := object.NewEnvironment()

	if env.RebindAllowed() {
		t.Errorf("expected RebindAllowed() to be false by default, got true")
	}
}

// TestAllowRebindEnablesRebindAllowed proves AllowRebind flips
// RebindAllowed() to true on the environment it's called on.
func TestAllowRebindEnablesRebindAllowed(t *testing.T) {
	env := object.NewEnvironment()
	env.AllowRebind()

	if !env.RebindAllowed() {
		t.Errorf("expected RebindAllowed() to be true after AllowRebind(), got false")
	}
}

// TestRebindAllowedInheritedByEnclosedEnvironment proves RebindAllowed
// delegates to outer the same way Registry() does, so an enclosed
// environment (e.g. a loop body or block) inherits the relaxation enabled
// on its outer/top-level environment.
func TestRebindAllowedInheritedByEnclosedEnvironment(t *testing.T) {
	outer := object.NewEnvironment()
	outer.AllowRebind()
	child := object.NewEnclosedEnvironment(outer)

	if !child.RebindAllowed() {
		t.Errorf("expected enclosed environment to inherit RebindAllowed() from outer, got false")
	}

	// An enclosed environment created from an outer environment that never
	// called AllowRebind must not report true on its own.
	other := object.NewEnclosedEnvironment(object.NewEnvironment())
	if other.RebindAllowed() {
		t.Errorf("expected enclosed environment to report false when outer never called AllowRebind(), got true")
	}
}
