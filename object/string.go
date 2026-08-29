package object

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"unicode"
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
				return NewInteger(strings.Count(s.Value, arg))
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
				return NewInteger(strings.Index(s.Value, arg))
			},
		},
		"format": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OverloadArg(ANY),
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
				return NewInteger(utf8.RuneCountInString(s.Value))
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
		"ascii": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				s := o.(*String)

				// One entry per rune, so ascii().size() matches size() and
				// reverse(), which both count runes rather than bytes. Sizing
				// this by len() instead used to leave the trailing slots of a
				// multi-byte string nil, and Inspect then dereferenced them.
				runes := []rune(s.Value)
				arr := make([]Object, len(runes))
				for idx, char := range runes {
					arr[idx] = NewInteger(int(char))
				}

				return NewArray(arr)
			},
		},
	}

	// Pure method and ! counterpart, each derived from a single transform. See
	// stringPair for why they are not written out twice.
	stringPair("reverse", nil, func(value string, _ []Object) string {
		return reverseString(value)
	})
	stringPair("upcase", nil, func(value string, _ []Object) string {
		return strings.ToUpper(value)
	})
	stringPair("downcase", nil, func(value string, _ []Object) string {
		return strings.ToLower(value)
	})
	stringPair("capitalize", nil, func(value string, _ []Object) string {
		return capitalizeString(value)
	})
	stringPair("swapcase", nil, func(value string, _ []Object) string {
		return swapCaseString(value)
	})
	stringPair("strip", nil, func(value string, _ []Object) string {
		return strings.TrimSpace(value)
	})
	stringPair("lstrip", nil, func(value string, _ []Object) string {
		return strings.TrimLeftFunc(value, unicode.IsSpace)
	})
	stringPair("rstrip", nil, func(value string, _ []Object) string {
		return strings.TrimRightFunc(value, unicode.IsSpace)
	})
	stringPair("chop", nil, func(value string, _ []Object) string {
		return chopString(value)
	})
	stringPair("chomp", Args(OptArg(STRING_OBJ)), func(value string, args []Object) string {
		if len(args) == 0 {
			return chompLineEnding(value)
		}

		return chompSeparator(value, args[0].(*String).Value)
	})
	stringPair("replace", Args(Arg(STRING_OBJ), Arg(STRING_OBJ)), func(value string, args []Object) string {
		return strings.ReplaceAll(value, args[0].(*String).Value, args[1].(*String).Value)
	})

	stringPredicate("empty?", nil, func(value string, _ []Object) bool {
		return value == ""
	})
	stringPredicate("include?", Args(Arg(STRING_OBJ)), func(value string, args []Object) bool {
		return strings.Contains(value, args[0].(*String).Value)
	})
	stringPredicate("start_with?", Args(OverloadArg(STRING_OBJ)), func(value string, args []Object) bool {
		return anyAffix(args, func(affix string) bool {
			return strings.HasPrefix(value, affix)
		})
	})
	stringPredicate("end_with?", Args(OverloadArg(STRING_OBJ)), func(value string, args []Object) bool {
		return anyAffix(args, func(affix string) bool {
			return strings.HasSuffix(value, affix)
		})
	})
}

// stringPair registers a method and its in-place counterpart from one
// transformation. Writing the pair out twice is what let Array#reverse mutate
// while Array#uniq did not: the two halves drifted because nothing tied them
// together. Here the ! method can only ever do what the plain one does.
//
// Both return a STRING. A ! method hands back the receiver so calls chain, which
// is a deliberate departure from Ruby, where String#upcase! returns nil when it
// changed nothing and "ABC".upcase!.reverse! therefore raises NoMethodError.
func stringPair(name string, argPattern []Argument, transform func(value string, args []Object) string) {
	layout := MethodLayout{
		ArgPattern:    argPattern,
		ReturnPattern: Args(Arg(STRING_OBJ)),
	}

	objectMethods[STRING_OBJ][name] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			return NewString(transform(o.(*String).Value, args))
		},
	}

	objectMethods[STRING_OBJ][name+"!"] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			s := o.(*String)
			s.Value = transform(s.Value, args)

			return s
		},
	}
}

// stringPredicate registers a method returning a BOOLEAN. Predicates have no !
// counterpart: there is nothing to change.
func stringPredicate(name string, argPattern []Argument, test func(value string, args []Object) bool) {
	objectMethods[STRING_OBJ][name] = ObjectMethod{
		Layout: MethodLayout{
			ArgPattern:    argPattern,
			ReturnPattern: Args(Arg(BOOLEAN_OBJ)),
		},
		method: func(o Object, args []Object, _ Environment) Object {
			if test(o.(*String).Value, args) {
				return TRUE
			}

			return FALSE
		},
	}
}

