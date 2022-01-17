package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var Builtins = map[string]*object.Builtin{}

func init() {
	RegisterFunction("puts", putsFunction)
	RegisterFunction("exit", exitFunction)
	RegisterFunction("raise", raiseFunction)
	RegisterFunction("open", openFunction)
}

func RegisterFunction(name string, function object.BuiltinFunction) {
	Builtins[name] = object.NewBuiltin(name, function)
}
