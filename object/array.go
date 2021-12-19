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

var arrayObjectMethods = map[string]ObjectMethod{
	"type": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			return &String{Value: string(o.Type())}
		},
	},
	"size": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			ao := o.(*Array)
			return &Integer{Value: int64(len(ao.Elements))}
		},
	},
	"yeet": ObjectMethod{
		method: func(o Object, _ []Object) Object {
			ao := o.(*Array)
			length := len(ao.Elements)

			newElements := make([]Object, length-1, length-1)
			copy(newElements, ao.Elements[:(length-1)])

			returnElement := ao.Elements[length-1]

			ao.Elements = newElements

			return returnElement
		},
	},
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "methods":
		return listObjectMethods(arrayObjectMethods)
	case "wat":
		return listObjectUsage(ao, arrayObjectMethods)
	default:
		if objMethod, ok := arrayObjectMethods[method]; ok {
			return objMethod.Call(ao, args)
		}
	}

	return nil
}
