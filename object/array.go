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

				if err := requireElements(ao.Elements, STRINGABLE); err != nil {
					return err
				}

				for i, element := range ao.Elements {
					arr[i] = element.(Stringable).ToStringObj().Value
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
				return NewArray(reversedElements(o.(*Array).Elements))
			},
		},
		"reverse!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = reversedElements(ao.Elements)

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
				return NewInteger(len(ao.Elements))
			},
		},
		"sort": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(ARRAY_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				sorted, err := sortedElements(o.(*Array).Elements)
				if err != nil {
					return err
				}

				return NewArray(sorted)
			},
		},
		"sort!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(ARRAY_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)

				sorted, err := sortedElements(ao.Elements)
				if err != nil {
					return err
				}
				ao.Elements = sorted

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

				if err := requireElements(ao.Elements, INTEGERABLE); err != nil {
					return err
				}

				for i, element := range ao.Elements {
					// Being convertible is not the same as converting: a string
					// is INTEGERABLE whether or not it parses as a number.
					integer, ok := element.(Integerable).ToIntegerObj().(*Integer)
					if !ok {
						return NewErrorFormat("element %d does not convert to a number, got %s", i, element.Type())
					}

					result += int(integer.Value)
				}
				return NewInteger(result)
			},
		},
		"uniq": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				unique, err := uniqueElements(o.(*Array).Elements)
				if err != nil {
					return err
				}

				return NewArray(unique)
			},
		},
		"uniq!": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)

				unique, err := uniqueElements(ao.Elements)
				if err != nil {
					return err
				}
				ao.Elements = unique

				return ao
			},
		},
		"index": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ArgPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				return NewInteger(ao.index(args[0]))
			},
		},
		"first": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				if len(args) > 0 {
					count := args[0].(*Integer).Value
					if count < 0 {
						return NewErrorFormat("negative count %d", count)
					}

					elements := ao.Elements
					if count > len(elements) {
						count = len(elements)
					}

					return NewArray(copyElements(elements[:count]))
				}

				if len(ao.Elements) == 0 {
					return NIL
				}

				return ao.Elements[0]
			},
		},
		"last": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				if len(args) > 0 {
					count := args[0].(*Integer).Value
					if count < 0 {
						return NewErrorFormat("negative count %d", count)
					}

					elements := ao.Elements
					if count > len(elements) {
						count = len(elements)
					}

					return NewArray(copyElements(elements[len(elements)-count:]))
				}

				if len(ao.Elements) == 0 {
					return NIL
				}

				return ao.Elements[len(ao.Elements)-1]
			},
		},
		"pop": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				if length == 0 {
					return NIL
				}

				newElements := make([]Object, length-1)
				copy(newElements, ao.Elements[:(length-1)])

				returnElement := ao.Elements[length-1]

				ao.Elements = newElements

				return returnElement
			},
		},
		"push": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
				ArgPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = append(ao.Elements, args[0])

				return ao
			},
		},
		"include?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
				ArgPattern: Args(
					Arg(ANY),
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
		"to_m": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(MATRIX_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				matrix, err := NewMatrixFromObjects(ao.Elements)
				if err != nil {
					return NewErrorFormat("failed to convert array to matrix: %s", err.Error())
				}
				return matrix
			},
		},
		"empty?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if len(o.(*Array).Elements) == 0 {
					return TRUE
				}

				return FALSE
			},
		},
		"count": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(ANY),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				// Without an argument this is size(). With one it counts how
				// often that element occurs, which is what index() cannot tell
				// you.
				if len(args) == 0 {
					return NewInteger(len(ao.Elements))
				}

				count := 0
				for _, element := range ao.Elements {
					if CompareObjects(element, args[0]) {
						count++
					}
				}

				return NewInteger(count)
			},
		},
		"rindex": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(ANY),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				// -1 when absent, the same answer index() gives.
				for i := len(ao.Elements) - 1; i >= 0; i-- {
					if CompareObjects(ao.Elements[i], args[0]) {
						return NewInteger(i)
					}
				}

				return NewInteger(-1)
			},
		},
		"min": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return extremeElement(o.(*Array).Elements, true)
			},
		},
		"max": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return extremeElement(o.(*Array).Elements, false)
			},
		},
		"shift": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)

				// The mirror of pop: it changes the array and hands back the
				// element, so it has no ! either.
				if len(ao.Elements) == 0 {
					return NIL
				}

				first := ao.Elements[0]
				ao.Elements = copyElements(ao.Elements[1:])

				return first
			},
		},
		"unshift": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(ANY),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = append([]Object{args[0]}, ao.Elements...)

				return ao
			},
		},
		"insert": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
					Arg(ANY),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				at := args[0].(*Integer).Value
				if at < 0 {
					at = length + at + 1
				}
				// Inserting at length appends. Anything past that would need
				// the array padded with nils, which is a surprise rather than a
				// convenience.
				if at < 0 || at > length {
					return NewErrorFormat("index out of range, got %d but array has only %d elements", args[0].(*Integer).Value, length)
				}

				elements := make([]Object, 0, length+1)
				elements = append(elements, ao.Elements[:at]...)
				elements = append(elements, args[1])
				elements = append(elements, ao.Elements[at:]...)
				ao.Elements = elements

				return ao
			},
		},
		"delete": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(ANY),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				kept := make([]Object, 0, len(ao.Elements))
				found := false
				for _, element := range ao.Elements {
					if CompareObjects(element, args[0]) {
						found = true
						continue
					}
					kept = append(kept, element)
				}
				ao.Elements = kept

				// The element when something went, nil when nothing did, so the
				// caller can tell the two apart.
				if !found {
					return NIL
				}

				return args[0]
			},
		},
		"delete_at": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				length := len(ao.Elements)

				at := args[0].(*Integer).Value
				if at < 0 {
					at = length + at
				}
				// nil rather than an error for a position that is not there,
				// the same answer first() and pop() give for an empty array.
				if at < 0 || at >= length {
					return NIL
				}

				removed := ao.Elements[at]
				elements := make([]Object, 0, length-1)
				elements = append(elements, ao.Elements[:at]...)
				elements = append(elements, ao.Elements[at+1:]...)
				ao.Elements = elements

				return removed
			},
		},
		"clear": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = make([]Object, 0)

				return ao
			},
		},
		"concat": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(ARRAY_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)
				ao.Elements = append(ao.Elements, args[0].(*Array).Elements...)

				return ao
			},
		},
		"take": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				count := args[0].(*Integer).Value
				if count < 0 {
					return NewErrorFormat("negative count %d", count)
				}
				if count > len(ao.Elements) {
					count = len(ao.Elements)
				}

				return NewArray(copyElements(ao.Elements[:count]))
			},
		},
		"drop": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				ao := o.(*Array)

				count := args[0].(*Integer).Value
				if count < 0 {
					return NewErrorFormat("negative count %d", count)
				}
				if count > len(ao.Elements) {
					count = len(ao.Elements)
				}

				return NewArray(copyElements(ao.Elements[count:]))
			},
		},
	}

	// Pure method and ! counterpart from one transformation each, the same way
	// string.go registers its pairs.
	arrayPair("compact", nil, func(elements []Object, _ []Object) ([]Object, Object) {
		return compactElements(elements), nil
	})
	arrayPair("flatten", Args(OptArg(INTEGER_OBJ)), func(elements []Object, args []Object) ([]Object, Object) {
		depth := -1
		if len(args) > 0 {
			depth = args[0].(*Integer).Value
			if depth < 0 {
				return nil, NewErrorFormat("negative depth %d", depth)
			}
		}

		return flattenElements(elements, depth), nil
	})
	arrayPair("rotate", Args(OptArg(INTEGER_OBJ)), func(elements []Object, args []Object) ([]Object, Object) {
		by := 1
		if len(args) > 0 {
			by = args[0].(*Integer).Value
		}

		return rotatedElements(elements, by), nil
	})
}

