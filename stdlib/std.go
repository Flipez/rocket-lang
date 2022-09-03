package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var Functions = map[string]*object.BuiltinFunction{}
var Modules = map[string]*object.BuiltinModule{}

func init() {
	RegisterFunction("puts", putsFunction)
	RegisterFunction("exit", exitFunction)
	RegisterFunction("raise", raiseFunction)
	RegisterFunction("open", openFunction)

	RegisterModule("Math", mathFunctions, mathProperties)
}

func RegisterFunction(name string, function func(object.Environment, ...object.Object) object.Object) {
	Functions[name] = object.NewBuiltinFunction(name, function)
}

func RegisterModule(name string, funcs map[string]*object.BuiltinFunction, props map[string]*object.BuiltinProperty) {
	Modules[name] = object.NewBuiltinModule(name, funcs, props)
}
