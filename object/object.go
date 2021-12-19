package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/flipez/rocket-lang/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
	InvokeMethod(method string, env Environment, args ...Object) Object
}

const (
	INTEGER_OBJ      = "INTEGER"
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
)

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType                                                   { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string                                                    { return "builtin function" }
func (b *Builtin) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (f *Function) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType                                                   { return ERROR_OBJ }
func (e *Error) Inspect() string                                                    { return "ERROR: " + e.Message }
func (e *Error) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Null struct{}

func (n *Null) Type() ObjectType                                                   { return NULL_OBJ }
func (n *Null) Inspect() string                                                    { return "null" }
func (n *Null) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}

type ObjectMethod struct {
	argsOptional   bool
	argOverloading bool
	argPattern     [][]string
	method         func(Object, []Object) Object
}

func (om *ObjectMethod) validateArgs(args []Object) error {
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

func (om *ObjectMethod) Usage(name string) string {
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

func (om *ObjectMethod) Call(o Object, args []Object) Object {
	if err := om.validateArgs(args); err != nil {
		return &Error{Message: err.Error()}
	}
	return om.method(o, args)
}

var objectMethods = make(map[ObjectType]map[string]ObjectMethod)

func init() {
	objectMethods["*"] = map[string]ObjectMethod{
		"methods": ObjectMethod{
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
			method: func(o Object, _ []Object) Object {
				return &String{Value: string(o.Type())}
			},
		},
	}
}

func objectMethodLookop(o Object, method string, args []Object) Object {
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

type Iterable interface {
	Reset()
	Next() (Object, Object, bool)
}
