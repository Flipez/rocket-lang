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
				Description: "Counts how often a given string or integer occurs in the string. Converts given integers to strings automatically.",
				Example: `🚀 > "test".count("t")
=> 2
🚀 > "test".count("f")
=> 0
🚀 > "test1".count("1")
=> 1
🚀 > "test1".count(1)
=> 1`,
				ArgPattern: Args(
					Arg(STRING_OBJ, INTEGER_OBJ), // first argument can be string or int
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
				Description: "Returns the character index of a given string if found. Otherwise returns `-1`",
				Example: `🚀 > "test".find("e")
=> 1
🚀 > "test".find("f")
=> -1`,
				ArgPattern: Args(
					Arg(STRING_OBJ, INTEGER_OBJ), // first argument can be string or int
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
				Description: "Formats according to a format specifier and returns the resulting string",
				Example: `🚀 » "test%9d".format(1)
» "test        1"
🚀 » "test%1.2f".format(1.5)
» "test1.50"
🚀 » "test%s".format("test")
» "testtest"`,
				ArgPattern: Args(
					OverloadArg(STRING_OBJ, INTEGER_OBJ, FLOAT_OBJ, BOOLEAN_OBJ), // first argument can be string or int
				),
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				nativeObjects := []interface{}{}
				for _, arg := range args {
					switch e := arg.(type) {
					case *String:
						nativeObjects = append(nativeObjects, e.Value)
					case *Integer:
						nativeObjects = append(nativeObjects, e.Value)
					case *Float:
						nativeObjects = append(nativeObjects, e.Value)
					case *Boolean:
						nativeObjects = append(nativeObjects, e.Value)
					}
				}

				return NewString(fmt.Sprintf(s.Value, nativeObjects...))
			},
		},
		"size": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns the amount of characters in the string.",
				Example: `🚀 > "test".size()
=> 4`,
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
				Description: "Interprets the string as an integer with an optional given base. The default base is `10` and switched to `8` if the string starts with `0x`.",
				Example: `🚀 > "1234".plz_i()
=> 1234

🚀 > "1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i()
=> 668

🚀 > "0x1234".plz_i(10)
=> 0`,
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
				Description: "Replaces the first string with the second string in the given string.",
				Example: `🚀 > "test".replace("t", "f")
=> "fesf"`,
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
				Description: "Returns a copy of the string with all characters reversed.",
				Example: `🚀 > "stressed".reverse()
=> "desserts"`,
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
				Description: "Replaces all the characters in a string in reverse order.",
				Example: `🚀 > a = "stressed"
=> "stressed"
🚀 > a.reverse!()
=> nil
🚀 > a
=> "desserts"`,
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
				Description: "Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `\" \"`",
				Example: `🚀 > "a,b,c,d".split(",")
=> ["a", "b", "c", "d"]

🚀 > "test and another test".split()
=> ["test", "and", "another", "test"]`,
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
				Description: "Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split(\"\\n\")`.",
				Example: `🚀 > "test\ntest2".lines()
=> ["test", "test2"]`,
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
				Description: "Returns a copy of the string with all leading and trailing whitespaces removed.",
				Example: `🚀 > " test ".strip()
=> "test"`,
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
				Description: "Removes all leading and trailing whitespaces in the string.",
				Example: `
🚀 > a = " test "
=> " test "
🚀 > a.strip!()
=> nil
🚀 > a
=> "test"`,
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
				Description: "Returns the string with all uppercase letters replaced with lowercase counterparts.",
				Example: `🚀 > "TeST".downcase()
=> test`,
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
				Description: "Replaces all upcase characters with lowercase counterparts.",
				Example: `
🚀 > a = "TeST"
=> TeST
🚀 > a.downcase!()
=> nil
🚀 > a
=> test`,
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
				Description: "Returns the string with all lowercase letters replaced with uppercase counterparts.",
				Example: `🚀 > "test".upcase()
=> TEST`,
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
				Description: "Replaces all lowercase characters with upcase counterparts.",
				Example: `
🚀 > a = "test"
=> test
🚀 > a.upcase!()
=> nil
🚀 > a
=> TEST`,
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
