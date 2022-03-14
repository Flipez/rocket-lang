package object

import (
	"encoding/json"
	"hash/fnv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type String struct {
	Value  string
	offset int
}

func NewString(s string) *String {
	return &String{Value: s}
}

func init() {
	objectMethods[STRING_OBJ] = map[string]ObjectMethod{
		"count": ObjectMethod{
			description: "Counts how often a given string or integer occurs in the string. Converts given integers to strings automatically.",
			example: `🚀 > "test".count("t")
=> 2
🚀 > "test".count("f")
=> 0
🚀 > "test1".count("1")
=> 1
🚀 > "test1".count(1)
=> 1`,
			argPattern: [][]string{
				[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				arg := args[0].(*String).Value
				return NewInteger(int64(strings.Count(s.Value, arg)))
			},
		},
		"find": ObjectMethod{
			description: "Returns the character index of a given string if found. Otherwise returns `-1`",
			example: `🚀 > "test".find("e")
=> 1
🚀 > "test".find("f")
=> -1`,
			argPattern: [][]string{
				[]string{STRING_OBJ, INTEGER_OBJ}, // first argument can be string or int
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				arg := args[0].(*String).Value
				return NewInteger(int64(strings.Index(s.Value, arg)))
			},
		},
		"size": ObjectMethod{
			description: "Returns the amount of characters in the string.",
			example: `🚀 > "test".size()
=> 4`,
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewInteger(int64(utf8.RuneCountInString(s.Value)))
			},
		},
		"plz_i": ObjectMethod{
			description: "Interprets the string as an integer with an optional given base. The default base is `10` and switched to `8` if the string starts with `0x`.",
			example: `🚀 > "1234".plz_i()
=> 1234

🚀 > "1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i()
=> 668

🚀 > "0x1234".plz_i(10)
=> 0`,
			argsOptional: true,
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
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
			description: "Replaces the first string with the second string in the given string.",
			example: `🚀 > "test".replace("t", "f")
=> "fesf"`,
			argPattern: [][]string{
				[]string{STRING_OBJ},
				[]string{STRING_OBJ},
			},
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				oldS := args[0].(*String).Value
				newS := args[1].(*String).Value
				return NewString(strings.Replace(s.Value, oldS, newS, -1))
			},
		},
		"reverse": ObjectMethod{
			description: "Returns a copy of the string with all characters reversed.",
			example: `🚀 > "stressed".reverse()
=> "desserts"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
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
			description: "Replaces all the characters in a string in reverse order.",
			example: `🚀 > a = "stressed"
=> "stressed"
🚀 > a.reverse!()
=> null
🚀 > a
=> "desserts"`,
			returnPattern: [][]string{
				[]string{NULL_OBJ},
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
				return NULL
			},
		},
		"split": ObjectMethod{
			description: "Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `\" \"`",
			example: `🚀 > "a,b,c,d".split(",")
=> ["a", "b", "c", "d"]

🚀 > "test and another test".split()
=> ["test", "and", "another", "test"]`,
			argsOptional: true,
			argPattern: [][]string{
				[]string{STRING_OBJ},
			},
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
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
			description: "Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split(\"\\n\")`.",
			example: `🚀 > "test\ntest2".lines()
=> ["test", "test2"]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
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
			description: "Returns a copy of the string with all leading and trailing whitespaces removed.",
			example: `🚀 > " test ".strip()
=> "test"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.TrimSpace(s.Value))
			},
		},
		"strip!": ObjectMethod{
			description: "Removes all leading and trailing whitespaces in the string.",
			example: `
🚀 > a = " test "
=> " test "
🚀 > a.strip!()
=> null
🚀 > a
=> "test"`,
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.TrimSpace(s.Value)
				return NULL
			},
		},
		"downcase": ObjectMethod{
			description: "Returns the string with all uppercase letters replaced with lowercase counterparts.",
			example: `🚀 > "TeST".downcase()
=> test`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.ToLower(s.Value))
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
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.ToLower(s.Value)
				return NULL
			},
		},
		"upcase": ObjectMethod{
			description: "Returns the string with all lowercase letters replaced with uppercase counterparts.",
			example: `🚀 > "test".upcase()
=> TEST`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				return NewString(strings.ToUpper(s.Value))
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
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)
				s.Value = strings.ToUpper(s.Value)
				return NULL
			},
		},
	}
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return `"` + s.Value + `"` }
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(s, method, env, args)
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
		val := NewString(string(chars[s.offset-1]))

		return val, NewInteger(int64(s.offset - 1)), true
	}

	return nil, NewInteger(0), false
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}
