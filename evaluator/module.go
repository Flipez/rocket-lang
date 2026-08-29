package evaluator

import (
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/utilities"
)

// EvalModuleFile evaluates an already-resolved module file in an isolated
// environment that shares reg, and returns its exported attributes.
func EvalModuleFile(filename string, reg *object.ModuleRegistry) object.Object {
	b, err := utilities.ReadFile(filename)

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
