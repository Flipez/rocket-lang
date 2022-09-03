package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var Functions = map[string]*object.BuiltinFunction{}
var Modules = map[string]*object.BuiltinModule{}

func init() {
	RegisterFunction("puts", object.MethodLayout{ArgPattern: [][]string{object.ANY_OBJ}}, putsFunction)
	RegisterFunction("exit", object.MethodLayout{ArgPattern: [][]string{[]string{object.INTEGER_OBJ}}}, exitFunction)
	RegisterFunction("raise", object.MethodLayout{ArgPattern: [][]string{[]string{object.INTEGER_OBJ}, []string{object.STRING_OBJ}}}, raiseFunction)
	RegisterFunction("open", object.MethodLayout{ArgPattern: [][]string{[]string{object.STRING_OBJ}, []string{object.STRING_OBJ}, []string{object.STRING_OBJ}}}, openFunction)

	RegisterModule("Math", "", mathFunctions, mathProperties)
	RegisterModule("HTTP", "", httpFunctions, httpProperties)
}

func RegisterFunction(name string, layout object.MethodLayout, function func(object.Environment, ...object.Object) object.Object) {
	Functions[name] = object.NewBuiltinFunction(name, layout, function)
}

func RegisterModule(name string, description string, funcs map[string]*object.BuiltinFunction, props map[string]*object.BuiltinProperty) {
	Modules[name] = object.NewBuiltinModule(name, description, funcs, props)
}
