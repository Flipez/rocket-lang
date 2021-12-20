package repl

import (
	//"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/abiosoft/ishell/v2"
	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

const PROMPT = ">> "

var buildVersion = "v0.9.0"
var buildDate = "2021-09-27T21:13:44Z"

func Start(in io.Reader, out io.Writer) {
	shell := ishell.New()
	shell.SetHomeHistoryPath(".rocket_history")
	shell.SetOut(out)
	shell.SetPrompt("🚀 > ")

	env := object.NewEnvironment()

	shell.Println(fmt.Sprintf(ROCKET, buildVersion, buildDate))
	shell.NotFound(func(ctx *ishell.Context) {

		l := lexer.New(strings.Join(ctx.RawArgs, " "))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(ctx, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			ctx.Println("=> " + evaluated.Inspect())
		}
	})

	shell.Run()
}

const ROCKET = `
   /\
  (  )     ___         _       _   _
  (  )    | _ \___  __| |_____| |_| |   __ _ _ _  __ _
 /|/\|\   |   / _ \/ _| / / -_)  _| |__/ _  | ' \/ _  |
/_||||_\  |_|_\___/\__|_\_\___|\__|____\__,_|_||_\__, |
              %10s | %-15s   |___/
`

func printParserErrors(ctx *ishell.Context, errors []string) {
	ctx.Println("🔥 Great, you broke it!")
	ctx.Println(" parser errors:")
	for _, msg := range errors {
		ctx.Printf("\t %s\n", msg)
	}
}
