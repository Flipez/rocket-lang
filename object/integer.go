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
func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "plz_s":
		return &String{Value: strconv.FormatInt(i.Value, 10)}
	}
	return nil
}
