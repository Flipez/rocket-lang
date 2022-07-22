package object

import (
	"encoding/json"
)

type JSON struct{}

func (j *JSON) Type() ObjectType { return JSON_OBJ }
func (j *JSON) Inspect() string  { return "JSON" }
func (j *JSON) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(j, method, env, args)
}

func init() {
	objectMethods[JSON_OBJ] = map[string]ObjectMethod{
		"parse": ObjectMethod{
			returnPattern: [][]string{
				[]string{HASH_OBJ},
			},
			argPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(_ Object, args []Object, _ Environment) Object {
				var i interface{}
				input := args[0].(*String).Value

				err := json.Unmarshal([]byte(input), &i)

				if err != nil {
					return NewErrorFormat("Error while parsing json: %s", err)
				}

				return interfaceToObject(i)
			},
		},
	}
}

func interfaceToObject(i interface{}) Object {
	switch v := i.(type) {
	case map[string]interface{}:
		jsonObject := NewHash(nil)
		for key, val := range v {
			hp := HashPair{
				Key:   NewString(key),
				Value: interfaceToObject(val),
			}

			jsonObject.Pairs[hp.Key.(Hashable).HashKey()] = hp
		}

		return jsonObject
	case []interface{}:
		jsonArray := NewArray(nil)
		for _, element := range v {
			jsonArray.Elements = append(jsonArray.Elements, interfaceToObject(element))
		}
		return jsonArray
	case string:
		return NewString(v)
	case float64:
		return NewFloat(v)
	case bool:
		if v {
			return TRUE
		}
		return FALSE
	}
	return NIL
}
