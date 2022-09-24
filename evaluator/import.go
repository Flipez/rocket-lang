package evaluator

import (
	"path/filepath"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalImport(ie *ast.Import, env *object.Environment) object.Object {
	location := Eval(ie.Location, env)

	if object.IsError(location) {
		return location
	}

	if s, ok := location.(*object.String); ok {
		attributes := EvalModule(s.Value)

		if object.IsError(attributes) {
			return attributes
		}

		if ie.Name != nil {
			name := Eval(ie.Name, env)
			if nameString, ok := name.(*object.String); ok {
				env.Set(nameString.Value, object.NewModule(s.Value, attributes))
			}
		} else {
			env.Set(filepath.Base(s.Value), object.NewModule(s.Value, attributes))
		}
		return object.NIL
	}

	return object.NewErrorFormat("Import Error: invalid import path '%s'", location)
}
