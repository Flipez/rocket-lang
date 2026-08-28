package evaluator

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
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

	attributes := EvalModule(s.Value)
	if object.IsError(attributes) {
		return attributes
	}

	if len(ie.Only) > 0 {
		filtered := filterExports(ie, attributes, s.Value)
		if object.IsError(filtered) {
			return filtered
		}
		attributes = filtered
	}

	name := ie.Alias
	if name == "" {
		name = filepath.Base(s.Value)
	}

	env.Set(name, object.NewModule(s.Value, attributes))

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
