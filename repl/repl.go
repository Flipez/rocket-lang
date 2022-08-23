package repl

import (
	//"bufio"
	"fmt"
	"io"
	//"log"
	"os"
	//"strings"

	"github.com/peterh/liner"

	//"github.com/abiosoft/ishell/v2"
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

const PROMPT = ">> "

var buildVersion = "v0.10.0"
var buildDate = "2021-12-27T21:13:44Z"

func Start(in io.Reader, out io.Writer) {
	line := liner.NewLiner()
	defer line.Close()

	//shell := ishell.New()
	//shell.SetHomeHistoryPath(".rocket_history")
	//shell.SetOut(out)
	//shell.SetPrompt("🚀 > ")

	env := object.NewEnvironment()
	imports := make(map[string]struct{})

	line.SetCtrlCAborts(true)

	//if f, err := os.Open(history); err == nil {
	//	line.ReadHistory(f)
	//	f.Close()
	//}

	//if f, err := os.Create(history); err != nil {
	//	log.Error("system error: unable to write to history file: %s", err)
	//} else {
	//	line.WriteHistory(f)
	//	f.Close()
	//}

	for {
		source, err := line.Prompt("🚀 > ")

		if err == liner.ErrPromptAborted {
			//log.Info("Exiting...")
			os.Exit(1)
		} else {
			l := lexer.New(source)
			p := parser.New(l, imports)

			object.AddEvaluator(evaluator.Eval)

			var program *ast.Program

			program, imports = p.ParseProgram()
			if len(p.Errors()) > 0 {
				//printParserErrors(source, p.Errors())
				return
			}

			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				//ctx.Println("=> " + evaluated.Inspect())
			}
		}
	}
}

const ROCKET = `
   /\
  (  )     ___         _       _   _
  (  )    | _ \___  __| |_____| |_| |   __ _ _ _  __ _
 /|/\|\   |   / _ \/ _| / / -_)  _| |__/ _  | ' \/ _  |
/_||||_\  |_|_\___/\__|_\_\___|\__|____\__,_|_||_\__, |
              %10s | %-15s   |___/
`

func SplashScreen() string {
	return fmt.Sprintf(ROCKET, buildVersion, buildDate)
}

func SplashVersion() string {
	return fmt.Sprintf("rocket-lang version %s (%s)\n", buildVersion, buildDate)
}

// func printParserErrors(ctx *ishell.Context, errors []string) {
// 	ctx.Println("🔥 Great, you broke it!")
// 	ctx.Println(" parser errors:")
// 	for _, msg := range errors {
// 		ctx.Printf("\t %s\n", msg)
// 	}
// }
