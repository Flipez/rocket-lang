package stdlib

import (
	"fmt"
	"os"

	"github.com/flipez/rocket-lang/object"
)

var osFunctions = map[string]*object.BuiltinFunction{}
var osProperties = map[string]*object.BuiltinProperty{}

func init() {
	osFunctions["exit"] = object.NewBuiltinFunction(
		"exit",
		object.MethodLayout{
			Description: "Terminates the program with the given exit code.",
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			Example: `🚀 > OS.exit(1)
exit status 1`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			os.Exit(int(args[0].(*object.Integer).Value))

			return nil
		})

	osFunctions["raise"] = object.NewBuiltinFunction(
		"raise",
		object.MethodLayout{
			Description: "Terminates the program with the given exit code and prints the error message.",
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
				object.Arg(object.STRING_OBJ),
			),
			Example: `🚀 > OS.raise(1, "broken")
🔥 RocketLang raised an error: "broken"
exit status 1`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			fmt.Printf("🔥 RocketLang raised an error: %s\n", args[1].Inspect())
			os.Exit(int(args[0].(*object.Integer).Value))

			return nil
		})
}
