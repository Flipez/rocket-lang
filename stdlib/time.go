package stdlib

import (
	"time"

	"github.com/flipez/rocket-lang/object"
)

var timeFunctions = map[string]*object.BuiltinFunction{}
var timeProperties = map[string]*object.BuiltinProperty{}

func init() {
	timeFunctions["sleep"] = object.NewBuiltinFunction(
		"sleep",
		object.MethodLayout{
			Description: "Stops the RocketLang routine for at least the stated duration in seconds",
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.NIL_OBJ),
			),
			Example: `🚀 > Time.sleep(2)`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			time.Sleep(time.Duration(args[0].(*object.Integer).Value) * time.Second)
			return object.NIL
		})
	timeFunctions["unix"] = object.NewBuiltinFunction(
		"unix",
		object.MethodLayout{
			Description: "Returns the current time as unix timestamp",
			ReturnPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			Example: `🚀 > Time.Unix()`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			return object.NewInteger(time.Now().Unix())
		})
}
