package stdlib

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/flipez/rocket-lang/object"
)

var ioFunctions = map[string]*object.BuiltinFunction{}
var ioProperties = map[string]*object.BuiltinProperty{}

func init() {
	ioFunctions["write"] = object.NewBuiltinFunction("write",
		object.MethodLayout{
			ArgPattern:    object.Args(object.OverloadArg(object.ANY)),
			ReturnPattern: object.Args(object.Arg(object.NIL_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			for _, arg := range args {
				// A string writes its value, not its quoted form -- the same
				// choice print makes, for the same reason.
				if str, ok := arg.(*object.String); ok {
					fmt.Print(str.Value)

					continue
				}

				fmt.Print(arg.Inspect())
			}

			return object.NIL
		})

	ioFunctions["read_line"] = object.NewBuiltinFunction("read_line",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.STRING_OBJ, object.NIL_OBJ)),
		},
		func(_ object.Environment, _ ...object.Object) object.Object {
			line, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil && line == "" {
				// End of input is not a failure: it is how a piped program
				// ends. nil says "nothing more", which the caller can test.
				return object.NIL
			}

			return object.NewString(strings.TrimRight(line, "\r\n"))
		})

	ioFunctions["open"] = object.NewBuiltinFunction(
		"open",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.STRING_OBJ),
				object.OptArg(object.STRING_OBJ),
				object.OptArg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.FILE_OBJ),
			),
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
				return object.NewErrorFormat("%s", err.Error())
			}
			return (file)
		})
}
