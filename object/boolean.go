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

var booleanObjectMethods = map[string]ObjectMethod{
	"type": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			return &String{Value: string(o.Type())}
		},
	},
	"plz_s": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			b := o.(*Boolean)
			return &String{Value: strconv.FormatBool(b.Value)}
		},
	},
}

func (b *Boolean) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "methods":
		return listObjectMethods(booleanObjectMethods)
	case "wat":
		return listObjectUsage(b, booleanObjectMethods)
	default:
		if objMethod, ok := booleanObjectMethods[method]; ok {
			return objMethod.Call(b, args)
		}
	}

	return nil
}
