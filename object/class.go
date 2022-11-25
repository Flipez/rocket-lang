package object

import (
	"bytes"
	"github.com/flipez/rocket-lang/ast"
)

type Class struct {
	Name  *ast.Identifier
	Env   *Environment
	Super *Class
}

func (c *Class) String() string   { return "class" }
func (c *Class) Type() ObjectType { return CLASS_OBJ }
func (c *Class) Inspect() string {
	var out bytes.Buffer

	out.WriteString("class ")
	out.WriteString(c.Name.Token.Literal)
	out.WriteString("\nend")

	return out.String()
}

func (c *Class) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(c, method, env, args)
}
func init() {
	objectMethods[CLASS_OBJ] = map[string]ObjectMethod{
		"new": ObjectMethod{
			Layout: MethodLayout{
				Description: "Create a new instance of the class",
				ArgPattern: Args(
					OverloadArg(ANY_OBJ...),
				),
				ReturnPattern: Args(
					Arg(INSTANCE_OBJ),
				),
			},
			method: func(o Object, args []Object, env Environment) Object {
				clazz := o.(*Class)
				instance := &Instance{Class: clazz, Environment: NewEnclosedEnvironment(clazz.Env)}

				if _, ok := instance.Environment.Get("initialize"); ok {
					instance.InvokeMethod("initialize", *instance.Environment, args...)
				}

				return instance
			},
		},
	}
}
