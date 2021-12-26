package object

import (
	"hash/fnv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type String struct {
	Value  string
	offset int
}

func init() {
	objectMethods[STRING_OBJ] = map[string]ObjectMethod{
		"count": ObjectMethod{
			argPattern: [][]string{
				[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
				arg := args[0].Inspect()
				return &Integer{Value: int64(strings.Count(s.Value, arg))}
			},
		},
		"find": ObjectMethod{
			argPattern: [][]string{
				[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
				arg := args[0].Inspect()
				return &Integer{Value: int64(strings.Index(s.Value, arg))}
			},
		},
		"size": ObjectMethod{
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
			},
		},
		"plz_i": ObjectMethod{
			argsOptional: true,
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
				value := s.Value
				base := 10

				if len(args) > 0 {
					base = int(args[0].(*Integer).Value)
				} else if strings.HasPrefix(value, "0x") {
					base = 8
				}
				if base == 8 {
					value = strings.TrimPrefix(value, "0x")
				}
				i, _ := strconv.ParseInt(value, base, 64)
				return &Integer{Value: i}
			},
		},
		"replace": ObjectMethod{
			argPattern: [][]string{
				[]string{STRING_OBJ},
				[]string{STRING_OBJ},
			},
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
				oldS := args[0].Inspect()
				newS := args[1].Inspect()
				return &String{Value: strings.Replace(s.Value, oldS, newS, -1)}
			},
		},
		"reverse": ObjectMethod{
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
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
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
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
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
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
		"lines": ObjectMethod{
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, args []Object) Object {
				s := o.(*String)
				sep := "\n"

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
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				return &String{Value: strings.TrimSpace(s.Value)}
			},
		},
		"strip!": ObjectMethod{
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				s.Value = strings.TrimSpace(s.Value)
				return &Null{}
			},
		},
		"downcase": ObjectMethod{
			description: "Returns the string with all uppercase letters replaced with lowercase counterparts.",
			example: `🚀 > "TeST".downcase()
=> test`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				return &String{Value: strings.ToLower(s.Value)}
			},
		},
		"downcase!": ObjectMethod{
			description: "Replaces all upcase characters with lowercase counterparts.",
			example: `
🚀 > a = "TeST"
=> TeST
🚀 > a.downcase!()
=> null
🚀 > a
=> test`,
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				s.Value = strings.ToLower(s.Value)
				return &Null{}
			},
		},
		"upcase": ObjectMethod{
			description: "Returns the string with all lowercase letters replaced with uppercase counterparts.",
			example: `🚀 > "test".upcase()
=> TEST`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				return &String{Value: strings.ToUpper(s.Value)}
			},
		},
		"upcase!": ObjectMethod{
			description: "Replaces all lowercase characters with upcase counterparts.",
			example: `
🚀 > a = "test"
=> test
🚀 > a.upcase!()
=> null
🚀 > a
=> TEST`,
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				s := o.(*String)
				s.Value = strings.ToUpper(s.Value)
				return &Null{}
			},
		},
	}
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(s, method, args)
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (s *String) Reset() {
	s.offset = 0
}

func (s *String) Next() (Object, Object, bool) {
	if s.offset < utf8.RuneCountInString(s.Value) {
		s.offset++

		chars := []rune(s.Value)
		val := &String{Value: string(chars[s.offset-1])}

		return val, &Integer{Value: int64(s.offset - 1)}, true
	}

	return nil, &Integer{Value: 0}, false
}
