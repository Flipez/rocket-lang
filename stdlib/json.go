package stdlib

import (
	"encoding/json"

	"github.com/flipez/rocket-lang/object"
)

var jsonFunctions = map[string]*object.BuiltinFunction{}
var jsonProperties = map[string]*object.BuiltinProperty{}

func init() {
	jsonFunctions["parse"] = object.NewBuiltinFunction(
		"parse",
		object.MethodLayout{
			Description: "Takes a STRING and parses it to a HASH or ARRAY. Numbers are always FLOAT.",
			Example: `🚀 > JSON.parse('{"test": 123}')
=> {"test": 123.0}
🚀 > JSON.parse('["test", 123]')
=> ["test", 123.0]`,
			ReturnPattern: object.Args(
				object.Arg(object.HASH_OBJ),
			),
			ArgPattern: object.Args(
				object.Arg(object.STRING_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			var i interface{}
			input := args[0].(*object.String).Value

			err := json.Unmarshal([]byte(input), &i)

			if err != nil {
				return object.NewErrorFormat("Error while parsing json: %s", err)
			}

			return object.AnyToObject(i)
		},
	)
}
