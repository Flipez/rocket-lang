package object

import (
	"encoding/json"
	"fmt"
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
				m := map[string]interface{}{}
				input := args[0].(*String).Value

				err := json.Unmarshal([]byte(input), &m)
				if err != nil {
					return NewErrorFormat("Error while parsing json: %s", err)
				}
				fmt.Println(m)

				return NIL
			},
		},
	}
}
