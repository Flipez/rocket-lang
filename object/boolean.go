package object

import (
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
	objectMethods[BOOLEAN_OBJ] = map[string]ObjectMethod{
		"plz_s": ObjectMethod{
			description: "Converts a boolean into a String representation and returns `\"true\"` or `\"false\"` based on the value.",
			example: `🚀 > true.plz_s()
=> "true"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				b := o.(*Boolean)
				return NewString(strconv.FormatBool(b.Value))
			},
		},
	}
}

func (b *Boolean) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, env, args)

}
