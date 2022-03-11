package object

import (
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
			description: "Returns self",
			returnPattern: [][]string{
				[]string{FLOAT_OBJ},
			},
			method: func(o Object, _ []Object) Object {
				return o
			},
		},
		"plz_s": ObjectMethod{
			description: "Returns a string representation of the float.",
			example: `🚀 > a = 123.456
=> 123.456
🚀 > a.plz_s()
=> "123.456"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(o Object, args []Object) Object {
				f := o.(*Float)
				return NewString(f.toString())
			},
		},
		"plz_i": ObjectMethod{
			description: "Converts the float into an integer.",
			example: `🚀 > a = 123.456
=> 123.456
🚀 > a.plz_i()
=> "123"`,
			method: func(o Object, args []Object) Object {
				f := o.(*Float)
				return NewInteger(int64(f.Value))
			},
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
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
