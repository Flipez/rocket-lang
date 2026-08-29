package evaluator

import "github.com/flipez/rocket-lang/object"

// The object package needs two things from here that it cannot reach on its
// own, because it is the package the evaluator is built on: evaluating a node,
// and calling a function with its parameters bound. Registering them here
// rather than at each entry point means a new one cannot forget, and a missed
// registration used to show up as a nil dereference far from the cause.
func init() {
	object.AddEvaluator(Eval)
	object.AddFunctionApplier(applyFunction)
}
