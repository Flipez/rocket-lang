package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalIf(ie *ast.If, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)

	if object.IsError(condition) {
		return condition
	}
	if object.IsTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return object.NULL
	}
}
