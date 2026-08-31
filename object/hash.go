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
	hashable, ok := keyObj.(Hashable)
	if !ok {
		// Called from Go rather than from the interpreter, so there is no
		// argument pattern guarding this one. Dropping the entry beats bringing
		// the process down over a key that could never be looked up anyway.
		return
	}

	h.Pairs[hashable.HashKey()] = HashPair{Key: keyObj, Value: valObj}
}

// hashKeyOf turns an argument into a hash key. The HASHABLE argument pattern
// already rejects anything that cannot be one, so the error is unreachable from
// the interpreter; it exists so that widening a pattern can never again turn
// into a panic. include? and get used to assert without checking, and
// {"a": 1}.get(nil, 0) brought the process down.
func hashKeyOf(o Object) (HashKey, Object) {
	hashable, ok := o.(Hashable)
	if !ok {
		return HashKey{}, NewErrorFormat("unusable as hash key: %s", o.Type())
	}

	return hashable.HashKey(), nil
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
					Arg(HASHABLE),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)

				key, err := hashKeyOf(args[0])
				if err != nil {
					return err
				}

				if _, ok := h.Pairs[key]; ok {
					return TRUE
				}

				return FALSE
			},
		},
		"get": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(HASHABLE),
					Arg(ANY),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)

				key, err := hashKeyOf(args[0])
				if err != nil {
					return err
				}

				if pair, ok := h.Pairs[key]; ok {
					return pair.Value
				}

				return args[1]
			},
		},
		"each": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(CALLABLE),
				),
				ReturnPattern: Args(
					Arg(HASH_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, env Environment) Object {
				h := o.(*Hash)

				// The order the entries arrive in is not defined and differs
				// between runs, the same caveat keys() carries. Use it for a
				// side effect on each entry, not to build something ordered.
				for _, pair := range h.Pairs {
					result := CallFunction(args[0], env, pair.Key, pair.Value)

					if IsError(result) {
						return result
					}

					if CallbackStopped(result) {
						break
					}
				}

				return h
			},
		},
		"size": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewInteger(len(o.(*Hash).Pairs))
			},
		},
		"empty?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if len(o.(*Hash).Pairs) == 0 {
					return TRUE
				}

				return FALSE
			},
		},
		"fetch": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(HASHABLE),
					OptArg(ANY),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)

				key, err := hashKeyOf(args[0])
				if err != nil {
					return err
				}

				if pair, found := h.Pairs[key]; found {
					return pair.Value
				}

				// Without a fallback a missing key is an error rather than nil.
				// That is the difference from get(), which always needs one.
				if len(args) > 1 {
					return args[1]
				}

				return NewErrorFormat("key not found: %s", args[0].Inspect())
			},
		},
		"delete": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(HASHABLE),
				),
				ReturnPattern: Args(
					Arg(ANY),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				h := o.(*Hash)

				hashed, err := hashKeyOf(args[0])
				if err != nil {
					return err
				}

				pair, found := h.Pairs[hashed]
				if !found {
					return NIL
				}
				delete(h.Pairs, hashed)

				// The value that went, so a delete can be told from a miss.
				return pair.Value
			},
		},
		"clear": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(HASH_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				h := o.(*Hash)
				h.Pairs = make(map[HashKey]HashPair)

				return h
			},
		},
		"invert": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(HASH_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				pairs := make(map[HashKey]HashPair)

				for _, pair := range o.(*Hash).Pairs {
					value, ok := pair.Value.(Hashable)
					if !ok {
						return NewErrorFormat("unusable as hash key: %s", pair.Value.Type())
					}
					// Duplicate values collapse into one entry, and which one
					// survives is not defined -- the same caveat Ruby carries.
					pairs[value.HashKey()] = HashPair{Key: pair.Value, Value: pair.Key}
				}

				return NewHash(pairs)
			},
		},
	}

	hashCallbackPair("select", func(pairs map[HashKey]HashPair, fn Object, env Environment) (map[HashKey]HashPair, Object) {
		return filteredPairs(pairs, fn, env, true)
	})
	hashCallbackPair("reject", func(pairs map[HashKey]HashPair, fn Object, env Environment) (map[HashKey]HashPair, Object) {
		return filteredPairs(pairs, fn, env, false)
	})
	hashCallbackPair("transform_values", func(pairs map[HashKey]HashPair, fn Object, env Environment) (map[HashKey]HashPair, Object) {
		out := make(map[HashKey]HashPair, len(pairs))

		for hashed, pair := range pairs {
			result := CallFunction(fn, env, pair.Value)

			switch ClassifyCallback(result) {
			case Failed:
				return nil, result
			case StopWalk:
				return out, nil
			case SkipElement:
				out[hashed] = HashPair{Key: pair.Key, Value: NIL}
			default:
				out[hashed] = HashPair{Key: pair.Key, Value: result}
			}
		}

		return out, nil
	})
	hashCallbackPair("transform_keys", func(pairs map[HashKey]HashPair, fn Object, env Environment) (map[HashKey]HashPair, Object) {
		out := make(map[HashKey]HashPair, len(pairs))

		for _, pair := range pairs {
			result := CallFunction(fn, env, pair.Key)

			switch ClassifyCallback(result) {
			case Failed:
				return nil, result
			case StopWalk:
				return out, nil
			case SkipElement:
				return nil, NewError("a key cannot be nothing: the callback of transform_keys ran next")
			}

			// A new key still has to be usable as one, and two keys mapping to
			// the same one collapse -- which of them survives is not defined.
			hashed, err := hashKeyOf(result)
			if err != nil {
				return nil, err
			}
			out[hashed] = HashPair{Key: result, Value: pair.Value}
		}

		return out, nil
	})

	hashPair("merge", Args(Arg(HASH_OBJ)), func(pairs map[HashKey]HashPair, args []Object) (map[HashKey]HashPair, Object) {
		merged := copyPairs(pairs)
		// The argument wins on a clash, as it does in Ruby.
		for hashed, pair := range args[0].(*Hash).Pairs {
			merged[hashed] = pair
		}

		return merged, nil
	})
	hashPair("compact", nil, func(pairs map[HashKey]HashPair, _ []Object) (map[HashKey]HashPair, Object) {
		kept := make(map[HashKey]HashPair, len(pairs))
		for hashed, pair := range pairs {
			if pair.Value.Type() == NIL_OBJ {
				continue
			}
			kept[hashed] = pair
		}

		return kept, nil
	})
}

