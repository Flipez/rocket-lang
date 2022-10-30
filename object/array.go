package object

import (
	"bytes"
	"encoding/json"
	"hash/fnv"
	"sort"
	"strings"
)

type Array struct {
	Elements []Object
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

func (ao *Array) Add(items ...any) {
	for _, item := range items {
		obj, ok := item.(Object)
		if !ok {
			obj = AnyToObject(item)
		}
		ao.Elements = append(ao.Elements, obj)
	}
}

func init() {
	objectMethods[ARRAY_OBJ] = map[string]ObjectMethod{
		"reverse": ObjectMethod{
			Layout: MethodLayout{
				Description: "Reverses the elements of the array",
				Example: `🚀 > ["a", "b", 1, 2].reverse()
=> [2, 1, "b", "a"]`,
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)

				for i, j := 0, len(ao.Elements)-1; i < j; i, j = i+1, j-1 {
					ao.Elements[i], ao.Elements[j] = ao.Elements[j], ao.Elements[i]
				}
				return ao
			},
		},
		"size": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns the amount of elements in the array.",
				Example: `🚀 > ["a", "b", 1, 2].size()
=> 4`,
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				return NewInteger(int64(len(ao.Elements)))
			},
		},
		"sort": ObjectMethod{
			// Can be refactored to generics once
			// https://github.com/golang/go/issues/48522
			// is fixed
			Layout: MethodLayout{
				Description: "Sorts the array if it contains only one type of STRING, INTEGER or FLOAT",
				Example: `🚀 » [3.4, 3.1, 2.0].sort()
» [2.0, 3.1, 3.4]`,
				ReturnPattern: Args(Arg(ARRAY_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				sortError := false

				if len(ao.Elements) == 0 {
					return ao
				}

				switch ao.Elements[0].(type) {
				case *Float:
					sort.SliceStable(ao.Elements, func(i, j int) bool {
						leftElement, ok := ao.Elements[i].(*Float)
						if !ok {
							sortError = true
							return false
						}

						rightElement, ok := ao.Elements[j].(*Float)
						if !ok {
							sortError = true
							return false
						}

						return leftElement.Value < rightElement.Value
					})
				case *Integer:
					sort.SliceStable(ao.Elements, func(i, j int) bool {
						leftElement, ok := ao.Elements[i].(*Integer)
						if !ok {
							sortError = true
							return false
						}

						rightElement, ok := ao.Elements[j].(*Integer)
						if !ok {
							sortError = true
							return false
						}

						return leftElement.Value < rightElement.Value
					})
				case *String:
					sort.SliceStable(ao.Elements, func(i, j int) bool {
						leftElement, ok := ao.Elements[i].(*String)
						if !ok {
							sortError = true
							return false
						}

						rightElement, ok := ao.Elements[j].(*String)
						if !ok {
							sortError = true
							return false
						}

						return leftElement.Value < rightElement.Value
					})
				default:
					sortError = true
				}

				if sortError {
					return NewError("Array does contain either an object not INTEGER, FLOAT or STRING or is mixed")
				}
				return ao
			},
		},
		"uniq": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns a copy of the array with deduplicated elements. Raises an error if a element is not hashable.",
				Example: `🚀 > ["a", 1, 1, 2].uniq()
=> [1, 2, "a"]`,
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
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
			Layout: MethodLayout{
				Description: "Returns the index of the given element in the array if found. Otherwise return `-1`.",
				Example: `🚀 > ["a", "b", 1, 2].index(1)
=> 2`,
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ArgPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FILE_OBJ),
				),
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
			Layout: MethodLayout{
				Description: "Returns the first element of the array. Shorthand for `array[0]`",
				Example: `🚀 > ["a", "b", 1, 2].first()
=> "a"`,
				ReturnPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FUNCTION_OBJ, FILE_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				if len(ao.Elements) == 0 {
					return NIL
				}
				return ao.Elements[0]
			},
		},
		"last": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns the last element of the array.",
				Example: `🚀 > ["a", "b", 1, 2].last()
=> 2`,
				ReturnPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FUNCTION_OBJ, FILE_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				if len(ao.Elements) == 0 {
					return NIL
				}
				return ao.Elements[len(ao.Elements)-1]
			},
		},
		"yeet": ObjectMethod{
			Layout: MethodLayout{
				Description: "Removes the last element of the array and returns it.",
				Example: `🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yeet()
=> 3
🚀 > a
=> [1, 2]`,
				ReturnPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FUNCTION_OBJ, FILE_OBJ),
				),
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
			Layout: MethodLayout{
				Description: "Adds the given object as last element to the array.",
				Example: `🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yoink("a")
=> nil
🚀 > a
=> [1, 2, 3, "a"]`,
				ReturnPattern: Args(
					Arg(NIL_OBJ),
				),
				ArgPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FUNCTION_OBJ, FILE_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = append(ao.Elements, args[0])
				return NIL
			},
		},
	}
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(ao, method, env, args)
}

func (ao *Array) GetIterator() Iterator {
	return &arrayIterator{items: ao.Elements}
}

func (ao *Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(ao.Elements)
}

type arrayIterator struct {
	items []Object
	index int
}

func (a *arrayIterator) Next() (Object, Object, bool) {
	if a.index < len(a.items) {
		val := a.items[a.index]
		idx := NewInteger(int64(a.index))
		a.index++
		return val, idx, true
	}
	return nil, NewInteger(0), false
}
