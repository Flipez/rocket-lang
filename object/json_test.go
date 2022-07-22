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
		{`a = open("tests/data.json").content(); JSON.parse(a)`, `{"phone": 12.3, "postalZip": ["11137", 1234.0], "name": "Kadeem Sawyer"`},
	}

	testInput(t, tests)
}
