package stdlib

import (
	"os"

	"github.com/flipez/rocket-lang/object"
)

var osFunctions = map[string]*object.BuiltinFunction{}
var osProperties = map[string]*object.BuiltinProperty{}

// Ending the process is reached through this rather than called directly, so
// that exit can be tested: a test calling the real os.Exit would take the
// test binary with it.
//
// Same shape as utilities.SetFileSystem and object.AddEvaluator.
var exitProcess = os.Exit

func init() {
	osFunctions["exit"] = object.NewBuiltinFunction(
		"exit",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			exitProcess(int(args[0].(*object.Integer).Value))

			return nil
		})
}
