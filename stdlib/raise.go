package stdlib

import (
	"fmt"
	"os"

	"github.com/flipez/rocket-lang/object"
)

func raiseFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.INTEGER_OBJ {
		return object.NewErrorFormat("first argument to `raise` must be INTEGER, got=%s", args[0].Type())
	}
	if args[1].Type() != object.STRING_OBJ {
		return object.NewErrorFormat("second argument to `raise` must be STRING, got=%s", args[1].Type())
	}

	fmt.Printf("🔥 RocketLang raised an error: %s\n", args[1].Inspect())
	os.Exit(int(args[0].(*object.Integer).Value))

	return nil
}
