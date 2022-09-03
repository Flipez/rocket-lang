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
}
