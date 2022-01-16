package evaluator

import (
	"github.com/flipez/rocket-lang/object"
)

func evalIntegerInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("devision by zero not allowed")
		}
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalFloatInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("devision by zero not allowed")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
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
			left = left.(*object.Integer).ToFloat()
		}
		if right.Type() == object.INTEGER_OBJ {
			right = right.(*object.Integer).ToFloat()
		}

		result := evalFloatInfix(operator, left, right)

		if object.IsNumber(result) && leftOrig.Type() == object.INTEGER_OBJ && rightOrig.Type() == object.INTEGER_OBJ {
			return result.(*object.Float).TryInteger()
		}
		return result
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ:
		return evalArrayInfixExpression(operator, left, right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalArrayInfixExpression(operator string, left, right object.Object) object.Object {
	leftArray := left.(*object.Array)
	rightArray := right.(*object.Array)

	switch operator {
	case "+":
		length := len(leftArray.Elements) + len(rightArray.Elements)
		elements := make([]object.Object, length, length)
		copy(elements, leftArray.Elements)
		copy(elements[len(leftArray.Elements):], rightArray.Elements)
		return &object.Array{Elements: elements}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
