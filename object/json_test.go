package object_test

import (
	"testing"
)

func TestJSONObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`JSON.nope()`, "undefined method `.nope()` for JSON"},
		{`JSON.parse("{a}")`, "Error while parsing json: invalid character 'a' looking for beginning of object key string"},
		{`JSON.parse("{}").type()`, "HASH"},
		{`JSON.parse("[]").type()`, "ARRAY"},
	}

	testInput(t, tests)
}
