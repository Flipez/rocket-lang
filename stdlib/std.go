package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var Builtins = map[string]*object.Builtin{}
var Clazzes = map[string]object.Object{}
var Constants = map[string]object.Object{}

func init() {
	RegisterFunction("puts", putsFunction)
	RegisterFunction("exit", exitFunction)
	RegisterFunction("raise", raiseFunction)
	RegisterFunction("open", openFunction)

	RegisterClass("HTTP", &object.HTTP{})
	RegisterClass("JSON", &object.JSON{})
	RegisterClass("Math", &object.Math{})

	registerConstants()
}

func RegisterFunction(name string, function object.BuiltinFunction) {
	Builtins[name] = object.NewBuiltin(name, function)
}

func RegisterClass(name string, class object.Object) {
	Clazzes[name] = class
}
