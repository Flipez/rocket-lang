package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

type Hash struct {
	Pairs  map[HashKey]HashPair
	offset int
}

func NewHash(pairs map[HashKey]HashPair) *Hash {
	if pairs == nil {
		pairs = make(map[HashKey]HashPair)
	}
	return &Hash{Pairs: pairs}
}

type HashPair struct {
	Key   Object
	Value Object
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	length := len(h.Pairs)
	pairs := make([]string, length)
	var index int
	for _, pair := range h.Pairs {
		pairs[index] = fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect())
		index++
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (h *Hash) HashKey() HashKey {
	ha := fnv.New64a()
	ha.Write([]byte(h.Inspect()))

	return HashKey{Type: h.Type(), Value: ha.Sum64()}
}

func init() {
	objectMethods[HASH_OBJ] = map[string]ObjectMethod{
		"keys": ObjectMethod{
			description: "Returns the keys of the hash.",
			example: `🚀 > {"a": "1", "b": "2"}.keys()
=> ["a", "b"]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				h := o.(*Hash)

				keys := make([]Object, len(h.Pairs))

				i := 0
				for _, k := range h.Pairs {
					keys[i] = k.Key
					i++
				}

				return NewArray(keys)
			},
		},
		"values": ObjectMethod{
			description: "Returns the values of the hash.",
			example: `🚀 > {"a": "1", "b": "2"}.values()
=> ["2", "1"]`,
			returnPattern: [][]string{
				[]string{ARRAY_OBJ},
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				h := o.(*Hash)

				values := make([]Object, len(h.Pairs))

				i := 0
				for _, k := range h.Pairs {
					values[i] = k.Value
					i++
				}

				return NewArray(values)
			},
		},
	}
}

func (h *Hash) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(h, method, env, args)

}

func (h *Hash) Reset() {
	h.offset = 0
}

func (h *Hash) Next() (Object, Object, bool) {
	if h.offset < len(h.Pairs) {
		idx := 0

		for _, pair := range h.Pairs {
			if h.offset == idx {
				h.offset++
				return pair.Key, pair.Value, true
			}
			idx++
		}
	}

	return nil, NewInteger(0), false
}
