package object

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/flipez/rocket-lang/ast"
)

type ObjectType string

var Evaluator func(node ast.Node, env *Environment) Object

type Object interface {
	Type() ObjectType
	Inspect() string
	InvokeMethod(method string, env Environment, args ...Object) Object
}

type Iterable interface {
	Reset()
	Next() (Object, Object, bool)
}

type Hashable interface {
	HashKey() HashKey
}

type Serializable interface {
	MarshalJSON() ([]byte, error)
}

const (
	INTEGER_OBJ          = "INTEGER"
	FLOAT_OBJ            = "FLOAT"
	BOOLEAN_OBJ          = "BOOLEAN"
	NIL_OBJ              = "NIL"
	RETURN_VALUE_OBJ     = "RETURN_VALUE"
	BREAK_VALUE_OBJ      = "BREAK_VALUE"
	NEXT_VALUE_OBJ       = "NEXT_VALUE"
	ERROR_OBJ            = "ERROR"
	FUNCTION_OBJ         = "FUNCTION"
	STRING_OBJ           = "STRING"
	ARRAY_OBJ            = "ARRAY"
	HASH_OBJ             = "HASH"
	FILE_OBJ             = "FILE"
	MODULE_OBJ           = "MODULE"
	HTTP_OBJ             = "HTTP"
	JSON_OBJ             = "JSON"
	BUILTIN_MODULE_OBJ   = "BUILTIN_MODULE"
	BUILTIN_FUNCTION_OBJ = "BUILTIN_FUNCTION"
	BUILTIN_PROPERTY_OBJ = "BUILTIN_PROPERTY"
)

type ObjectMethod struct {
	argsOptional   bool
	argOverloading bool
	argPattern     [][]string
	returnPattern  [][]string
	description    string
	example        string
	method         func(Object, []Object, Environment) Object
}

