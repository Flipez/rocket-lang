package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalThis(node *ast.This, env *object.Environment) object.Object {
	if env.Self != nil {
		return env.Self
	}

	return object.NewErrorFormat("%d:%d: runtime error: cannot call 'this' outside of class", node.Token.LineNumber, node.Token.LinePosition)
}
