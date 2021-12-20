package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
	offset   int
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

func init() {
	objectMethods[ARRAY_OBJ] = map[string]ObjectMethod{
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
		"yoink": ObjectMethod{
			argPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FUNCTION_OBJ, FILE_OBJ},
			},
			method: func(o Object, args []Object) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				newElements := make([]Object, length+1, length+1)
				copy(newElements, ao.Elements)
				newElements[length] = args[0]

				ao.Elements = newElements
				return &Null{}
			},
		},
	}
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(ao, method, args)
}

func (ao *Array) Reset() {
	ao.offset = 0
}
func (ao *Array) Next() (Object, Object, bool) {
	if ao.offset < len(ao.Elements) {
		ao.offset++

		element := ao.Elements[ao.offset-1]
		return element, &Integer{Value: int64(ao.offset - 1)}, true
	}

	return nil, &Integer{Value: 0}, false
}
