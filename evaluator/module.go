package evaluator

import (
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/utilities"

	"io/ioutil"
)

func EvalModule(name string) object.Object {
	filename := utilities.FindModule(name)

	if filename == "" {
		return object.NewErrorFormat("Import Error: no module named '%s' found", name)
	}

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return object.NewErrorFormat("IO Error: error reading module '%s': %s", name, err)
	}

	l := lexer.New(string(b))
	imports := make(map[string]struct{})
	p := parser.New(l, imports)

	module, _ := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return object.NewErrorFormat("Parse Error: %s", p.Errors())
	}

	env := object.NewEnvironment()
	Eval(module, env)

	return env.Exported()
}