// arrayPair registers a method and its in-place counterpart from one
// transformation, so the two halves cannot drift apart. See stringPair in
// string.go for why this is not written out twice.
func arrayPair(name string, argPattern []Argument, transform func(elements []Object, args []Object) ([]Object, Object)) {
	layout := MethodLayout{
		ArgPattern:    argPattern,
		ReturnPattern: Args(Arg(ARRAY_OBJ, ERROR_OBJ)),
	}

	objectMethods[ARRAY_OBJ][name] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			elements, err := transform(o.(*Array).Elements, args)
			if err != nil {
				return err
			}

			return NewArray(elements)
		},
	}

	objectMethods[ARRAY_OBJ][name+"!"] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			ao := o.(*Array)

			elements, err := transform(ao.Elements, args)
			if err != nil {
				return err
			}
			ao.Elements = elements

			return ao
		},
	}
}

// copyElements returns a fresh slice, so a method handing back a sub-slice
// cannot be written through to the array it came from.
func copyElements(src []Object) []Object {
	out := make([]Object, len(src))
	copy(out, src)

	return out
}

func compactElements(src []Object) []Object {
	out := make([]Object, 0, len(src))
	for _, element := range src {
		if element.Type() == NIL_OBJ {
			continue
		}
		out = append(out, element)
	}

	return out
}

// flattenElements inlines nested arrays. A depth below zero means all the way
// down; zero means leave them alone.
func flattenElements(src []Object, depth int) []Object {
	out := make([]Object, 0, len(src))
	for _, element := range src {
		nested, ok := element.(*Array)
		if !ok || depth == 0 {
			out = append(out, element)
			continue
		}

		next := depth - 1
		if depth < 0 {
			next = depth
		}
		out = append(out, flattenElements(nested.Elements, next)...)
	}

	return out
}

