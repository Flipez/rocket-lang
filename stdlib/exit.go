package stdlib

import (
	"os"

	"github.com/flipez/rocket-lang/object"
)

func exitFunction(env object.Environment, args ...object.Object) object.Object {
	os.Exit(int(args[0].(*object.Integer).Value))

	return nil
}
