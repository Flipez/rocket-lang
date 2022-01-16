package evaluator

import (
	"path/filepath"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalImport(ie *ast.Import, env *object.Environment) object.Object {
	name := Eval(ie.Name, env)

	if isError(name) {
		return name
	}

	if s, ok := name.(*object.String); ok {
		attributes := EvalModule(s.Value)

		if isError(attributes) {
			return attributes
		}

		env.Set(filepath.Base(s.Value), &object.Module{Name: s.Value, Attributes: attributes})
		return &object.Null{}
	}

	return newError("Import Error: invalid import path '%s'", name)
}
