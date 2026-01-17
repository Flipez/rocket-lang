package stdlib

import (
	"fmt"

	"github.com/flipez/rocket-lang/object"
)

func putsFunction(env object.Environment, args ...object.Object) object.Object {
	for _, arg := range args {
		// For strings, print the actual value (not the quoted representation)
		if str, ok := arg.(*object.String); ok {
			fmt.Println(str.Value)
		} else {
			fmt.Println(arg.Inspect())
		}
	}

	return object.NIL
}