// rotatedElements moves the first by elements to the end. A negative count
// rotates the other way, and the count wraps, so rotating a 3-element array by
// 4 is the same as by 1.
func rotatedElements(src []Object, by int) []Object {
	if len(src) == 0 {
		return copyElements(src)
	}

	at := by % len(src)
	if at < 0 {
		at += len(src)
	}

	return append(copyElements(src[at:]), src[:at]...)
}

// extremeElement returns the smallest or largest element. It sorts a copy so
// that the rule about what can be compared, and the error when it cannot, are
// exactly sort's.
func extremeElement(src []Object, min bool) Object {
	if len(src) == 0 {
		return NIL
	}

	sorted, err := sortedElements(src)
	if err != nil {
		return err
	}

	if min {
		return sorted[0]
	}

	return sorted[len(sorted)-1]
}

func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(ao, method, env, args)
}

func (ao *Array) GetIterator(start, step int, _ bool) Iterator {
	return &arrayIterator{items: ao.Elements, index: start, step: step}
}

// ToStringObj renders the array the way Inspect does, which is what Ruby's
// Array#to_s does too. Without it the generic to_s fell through to an empty
// string, so [1,2].to_s() was "" while to_json() worked.
func (ao *Array) ToStringObj() *String {
	return NewString(ao.Inspect())
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
		idx := NewInteger(a.index)
		a.index += a.step
		return val, idx, true
	}
	return nil, NewInteger(0), false
}

// requireElements returns an error naming the first element outside the group,
// or nil. The four requirements on elements -- STRINGABLE for join, INTEGERABLE
// for sum, HASHABLE for uniq and COMPARABLE for sort -- used to be four checks
// with four unrelated messages, one of which did not say which element was at
// fault.
func requireElements(elements []Object, group string) Object {
	for i, element := range elements {
		if !InGroup(group, element) {
			return NewErrorFormat("element %d is not %s, got %s", i, group, element.Type())
		}
	}

	return nil
}

// reversedElements returns a reversed copy, leaving src untouched.
func reversedElements(src []Object) []Object {
	out := make([]Object, len(src))
	for i, element := range src {
		out[len(src)-1-i] = element
	}

	return out
}

// sortedElements returns a sorted copy of src. The second value is an error
// object when the elements are not a single COMPARABLE type, in which case the
// copy must not be used -- working on a copy means a failed sort cannot leave a
// half-ordered array behind.
//
// Both requirements are checked before any comparison runs, so the comparison
// functions can assert without a fallback. They used to carry an ok check each
// and set a flag that sort.SliceStable might or might not have reached.
//
// Can be refactored to generics once
// https://github.com/golang/go/issues/48522 is fixed.
func sortedElements(src []Object) ([]Object, Object) {
	out := make([]Object, len(src))
	copy(out, src)

	if len(out) == 0 {
		return out, nil
	}

	// Two separate requirements, and they used to share one message that named
	// neither the element nor which of the two had failed: [1, nil].sort() and
	// [1, 2.5].sort() both said "an object not INTEGER, FLOAT or STRING or is
	// mixed".
	if err := requireElements(out, COMPARABLE); err != nil {
		return nil, err
	}

	wanted := out[0].Type()
	for i, element := range out {
		if element.Type() != wanted {
			return nil, NewErrorFormat("elements must all be one %s type, got %s at 0 and %s at %d", COMPARABLE, wanted, element.Type(), i)
		}
	}

	switch wanted {
	case FLOAT_OBJ:
		sort.SliceStable(out, func(i, j int) bool {
			return out[i].(*Float).Value < out[j].(*Float).Value
		})
	case INTEGER_OBJ:
		sort.SliceStable(out, func(i, j int) bool {
			return out[i].(*Integer).Value < out[j].(*Integer).Value
		})
	case STRING_OBJ:
		sort.SliceStable(out, func(i, j int) bool {
			return out[i].(*String).Value < out[j].(*String).Value
		})
	}

	return out, nil
}

// uniqueElements returns a copy with duplicates removed, keeping the order of
// first appearance. Building the result from a map, as this once did, made the
// order depend on Go's map iteration and therefore vary between runs.
func uniqueElements(src []Object) ([]Object, Object) {
	seen := make(map[HashKey]struct{}, len(src))
	out := make([]Object, 0, len(src))

	if err := requireElements(src, HASHABLE); err != nil {
		return nil, err
	}

	for _, element := range src {
		key := element.(Hashable).HashKey()
		if _, duplicate := seen[key]; duplicate {
			continue
		}

		seen[key] = struct{}{}
		out = append(out, element)
	}

	return out, nil
}
