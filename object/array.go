package object

import (
	"bytes"
	"encoding/json"
	"hash/fnv"
	"strings"
)

type Array struct {
	Elements []Object
	offset   int
}

func NewArray(slice []Object) *Array {
	return &Array{Elements: slice}
}

func NewArrayWithObjects(objs ...Object) *Array {
	slice := make([]Object, len(objs))
	copy(slice, objs)
	return NewArray(slice)
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	length := len(ao.Elements)
	elements := make([]string, length)
	for index, element := range ao.Elements {
		elements[index] = element.Inspect()
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ao *Array) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(ao.Inspect()))

	return HashKey{Type: ao.Type(), Value: h.Sum64()}
}

func init() {
	objectMethods[ARRAY_OBJ] = map[string]ObjectMethod{
		"size": ObjectMethod{
			description: "Returns the amount of elements in the array.",
			example: `🚀 > ["a", "b", 1, 2].size()
=> 4`,
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				return NewInteger(int64(len(ao.Elements)))
			},
		},
		"uniq": ObjectMethod{
			description: "Returns a copy of the array with deduplicated elements. Raises an error if a element is not hashable.",
			example: `🚀 > ["a", 1, 1, 2].uniq()
=> [1, 2, "a"]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ, ERROR_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)

				items := make(map[HashKey]Object)
				for _, element := range ao.Elements {
					helper, ok := element.(Hashable)
					if !ok {
						return NewErrorFormat("failed because element %s is not hashable", element.Type())
					}
					items[helper.HashKey()] = element
				}

				length := len(items)
				newElements := make([]Object, length)
				var idx int
				for _, item := range items {
					newElements[idx] = item
					idx++
				}

				return NewArray(newElements)
			},
		},
		"index": ObjectMethod{
			description: "Returns the index of the given element in the array if found. Otherwise return `-1`.",
			example: `🚀 > ["a", "b", 1, 2].index(1)
=> 2`,
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			argPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FILE_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				index := -1
				for idx, element := range ao.Elements {
					if CompareObjects(element, args[0]) {
						index = idx
						break
					}
				}

				return NewInteger(int64(index))
			},
		},
		"first": ObjectMethod{
			description: "Returns the first element of the array. Shorthand for `array[0]`",
			example: `🚀 > ["a", "b", 1, 2].first()
=> "a"`,
			returnPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FUNCTION_OBJ, FILE_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				if len(ao.Elements) == 0 {
					return NULL
				}
				return ao.Elements[0]
			},
		},
		"last": ObjectMethod{
			description: "Returns the last element of the array.",
			example: `🚀 > ["a", "b", 1, 2].last()
=> 2`,
			returnPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FUNCTION_OBJ, FILE_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				if len(ao.Elements) == 0 {
					return NULL
				}
				return ao.Elements[len(ao.Elements)-1]
			},
		},
		"yeet": ObjectMethod{
			description: "Removes the last element of the array and returns it.",
			example: `🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yeet()
=> 3
🚀 > a
=> [1, 2]`,
			returnPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FUNCTION_OBJ, FILE_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				newElements := make([]Object, length-1)
				copy(newElements, ao.Elements[:(length-1)])

				returnElement := ao.Elements[length-1]

				ao.Elements = newElements

				return returnElement
			},
		},
		"yoink": ObjectMethod{
			description: "Adds the given object as last element to the array.",
			example: `🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yoink("a")
=> null
🚀 > a
=> [1, 2, 3, "a"]`,
			returnPattern: [][]string{
				[]string{NULL_OBJ},
			},
			argPattern: [][]string{
				[]string{STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NULL_OBJ, FUNCTION_OBJ, FILE_OBJ},
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				newElements := make([]Object, length+1)
				copy(newElements, ao.Elements)
				newElements[length] = args[0]

				ao.Elements = newElements
				return NULL
			},
		},
	}
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(ao, method, env, args)
}

func (ao *Array) Reset() {
	ao.offset = 0
}
func (ao *Array) Next() (Object, Object, bool) {
	if ao.offset < len(ao.Elements) {
		ao.offset++

		element := ao.Elements[ao.offset-1]
		return element, NewInteger(int64(ao.offset - 1)), true
	}

	return nil, NewInteger(0), false
}

func (ao *Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(ao.Elements)
}
