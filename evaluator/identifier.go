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

	if fn, ok := stdlib.Functions[node.Value]; ok {
		return fn
	}

	if mod, ok := stdlib.Modules[node.Value]; ok {
		return mod
	}

	return object.NewErrorFormat("%s:%d:%d: identifier not found: "+node.Value, node.Token.File, node.Token.LineNumber, node.Token.LinePosition-len(node.TokenLiteral()))
}
