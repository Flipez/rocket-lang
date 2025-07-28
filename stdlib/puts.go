package stdlib

import (
	"fmt"

	"github.com/flipez/rocket-lang/object"
)

func putsFunction(env object.Environment, args ...object.Object) object.Object {
	for _, arg := range args {
		if val, ok := arg.(object.Stringable); ok {
			fmt.Println("stringable")
			fmt.Println(val.ToStringObj(nil).Value)
		} else {
			fmt.Println(arg.Inspect())
		}
	}

	return object.NIL
}
