package evaluator

import (
	"github.com/flipez/rocket-lang/object"
)

func evalIndex(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	case left.Type() == object.STRING_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalStringIndexExpression(left, index)
	case left.Type() == object.MODULE_OBJ:
		return evalModuleIndexExpression(left, index)
	default:
		return object.NewErrorFormat("index operator not supported: %s", left.Type())
	}
}

func evalRangeIndex(left, firstIndex, secondIndex object.Object) object.Object {
	if firstIndex != nil {
		if objType := firstIndex.Type(); objType != object.INTEGER_OBJ {
			return object.NewErrorFormat("invalid type for first index: %s", objType)
		}
	}
	if secondIndex != nil {
		if objType := secondIndex.Type(); objType != object.INTEGER_OBJ {
			return object.NewErrorFormat("invalid type for second index: %s", objType)
		}
	}

	switch {
	case left.Type() == object.ARRAY_OBJ:
		return evalArrayRangeIndexExpression(left, firstIndex, secondIndex)
	case left.Type() == object.STRING_OBJ:
		return evalStringRangeIndexExpression(left, firstIndex, secondIndex)
	default:
		return object.NewErrorFormat("range index operator not supported: %s", left.Type())
	}
}

func evalModuleIndexExpression(module, index object.Object) object.Object {
	moduleObject := module.(*object.Module)

	return evalHashIndexExpression(moduleObject.Attributes, index)
}

func evalStringIndexExpression(left, index object.Object) object.Object {
	obj := left.(*object.String)
	max := int64(len(obj.Value) - 1)
	idx := transformIndex(index.(*object.Integer).Value, max)

	if idx > max {
		return object.NULL
	}

	return object.NewString(string(obj.Value[idx]))
}

func evalStringRangeIndexExpression(left, firstIndex, secondIndex object.Object) object.Object {
	obj := left.(*object.String)
	max := int64(len(obj.Value) - 1)

	if firstIndex == nil && secondIndex == nil {
		return object.NewString(obj.Value)
	} else if firstIndex != nil && secondIndex != nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if first <= max && second <= max && first <= second {
			return object.NewString(obj.Value[first:second])
		}
	} else if firstIndex != nil && secondIndex == nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)

		if first <= max {
			return object.NewString(obj.Value[first:])
		}
	} else if firstIndex == nil && secondIndex != nil {
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if second <= max {
			return object.NewString(obj.Value[:second])
		}
	}

	return object.NULL
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return object.NewErrorFormat("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return object.NULL
	}

	return pair.Value
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	obj := array.(*object.Array)
	max := int64(len(obj.Elements) - 1)
	idx := transformIndex(index.(*object.Integer).Value, max)

	if idx > max {
		return object.NULL
	}

	return obj.Elements[idx]
}

func evalArrayRangeIndexExpression(left, firstIndex, secondIndex object.Object) object.Object {
	obj := left.(*object.Array)
	max := int64(len(obj.Elements) - 1)

	if firstIndex == nil && secondIndex == nil {
		return object.NewArray(obj.Elements)
	} else if firstIndex != nil && secondIndex != nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if first <= max && second <= max && first <= second {
			return object.NewArray(obj.Elements[first:second])
		}
	} else if firstIndex != nil && secondIndex == nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)

		if first <= max {
			return object.NewArray(obj.Elements[first:])
		}
	} else if firstIndex == nil && secondIndex != nil {
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if second <= max {
			return object.NewArray(obj.Elements[:second])
		}
	}

	return object.NULL
}

func transformIndex(idx, max int64) int64 {
	if idx < 0 {
		idx += max + 1
	}
	return idx
}
