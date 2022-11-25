package object

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Integer struct {
	Value int64
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
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
				ArgPattern: Args(
					OptArg(INTEGER_OBJ),
				),
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
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return o
			},
		},
		"plz_f": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
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

func (i *Integer) GetIterator() Iterator {
	return &integerIterator{max: i.Value}
}

func (i *Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

type integerIterator struct {
	current, max int64
}

func (i *integerIterator) Next() (Object, Object, bool) {
	if i.current < i.max {
		obj := NewInteger(i.current)
		i.current++
		return obj, obj, true
	}
	return nil, NewInteger(0), false
}
