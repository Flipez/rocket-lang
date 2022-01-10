package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		file, err := ioutil.ReadFile(os.Args[1])
		if err == nil {
			env := object.NewEnvironment()
			l := lexer.New(string(file))
			p := parser.New(l, make(map[string]struct{}))

			program, _ := p.ParseProgram()
			if len(p.Errors()) > 0 {
				printParserErrors(p.Errors())
				return
			}

			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				fmt.Println(evaluated.Inspect())
			}
		}
	}
}

func printParserErrors(errors []string) {
	fmt.Println("🔥 Great, you broke it!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
