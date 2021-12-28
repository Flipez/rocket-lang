package stdlib

import (
	"fmt"

	"github.com/flipez/rocket-lang/object"
)

var Builtins = map[string]*object.Builtin{}

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func init() {
	RegisterFunction("len", lenFunction)
	RegisterFunction("puts", putsFunction)
	RegisterFunction("exit", exitFunction)
	RegisterFunction("raise", raiseFunction)
	RegisterFunction("open", openFunction)
}

func RegisterFunction(name string, function object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: function}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
