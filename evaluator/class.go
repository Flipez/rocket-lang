package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalClass(clazz *ast.Class, env *object.Environment) object.Object {

	class := &object.Class{
		Name:  clazz.Name,
		Env:   object.NewEnvironment(),
		Super: nil,
	}

	//classEnv := object.NewEnclosedEnvironment(env)

	Eval(clazz.Body, class.Env)

	env.Set(clazz.Name.Value, class)

	return class

}
