package object

import (
	"strings"
	"unicode"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	if e.outer != nil {
		_, ok := e.outer.Get(name)
		if ok {
			e.outer.Set(name, val)
			return val
		}
	}
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Names(prefix string) []string {
	var ret []string

	for key := range e.store {
		if strings.HasPrefix(key, prefix) {
			ret = append(ret, key)
		}

		// Functions with an "object." prefix are available
		// to all object-methods.
		if strings.HasPrefix(key, "object.") {
			ret = append(ret, key)
		}
	}
	return ret
}

func (e *Environment) Exported() *Hash {
	pairs := make(map[HashKey]HashPair)

	for k, v := range e.store {
		// Replace this with checking for "Import" token
		if unicode.IsUpper(rune(k[0])) {
			s := NewString(k)
			pairs[s.HashKey()] = HashPair{Key: s, Value: v}
		}
	}

	return NewHash(pairs)
}
