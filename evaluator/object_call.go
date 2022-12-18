package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalObjectCall(call *ast.ObjectCall, env *object.Environment) object.Object {
	obj := Eval(call.Object, env)
	if method, ok := call.Call.(*ast.Call); ok {
		args := evalExpressions(call.Call.(*ast.Call).Arguments, env)
		ret := obj.InvokeMethod(method.Callable.String(), *env, args...)
		if ret != nil {
			return ret
		}
	}

	return object.NewErrorFormat("%s:%d:%d: undefined method `.%s()` for %s", call.StartToken.File, call.StartToken.LineNumber, call.StartToken.LinePosition, call.Call.(*ast.Call).Callable.String(), obj.Type())
}
