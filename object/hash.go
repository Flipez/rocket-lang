package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Hash struct {
	Pairs map[HashKey]HashPair
}

type HashPair struct {
	Key   Object
	Value Object
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (h *Hash) InvokeMethod(method string, env Environment, args ...Object) Object {
	switch method {
	case "keys":
		keys := make([]Object, len(h.Pairs))

		i := 0
		for _, k := range h.Pairs {
			keys[i] = k.Key
			i++
		}

		return &Array{Elements: keys}
	}

	return nil
}
