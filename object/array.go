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

func (ao *Array) index(obj Object) int {
	for idx, element := range ao.Elements {
		if CompareObjects(element, obj) {
			return idx
		}
	}
	return -1
}

func init() {
	objectMethods[ARRAY_OBJ] = map[string]ObjectMethod{
		"join": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
				ArgPattern: Args(
					OptArg(STRING_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				arr := make([]string, len(ao.Elements))
				join := ""

				if len(args) > 0 {
					join = args[0].(*String).Value
				}

				for i, element := range ao.Elements {
					if e, ok := element.(Stringable); ok {
						arr[i] = e.ToStringObj(nil).Value
					} else {
						return NewErrorFormat("Found non stringable element %s on index %d", element.Type(), i)
					}
				}

				return NewString(strings.Join(arr, join))
			},
		},
		"reverse": ObjectMethod{
			Layout: MethodLayout{
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
		"sum": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				var result int

				for i, element := range ao.Elements {
					switch val := element.(type) {
					case Integerable:
						result += int(val.ToIntegerObj(nil).Value)
					case Floatable:
						result += int(val.ToFloatObj().Value)
					default:
						return NewErrorFormat("Found non number element %s on index %d", val.Type(), i)
					}
				}
				return NewInteger(int64(result))
			},
		},
		"uniq": ObjectMethod{
			Layout: MethodLayout{
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
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ArgPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FILE_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				return NewInteger(int64(ao.index(args[0])))
			},
		},
		"first": ObjectMethod{
			Layout: MethodLayout{
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
		"include?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
				ArgPattern: Args(
					Arg(STRING_OBJ, ARRAY_OBJ, HASH_OBJ, BOOLEAN_OBJ, INTEGER_OBJ, NIL_OBJ, FILE_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				if ao.index(args[0]) == -1 {
					return FALSE
				}
				return TRUE
			},
		},
		"slices": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				size := int(args[0].(*Integer).Value)
				if size == 0 {
					return NewError("invalid slice size, needs to be > 0")
				}

				length := len(ao.Elements)

				slices := NewArray(make([]Object, 0))
				for i := 0; i < length; i += size {
					end := i + size
					if end > length {
						end = length
					}
					slices.Add(NewArray(ao.Elements[i:end]))
				}

				return slices
			},
		},
	}
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(ao, method, env, args)
}

func (ao *Array) GetIterator(start, step int, _ bool) Iterator {
	return &arrayIterator{items: ao.Elements, index: start, step: step}
}

func (ao *Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(ao.Elements)
}

type arrayIterator struct {
	items []Object
	index int
	step  int
}

func (a *arrayIterator) Next() (Object, Object, bool) {
	if a.index < len(a.items) {
		val := a.items[a.index]
		idx := NewInteger(int64(a.index))
		a.index += a.step
		return val, idx, true
	}
	return nil, NewInteger(0), false
}