// anyAffix reports whether any of the given strings satisfies match, which is
// how Ruby's start_with? and end_with? take more than one candidate.
func anyAffix(args []Object, match func(affix string) bool) bool {
	for _, arg := range args {
		if match(arg.(*String).Value) {
			return true
		}
	}

	return false
}

func reverseString(value string) string {
	out := make([]rune, utf8.RuneCountInString(value))
	i := len(out)
	for _, c := range value {
		i--
		out[i] = c
	}

	return string(out)
}

// capitalizeString upcases the first character and downcases the rest, as
// Ruby's capitalize does -- 'hello World!' becomes 'Hello world!' rather than
// keeping the W.
func capitalizeString(value string) string {
	if value == "" {
		return value
	}

	runes := []rune(value)
	out := make([]rune, len(runes))
	out[0] = unicode.ToUpper(runes[0])
	for i, c := range runes[1:] {
		out[i+1] = unicode.ToLower(c)
	}

	return string(out)
}

func swapCaseString(value string) string {
	return strings.Map(func(c rune) rune {
		switch {
		case unicode.IsUpper(c):
			return unicode.ToLower(c)
		case unicode.IsLower(c):
			return unicode.ToUpper(c)
		default:
			return c
		}
	}, value)
}

// chopString removes the last character, or both characters of a trailing
// CRLF so that a line ending is never left half there. An empty string chops
// to an empty string rather than erroring.
func chopString(value string) string {
	if value == "" {
		return value
	}

	if strings.HasSuffix(value, "\r\n") {
		return value[:len(value)-2]
	}

	runes := []rune(value)

	return string(runes[:len(runes)-1])
}

// chompLineEnding removes one trailing line ending: CRLF, LF or CR. A trailing
// "\n\r" loses only the "\r", which is what Ruby does.
func chompLineEnding(value string) string {
	for _, ending := range []string{"\r\n", "\n", "\r"} {
		if strings.HasSuffix(value, ending) {
			return value[:len(value)-len(ending)]
		}
	}

	return value
}

// chompSeparator removes one trailing occurrence of separator. The empty
// separator is Ruby's special case for "strip every trailing blank line": it
// removes any number of trailing CRLF or LF, but leaves a bare CR alone.
func chompSeparator(value, separator string) string {
	if separator == "" {
		for {
			switch {
			case strings.HasSuffix(value, "\r\n"):
				value = value[:len(value)-2]
			case strings.HasSuffix(value, "\n"):
				value = value[:len(value)-1]
			default:
				return value
			}
		}
	}

	return strings.TrimSuffix(value, separator)
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string {
	var output string

	for _, char := range s.Value {
		switch char {
		case '"':
			output += `\"`
		case '\n':
			output += `\n`
		case '\t':
			output += `\t`
		case '\r':
			output += `\r`
		case '\\':
			output += `\\`
		default:
			output += string(char)
		}
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

func (s *String) GetIterator(start, step int, _ bool) Iterator {
	return &stringIterator{chars: []rune(s.Value), index: start, step: step}
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s *String) ToStringObj() *String {
	return s
}

func (s *String) ToIntegerObj() Object {
	base := 10
	digits := s.Value

	sign := ""
	if strings.HasPrefix(digits, "-") || strings.HasPrefix(digits, "+") {
		sign, digits = digits[:1], digits[1:]
	}

	lower := strings.ToLower(digits)

	switch {
	case strings.HasPrefix(lower, "0b"):
		base, digits = 2, digits[2:]
	case strings.HasPrefix(lower, "0o"):
		base, digits = 8, digits[2:]
	case strings.HasPrefix(lower, "0x"):
		base, digits = 16, digits[2:]
	case isLegacyOctal(digits):
		base, digits = 8, digits[1:]
	}

	i, err := strconv.ParseInt(sign+digits, base, 64)

	if err != nil {
		return NIL
	}

	return NewIntegerWithBase(int(i), base)
}

// isLegacyOctal reports whether digits is a leading-zero octal literal such
// as "0125". A bare "0" is decimal zero rather than an octal prefix with
// nothing following it, and a leading zero followed by a non-octal digit
// ("08") is a zero-padded decimal.
func isLegacyOctal(digits string) bool {
	if len(digits) < 2 || digits[0] != '0' {
		return false
	}

	for _, c := range digits[1:] {
		if c < '0' || c > '7' {
			return false
		}
	}

	return true
}

func (s *String) ToFloatObj() Object {
	f, err := strconv.ParseFloat(s.Value, 64)

	if err != nil {
		return NIL
	}

	return NewFloat(f)
}

type stringIterator struct {
	chars []rune
	index int
	step  int
}

func (s *stringIterator) Next() (Object, Object, bool) {
	if s.index < len(s.chars) {
		val := NewString(string(s.chars[s.index]))
		idx := NewInteger(s.index)
		s.index += s.step
		return val, idx, true
	}
	return nil, NewInteger(0), false
}
