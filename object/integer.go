package object

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Integer struct {
	Value int64
	index int64
}

func NewInteger(i int64) *Integer {
	return &Integer{Value: i}
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func init() {
	objectMethods[INTEGER_OBJ] = map[string]ObjectMethod{
		"plz_s": ObjectMethod{
			description: "Returns a string representation of the integer. Also takes an argument which represents the integer base to convert between different number systems",
			example: `🚀 > a = 456
=> 456
🚀 > a.plz_s()
=> "456"

🚀 > 1234.plz_s(2)
=> "10011010010"
🚀 > 1234.plz_s(8)
=> "2322"
🚀 > 1234.plz_s(10)
=> "1234"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			argsOptional: true,
			argPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				i := o.(*Integer)

				base := 10
				if len(args) > 0 {
					base = int(args[0].(*Integer).Value)
				}

				return NewString(strconv.FormatInt(i.Value, base))
			},
		},
		"plz_i": ObjectMethod{
			description: "Returns self",
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return o
			},
		},
		"plz_f": ObjectMethod{
			description: "Converts the integer into a float.",
			example: `🚀 > a = 456
=> 456
🚀 > a.plz_f()
=> 456.0

🚀 > 1234.plz_f()
=> 1234.0`,
			returnPattern: [][]string{
				[]string{FLOAT_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				i := o.(*Integer)
				return NewFloat(float64(i.Value))
			},
		},
	}
}

func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(i, method, env, args)
}

func (i *Integer) ToFloat() Object {
	return NewFloat(float64(i.Value))
}

func (i *Integer) Reset() {
	i.index = 0
}

func (i *Integer) Next() (Object, Object, bool) {
	if i.index < i.Value {
		index := NewInteger(i.index)
		i.index++
		return index, index, true
	}
	return nil, NewInteger(0), false
}

func (i *Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}
