package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements

	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.Block:
		return evalBlock(node, env)
	case *ast.Foreach:
		return evalForeach(node, env)
	case *ast.While:
		return evalWhile(node, env)
	case *ast.Return:
		val := Eval(node.ReturnValue, env)
		if object.IsError(val) {
			return val
		}
		return object.NewReturnValue(val)

	// Expressions
	case *ast.Integer:
		return object.NewInteger(node.Value)
	case *ast.Float:
		return object.NewFloat(node.Value)
	case *ast.Function:
		function := object.NewFunction(
			node.Parameters,
			env,
			node.Body,
		)

		if node.Name != "" {
			env.Set(node.Name, function)
		}

		return function
	case *ast.Import:
		return evalImport(node, env)
	case *ast.String:
		return object.NewString(node.Value)
	case *ast.Array:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && object.IsError(elements[0]) {
			return elements[0]
		}
		return object.NewArray(elements)
	case *ast.Hash:
		return evalHash(node, env)

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.Prefix:
		right := Eval(node.Right, env)
		if object.IsError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.Infix:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if object.IsError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.If:
		return evalIf(node, env)
	case *ast.Ternary:
		return evalTernary(node, env)

	case *ast.Call:
		function := Eval(node.Callable, env)
		if object.IsError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

	case *ast.Index:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if object.IsError(index) {
			return index
		}
		return evalIndex(left, index)

	case *ast.ObjectCall:
		res := evalObjectCall(node, env)
		return (res)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.Assign:
		return evalAssign(node, env)
	}

	return nil
}

func applyFunction(def object.Object, args []object.Object) object.Object {
	switch def := def.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(def, args)
		evaluated := Eval(def.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return def.Fn(args...)

	default:
		return object.NewErrorFormat("not a function: %s", def.Type())
	}
}

func extendFunctionEnv(def *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(def.Env)

	for paramIdx, param := range def.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return object.NewErrorFormat("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return object.NewInteger(-value)
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return object.TRUE
	}
	return object.FALSE
}
