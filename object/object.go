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
	GetIterator() Iterator
}

type Iterator interface {
	Next() (Object, Object, bool)
}

type Hashable interface {
	HashKey() HashKey
}

type Serializable interface {
	MarshalJSON() ([]byte, error)
}

var ANY_OBJ = []string{
	INTEGER_OBJ,
	STRING_OBJ,
	BOOLEAN_OBJ,
	ARRAY_OBJ,
	HASH_OBJ,
	FLOAT_OBJ,
	ERROR_OBJ,
	NIL_OBJ,
}

var NUMBER_OBJ = []string{
	INTEGER_OBJ,
	FLOAT_OBJ,
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
	BUILTIN_MODULE_OBJ   = "BUILTIN_MODULE"
	BUILTIN_FUNCTION_OBJ = "BUILTIN_FUNCTION"
	BUILTIN_PROPERTY_OBJ = "BUILTIN_PROPERTY"
)

type Argument struct {
	Types    []string
	Optional bool
	Overload bool
}

func (a Argument) String() string {
	return strings.Join(a.Types, "|")
}

func (a Argument) Check(o Object) bool {
	for _, t := range a.Types {
		if ObjectType(t) == o.Type() {
			return true
		}
	}
	return false
}

func Arg(types ...string) Argument {
	return Argument{Types: types}
}

func OptArg(types ...string) Argument {
	return Argument{Types: types, Optional: true}
}

func OverloadArg(types ...string) Argument {
	return Argument{Types: types, Overload: true}
}

func Args(args ...Argument) []Argument {
	return args
}

type MethodLayout struct {
	ArgPattern    []Argument
	ReturnPattern []Argument
	Description   string
	Example       string
}

func (ml MethodLayout) requiredArgs() []Argument {
	args := make([]Argument, 0)
	for _, arg := range ml.ArgPattern {
		if arg.Optional {
			continue
		}
		args = append(args, arg)
	}
	return args
}

type ObjectMethod struct {
	Layout MethodLayout
	method func(Object, []Object, Environment) Object
}

func (ml MethodLayout) validateArgs(args []Object) error {
	requiredArgs := ml.requiredArgs()
	requiredArgsLen := len(requiredArgs)
	possibleArgsLen := len(ml.ArgPattern)
	givenArgsLen := len(args)
	overloadArgs := possibleArgsLen > 0 && ml.ArgPattern[possibleArgsLen-1].Overload

	if givenArgsLen < requiredArgsLen {
		return fmt.Errorf("to few arguments: got=%d, want=%d", givenArgsLen, requiredArgsLen)
	}

	if (givenArgsLen > possibleArgsLen) && !overloadArgs {
		return fmt.Errorf("to many arguments: got=%d, want=%d", givenArgsLen, possibleArgsLen)
	}

	if givenArgsLen > 0 {
		for idx, arg := range args {
			var argToCheck Argument
			if (idx >= possibleArgsLen) && overloadArgs {
				argToCheck = ml.ArgPattern[possibleArgsLen-1]
			} else {
				argToCheck = ml.ArgPattern[idx]
			}

			if !argToCheck.Check(arg) {
				return fmt.Errorf("wrong argument type on position %d: got=%s, want=%s", idx+1, arg.Type(), argToCheck)
			}
		}
	}

	return nil
}

func (ml MethodLayout) DocsReturnPattern() string {
	types := make([]string, len(ml.ReturnPattern))
	for idx, pattern := range ml.ReturnPattern {
		types[idx] = strings.Join(pattern.Types, "|")
	}
	return strings.Join(types, ", ")
}

func (ml MethodLayout) Usage(name string) string {
	var args string

	if len(ml.ArgPattern) > 0 {
		types := make([]string, len(ml.ArgPattern))
		for idx, pattern := range ml.ArgPattern {
			types[idx] = strings.Join(pattern.Types, "|")
		}
		args = strings.Join(types, ", ")
	}

	return fmt.Sprintf("%s(%s)", name, args)
}

func (om ObjectMethod) Call(o Object, args []Object, env Environment) Object {
	if err := om.Layout.validateArgs(args); err != nil {
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
			Layout: MethodLayout{
				Description: "Returns the object as json notation.",
				Example: `🚀 > a = {"test": 1234}
=> {"test": 1234}
🚀 > a.to_json()
=> "{"test":1234}"`,
				ReturnPattern: Args(
					Arg(STRING_OBJ, ERROR_OBJ),
				),
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
			Layout: MethodLayout{
				Description: "Returns an array of all supported methods names.",
				Example: `🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]`,
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
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
			Layout: MethodLayout{
				Description: "Returns the supported methods with usage information.",
				Example: `🚀 > true.wat()
=> BOOLEAN supports the following methods:
				plz_s()`,
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				oms := objectMethods[o.Type()]
				result := make([]string, len(oms))
				var i int
				for name, objectMethod := range oms {
					result[i] = fmt.Sprintf("\t%s", objectMethod.Layout.Usage(name))
					i++
				}
				return NewString(fmt.Sprintf("%s supports the following methods:\n%s", o.Type(), strings.Join(result, "\n")))
			},
		},
		"type": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns the type of the object.",
				Example: `🚀 > "test".type()
=> "STRING"`,
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
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

func AnyToObject(a any) Object {
	switch v := a.(type) {
	case bool:
		if v {
			return TRUE
		}
		return FALSE
	case string:
		return NewString(v)
	case int:
		return NewInteger(int64(v))
	case int64:
		return NewInteger(v)
	case float64:
		return NewFloat(v)
	case []any:
		arr := make([]Object, len(v))
		for idx, item := range v {
			arr[idx] = AnyToObject(item)
		}
		return NewArray(arr)
	case map[string]any:
		hash := NewHash(nil)
		for key, val := range v {
			hash.Set(AnyToObject(key), AnyToObject(val))
		}
		return hash
	}
	return NIL
}
