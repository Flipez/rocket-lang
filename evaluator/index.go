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
	case left.Type() == object.MATRIX_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalMatrixIndexExpression(left, index)
	case left.Type() == object.MODULE_OBJ:
		return evalModuleIndexExpression(left, index)
	case left.Type() == object.BUILTIN_MODULE_OBJ:
		return evalBuiltinModuleIndexExpression(left, index)
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

	hash, ok := moduleObject.Attributes.(*object.Hash)
	if !ok {
		return object.NewErrorFormat("module '%s' has no exports", moduleObject.Name)
	}

	key, ok := index.(object.Hashable)
	if !ok {
		return object.NewErrorFormat("unusable as module member: %s", index.Type())
	}

	pair, ok := hash.Pairs[key.HashKey()]
	if !ok {
		// Member lookups are always by name, so a *object.String is
		// unwrapped to its bare value; anything else falls back to
		// Inspect(). Either way the format string below supplies the
		// single-quote wrapping.
		memberName := index.Inspect()
		if s, ok := index.(*object.String); ok {
			memberName = s.Value
		}
		return object.NewErrorFormat("module '%s' has no export '%s'", moduleObject.Name, memberName)
	}

	return pair.Value
}

func evalBuiltinModuleIndexExpression(module, index object.Object) object.Object {
	moduleObject := module.(*object.BuiltinModule)

	name := index.(*object.String).Value
	if val, ok := moduleObject.Properties[name]; ok {
		return val.Value
	}

	return object.NewErrorFormat("property `%s` not found for builtin-module `%s`", name, moduleObject.Name)
}

func evalStringIndexExpression(left, index object.Object) object.Object {
	// Characters, not bytes. Indexing the string directly cuts a multi-byte
	// character in half: "тест"[0] answered a single byte, which is not a
	// character at all. size(), reverse() and the rest have always counted
	// characters, so indexing has to agree with them.
	runes := []rune(left.(*object.String).Value)
	max := len(runes) - 1
	idx := transformIndex(index.(*object.Integer).Value, max)

	if idx > max {
		return object.NIL
	}

	return object.NewString(string(runes[idx]))
}

func evalMatrixIndexExpression(left, index object.Object) object.Object {
	matrix := left.(*object.Matrix)
	idx := int(index.(*object.Integer).Value)

	// Support negative indexing
	if idx < 0 {
		idx = matrix.Rows + idx
	}

	if idx < 0 || idx >= matrix.Rows {
		return object.NewErrorFormat("row index %d out of bounds [0, %d)", idx, matrix.Rows)
	}

	// Return the row as an array
	row, err := matrix.Row(idx)
	if err != nil {
		return object.NewErrorFormat("%s", err.Error())
	}
	return row
}

func evalStringRangeIndexExpression(left, firstIndex, secondIndex object.Object) object.Object {
	// Characters, for the same reason as evalStringIndexExpression: slicing the
	// bytes of "тест" landed between the halves of a character.
	runes := []rune(left.(*object.String).Value)
	max := len(runes) - 1

	if firstIndex == nil && secondIndex == nil {
		return object.NewString(string(runes))
	} else if firstIndex != nil && secondIndex != nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if first <= max && second <= max && first <= second {
			return object.NewString(string(runes[first:second]))
		}
	} else if firstIndex != nil && secondIndex == nil {
		first := transformIndex(firstIndex.(*object.Integer).Value, max)

		if first <= max {
			return object.NewString(string(runes[first:]))
		}
	} else if firstIndex == nil && secondIndex != nil {
		second := transformIndex(secondIndex.(*object.Integer).Value, max)

		if second <= max {
			return object.NewString(string(runes[:second]))
		}
	}

	return object.NIL
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return object.NewErrorFormat("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return object.NIL
	}

	return pair.Value
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	obj := array.(*object.Array)
	max := len(obj.Elements) - 1
	idx := transformIndex(index.(*object.Integer).Value, max)

	if idx > max {
		return object.NIL
	}

	return obj.Elements[idx]
}

func evalArrayRangeIndexExpression(left, firstIndex, secondIndex object.Object) object.Object {
	obj := left.(*object.Array)
	max := len(obj.Elements) - 1

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

	return object.NIL
}

func transformIndex(idx, max int) int {
	if idx < 0 {
		idx += max + 1
	}
	return idx
}
