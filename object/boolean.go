package object

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func init() {
	objectMethods[BOOLEAN_OBJ] = map[string]ObjectMethod{}
}

func (b *Boolean) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, env, args)

}

func (b *Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Value)
}

func (b *Boolean) ToStringObj() *String {
	return NewString(strconv.FormatBool(b.Value))
}

func (b *Boolean) ToIntegerObj() *Integer {
	if b.Value {
		return NewInteger(1)
	}
	return NewInteger(0)
}

func (b *Boolean) ToFloatObj() *Float {
	if b.Value {
		return NewFloat(1.0)
	}
	return NewFloat(0.0)
}
