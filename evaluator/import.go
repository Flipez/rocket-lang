package evaluator

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
	"github.com/flipez/rocket-lang/token"
	"github.com/flipez/rocket-lang/utilities"
)

func evalImport(ie *ast.Import, env *object.Environment) object.Object {
	reg := env.Registry()

	// A `./` or `../` path resolves against the directory of the file doing
	// the importing. The REPL and `rocket-lang -e` have no such file, so they
	// anchor on the working directory: that is where their source effectively
	// lives, and it is already where a bare module name resolves from.
	importerDir := ""
	if ie.Token.File != "" {
		importerDir = filepath.Dir(ie.Token.File)
	} else if wd, err := os.Getwd(); err == nil {
		importerDir = wd
	}

	filename := utilities.FindModuleFrom(ie.Path, importerDir)
	if filename == "" {
		return object.NewErrorFormat("%s:%d:%d: Import Error: no module named '%s' found", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, ie.Path)
	}

	name := ie.Alias
	if name == "" {
		name = filepath.Base(ie.Path)

		// An implicit name comes from the path, which may contain characters
		// that cannot be referenced afterwards. "my-lib" binds, but my-lib.Foo
		// parses as subtraction, so the module would be unreachable. Planet
		// names routinely contain hyphens, which makes this common.
		if !isUsableBinding(name) {
			return object.NewErrorFormat("%s:%d:%d: Import Error: cannot bind module as '%s': not a usable name, use 'as' to choose one", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, name)
		}
	}

	if cached, ok := reg.Get(filename); ok {
		return bindModule(ie, env, name, cached, ie.Path)
	}

	if reg.InProgress(filename) {
		return object.NewErrorFormat("%s:%d:%d: Import Error: circular import\n  %s", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, reg.Chain(filename))
	}

	attributes := evalModuleOnce(filename, reg)

	if object.IsError(attributes) {
		return attributes
	}

	mod := object.NewModule(ie.Path, attributes)
	reg.Put(filename, mod)

	return bindModule(ie, env, name, mod, ie.Path)
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
	if nameInUse(env, name) {
		existing, _ := env.Get(name)
		existingMod, isModule := existing.(*object.Module)
		// A re-import is a no-op only if it would bind the exact same
		// namespace that's already bound: same resolved module, and the
		// same narrowing. A plain (non-`only`) import must be
		// pointer-identical to the cached instance (the module cache's
		// "same instance" contract); an `only` import always constructs a
		// fresh *object.Module, so it is compared by its resulting
		// attribute keys instead.
		sameBinding := isModule && sameNarrowing(ie, existingMod, mod)

		if !(env.RebindAllowed() && sameBinding) {
			return object.NewErrorFormat("%s:%d:%d: Import Error: cannot bind module as '%s', name already in use", ie.Token.File, ie.Token.LineNumber, ie.Token.LinePosition, name)
		}
	}

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

// sameNarrowing reports whether re-running this import would bind the exact
// same namespace as what's already bound under name: same resolved module
// (mod is always the freshly resolved/cached instance for this import) and
// the same narrowing as the existing binding. A no-`only` import must be
// pointer-identical to the cached module instance; an `only` import must
// match the existing binding's attribute keys exactly.
func sameNarrowing(ie *ast.Import, existing *object.Module, mod *object.Module) bool {
	if len(ie.Only) == 0 {
		return existing == mod
	}

	return sameExportedKeys(existing.Attributes, ie.Only)
}

// sameExportedKeys reports whether attrs (an *object.Hash of exported
// members) has exactly the keys named in only, no more and no fewer.
func sameExportedKeys(attrs object.Object, only []string) bool {
	hash, ok := attrs.(*object.Hash)
	if !ok {
		return false
	}

	if len(hash.Pairs) != len(only) {
		return false
	}

	for _, want := range only {
		key := object.NewString(want)
		if _, found := hash.Pairs[key.HashKey()]; !found {
			return false
		}
	}

	return true
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

// nameInUse reports whether name is already bound in scope or shadows a
// builtin function or module.
func nameInUse(env *object.Environment, name string) bool {
	if _, ok := env.Get(name); ok {
		return true
	}
	if _, ok := stdlib.Functions[name]; ok {
		return true
	}
	if _, ok := stdlib.Modules[name]; ok {
		return true
	}
	return false
}

// exportNames renders a module's export list for error messages, sorted so
// the output is stable across runs.
func exportNames(hash *object.Hash) string {
	names := make([]string, 0, len(hash.Pairs))
	for _, pair := range hash.Pairs {
		name := pair.Key.Inspect()
		if s, ok := pair.Key.(*object.String); ok {
			name = s.Value
		}
		names = append(names, "'"+name+"'")
	}
	sort.Strings(names)

	if len(names) == 0 {
		return "(none)"
	}

	return strings.Join(names, ", ")
}

// isUsableBinding reports whether name can be written in source and referenced.
// Rather than restate the lexer's rules, it asks the lexer: a usable name is one
// that lexes to exactly one identifier and nothing else.
func isUsableBinding(name string) bool {
	l := lexer.New(name, "")

	first := l.NextToken()
	if first.Type != token.IDENT || first.Literal != name {
		return false
	}

	return l.NextToken().Type == token.EOF
}
