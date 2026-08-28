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
		env.Set(node.Name, val)
	} else if _, ok := env.Get(node.Name); !ok {
		return object.NewErrorFormat("%s:%d:%d: Export Error: '%s' is not defined", node.Token.File, node.Token.LineNumber, node.Token.LinePosition, node.Name)
	}

	env.MarkExport(node.Name)

	return object.NIL
}
