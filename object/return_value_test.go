package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestReturnValue(t *testing.T) {
	rv := object.NewReturnValue(object.NewString("a"))

	if rv.Type() != object.RETURN_VALUE_OBJ {
		t.Errorf("returnValue.Type() returns wrong type")
	}
	if rv.Inspect() != `"a"` {
		t.Errorf("returnValue.Inspect() returns wrong type")
	}
}
