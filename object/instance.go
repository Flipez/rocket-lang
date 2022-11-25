package object

import (
	"fmt"
)

type Instance struct {
	Class       *Class
	Environment *Environment
}

func (i *Instance) String() string   { return fmt.Sprintf("class instance %s", i.Class.Name.Value) }
func (i *Instance) Type() ObjectType { return INSTANCE_OBJ }
func (i *Instance) Inspect() string  { return "instance" }
func (i *Instance) InvokeMethod(method string, env Environment, args ...Object) Object {
	if function, ok := i.Environment.Get(method); ok {
		if method, ok := function.(*Function); ok {
			methodEnv := createMethodEnvironment(method, args)
			methodEnv.Self = i

			return Evaluator(method.Body, methodEnv)
		}
	}
	return objectMethodLookup(i, method, env, args)

}

func createMethodEnvironment(fn *Function, arguments []Object) *Environment {
	env := NewEnclosedEnvironment(fn.Env)

	for index, parameter := range fn.Parameters {
		if index < len(arguments) {
			env.Set(parameter.Value, arguments[index])
		}
	}

	return env
}
