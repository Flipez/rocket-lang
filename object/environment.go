package object

import (
	"fmt"
	"os"
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
	// permit stores the names of variables we can set in this
	// environment, if any
	permit []string
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	if len(e.permit) > 0 {
		for _, v := range e.permit {
			// we're permitted to store this variable
			if v == name {
				e.store[name] = val
				return val
			}
		}
		// ok we're not permitted, we must store in the parent
		if e.outer != nil {
			return e.outer.Set(name, val)
		} else {
			fmt.Printf("scoping weirdness; please report a bug\n")
			os.Exit(5)
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

// NewTemporaryScope creates a temporary scope where some values
// are ignored.
//
// This is used as a sneaky hack to allow `foreach` to access all
// global values as if they were local, but prevent the index/value
// keys from persisting.
func NewTemporaryScope(outer *Environment, keys []string) *Environment {
	env := NewEnvironment()
	env.outer = outer
	env.permit = keys
	return env
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
