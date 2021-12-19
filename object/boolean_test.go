package object_test

import (
	"testing"
)

func TestBooleanObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`true.plz_s()`, "true"},
		{`false.plz_s()`, "false"},
	}

	testInput(t, tests)
}
