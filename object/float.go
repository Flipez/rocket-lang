package object

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
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
		"abs": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewFloat(math.Abs(o.(*Float).Value))
			},
		},
		"divmod": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(FLOAT_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				value := o.(*Float).Value
				divisor := args[0].(*Float).Value

				if divisor == 0 {
					return NewError("division by zero not allowed")
				}

				// Truncated, so that this agrees with Integer#divmod and with
				// the / operator. Ruby floors instead.
				quotient := math.Trunc(value / divisor)

				return NewArrayWithObjects(
					NewFloat(quotient),
					NewFloat(value-quotient*divisor),
				)
			},
		},
		"infinite?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				// 1, -1 or nil rather than a boolean, so the direction is not
				// lost. finite? is the plain yes-or-no question.
				switch {
				case math.IsInf(o.(*Float).Value, 1):
					return NewInteger(1)
				case math.IsInf(o.(*Float).Value, -1):
					return NewInteger(-1)
				default:
					return NIL
				}
			},
		},
	}

	// Rounding a float takes an optional number of decimal places. A negative
	// count rounds to a power of ten, so 555.5.round(-1) is 560.0. The result
	// stays a FLOAT: a numeric method returns the type it was given, which is
	// why this differs from Ruby, where 1.5.round is an Integer.
	floatRounding("ceil", math.Ceil)
	floatRounding("floor", math.Floor)
	floatRounding("round", math.Round)
	floatRounding("truncate", math.Trunc)

	floatPredicate("zero?", func(value float64) bool { return value == 0 })
	floatPredicate("positive?", func(value float64) bool { return value > 0 })
	floatPredicate("negative?", func(value float64) bool { return value < 0 })
	floatPredicate("nan?", math.IsNaN)
	floatPredicate("finite?", func(value float64) bool {
		return !math.IsInf(value, 0) && !math.IsNaN(value)
	})
}

// floatPredicate registers a method returning a BOOLEAN about the value.
func floatPredicate(name string, test func(value float64) bool) {
	objectMethods[FLOAT_OBJ][name] = ObjectMethod{
		Layout: MethodLayout{
			ReturnPattern: Args(Arg(BOOLEAN_OBJ)),
		},
		method: func(o Object, _ []Object, _ Environment) Object {
			if test(o.(*Float).Value) {
				return TRUE
			}

			return FALSE
		},
	}
}

// floatRounding registers ceil, floor, round or truncate, each taking an
// optional number of decimal places.
func floatRounding(name string, apply func(float64) float64) {
	objectMethods[FLOAT_OBJ][name] = ObjectMethod{
		Layout: MethodLayout{
			ArgPattern:    Args(OptArg(INTEGER_OBJ)),
			ReturnPattern: Args(Arg(FLOAT_OBJ)),
		},
		method: func(o Object, args []Object, _ Environment) Object {
			value := o.(*Float).Value

			if len(args) == 0 {
				return NewFloat(apply(value))
			}

			// Scaling by a power of ten and back is what gives the digit
			// count its meaning. Guard the non-finite cases, where scaling
			// produces NaN instead of leaving the value alone.
			if math.IsInf(value, 0) || math.IsNaN(value) {
				return NewFloat(value)
			}

			scale := math.Pow(10, float64(args[0].(*Integer).Value))

			return NewFloat(apply(value*scale) / scale)
		},
	}
}

func (f *Float) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(f, method, env, args)
}

func (f *Float) TryInteger() Object {
	if i := int(f.Value); f.Value == float64(i) {
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

func (f *Float) ToStringObj() *String {
	return NewString(f.toString())
}

func (f *Float) ToIntegerObj() Object {
	return NewInteger(int(f.Value))
}

func (f *Float) ToFloatObj() Object {
	return f
}
