package object

import (
	"hash/fnv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type String struct {
	Value string
}

var stringObjectMethods = map[string]ObjectMethod{
	"count": ObjectMethod{
		argPattern: [][]string{
			[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
		},
		method: func(s *String, args []Object) Object {
			arg := args[0].Inspect()
			return &Integer{Value: int64(strings.Count(s.Value, arg))}
		},
	},

	"find": ObjectMethod{
		argPattern: [][]string{
			[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
		},
		method: func(s *String, args []Object) Object {
			arg := args[0].Inspect()
			return &Integer{Value: int64(strings.Index(s.Value, arg))}
		},
	},

	"size": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
		},
	},
	"type": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			return &String{Value: STRING_OBJ}
		},
	},
	"plz_i": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			i, _ := strconv.ParseInt(s.Value, 10, 64)
			return &Integer{Value: i}
		},
	},
	"replace": ObjectMethod{
		argPattern: [][]string{
			[]string{STRING_OBJ},
			[]string{STRING_OBJ},
		},
		method: func(s *String, args []Object) Object {
			oldS := args[0].Inspect()
			newS := args[1].Inspect()
			return &String{Value: strings.Replace(s.Value, oldS, newS, -1)}
		},
	},
	"reverse": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			out := make([]rune, utf8.RuneCountInString(s.Value))
			i := len(out)
			for _, c := range s.Value {
				i--
				out[i] = c
			}
			return &String{Value: string(out)}
		},
	},
	"reverse!": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			out := make([]rune, utf8.RuneCountInString(s.Value))
			i := len(out)
			for _, c := range s.Value {
				i--
				out[i] = c
			}
			s.Value = string(out)
			return &Null{}
		},
	},
	"split": ObjectMethod{
		argsOptional: true,
		argPattern: [][]string{
			[]string{STRING_OBJ},
		},
		method: func(s *String, args []Object) Object {
			sep := " "

			if len(args) > 0 {
				sep = args[0].(*String).Value
			}

			fields := strings.Split(s.Value, sep)

			l := len(fields)
			result := make([]Object, l, l)
			for i, txt := range fields {
				result[i] = &String{Value: txt}
			}
			return &Array{Elements: result}
		},
	},
	"strip": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			return &String{Value: strings.TrimSpace(s.Value)}
		},
	},
	"strip!": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			s.Value = strings.TrimSpace(s.Value)
			return &Null{}
		},
	},
	"downcase": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			return &String{Value: strings.ToLower(s.Value)}
		},
	},
	"downcase!": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			s.Value = strings.ToLower(s.Value)
			return &Null{}
		},
	},
	"upcase": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			return &String{Value: strings.ToUpper(s.Value)}
		},
	},
	"upcase!": ObjectMethod{
		method: func(s *String, _ []Object) Object {
			s.Value = strings.ToUpper(s.Value)
			return &Null{}
		},
	},
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "methods":
		result := make([]Object, len(stringObjectMethods), len(stringObjectMethods))
		var i int
		for name := range stringObjectMethods {
			result[i] = &String{Value: name}
			i++
		}
		return &Array{Elements: result}
	default:
		if objMethod, ok := stringObjectMethods[method]; ok {
			return objMethod.Call(s, args)
		}
	}

	return nil
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
