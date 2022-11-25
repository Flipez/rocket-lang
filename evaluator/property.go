package evaluator

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func evalProperty(property *ast.Property, env *object.Environment) object.Object {
	left := Eval(property.Left, env)

	if object.IsError(left) {
		return left
	}

	switch left.(type) {
	case *object.Instance:
		p := property.Property.(*ast.Identifier)
		i := left.(*object.Instance)

		if v, ok := i.Environment.Get(p.Value); ok {
			return v
		}
		//switch prop := property.Property.(type) {
		//case *ast.Assign:
		//	p := prop.Value.(*ast.Identifier)
		//	i := left.(*object.Instance)

		//	if v, ok := i.Environment.Get(p.Value); ok {
		//		return v
		//	}
		//case *ast.Identifier:
		//	i := left.(*object.Instance)

		//	if v, ok := i.Environment.Get(prop.Value); ok {
		//		return v
		//	}
		//}
	}

	return object.NewErrorFormat("%d:%d: runtime error: unknown property: %s.%s", property.Token.LineNumber, property.Token.LinePosition, left.(*object.Instance).String(), property.Property)
}
