package stdlib

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/flipez/rocket-lang/object"
)

var ioFunctions = map[string]*object.BuiltinFunction{}
var ioProperties = map[string]*object.BuiltinProperty{}

// stdinReader wraps os.Stdin once for the process lifetime. A reader built
// fresh inside read_line's closure would, for piped input, buffer everything
// still on the descriptor on its first call and then discard that buffer
// when the closure returned -- so only the first line was ever visible and
// every later call saw a descriptor already drained past it.
var stdinReader = bufio.NewReader(os.Stdin)

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
			ReturnPattern: object.Args(object.Arg(object.STRING_OBJ, object.NIL_OBJ, object.ERROR_OBJ)),
		},
		func(_ object.Environment, _ ...object.Object) object.Object {
			line, err := stdinReader.ReadString('\n')
			if err != nil {
				if !errors.Is(err, io.EOF) {
					// A real read failure is a fault, unlike EOF, and the
					// caller needs to be able to tell the two apart.
					return object.NewErrorFormat("%s", err.Error())
				}

				if line == "" {
					// End of input is not a failure: it is how a piped program
					// ends. nil says "nothing more", which the caller can test.
					return object.NIL
				}

				// The final line before EOF has no trailing newline but is
				// still real data -- fall through and return it.
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
