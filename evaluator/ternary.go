package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalTernary(t *ast.Ternary, env *object.Environment) object.Object {
	condition := Eval(t.Condition, env)

	if object.IsError(condition) {
		return condition
	}

	if object.IsTruthy(condition) {
		return Eval(t.Consequence, env)
	} else if t.Alternative != nil {
		return Eval(t.Alternative, env)
	} else {
		return object.NULL
	}
}
