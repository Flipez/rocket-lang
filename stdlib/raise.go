package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func raiseFunction(_ object.Environment, args ...object.Object) object.Object {
	return object.NewError(args[0].(*object.String).Value)
}
