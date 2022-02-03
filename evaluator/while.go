package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalWhile(w *ast.While, env *object.Environment) object.Object {
	child := object.NewEnclosedEnvironment(env)

	v := Eval(w.Condition, child)
	for object.IsTruthy(v) {
		rt := Eval(w.Body, child)
		if rt != nil && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
			return rt
		}
		v = Eval(w.Condition, env)
	}

	return object.NULL
}
