package evaluator

import (
	"fmt"
	"github.com/flipez/rocket-lang/object"
	"os"
)

var array_pop = &object.Builtin{
	Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("argument to `pop` must be ARRAY, got=%s", args[0].Type())
		}

		arr := args[0].(*object.Array)
		length := len(arr.Elements)

		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[:(length-1)])

		returnArray := make([]object.Object, 2)

		returnArray[0] = &object.Array{Elements: newElements}
		returnArray[1] = arr.Elements[length-1]

		return &object.Array{Elements: returnArray}
	},
}

var array_push = &object.Builtin{
	Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("argument to `push` must be ARRAY, got=%s", args[0].Type())
		}

		arr := args[0].(*object.Array)
		length := len(arr.Elements)

		newElements := make([]object.Object, length+1, length+1)
		copy(newElements, arr.Elements)
		newElements[length] = args[1]

		return &object.Array{Elements: newElements}
	},
}

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
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got=%s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push":  array_push,
	"yoink": array_push,
	"pop":   array_pop,
	"yeet":  array_pop,
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
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

			return NULL
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

			return NULL
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
