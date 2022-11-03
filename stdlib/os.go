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
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			os.Exit(int(args[0].(*object.Integer).Value))

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
			fmt.Printf("🔥 RocketLang raised an error: %s\n", args[1].Inspect())
			os.Exit(int(args[0].(*object.Integer).Value))

			return nil
		})
}
