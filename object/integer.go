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
			Layout: MethodLayout{
				Description: "Returns a string representation of the integer. Also takes an argument which represents the integer base to convert between different number systems",
				Example: `🚀 > a = 456
=> 456
🚀 > a.plz_s()
=> "456"

🚀 > 1234.plz_s(2)
=> "10011010010"
🚀 > 1234.plz_s(8)
=> "2322"
🚀 > 1234.plz_s(10)
=> "1234"`,
				ReturnPattern: [][]string{
					[]string{STRING_OBJ},
				},
				ArgsOptional: true,
				ArgPattern: [][]string{
					[]string{INTEGER_OBJ},
				},
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
			Layout: MethodLayout{
				Description: "Returns self",
				ReturnPattern: [][]string{
					[]string{INTEGER_OBJ},
				},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return o
			},
		},
		"plz_f": ObjectMethod{
			Layout: MethodLayout{
				Description: "Converts the integer into a float.",
				Example: `🚀 > a = 456
=> 456
🚀 > a.plz_f()
=> 456.0

🚀 > 1234.plz_f()
=> 1234.0`,
				ReturnPattern: [][]string{
					[]string{FLOAT_OBJ},
				},
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
