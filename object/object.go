package object

import (
	"bytes"
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
