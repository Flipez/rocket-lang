package evaluator

import (
	"path/filepath"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalImport(ie *ast.Import, env *object.Environment) object.Object {
	name := Eval(ie.Name, env)

	if object.IsError(name) {
		return name
	}

	if s, ok := name.(*object.String); ok {
		attributes := EvalModule(s.Value)

		if object.IsError(attributes) {
			return attributes
		}

		env.Set(filepath.Base(s.Value), object.NewModule(s.Value, attributes))
		return object.NULL
	}

	return object.NewErrorFormat("Import Error: invalid import path '%s'", name)
}
