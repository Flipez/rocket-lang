package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var Functions = map[string]*object.BuiltinFunction{}
var Modules = map[string]*object.BuiltinModule{}

func init() {
	RegisterFunction("puts", object.MethodLayout{ArgPattern: object.Args(object.Arg(object.ANY_OBJ...))}, putsFunction)

	RegisterModule("Math", "", mathFunctions, mathProperties)
	RegisterModule("HTTP", "", httpFunctions, httpProperties)
	RegisterModule("JSON", "", jsonFunctions, jsonProperties)
	RegisterModule("IO", "", ioFunctions, ioProperties)
	RegisterModule("OS", "", osFunctions, osProperties)
	RegisterModule("Time", "", timeFunctions, timeProperties)
}

func RegisterFunction(name string, layout object.MethodLayout, function func(object.Environment, ...object.Object) object.Object) {
	Functions[name] = object.NewBuiltinFunction(name, layout, function)
}

func RegisterModule(name string, description string, funcs map[string]*object.BuiltinFunction, props map[string]*object.BuiltinProperty) {
	Modules[name] = object.NewBuiltinModule(name, description, funcs, props)
}
