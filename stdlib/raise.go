package stdlib

import (
	"fmt"
	"os"

	"github.com/flipez/rocket-lang/object"
)

func raiseFunction(env object.Environment, args ...object.Object) object.Object {
	fmt.Printf("🔥 RocketLang raised an error: %s\n", args[1].Inspect())
	os.Exit(int(args[0].(*object.Integer).Value))

	return nil
}
