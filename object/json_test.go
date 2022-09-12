package object_test

import (
	"testing"
)

func TestJSONObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`JSON.nope()`, "undefined method `.nope()` for BUILTIN_MODULE"},
		{`JSON.parse("{a}")`, "Error while parsing json: invalid character 'a' looking for beginning of object key string"},
		{`JSON.parse("{}").type()`, "HASH"},
		{`JSON.parse("[]").type()`, "ARRAY"},
		{`a = IO.open("../fixtures/data.json", "r", "0644").content(); JSON.parse(a)["string"]`, "Kadeem Sawyer"},
		{`a = IO.open("../fixtures/data.json", "r", "0644").content(); JSON.parse(a)["bool"]`, true},
		{`a = IO.open("../fixtures/data.json", "r", "0644").content(); JSON.parse(a)["float"]`, 12.3},
		{`a = IO.open("../fixtures/data.json", "r", "0644").content(); JSON.parse(a)["array"][1]`, 1234.0},
	}

	testInput(t, tests)
}
