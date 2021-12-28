package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
			l := lexer.New(processImports(string(file)))
			p := parser.New(l)

			program := p.ParseProgram()
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

func processImports(file string) string {
	lines := strings.Split(file, "\n")
	didReplace := true
	basePath := os.Args[1]
	r := regexp.MustCompile(`import "(?P<Month>.*)"`)

	for didReplace {
		didReplace = false
		for _, line := range lines {
			if strings.HasPrefix(line, "import") {
				importPath := r.FindStringSubmatch(line)[1]
				dir, _ := filepath.Split(basePath)
				importFile, err := ioutil.ReadFile(dir + importPath)
				if err != nil {
					panic(err)
				}
				file = strings.Replace(file, line, string(importFile), 1)
				didReplace = true
				break
			}
		}
		lines = strings.Split(file, "\n")
	}
	return file
}

func printParserErrors(errors []string) {
	fmt.Println("🔥 Great, you broke it!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
