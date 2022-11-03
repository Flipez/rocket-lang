package stdlib

import (
	"math"
	"math/rand"
	"time"

	"github.com/flipez/rocket-lang/object"
)

var mathFunctions = map[string]*object.BuiltinFunction{}
var mathProperties = map[string]*object.BuiltinProperty{}

func init() {
	// Randomize seed each runtime to avoid predictable results
	randSource := rand.NewSource(time.Now().UnixNano())
	randomizedRand := rand.New(randSource)

	mathFunctions["abs"] = object.NewBuiltinFunction("abs",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Abs(f.Value))
		})

	mathFunctions["acos"] = object.NewBuiltinFunction("acos",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Acos(f.Value))
		},
	)
	mathFunctions["asin"] = object.NewBuiltinFunction("asin",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Asin(f.Value))
		},
	)
	mathFunctions["atan"] = object.NewBuiltinFunction("atan",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Atan(f.Value))
		},
	)
	mathFunctions["ceil"] = object.NewBuiltinFunction("ceil",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Ceil(f.Value))
		},
	)
	mathFunctions["copysign"] = object.NewBuiltinFunction("copysign",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ), object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			sign := args[1].(*object.Float)
			return object.NewFloat(math.Copysign(f.Value, sign.Value))
		},
	)
	mathFunctions["cos"] = object.NewBuiltinFunction("cos",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Cos(f.Value))
		},
	)
	mathFunctions["exp"] = object.NewBuiltinFunction("exp",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Exp(f.Value))
		},
	)
	mathFunctions["floor"] = object.NewBuiltinFunction("floor",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Floor(f.Value))
		},
	)
	mathFunctions["log"] = object.NewBuiltinFunction("log",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Log(f.Value))
		},
	)
	mathFunctions["log10"] = object.NewBuiltinFunction("log10",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Log10(f.Value))
		},
	)
	mathFunctions["log2"] = object.NewBuiltinFunction("log2",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Log2(f.Value))
		},
	)
	mathFunctions["max"] = object.NewBuiltinFunction("max",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ), object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			x := args[0].(*object.Float)
			y := args[1].(*object.Float)
			return object.NewFloat(math.Max(x.Value, y.Value))
		},
	)
	mathFunctions["min"] = object.NewBuiltinFunction("min",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ), object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			x := args[0].(*object.Float)
			y := args[1].(*object.Float)
			return object.NewFloat(math.Min(x.Value, y.Value))
		},
	)
	mathFunctions["pow"] = object.NewBuiltinFunction("pow",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ), object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			x := args[0].(*object.Float)
			y := args[1].(*object.Float)
			return object.NewFloat(math.Pow(x.Value, y.Value))
		},
	)
	mathFunctions["rand"] = object.NewBuiltinFunction("rand",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			return object.NewFloat(randomizedRand.Float64())
		},
	)
	mathFunctions["remainder"] = object.NewBuiltinFunction("remainder",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ), object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			x := args[0].(*object.Float)
			y := args[1].(*object.Float)
			return object.NewFloat(math.Remainder(x.Value, y.Value))
		},
	)
	mathFunctions["round"] = object.NewBuiltinFunction("round",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Round(f.Value))
		},
	)
	mathFunctions["sin"] = object.NewBuiltinFunction("sin",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Sin(f.Value))
		},
	)
	mathFunctions["sqrt"] = object.NewBuiltinFunction("sqrt",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Sqrt(f.Value))
		},
	)
	mathFunctions["tan"] = object.NewBuiltinFunction("tan",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.FLOAT_OBJ)),
			ArgPattern:    object.Args(object.Arg(object.FLOAT_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			f := args[0].(*object.Float)
			return object.NewFloat(math.Tan(f.Value))
		},
	)

	mathProperties["E"] = object.NewBuiltinProperty("E", object.NewFloat(math.E))
	mathProperties["Pi"] = object.NewBuiltinProperty("Pi", object.NewFloat(math.Pi))
	mathProperties["Phi"] = object.NewBuiltinProperty("Phi", object.NewFloat(math.Phi))
	mathProperties["Sqrt2"] = object.NewBuiltinProperty("Sqrt2", object.NewFloat(math.Sqrt2))
	mathProperties["SqrtE"] = object.NewBuiltinProperty("SqrtE", object.NewFloat(math.SqrtE))
	mathProperties["SqrtPi"] = object.NewBuiltinProperty("SqrtPi", object.NewFloat(math.SqrtPi))
	mathProperties["SqrtPhi"] = object.NewBuiltinProperty("SqrtPhi", object.NewFloat(math.SqrtPhi))
	mathProperties["Ln2"] = object.NewBuiltinProperty("Ln2", object.NewFloat(math.Ln2))
	mathProperties["Log2E"] = object.NewBuiltinProperty("Log2E", object.NewFloat(math.Log2E))
	mathProperties["Ln10"] = object.NewBuiltinProperty("Ln10", object.NewFloat(math.Ln10))
	mathProperties["Log10E"] = object.NewBuiltinProperty("Log10E", object.NewFloat(math.Log10E))
}
