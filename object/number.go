package object

import "math"

// numberMath registers a method that treats the receiver as a float and
// returns a float. Written as a table because thirteen near-identical blocks
// is thirteen chances for one of them to call the wrong function -- which is
// how Math.log2 and Math.log10 would differ only by a character.
func numberMath(objectType ObjectType, toFloat func(Object) float64) {
	unary := map[string]func(float64) float64{
		"sqrt":  math.Sqrt,
		"exp":   math.Exp,
		"log":   math.Log,
		"log2":  math.Log2,
		"log10": math.Log10,
		"sin":   math.Sin,
		"cos":   math.Cos,
		"tan":   math.Tan,
		"asin":  math.Asin,
		"acos":  math.Acos,
		"atan":  math.Atan,
	}

	for name, fn := range unary {
		apply := fn
		objectMethods[objectType][name] = ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewFloat(apply(toFloat(o)))
			},
		}
	}

	binary := map[string]func(float64, float64) float64{
		"copysign":  math.Copysign,
		"remainder": math.Remainder,
	}

	for name, fn := range binary {
		apply := fn
		objectMethods[objectType][name] = ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(NUMERIC),
				),
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				other, ok := args[0].(Floatable)
				if !ok {
					return NewErrorFormat("expected a number, got %s", args[0].Type())
				}

				value, ok := other.ToFloatObj().(*Float)
				if !ok {
					return NewErrorFormat("expected a number, got %s", args[0].Type())
				}

				return NewFloat(apply(toFloat(o), value.Value))
			},
		}
	}
}
