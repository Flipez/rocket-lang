package stdlib

import (
	"math"

	"github.com/flipez/rocket-lang/object"
)

var mathFunctions = map[string]*object.BuiltinFunction{}
var mathProperties = map[string]*object.BuiltinProperty{}

func init() {
	mathFunctions["abs"] = object.NewBuiltinFunction("abs",
		object.MethodLayout{},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Abs(f.Value))
		})

	mathProperties["E"] = object.NewBuiltinProperty("E", object.NewFloat(math.E))
	mathProperties["Pi"] = object.NewBuiltinProperty("Pi", object.NewFloat(math.Pi))
	mathProperties["Phi"] = object.NewBuiltinProperty("Phi", object.NewFloat(math.Phi))
	mathProperties["Sqrt2"] = object.NewBuiltinProperty("Sqrt2", object.NewFloat(math.Sqrt2))
	mathProperties["SqrtE"] = object.NewBuiltinProperty("SqrtE", object.NewFloat(math.SqrtE))
	mathProperties["SqrtPi"] = object.NewBuiltinProperty("SqrtPi", object.NewFloat(math.SqrtPi))
	mathProperties["SqrtPhi"] = object.NewBuiltinProperty("SqrtPhi", object.NewFloat(math.SqrtPhi))
	mathProperties["Ln2"] = object.NewBuiltinProperty("Ln2", object.NewFloat(math.Ln2))
	mathProperties["Log2E"] = object.NewBuiltinProperty("Log2E", object.NewFloat(math.Log2E))
	mathProperties["Ln10"] = object.NewBuiltinProperty("Ln10", object.NewFloat(math.Ln10))
	mathProperties["Log10E"] = object.NewBuiltinProperty("Log10E", object.NewFloat(math.Log10E))
}
