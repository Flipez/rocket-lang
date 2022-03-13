package stdlib

import (
	"os"

	"github.com/flipez/rocket-lang/object"
)

func exitFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.INTEGER_OBJ {
		return object.NewErrorFormat("argument to `exit` must be INTEGER, got=%s", args[0].Type())
	}

	os.Exit(int(args[0].(*object.Integer).Value))

	return nil
}
