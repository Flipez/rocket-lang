package object

import (
	"fmt"
	"strconv"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

var integerObjectMethods = map[string]ObjectMethod{
	"type": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			return &String{Value: string(o.Type())}
		},
	},
	"plz_s": ObjectMethod{
		argsOptional: true,
		argPattern: [][]string{
			[]string{INTEGER_OBJ},
		},
		method: func(o Object, args []Object) Object {
			i := o.(*Integer)

			base := 10
			if len(args) > 0 {
				base = int(args[0].(*Integer).Value)
			}

			return &String{Value: strconv.FormatInt(i.Value, base)}
		},
	},
}

func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "methods":
		return listObjectMethods(integerObjectMethods)
	case "wat":
		return listObjectUsage(i, integerObjectMethods)
	default:
		if objMethod, ok := integerObjectMethods[method]; ok {
			return objMethod.Call(i, args)
		}
	}

	return nil
}
