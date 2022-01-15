package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func openFunction(args ...object.Object) object.Object {
	path := ""
	mode := "r"
	perm := "0644"

	if len(args) < 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch args[0].(type) {
	case *object.String:
		path = args[0].(*object.String).Value
	default:
		return newError("argument to `file` not supported, got=%s", args[0].Type())
	}

	if len(args) == 2 {
		switch args[1].(type) {
		case *object.String:
			mode = args[1].(*object.String).Value
		default:
			return newError("argument mode to `file` not supported, got=%s", args[1].Type())
		}
	}

	if len(args) == 3 {
		switch args[2].(type) {
		case *object.String:
			perm = args[2].(*object.String).Value
		default:
			return newError("argument perm to `file` not supported, got=%s", args[2].Type())
		}
	}

	file := &object.File{Filename: path}
	err := file.Open(mode, perm)
	if err != nil {
		return newError(err.Error())
	}
	return (file)
}
