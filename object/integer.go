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
			description: "Returns a string representation of the integer. Also takes an argument which represents the integer base to convert between different number systems",
			example: `🚀 > a = 456
=> 456
🚀 > a.plz_s()
=> "456"

🚀 > 1234.plz_s(2)
=> "10011010010"
🚀 > 1234.plz_s(8)
=> "2322"
🚀 > 1234.plz_s(10)
=> "1234"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
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
	return objectMethodLookup(i, method, args)

}
