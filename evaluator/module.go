package evaluator

import (
	"os"

	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/utilities"
)

// EvalModule resolves name against the search paths and evaluates it.
// Kept for callers that only have a module name.
func EvalModule(name string) object.Object {
	filename := utilities.FindModule(name)

	if filename == "" {
		return object.NewErrorFormat("Import Error: no module named '%s' found", name)
	}

	return EvalModuleFile(filename, object.NewModuleRegistry())
}

// EvalModuleFile evaluates an already-resolved module file in an isolated
// environment that shares reg, and returns its exported attributes.
func EvalModuleFile(filename string, reg *object.ModuleRegistry) object.Object {
	b, err := os.ReadFile(filename)

	if err != nil {
		return object.NewErrorFormat("IO Error: error reading module '%s': %s", filename, err)
	}

	l := lexer.New(string(b), filename)
	p := parser.New(l)

	module := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return object.NewErrorFormat("Parse Error: %s", p.Errors())
	}

	env := object.NewModuleEnvironment(reg)
	if result := Eval(module, env); object.IsError(result) {
		return result
	}

	return env.Exported()
}
