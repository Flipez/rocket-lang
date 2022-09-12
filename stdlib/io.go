package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

var ioFunctions = map[string]*object.BuiltinFunction{}
var ioProperties = map[string]*object.BuiltinProperty{}

func init() {
	ioFunctions["open"] = object.NewBuiltinFunction(
		"open",
		object.MethodLayout{
			Description: "Opens a file pointer to the file at the path, mode and permission can be set optionally.",
			ArgPattern: object.Args(
				object.Arg(object.STRING_OBJ),
				object.OptArg(object.STRING_OBJ),
				object.OptArg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.FILE_OBJ),
			),
			Example: `🚀 > IO.open("main.go", "r", "0644")
=> <file:main.go>`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			mode := "r"
			perm := "0644"

			path := args[0].(*object.String).Value
			if len(args) > 1 {
				mode = args[1].(*object.String).Value
			}
			if len(args) > 2 {
				perm = args[2].(*object.String).Value
			}

			file := object.NewFile(path)
			err := file.Open(mode, perm)
			if err != nil {
				return object.NewErrorFormat(err.Error())
			}
			return (file)
		})
}
