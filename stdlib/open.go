package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func openFunction(args ...object.Object) object.Object {
	path := ""
	mode := "r"

	if len(args) < 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch args[0].(type) {
	case *object.String:
		path = args[0].(*object.String).Value
	default:
		return newError("argument to `file` not supported, got=%s", args[0].Type())
	}

	if len(args) > 1 {
		switch args[1].(type) {
		case *object.String:
			mode = args[1].(*object.String).Value
		default:
			return newError("argument mode to `file` not supported, got=%s", args[1].Type())
		}
	}

	file := &object.File{Filename: path}
	file.Open(mode)
	return (file)
}
