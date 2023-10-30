package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalIf(ie *ast.If, env *object.Environment) object.Object {
	for _, pair := range ie.ConConPairs {
		condition := Eval(pair.Condition, env)
		if object.IsError(condition) {
			return condition
		}
		if object.IsTruthy(condition) {
			return Eval(pair.Consequence, env)
		}
	}
	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}
	return object.NIL
}
