package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalBlock(block *ast.Block, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.ERROR_OBJ {
				if block.Rescue != nil {
					env.Set("error", result)
					result = evalBlock(block.Rescue, env)
				}
				return result
			}
			if rt == object.RETURN_VALUE_OBJ || rt == object.BREAK_VALUE_OBJ || rt == object.NEXT_VALUE_OBJ {
				return result
			}
		}
	}

	return result
}
