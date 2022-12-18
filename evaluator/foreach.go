package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalForeach(fle *ast.Foreach, env *object.Environment) object.Object {
	val := Eval(fle.Value, env)

	if val.Type() != object.INTEGER_OBJ && fle.Start != nil {
		return object.NewErrorFormat("%s:%d:%d: unsupported range rocket value, got %s", fle.Token.File, fle.Token.LineNumber, fle.Token.LinePosition, val.Type())
	}

	var start, step int

	if fle.Start != nil {
		o := Eval(fle.Start, env)
		if i, ok := o.(*object.Integer); ok {
			start = int(i.Value)
		} else {
			return object.NewErrorFormat("%s:%d:%d: range rocket start has to be an integer, got %s", fle.Token.File, fle.Token.LineNumber, fle.Token.LinePosition, o.Type())
		}
	}

	if fle.Step != nil {
		o := Eval(fle.Step, env)
		if i, ok := o.(*object.Integer); ok {
			step = int(i.Value)
		} else {
			return object.NewErrorFormat("%s:%d:%d: range rocket step has to be an integer, got %s", fle.Token.File, fle.Token.LineNumber, fle.Token.LinePosition, o.Type())
		}
	} else {
		step = 1
	}

	helper, ok := val.(object.Iterable)
	if !ok {
		return object.NewErrorFormat("%s:%d:%d: %s object doesn't implement the Iterable interface", fle.Token.File, fle.Token.LineNumber, fle.Token.LinePosition, val.Type())
	}

	child := object.NewEnclosedEnvironment(env)

	iterator := helper.GetIterator(start, step, fle.Inclusive)

	ret, idx, ok := iterator.Next()

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

		if rt != nil && rt.Type() == object.BREAK_VALUE_OBJ {
			return rt.(*object.BreakValue).Value
		}

		ret, idx, ok = iterator.Next()
	}

	return val
}
