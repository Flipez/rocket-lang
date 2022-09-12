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

			return interfaceToObject(i)
		},
	)
}

func interfaceToObject(i interface{}) object.Object {
	switch v := i.(type) {
	case map[string]interface{}:
		jsonObject := object.NewHash(nil)
		for key, val := range v {
			hp := object.HashPair{
				Key:   object.NewString(key),
				Value: interfaceToObject(val),
			}

			jsonObject.Pairs[hp.Key.(object.Hashable).HashKey()] = hp
		}

		return jsonObject
	case []interface{}:
		jsonArray := object.NewArray(nil)
		for _, element := range v {
			jsonArray.Elements = append(jsonArray.Elements, interfaceToObject(element))
		}
		return jsonArray
	case string:
		return object.NewString(v)
	case float64:
		return object.NewFloat(v)
	case bool:
		if v {
			return object.TRUE
		}
		return object.FALSE
	}
	return object.NIL
}
