package evaluator

import (
	"strings"

	"github.com/flipez/rocket-lang/object"
)

func evalIntegerInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return object.NewInteger(leftVal + rightVal)
	case "-":
		return object.NewInteger(leftVal - rightVal)
	case "*":
		return object.NewInteger(leftVal * rightVal)
	case "/":
		if rightVal == 0 {
			return object.NewErrorFormat("division by zero not allowed")
		}
		return object.NewInteger(leftVal / rightVal)
	case "%":
		if rightVal == 0 {
			return object.NewErrorFormat("division by zero not allowed")
		}
		return object.NewInteger(leftVal % rightVal)
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalFloatInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return object.NewFloat(leftVal + rightVal)
	case "-":
		return object.NewFloat(leftVal - rightVal)
	case "*":
		return object.NewFloat(leftVal * rightVal)
	case "/":
		if rightVal == 0 {
			return object.NewErrorFormat("division by zero not allowed")
		}
		return object.NewFloat(leftVal / rightVal)
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case operator == "and":
		if left == object.FALSE || left.Type() == object.NIL_OBJ {
			return left
		}
		return right
	case operator == "or":
		if left == object.FALSE || left.Type() == object.NIL_OBJ {
			return right
		}
		return left
	case operator == "==":
		return nativeBoolToBooleanObject(object.CompareObjects(left, right))
	case operator == "!=":
		return nativeBoolToBooleanObject(!object.CompareObjects(left, right))
	case object.IsNumber(left) && object.IsNumber(right):
		if left.Type() == right.Type() && operator != "/" {
			if left.Type() == object.INTEGER_OBJ {
				return evalIntegerInfix(operator, left, right)
			} else if left.Type() == object.FLOAT_OBJ {
				return evalFloatInfix(operator, left, right)
			}
		}

		leftOrig, rightOrig := left, right
		if left.Type() == object.INTEGER_OBJ {
			left = left.(*object.Integer).ToFloatObj()
		}
		if right.Type() == object.INTEGER_OBJ {
			right = right.(*object.Integer).ToFloatObj()
		}

		result := evalFloatInfix(operator, left, right)

		if object.IsNumber(result) && leftOrig.Type() == object.INTEGER_OBJ && rightOrig.Type() == object.INTEGER_OBJ {
			return result.(*object.Float).TryInteger()
		}
		return result
	case ((left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ) || (right.Type() == object.STRING_OBJ && left.Type() == object.INTEGER_OBJ)) && operator == "*":
		var stringObj string
		var intObj int64
		if left.Type() == object.STRING_OBJ {
			stringObj = left.(*object.String).Value
			intObj = right.(*object.Integer).Value
		} else {
			stringObj = right.(*object.String).Value
			intObj = left.(*object.Integer).Value
		}

		return object.NewString(strings.Repeat(stringObj, int(intObj)))
	case left.Type() != right.Type():
		return object.NewErrorFormat("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ:
		return evalArrayInfixExpression(operator, left, right)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return object.NewString(leftVal + rightVal)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalArrayInfixExpression(operator string, left, right object.Object) object.Object {
	leftArray := left.(*object.Array)
	rightArray := right.(*object.Array)

	switch operator {
	case "+":
		length := len(leftArray.Elements) + len(rightArray.Elements)
		elements := make([]object.Object, length)
		copy(elements, leftArray.Elements)
		copy(elements[len(leftArray.Elements):], rightArray.Elements)
		return object.NewArray(elements)
	default:
		return object.NewErrorFormat("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
