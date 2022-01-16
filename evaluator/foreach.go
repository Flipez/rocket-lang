package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalForeach(fle *ast.Foreach, env *object.Environment) object.Object {
	val := Eval(fle.Value, env)

	helper, ok := val.(object.Iterable)
	if !ok {
		return newError("%s object doesn't implement the Iterable interface", val.Type())
	}

	var permit []string
	permit = append(permit, fle.Ident)
	if fle.Index != "" {
		permit = append(permit, fle.Index)
	}

	//
	// This will allow writing EVERYTHING to the parent scope,
	// except the two variables named in the permit-array
	child := object.NewTemporaryScope(env, permit)

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
		if rt != nil && !isError(rt) && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
			return rt
		}

		ret, idx, ok = helper.Next()
	}

	return val
}
