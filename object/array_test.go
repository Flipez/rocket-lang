package object_test

import (
	"testing"
)

func TestArrayObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3].size()`, 3},
	}

	testInput(t, tests)
}
