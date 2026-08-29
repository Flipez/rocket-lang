package object

// functionApplier calls a callable with arguments bound to its parameters. This
// package cannot do that itself: it means building an environment from the
// parameter list and evaluating the body, and the code for that lives in the
// evaluator, which imports this package -- so the dependency can only run the
// other way. It is injected exactly as Evaluator is.
//
// Without it, a method could reach a function's Body but had no way to pass it
// anything. HTTP.handle worked around that by setting `request` and `response`
// as variables in the environment and evaluating the body directly, which is
// why an HTTP handler takes no parameters.
var functionApplier func(fn Object, args []Object, env *Environment) Object

// AddFunctionApplier wires up function application. The evaluator registers it
// from its own init, so no caller has to remember to.
func AddFunctionApplier(applier func(fn Object, args []Object, env *Environment) Object) {
	functionApplier = applier
}

// CallFunction runs a callback and returns whatever it produced: a value, an
// error, or one of the control-flow objects that break and next leave behind.
//
// Arity is the applier's business, and it reports a mismatch in the same words
// a call written out in full would.
func CallFunction(fn Object, env Environment, args ...Object) Object {
	if functionApplier == nil {
		// Only reachable from a program that uses this package without the
		// evaluator. Better than the nil dereference it replaces.
		return NewError("cannot call a callback: no function applier is registered")
	}

	return functionApplier(fn, args, &env)
}

// CallbackStopped reports a callback that ran `break`, and a method walking a
// collection should stop -- the way break ends a foreach.
//
// break and next inside a function are not consumed by it, so a callback can
// hand one back: `def() break end` returns a BREAK_VALUE. Treating that as an
// ordinary value would put a BREAK_VALUE inside a mapped array, which means
// nothing to anyone. Every method taking a callback has to answer for it, so
// the rule lives here rather than in each of them.
func CallbackStopped(o Object) bool {
	return o != nil && o.Type() == BREAK_VALUE_OBJ
}

// CallbackSkipped reports a callback that ran `next`. The element produced
// nothing, the way next moves a foreach along.
func CallbackSkipped(o Object) bool {
	return o != nil && o.Type() == NEXT_VALUE_OBJ
}
