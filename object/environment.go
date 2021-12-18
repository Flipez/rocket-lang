package object

import (
	"strings"
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
