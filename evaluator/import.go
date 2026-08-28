package evaluator

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/utilities"
)

func evalImport(ie *ast.Import, env *object.Environment) object.Object {
	location := Eval(ie.Location, env)

	if object.IsError(location) {
		return location
	}

	s, ok := location.(*object.String)
	if !ok {
		return object.NewErrorFormat("%s:%d:%d: Import Error: invalid import path '%s'", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, location.Inspect())
	}

	reg := env.Registry()

	filename := utilities.FindModule(s.Value)
	if filename == "" {
		return object.NewErrorFormat("%s:%d:%d: Import Error: no module named '%s' found", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, s.Value)
	}

	name := ie.Alias
	if name == "" {
		name = filepath.Base(s.Value)
	}

	if cached, ok := reg.Get(filename); ok {
		return bindModule(ie, env, name, cached, s.Value)
	}

	if reg.InProgress(filename) {
		return object.NewErrorFormat("%s:%d:%d: Import Error: circular import\n  %s", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, reg.Chain(filename))
	}

	attributes := evalModuleOnce(filename, reg)

	if object.IsError(attributes) {
		return attributes
	}

	mod := object.NewModule(s.Value, attributes)
	reg.Put(filename, mod)

	return bindModule(ie, env, name, mod, s.Value)
}

// evalModuleOnce evaluates a module file, bracketing the evaluation with
// reg.Begin/reg.End so the in-progress entry is always cleared as soon as
// evaluation finishes -- including on any early return inside
// EvalModuleFile -- without delaying the End past other imports handled by
// the caller.
func evalModuleOnce(filename string, reg *object.ModuleRegistry) object.Object {
	reg.Begin(filename)
	defer reg.End(filename)

	return EvalModuleFile(filename, reg)
}

// bindModule applies the import's `only` filter and binds the resulting
// namespace object into env. When the import has no `only` clause, the
// cached module instance itself is bound so that repeated imports of the
// same file share one *object.Module (per the module cache's "same
// instance" contract). A narrowed `only` import always gets its own
// *object.Module wrapping a filtered attributes hash.
func bindModule(ie *ast.Import, env *object.Environment, name string, mod *object.Module, path string) object.Object {
	if len(ie.Only) > 0 {
		filtered := filterExports(ie, mod.Attributes, path)
		if object.IsError(filtered) {
			return filtered
		}

		env.Set(name, object.NewModule(path, filtered))

		return object.NIL
	}

	env.Set(name, mod)

	return object.NIL
}

// filterExports narrows a module's attribute hash to the names listed in
// the import's `only` clause, erroring if the module does not export one.
func filterExports(ie *ast.Import, attributes object.Object, path string) object.Object {
	hash, ok := attributes.(*object.Hash)
	if !ok {
		return object.NewErrorFormat("%s:%d:%d: Import Error: module '%s' has no exports", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, path)
	}

	pairs := make(map[object.HashKey]object.HashPair)

	for _, want := range ie.Only {
		key := object.NewString(want)
		pair, found := hash.Pairs[key.HashKey()]
		if !found {
			return object.NewErrorFormat("%s:%d:%d: Import Error: '%s' does not export '%s'; exported: %s", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, path, want, exportNames(hash))
		}
		pairs[key.HashKey()] = pair
	}

	return object.NewHash(pairs)
}

// exportNames renders a module's export list for error messages, sorted so
// the output is stable across runs.
func exportNames(hash *object.Hash) string {
	names := make([]string, 0, len(hash.Pairs))
	for _, pair := range hash.Pairs {
		names = append(names, pair.Key.Inspect())
	}
	sort.Strings(names)

	if len(names) == 0 {
		return "(none)"
	}

	return strings.Join(names, ", ")
}
