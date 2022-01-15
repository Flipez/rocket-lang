package object

import (
	"fmt"
	"strings"
)

type ObjectType string

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

const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	FILE_OBJ         = "FILE"
	MODULE_OBJ       = "MODULE"
)

type ObjectMethod struct {
	argsOptional   bool
	argOverloading bool
	argPattern     [][]string
	returnPattern  [][]string
	description    string
	example        string
	method         func(Object, []Object) Object
}

func (om ObjectMethod) validateArgs(args []Object) error {
	if (len(args) < len(om.argPattern)) && !om.argsOptional {
		return fmt.Errorf("To few arguments: want=%d, got=%d", len(om.argPattern), len(args))
	}

	if len(args) > len(om.argPattern) && !om.argOverloading {
		return fmt.Errorf("To many arguments: want=%d, got=%d", len(om.argPattern), len(args))
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
				return fmt.Errorf("Wrong argument type on position %d: got=%s, want=%s", idx, args[idx].Type(), strings.Join(pattern, "|"))
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

func (om ObjectMethod) Call(o Object, args []Object) Object {
	if err := om.validateArgs(args); err != nil {
		return &Error{Message: err.Error()}
	}
	return om.method(o, args)
}

var objectMethods = make(map[ObjectType]map[string]ObjectMethod)

func ListObjectMethods() map[ObjectType]map[string]ObjectMethod {
	return objectMethods
}

func init() {
	objectMethods["*"] = map[string]ObjectMethod{
		"methods": ObjectMethod{
			description: "Returns an array of all supported methods names.",
			example: `🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				oms := objectMethods[o.Type()]
				result := make([]Object, len(oms), len(oms))
				var i int
				for name := range oms {
					result[i] = &String{Value: name}
					i++
				}
				return &Array{Elements: result}
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
			method: func(o Object, _ []Object) Object {
				oms := objectMethods[o.Type()]
				result := make([]string, len(oms), len(oms))
				var i int
				for name, objectMethod := range oms {
					result[i] = fmt.Sprintf("\t%s", objectMethod.Usage(name))
					i++
				}
				return &String{Value: fmt.Sprintf("%s supports the following methods:\n%s", o.Type(), strings.Join(result, "\n"))}
			},
		},
		"type": ObjectMethod{
			description: "Returns the type of the object.",
			example: `🚀 > "test".type()
=> "STRING"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				return &String{Value: string(o.Type())}
			},
		},
	}
}

func objectMethodLookup(o Object, method string, args []Object) Object {
	if oms, ok := objectMethods[o.Type()]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args)
		}
	}

	if oms, ok := objectMethods["*"]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args)
		}
	}

	return nil
}

func CompareObjects(ao, bo Object) bool {
	switch ao.Type() {
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

func IsNumber(o Object) bool {
	return o.Type() == INTEGER_OBJ || o.Type() == FLOAT_OBJ
}
