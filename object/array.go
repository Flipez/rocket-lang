package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "size":
		return &Integer{Value: int64(len(ao.Elements))}
	case "yeet":
		length := len(ao.Elements)

		newElements := make([]Object, length-1, length-1)
		copy(newElements, ao.Elements[:(length-1)])

		returnElement := ao.Elements[length-1]

		ao.Elements = newElements

		return returnElement
	}

	return nil
}
