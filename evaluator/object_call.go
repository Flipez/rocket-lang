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
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

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

	// A hash is a namespace as well as a collection, so a callable stored under
	// a name can be called as a method. That is what makes the closure-and-hash
	// pattern read like the object it already is: a hash of functions closing
	// over a constructor's locals had to be called as h["deposit"](50).
	//
	// This runs only after InvokeMethod found nothing, so a real Hash method
	// always wins and a hash of data -- from JSON, say -- cannot hijack size()
	// or keys().
	if hash, isHash := obj.(*object.Hash); isHash {
		name := method.Callable.String()

		if member, found := hash.Get(name); found {
			if !object.InGroup(object.CALLABLE, member) {
				return object.NewErrorFormat("%s:%d:%d: `%s` is not callable for HASH, it is %s", call.StartToken.File, call.StartToken.LineNumber, call.StartToken.LinePosition, name, member.Type())
			}

			return applyFunction(member, args, env)
		}
	}

	return object.NewErrorFormat("%s:%d:%d: undefined method `.%s()` for %s", call.StartToken.File, call.StartToken.LineNumber, call.StartToken.LinePosition, method.Callable.String(), obj.Type())
}
