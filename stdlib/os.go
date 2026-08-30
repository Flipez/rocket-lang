package stdlib

import (
	"fmt"
	"io"
	"os"

	"github.com/flipez/rocket-lang/object"
)

var osFunctions = map[string]*object.BuiltinFunction{}
var osProperties = map[string]*object.BuiltinProperty{}

// Ending the process is reached through these two rather than called directly,
// so that the only functions in this package that cannot otherwise be tested
// can be: a test calling the real os.Exit would take the test binary with it.
//
// Same shape as utilities.SetFileSystem and object.AddEvaluator.
var (
	exitProcess           = os.Exit
	exitOutput  io.Writer = os.Stdout
)

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

	osFunctions["raise"] = object.NewBuiltinFunction(
		"raise",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
				object.Arg(object.STRING_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			fmt.Fprintf(exitOutput, "🔥 RocketLang raised an error: %s\n", args[1].Inspect())
			exitProcess(int(args[0].(*object.Integer).Value))

			return nil
		})
}
