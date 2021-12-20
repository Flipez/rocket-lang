package object

import (
	"fmt"
	"strconv"
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
	objectMethods[BOOLEAN_OBJ] = map[string]ObjectMethod{
		"plz_s": ObjectMethod{
			method: func(o Object, _ []Object) Object {
				b := o.(*Boolean)
				return &String{Value: strconv.FormatBool(b.Value)}
			},
		},
	}
}

func (b *Boolean) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, args)

}
