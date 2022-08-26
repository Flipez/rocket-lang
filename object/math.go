package object

import (
	"math"
)

type Math struct{}

func (m *Math) Type() ObjectType { return MATH_OBJ }
func (m *Math) Inspect() string  { return "MATH" }
func (m *Math) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(m, method, env, args)
}

func init() {
	objectMethods[MATH_OBJ] = map[string]ObjectMethod{
		"abs": ObjectMethod{
			description: "Returns the absolute value of the given number as FLOAT",
			example: `🚀 > Math.abs(-2)
=> 2.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Abs(f.Value))
			},
		},
		"acos": ObjectMethod{
			description: "Returns the arccosine, in radians, of the argument",
			example: `🚀 > Math.acos(1.0)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Acos(f.Value))
			},
		},
		"asin": ObjectMethod{
			description: "Returns the arcsine, in radians, of the argument",
			example: `🚀 > Math.asin(0.0)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Asin(f.Value))
			},
		},
		"atan": ObjectMethod{
			description: "Returns the arctangent, in radians, of the argument",
			example: `🚀 > Math.atan(0.0)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Atan(f.Value))
			},
		},
		"ceil": ObjectMethod{
			description: "Returns the least integer value greater or equal to the argument",
			example: `🚀 > Math.ceil(1.49)
=> 2.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Ceil(f.Value))
			},
		},
		"copysign": ObjectMethod{
			description: "Returns a value with the magnitude of first argument and sign of second argument",
			example: `🚀 > Math.copysign(3.2, -1.0)
=> -3.2`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}, []string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				sign := args[1].(*Float)
				return NewFloat(math.Copysign(f.Value, sign.Value))
			},
		},
		"cos": ObjectMethod{
			description: "Returns the cosine of the radion argument",
			example: `🚀 > Math.cos(Pi/2)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Cos(f.Value))
			},
		},
		"exp": ObjectMethod{
			description: "Returns e**argument, the base-e exponential of argument",
			example: `🚀 > Math.exp(1.0)
=> 2.72`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Exp(f.Value))
			},
		},
		"floor": ObjectMethod{
			description: "Returns the greatest integer value less than or equal to argument",
			example: `🚀 > Math.floor(1.51)
=> 1.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Floor(f.Value))
			},
		},
		"log": ObjectMethod{
			description: "Returns the natural logarithm of argument",
			example: `🚀 > Math.log(2.7183)
=> 1.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Log(f.Value))
			},
		},
		"log10": ObjectMethod{
			description: "Returns the decimal logarithm of argument",
			example: `🚀 > Math.log(100.0)
=> 2.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Log10(f.Value))
			},
		},
		"log2": ObjectMethod{
			description: "Returns the binary logarithm of argument",
			example: `🚀 > Math.log2(256.0)
=> 8.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Log2(f.Value))
			},
		},
		"pow": ObjectMethod{
			description: "Returns argument1**argument2, the base-argument1 exponential of argument2",
			example: `🚀 > Math.pow(2.0, 3.0)
=> 8.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}, []string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				x := args[0].(*Float)
				y := args[1].(*Float)
				return NewFloat(math.Pow(x.Value, y.Value))
			},
		},
		"remainder": ObjectMethod{
			description: "Returns the IEEE 754 floating-point remainder of argument1/argument2",
			example: `🚀 > Math.remainder(100.0, 30.0)
=> 10.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}, []string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				x := args[0].(*Float)
				y := args[1].(*Float)
				return NewFloat(math.Remainder(x.Value, y.Value))
			},
		},
		"sin": ObjectMethod{
			description: "Returns the sine of the radion argument",
			example: `🚀 > Math.sin(Pi)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Sin(f.Value))
			},
		},
		"sqrt": ObjectMethod{
			description: "Returns the square root of argument",
			example: `🚀 > Math.sqrt(3.0 * 3.0 + 4.0 * 4.0)
=> 5.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Sqrt(f.Value))
			},
		},
		"tan": ObjectMethod{
			description: "Returns the tangent of the radion argument",
			example: `🚀 > Math.tan(0.0)
=> 0.0`,
			returnPattern: [][]string{[]string{FLOAT_OBJ}},
			argPattern:    [][]string{[]string{FLOAT_OBJ}},
			method: func(_ Object, args []Object, _ Environment) Object {
				f := args[0].(*Float)
				return NewFloat(math.Tan(f.Value))
			},
		},
	}
}
