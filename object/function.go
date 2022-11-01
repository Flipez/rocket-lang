package object

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/ast"
)

type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Body       *ast.Block
	Env        *Environment
}

func NewFunction(name string, params []*ast.Identifier, env *Environment, body *ast.Block) *Function {
	return &Function{
		Name:       name,
		Parameters: params,
		Env:        env,
		Body:       body,
	}
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("def ")
	out.WriteString(f.Name)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") \n")
	out.WriteString(f.Body.String())
	out.WriteString("\nend")

	return out.String()
}
func (f *Function) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(f, method, env, args)
}
