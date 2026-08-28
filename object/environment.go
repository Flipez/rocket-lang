package object

import (
	"strings"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store:    s,
		outer:    nil,
		exports:  make(map[string]struct{}),
		registry: NewModuleRegistry(),
	}
}

// NewModuleEnvironment creates the isolated top-level environment a module
// body is evaluated in. It shares the importer's registry so nested imports
// see the same cache and the same in-progress set.
func NewModuleEnvironment(reg *ModuleRegistry) *Environment {
	env := NewEnvironment()
	env.registry = reg
	return env
}

type Environment struct {
	store       map[string]Object
	outer       *Environment
	exports     map[string]struct{}
	registry    *ModuleRegistry
	allowRebind bool
}

// AllowRebind relaxes the import shadowing check, for the REPL where
// re-entering the same import line is expected.
func (e *Environment) AllowRebind() { e.allowRebind = true }

func (e *Environment) RebindAllowed() bool { return e.allowRebind }

func (e *Environment) Registry() *ModuleRegistry {
	if e.outer != nil {
		return e.outer.Registry()
	}
	return e.registry
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

// MarkExport records name as part of this environment's public surface.
func (e *Environment) MarkExport(name string) {
	e.exports[name] = struct{}{}
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

	for name := range e.exports {
		val, ok := e.store[name]
		if !ok {
			continue
		}

		s := NewString(name)
		pairs[s.HashKey()] = HashPair{Key: s, Value: val}
	}

	return NewHash(pairs)
}
