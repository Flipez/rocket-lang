package evaluator

import (
	"github.com/flipez/rocket-lang/object"
)

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return object.NewErrorFormat("unknown operator: %s%s", operator, right.Type())
	}
}
