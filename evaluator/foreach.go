package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalForeach(fle *ast.Foreach, env *object.Environment) object.Object {
	val := Eval(fle.Value, env)

	helper, ok := val.(object.Iterable)
	if !ok {
		return object.NewErrorFormat("%s object doesn't implement the Iterable interface", val.Type())
	}

	child := object.NewEnclosedEnvironment(env)

	helper.Reset()

	ret, idx, ok := helper.Next()

	for ok {

		child.Set(fle.Ident, ret)

		idxName := fle.Index
		if idxName != "" {
			child.Set(fle.Index, idx)
		}

		rt := Eval(fle.Body, child)

		//
		// If we got an error/return then we handle it.
		//
		if rt != nil && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
			return rt
		}

		ret, idx, ok = helper.Next()
	}

	return val
}