func (om ObjectMethod) validateArgs(args []Object) error {
	if (len(args) < len(om.argPattern)) && !om.argsOptional {
		return fmt.Errorf("to few arguments: want=%d, got=%d", len(om.argPattern), len(args))
	}

	if len(args) > len(om.argPattern) && !om.argOverloading {
		return fmt.Errorf("to many arguments: want=%d, got=%d", len(om.argPattern), len(args))
	}

	if !om.argsOptional || (om.argsOptional && len(args) > 0) {
		for idx, pattern := range om.argPattern {
			var valid bool
			for _, argType := range pattern {
				if ObjectType(argType) == args[idx].Type() {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("wrong argument type on position %d: got=%s, want=%s", idx, args[idx].Type(), strings.Join(pattern, "|"))
			}
		}
	}

	/* not needed for now, as there is no string method with flexible argument count
	if om.argOverloading {
		lastPattern := om.argPattern[len(om.argPattern)-1]
		for idx, arg := range args[len(om.argPattern)-1:] {
			var valid bool
			for _, argType := range lastPattern {
				if argType == args.Type() {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("FALSCHER TYPE AUF POSITION %d got=%s, want=%s", idx, arg.Type(), strings.Join(lastPattern, "|"))
			}
		}
	}
	*/

	return nil
}

func (om ObjectMethod) ReturnPattern() string {
	types := make([]string, len(om.returnPattern))
	for idx, pattern := range om.returnPattern {
		types[idx] = strings.Join(pattern, "|")
	}
	return strings.Join(types, ", ")
}

func (om ObjectMethod) Description() string {
	return om.description
}

func (om ObjectMethod) Example() string {
	return om.example
}

func (om ObjectMethod) Usage(name string) string {
	var args string

	if len(om.argPattern) > 0 {
		types := make([]string, len(om.argPattern))
		for idx, pattern := range om.argPattern {
			types[idx] = strings.Join(pattern, "|")
		}
		args = strings.Join(types, ", ")

		if om.argOverloading {
			args += "..."
		}
	}

	return fmt.Sprintf("%s(%s)", name, args)
}

func (om ObjectMethod) Call(o Object, args []Object, env Environment) Object {
	if err := om.validateArgs(args); err != nil {
		return NewError(err)
	}
	return om.method(o, args, env)
}

var objectMethods = make(map[ObjectType]map[string]ObjectMethod)

func ListObjectMethods() map[ObjectType]map[string]ObjectMethod {
	return objectMethods
}

func init() {
	objectMethods["*"] = map[string]ObjectMethod{
		"to_json": ObjectMethod{
			description: "Returns the object as json notation.",
			example: `🚀 > a = {"test": 1234}
=> {"test": 1234}
🚀 > a.to_json()
=> "{"test":1234}"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ, ERROR_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if serializeableObject, ok := o.(Serializable); ok {
					j, err := json.Marshal(serializeableObject)
					if err != nil {
						return NewErrorFormat("Error while marshal value: %s", err.Error())
					}
					return NewString(string(j))
				}

				return NewErrorFormat("%s is not serializable", o.Type())
			},
		},
		"methods": ObjectMethod{
			description: "Returns an array of all supported methods names.",
			example: `🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				oms := objectMethods[o.Type()]
				result := make([]Object, len(oms))
				var i int
				for name := range oms {
					result[i] = NewString(name)
					i++
				}
				return NewArray(result)
			},
		},
		"wat": ObjectMethod{
			description: "Returns the supported methods with usage information.",
			example: `🚀 > true.wat()
=> BOOLEAN supports the following methods:
				plz_s()`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				oms := objectMethods[o.Type()]
				result := make([]string, len(oms))
				var i int
				for name, objectMethod := range oms {
					result[i] = fmt.Sprintf("\t%s", objectMethod.Usage(name))
					i++
				}
				return NewString(fmt.Sprintf("%s supports the following methods:\n%s", o.Type(), strings.Join(result, "\n")))
			},
		},
		"type": ObjectMethod{
			description: "Returns the type of the object.",
			example: `🚀 > "test".type()
=> "STRING"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewString(string(o.Type()))
			},
		},
	}
}

func objectMethodLookup(o Object, method string, env Environment, args []Object) Object {
	if oms, ok := objectMethods[o.Type()]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args, env)
		}
	}

	if oms, ok := objectMethods["*"]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args, env)
		}
	}

	return nil
}

func CompareObjects(ao, bo Object) bool {
	switch ao.Type() {
	case NIL_OBJ:
		return bo.Type() == NIL_OBJ
	case INTEGER_OBJ:
		if b, ok := bo.(*Integer); ok {
			return ao.(*Integer).Value == b.Value
		}
		return false
	case FLOAT_OBJ:
		if b, ok := bo.(*Float); ok {
			return ao.(*Float).Value == b.Value
		}
		return false
	case BOOLEAN_OBJ:
		if b, ok := bo.(*Boolean); ok {
			return ao.(*Boolean).Value == b.Value
		}
		return false
	case ERROR_OBJ:
		if b, ok := bo.(*Error); ok {
			return ao.(*Error).Message == b.Message
		}
		return false
	case STRING_OBJ:
		if b, ok := bo.(*String); ok {
			return ao.(*String).Value == b.Value
		}
		return false
	case ARRAY_OBJ:
		if b, ok := bo.(*Array); ok {
			a, _ := ao.(*Array)

			if len(a.Elements) != len(b.Elements) {
				return false
			}

			for idx, element := range a.Elements {
				if !CompareObjects(element, b.Elements[idx]) {
					return false
				}
			}

			return true
		}
		return false
	case HASH_OBJ:
		if b, ok := bo.(*Hash); ok {
			a, _ := ao.(*Hash)

			if len(a.Pairs) != len(b.Pairs) {
				return false
			}

			for aKey, aPair := range a.Pairs {
				bPair, ok := b.Pairs[aKey]
				if !ok {
					return false
				}
				if !CompareObjects(aPair.Key, bPair.Key) {
					return false
				}
				if !CompareObjects(aPair.Value, bPair.Value) {
					return false
				}
			}

			return true
		}
		return false
	}

	return false
}

func IsError(o Object) bool {
	return o != nil && o.Type() == ERROR_OBJ
}

func IsNumber(o Object) bool {
	return o != nil && (o.Type() == INTEGER_OBJ || o.Type() == FLOAT_OBJ)
}

func IsTruthy(o Object) bool {
	switch o {
	case NIL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func IsFalsy(o Object) bool {
	return !IsTruthy(o)
}

func AddEvaluator(e func(node ast.Node, env *Environment) Object) {
	Evaluator = e
}
