package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalAssign(a *ast.Assign, env *object.Environment) (val object.Object) {
	evaluated := Eval(a.Value, env)
	if object.IsError(evaluated) {
		return evaluated
	}

	env.Set(a.Name.String(), evaluated)
	return evaluated
}
