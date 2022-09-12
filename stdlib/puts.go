package stdlib

import (
	"fmt"

	"github.com/flipez/rocket-lang/object"
)

func putsFunction(env object.Environment, args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return object.NIL
}
