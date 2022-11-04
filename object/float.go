package object

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strconv"
)

type Float struct {
	Value float64
}

func NewFloat(f float64) *Float {
	return &Float{Value: f}
}

func (f *Float) Inspect() string  { return f.toString() }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }
func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%f", f.Value)))

	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

func init() {
	objectMethods[FLOAT_OBJ] = map[string]ObjectMethod{
		"plz_f": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return o
			},
		},
		"plz_s": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				f := o.(*Float)
				return NewString(f.toString())
			},
		},
		"plz_i": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				f := o.(*Float)
				return NewInteger(int64(f.Value))
			},
		},
	}
}

func (f *Float) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(f, method, env, args)
}

func (f *Float) TryInteger() Object {
	if i := int64(f.Value); f.Value == float64(i) {
		return NewInteger(i)
	}
	return f
}

func (f *Float) toString() string {
	if f.Value == float64(int64(f.Value)) {
		return fmt.Sprintf("%.1f", f.Value)
	}
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

func (f *Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}