// hashCallbackPair registers a method taking a callback and its in-place
// counterpart, the same arrangement as hashPair.
func hashCallbackPair(name string, transform func(pairs map[HashKey]HashPair, fn Object, env Environment) (map[HashKey]HashPair, Object)) {
	layout := MethodLayout{
		ArgPattern:    Args(Arg(CALLABLE)),
		ReturnPattern: Args(Arg(HASH_OBJ, ERROR_OBJ)),
	}

	objectMethods[HASH_OBJ][name] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, env Environment) Object {
			pairs, err := transform(o.(*Hash).Pairs, args[0], env)
			if err != nil {
				return err
			}

			return NewHash(pairs)
		},
	}

	objectMethods[HASH_OBJ][name+"!"] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, env Environment) Object {
			h := o.(*Hash)

			pairs, err := transform(h.Pairs, args[0], env)
			if err != nil {
				return err
			}
			h.Pairs = pairs

			return h
		},
	}
}

// filteredPairs keeps the entries the callback answers for, which way round
// decided by keep -- select keeps a yes, reject keeps a no.
func filteredPairs(pairs map[HashKey]HashPair, fn Object, env Environment, keep bool) (map[HashKey]HashPair, Object) {
	out := make(map[HashKey]HashPair, len(pairs))

	for hashed, pair := range pairs {
		result := CallFunction(fn, env, pair.Key, pair.Value)

		switch ClassifyCallback(result) {
		case Failed:
			return nil, result
		case StopWalk:
			return out, nil
		case SkipElement:
			continue
		default:
			if IsTruthy(result) == keep {
				out[hashed] = pair
			}
		}
	}

	return out, nil
}

// hashPair registers a method and its in-place counterpart from one
// transformation. See stringPair in string.go for why the two halves are not
// written out separately.
func hashPair(name string, argPattern []Argument, transform func(pairs map[HashKey]HashPair, args []Object) (map[HashKey]HashPair, Object)) {
	layout := MethodLayout{
		ArgPattern:    argPattern,
		ReturnPattern: Args(Arg(HASH_OBJ, ERROR_OBJ)),
	}

	objectMethods[HASH_OBJ][name] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			pairs, err := transform(o.(*Hash).Pairs, args)
			if err != nil {
				return err
			}

			return NewHash(pairs)
		},
	}

	objectMethods[HASH_OBJ][name+"!"] = ObjectMethod{
		Layout: layout,
		method: func(o Object, args []Object, _ Environment) Object {
			h := o.(*Hash)

			pairs, err := transform(h.Pairs, args)
			if err != nil {
				return err
			}
			h.Pairs = pairs

			return h
		},
	}
}

func copyPairs(src map[HashKey]HashPair) map[HashKey]HashPair {
	out := make(map[HashKey]HashPair, len(src))
	for hashed, pair := range src {
		out[hashed] = pair
	}

	return out
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

// ToStringObj renders the hash the way Inspect does. See Array.ToStringObj
// for why this is not left to the generic to_string.
func (h *Hash) ToStringObj() *String {
	return NewString(h.Inspect())
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
