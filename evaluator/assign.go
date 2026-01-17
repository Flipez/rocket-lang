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

	// Handle multiple assignment (array unpacking)
	if len(a.Names) > 1 {
		return evalMultipleAssign(a.Names, evaluated, env)
	}

	// Single assignment
	if len(a.Names) == 0 {
		return object.NewError("no assignment target specified")
	}

	return evalSingleAssign(a.Names[0], evaluated, env)
}

// evalMultipleAssign handles a, b, c = [1, 2, 3]
func evalMultipleAssign(names []ast.Expression, value object.Object, env *object.Environment) object.Object {
	// Check if value is an array
	arr, ok := value.(*object.Array)
	if !ok {
		return object.NewErrorFormat("cannot unpack %s into multiple variables (expected ARRAY)", value.Type())
	}

	// Check if array has enough elements
	if len(arr.Elements) < len(names) {
		return object.NewErrorFormat("not enough values to unpack (expected %d, got %d)", len(names), len(arr.Elements))
	}

	// Assign each element to the corresponding variable
	for i, name := range names {
		ident, ok := name.(*ast.Identifier)
		if !ok {
			return object.NewErrorFormat("multiple assignment only supports simple identifiers, got %T", name)
		}

		env.Set(ident.Value, arr.Elements[i])
	}

	return value
}

// evalSingleAssign handles single assignment including indexed assignments
func evalSingleAssign(name ast.Expression, evaluated object.Object, env *object.Environment) object.Object {
	switch v := name.(type) {
	case *ast.Identifier:
		env.Set(v.String(), evaluated)
	case *ast.Index:
		obj := Eval(v.Left, env)
		switch o := obj.(type) {
		case *object.Array:
			idx, err := handleIntegerIndex(v, env)
			if err != nil {
				return object.NewError(err)
			}

			if l := len(o.Elements); int(math.Abs(float64(idx))) >= l {
				return object.NewErrorFormat(
					"index out of range, got %d but array has only %d elements", idx, l,
				)
			}

			if idx < 0 {
				idx = len(o.Elements) + idx
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

			if l := len(o.Value); int(math.Abs(float64(idx))) >= l {
				return object.NewErrorFormat(
					"index out of range, got %d but string is only %d long", idx, l,
				)
			}

			if idx < 0 {
				idx = len(o.Value) + idx
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

func handleIntegerIndex(ai *ast.Index, env *object.Environment) (int, error) {
	obj := Eval(ai.Index, env)
	num, ok := obj.(*object.Integer)
	if !ok {
		return 0, fmt.Errorf("expected index to be an INTEGER, got %s", obj.Type())
	}
	return num.Value, nil
}
