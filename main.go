package main

import (
	"fmt"
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: rocket-lang [flags] [program file] [arguments]\n\nAvailable flags:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		print(repl.SplashVersion())
		return
	}

	if len(*exec) > 0 {
		runProgram(*exec, "")
		return
	}

	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		file, err := os.ReadFile(os.Args[1])
		if err == nil {
			runProgram(string(file), os.Args[1])
		}
	}
}

func runProgram(input string, file string) {
	env := object.NewEnvironment()
	l := lexer.New(input, file)
	p := parser.New(l, make(map[string]struct{}))

	object.AddEvaluator(evaluator.Eval)

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
