package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := stdlib.Builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}
