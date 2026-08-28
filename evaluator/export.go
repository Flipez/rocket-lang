package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalExport(node *ast.Export, env *object.Environment) object.Object {
	if node.Value != nil {
		val := Eval(node.Value, env)
		if object.IsError(val) {
			return val
		}
		if _, ok := val.(*object.Module); ok {
			return object.NewErrorFormat("%s:%d:%d: Export Error: cannot export '%s': a module cannot be exported", node.Token.File, node.Token.LineNumber, node.Token.LinePosition, node.Name)
		}
		env.Set(node.Name, val)
	} else {
		existing, ok := env.Get(node.Name)
		if !ok {
			return object.NewErrorFormat("%s:%d:%d: Export Error: '%s' is not defined", node.Token.File, node.Token.LineNumber, node.Token.LinePosition, node.Name)
		}
		if _, ok := existing.(*object.Module); ok {
			return object.NewErrorFormat("%s:%d:%d: Export Error: cannot export '%s': a module cannot be exported", node.Token.File, node.Token.LineNumber, node.Token.LinePosition, node.Name)
		}
	}

	env.MarkExport(node.Name)

	return object.NIL
}
