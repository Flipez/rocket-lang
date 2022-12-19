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
	objectMethods[INTEGER_OBJ] = map[string]ObjectMethod{}
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

func (i *Integer) ToStringObj(base *Integer) *String {
	defaultBase := 10

	if base != nil {
		defaultBase = int(base.Value)
	}

	return NewString(strconv.FormatInt(i.Value, defaultBase))
}

func (i *Integer) ToIntegerObj(_ *Integer) *Integer {
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

	obj := NewInteger(int64(i.current))
	i.current += i.step
	return obj, obj, true
}
