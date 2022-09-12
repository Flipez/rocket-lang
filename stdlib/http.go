package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var httpFunctions = map[string]*object.BuiltinFunction{}
var httpProperties = map[string]*object.BuiltinProperty{}

func init() {
	httpFunctions["new"] = object.NewBuiltinFunction(
		"new",
		object.MethodLayout{
			Description: "Creates a new instance of HTTP",
			ReturnPattern: object.Args(
				object.Arg(object.HTTP_OBJ),
			),
		},
		func(_ object.Environment, _ ...object.Object) object.Object {
			return object.NewHTTP()
		})
}
