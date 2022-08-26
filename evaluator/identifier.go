package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := stdlib.Constants[node.Value]; ok {
		return val
	}

	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := stdlib.Builtins[node.Value]; ok {
		return builtin
	}

	if clazz, ok := stdlib.Clazzes[node.Value]; ok {
		return clazz
	}

	return object.NewErrorFormat("identifier not found: " + node.Value)
}
