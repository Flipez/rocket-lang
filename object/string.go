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

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "count":
		if len(args) < 1 {
			return &Error{Message: "Missing argument to count()!"}
		}
		// Note that this coerces into a string :)
		arg := args[0].Inspect()
		return &Integer{Value: int64(strings.Count(s.Value, arg))}
	case "find":
		if len(args) < 1 {
			return &Error{Message: "Missing argument to find()!"}
		}

		// Note that this coerces into a string :)
		arg := args[0].Inspect()
		return &Integer{Value: int64(strings.Index(s.Value, arg))}
	case "size":
		return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
	case "plz_i":
		i, err := strconv.ParseInt(s.Value, 10, 64)
		if err != nil {
			i = 0
		}
		return &Integer{Value: i}
	case "methods":
		names := []string{"count", "find", "size", "methods", "replace", "reverse", "split", "toupper", "tolower", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	case "replace":
		if len(args) < 2 {
			return &Error{Message: "Missing arguments to replace()!"}
		}
		// Note that this coerces into strings :)
		oldS := args[0].Inspect()
		newS := args[1].Inspect()
		return &String{Value: strings.Replace(s.Value, oldS, newS, -1)}
	case "reverse":
		out := make([]rune, utf8.RuneCountInString(s.Value))
		i := len(out)
		for _, c := range s.Value {
			i--
			out[i] = c
		}
		return &String{Value: string(out)}
	case "split":
		sep := " "

		if len(args) >= 1 {
			sep = args[0].(*String).Value
		}

		fields := strings.Split(s.Value, sep)

		l := len(fields)
		result := make([]Object, l, l)
		for i, txt := range fields {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	case "strip":
		return &String{Value: strings.TrimSpace(s.Value)}
	case "tolower":
		return &String{Value: strings.ToLower(s.Value)}
	case "toupper":
		return &String{Value: strings.ToUpper(s.Value)}
	case "type":
		return &String{Value: "string"}
	}

	return nil
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
