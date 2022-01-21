package main

import (
	"fmt"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/repl"
)

func main() {
	version := flag.BoolP("version", "v", false, "Prints the version and build date.")
	exec := flag.StringP("exec", "e", "", "Runs the given code.")

	flag.Parse()

	if *version {
		print(repl.SplashVersion())
		return
	}

	if len(*exec) > 0 {
		runProgram(*exec)
		return
	}

	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		file, err := ioutil.ReadFile(os.Args[1])
		if err == nil {
			runProgram(string(file))
		}
	}
}

func runProgram(input string) {
	env := object.NewEnvironment()
	l := lexer.New(input)
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

func printParserErrors(errors []string) {
	fmt.Println("🔥 Great, you broke it!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
