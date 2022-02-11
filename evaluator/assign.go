package evaluator

import (
	"fmt"
	"math"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalAssign(a *ast.Assign, env *object.Environment) (val object.Object) {
	evaluated := Eval(a.Value, env)
	if object.IsError(evaluated) {
		return evaluated
	}

	switch v := a.Name.(type) {
	case *ast.Identifier:
		env.Set(v.String(), evaluated)
	case *ast.Index:
		obj, _ := env.Get(v.Left.String())
		switch o := obj.(type) {
		case *object.Array:
			idx, err := handleIntegerIndex(v, env)
			if err != nil {
				return object.NewError(err)
			}

			if l := int64(len(o.Elements)); int64(math.Abs(float64(idx))) >= l {
				return object.NewErrorFormat(
					"index out of range, got %d but array has only %d elements", idx, l,
				)
			}

			if idx < 0 {
				idx = int64(len(o.Elements)) + idx
			}

			o.Elements[idx] = evaluated
		case *object.Hash:
			obj := Eval(v.Index, env)
			h, ok := obj.(object.Hashable)
			if !ok {
				return object.NewErrorFormat("expected index to be hashable")
			}

			key := h.HashKey()
			o.Pairs[key] = object.HashPair{Key: obj, Value: evaluated}
		case *object.String:
			idx, err := handleIntegerIndex(v, env)
			if err != nil {
				return object.NewError(err)
			}

			if l := int64(len(o.Value)); int64(math.Abs(float64(idx))) >= l {
				return object.NewErrorFormat(
					"index out of range, got %d but string is only %d long", idx, l,
				)
			}

			if idx < 0 {
				idx = int64(len(o.Value)) + idx
			}

			strEval, ok := evaluated.(*object.String)
			if !ok {
				return object.NewErrorFormat("expected STRING object, got %s", evaluated.Type())
			}
			if l := len(strEval.Value); l != 1 {
				return object.NewErrorFormat(
					"expected STRING object to have a length of 1, got %d", l,
				)
			}

			o.Value = o.Value[:idx] + strEval.Value + o.Value[idx+1:]
		default:
			return object.NewErrorFormat("expected object to be indexable")
		}
	}
	return evaluated
}

func handleIntegerIndex(ai *ast.Index, env *object.Environment) (int64, error) {
	obj := Eval(ai.Index, env)
	num, ok := obj.(*object.Integer)
	if !ok {
		return 0, fmt.Errorf("expected index to be an INTEGER, got %s", obj.Type())
	}
	return num.Value, nil
}
