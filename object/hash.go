package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strings"
)

type Hash struct {
	Pairs map[HashKey]HashPair
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

func (h *Hash) Get(name string) (Object, bool) {
	pair, ok := h.Pairs[NewString(name).HashKey()]

	if ok {
		return pair.Value, ok
	}

	return nil, ok
}

func (h *Hash) Set(key, value any) {
	var keyObj, valObj Object
	if obj, ok := key.(Object); ok {
		keyObj = obj
	} else {
		keyObj = AnyToObject(key)
	}
	if obj, ok := value.(Object); ok {
		valObj = obj
	} else {
		valObj = AnyToObject(value)
	}
	h.Pairs[keyObj.(Hashable).HashKey()] = HashPair{Key: keyObj, Value: valObj}
}

func init() {
	objectMethods[HASH_OBJ] = map[string]ObjectMethod{
		"keys": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
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
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
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
		"include?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
				ArgPattern: Args(
					Arg(BOOLEAN_OBJ, STRING_OBJ, INTEGER_OBJ, FLOAT_OBJ, ARRAY_OBJ, HASH_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)
				key := args[0].(Hashable)
				if _, ok := h.Pairs[key.HashKey()]; ok {
					return TRUE
				}
				return FALSE
			},
		},
		"get": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(ANY_OBJ...),
					Arg(ANY_OBJ...),
				),
				ReturnPattern: Args(
					Arg(ANY_OBJ...),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)
				k := args[0].(Hashable)
				if pair, ok := h.Pairs[k.HashKey()]; ok {
					return pair.Value
				}
				return args[1]
			},
		},
	}
}

func (h *Hash) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(h, method, env, args)

}

func (h *Hash) GetIterator(_, _ int, _ bool) Iterator {
	pairs := make([]HashPair, 0)
	for _, val := range h.Pairs {
		pairs = append(pairs, val)
	}
	return &hashIterator{pairs: pairs}
}

func (h *Hash) MarshalJSON() ([]byte, error) {
	tempHash := make(map[string]Serializable)
	for _, pair := range h.Pairs {
		_, ok := pair.Key.(Serializable)
		if !ok {
			return nil, fmt.Errorf("unable to serialize key: %s", pair.Key.Inspect())
		}
		serializableValue, ok := pair.Value.(Serializable)
		if !ok {
			return nil, fmt.Errorf("unable to serialize value: %s", pair.Key.Inspect())
		}

		if str, ok := pair.Key.(*String); ok {
			tempHash[str.Value] = serializableValue
		} else {
			tempHash[pair.Key.Inspect()] = serializableValue
		}
	}

	return json.Marshal(tempHash)
}

type hashIterator struct {
	pairs []HashPair
	index int
}

func (h *hashIterator) Next() (Object, Object, bool) {
	if h.index < len(h.pairs) {
		pair := h.pairs[h.index]
		h.index++
		return pair.Value, pair.Key, true
	}
	return nil, NewInteger(0), false
}
