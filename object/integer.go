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

func init() {
	objectMethods[INTEGER_OBJ] = map[string]ObjectMethod{
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
}

func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	if oms, ok := objectMethods[i.Type()]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(i, args)
		}
	}

	if oms, ok := objectMethods["*"]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(i, args)
		}
	}

	return nil
}
