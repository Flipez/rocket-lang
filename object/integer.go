package object

import (
	"encoding/json"
	"fmt"
)

type Integer struct {
	Value int
	Base  int
}

func NewInteger(i int) *Integer {
	return &Integer{Value: i, Base: 10}
}

func NewIntegerWithBase(i, b int) *Integer {
	return &Integer{Value: i, Base: b}
}

func (i *Integer) Inspect() string {
	switch i.Base {
	case 2:
		return fmt.Sprintf("0b%b", i.Value)
	case 8:
		return fmt.Sprintf("0o%o", i.Value)
	case 16:
		return fmt.Sprintf("0x%x", i.Value)
	}
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func init() {
	objectMethods[INTEGER_OBJ] = map[string]ObjectMethod{
		"base": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewInteger(o.(*Integer).Base)
			},
		},
		"to_base": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				return NewIntegerWithBase(o.(*Integer).Value, args[0].(*Integer).Value)
			},
		},
	}
}

func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(i, method, env, args)
}

func (i *Integer) GetIterator(start, step int, inclusive bool) Iterator {
	val := int(i.Value)
	if val < start {
		step *= -1
		if inclusive {
			val--
		}
	} else if inclusive {
		val++
	}

	return &integerIterator{max: val, step: step, current: start}
}

func (i *Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

func (i *Integer) ToStringObj() *String {
	return NewString(i.Inspect())
}

func (i *Integer) ToIntegerObj() *Integer {
	return i
}

func (i *Integer) ToFloatObj() *Float {
	return NewFloat(float64(i.Value))
}

type integerIterator struct {
	current, max, step int
}

func (i *integerIterator) Next() (Object, Object, bool) {
	if (i.step < 0 && i.current <= i.max) || (i.step > 0 && i.current >= i.max) {
		return nil, NewInteger(0), false
	}

	obj := NewInteger(i.current)
	i.current += i.step
	return obj, obj, true
}
