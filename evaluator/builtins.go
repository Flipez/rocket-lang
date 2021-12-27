package evaluator

import (
	"fmt"
	"github.com/flipez/rocket-lang/object"
	"os"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return nil
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to `exit` must be INTEGER, got=%s", args[0].Type())
			}

			os.Exit(int(args[0].(*object.Integer).Value))

			return nil
		},
	},
	"raise": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("first argument to `raise` must be INTEGER, got=%s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `raise` must be STRING, got=%s", args[1].Type())
			}

			fmt.Printf("🔥 RocketLang raised an error: %s\n", args[1].Inspect())
			os.Exit(int(args[0].(*object.Integer).Value))

			return nil
		},
	},
	"open": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			path := ""
			mode := "r"

			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch args[0].(type) {
			case *object.String:
				path = args[0].(*object.String).Value
			default:
				return newError("argument to `file` not supported, got=%s", args[0].Type())
			}

			if len(args) > 1 {
				switch args[1].(type) {
				case *object.String:
					mode = args[1].(*object.String).Value
				default:
					return newError("argument mode to `file` not supported, got=%s", args[1].Type())
				}
			}

			file := &object.File{Filename: path}
			file.Open(mode)
			return (file)
		},
	},
}
