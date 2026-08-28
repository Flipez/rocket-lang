package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalObjectCall(call *ast.ObjectCall, env *object.Environment) object.Object {
	obj := Eval(call.Object, env)

	method, ok := call.Call.(*ast.Call)
	if !ok {
		return object.NewErrorFormat("%s:%d:%d: undefined method `.%s()` for %s", call.StartToken.File, call.StartToken.LineNumber, call.StartToken.LinePosition, call.Call.String(), obj.Type())
	}

	args := evalExpressions(method.Arguments, env)

	// A user module is a namespace, not a method receiver. Resolve the
	// member from its attributes and apply it as a function. Builtin
	// modules keep using InvokeMethod, which dispatches via their own
	// Functions map.
	if obj.Type() == object.MODULE_OBJ {
		member := evalModuleIndexExpression(obj, object.NewString(method.Callable.String()))
		if object.IsError(member) {
			return member
		}
		return applyFunction(member, args, env)
	}

	if ret := obj.InvokeMethod(method.Callable.String(), *env, args...); ret != nil {
		return ret
	}

	return object.NewErrorFormat("%s:%d:%d: undefined method `.%s()` for %s", call.StartToken.File, call.StartToken.LineNumber, call.StartToken.LinePosition, method.Callable.String(), obj.Type())
}
