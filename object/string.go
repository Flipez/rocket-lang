package object

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type String struct {
	Value string
}

func NewString(s string) *String {
	return &String{Value: s}
}

func init() {
	objectMethods[STRING_OBJ] = map[string]ObjectMethod{
		"count": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				arg := args[0].(*String).Value
				return NewInteger(int64(strings.Count(s.Value, arg)))
			},
		},
		"find": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				arg := args[0].(*String).Value
				return NewInteger(int64(strings.Index(s.Value, arg)))
			},
		},
		"format": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OverloadArg(STRING_OBJ, INTEGER_OBJ, FLOAT_OBJ, BOOLEAN_OBJ, ARRAY_OBJ, HASH_OBJ),
				),
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				nativeObjects := make([]any, len(args))
				for idx, arg := range args {
					nativeObjects[idx] = ObjectToAny(arg)
				}
				return NewString(fmt.Sprintf(s.Value, nativeObjects...))
			},
		},
		"size": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewInteger(int64(utf8.RuneCountInString(s.Value)))
			},
		},
		"plz_i": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
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
				return NewInteger(i)
			},
		},
		"replace": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(STRING_OBJ),
					Arg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				oldS := args[0].(*String).Value
				newS := args[1].(*String).Value
				return NewString(strings.Replace(s.Value, oldS, newS, -1))
			},
		},
		"reverse": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				out := make([]rune, utf8.RuneCountInString(s.Value))
				i := len(out)
				for _, c := range s.Value {
					i--
					out[i] = c
				}
				return NewString(string(out))
			},
		},
		"reverse!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				out := make([]rune, utf8.RuneCountInString(s.Value))
				i := len(out)
				for _, c := range s.Value {
					i--
					out[i] = c
				}
				s.Value = string(out)
				return NIL
			},
		},
		"split": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				sep := " "

				if len(args) > 0 {
					sep = args[0].(*String).Value
				}

				fields := strings.Split(s.Value, sep)

				l := len(fields)
				result := make([]Object, l)
				for i, txt := range fields {
					result[i] = NewString(txt)
				}
				return NewArray(result)
			},
		},
		"lines": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				sep := "\n"

				fields := strings.Split(s.Value, sep)

				l := len(fields)
				result := make([]Object, l)
				for i, txt := range fields {
					result[i] = NewString(txt)
				}
				return NewArray(result)
			},
		},
		"strip": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.TrimSpace(s.Value))
			},
		},
		"strip!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.TrimSpace(s.Value)
				return NIL
			},
		},
		"downcase": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.ToLower(s.Value))
			},
		},
		"downcase!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.ToLower(s.Value)
				return NIL
			},
		},
		"upcase": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.ToUpper(s.Value))
			},
		},
		"upcase!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.ToUpper(s.Value)
				return NIL
			},
		},
		"ascii": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				length := len(s.Value)
				var result Object
				switch length {
				case 0:
					result = NewInteger(-1)
				case 1:
					result = NewInteger(int64(s.Value[0]))
				default:
					arr := make([]Object, length)
					for idx, char := range s.Value {
						arr[idx] = NewInteger(int64(char))
					}
					result = NewArray(arr)
				}
				return result
			},
		},
	}
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string {
	var output string

	for _, char := range s.Value {
		if char == '"' {
			output += string('\\')
		}

		output += string(char)
	}

	return `"` + output + `"`
}

func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(s, method, env, args)
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (s *String) GetIterator() Iterator {
	return &stringIterator{chars: []rune(s.Value)}
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

type stringIterator struct {
	chars []rune
	index int
}

func (s *stringIterator) Next() (Object, Object, bool) {
	if s.index < len(s.chars) {
		val := NewString(string(s.chars[s.index]))
		idx := NewInteger(int64(s.index))
		s.index++
		return val, idx, true
	}
	return nil, NewInteger(0), false
}
