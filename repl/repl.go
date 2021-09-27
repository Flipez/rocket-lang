package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

const PROMPT = ">> "

var buildVersion = "v0.9.0"
var buildDate = "2021-09-27"

func Start(in io.Reader, out io.Writer) {
	io.WriteString(out, fmt.Sprintf(ROCKET, buildVersion, buildDate))

	scanner := bufio.NewScanner((in))
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const ROCKET = `
   /\
  (  )     ___         _       _   _
  (  )    | _ \___  __| |_____| |_| |   __ _ _ _  __ _
 /|/\|\   |   / _ \/ _| / / -_)  _| |__/ _  | ' \/ _  |
/_||||_\  |_|_\___/\__|_\_\___|\__|____\__,_|_||_\__, |
                %10s | %-15s      |___/
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "🔥 Great, you broke it!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
