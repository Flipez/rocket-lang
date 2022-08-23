//go:build !wasm

package repl

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

var buildVersion = "v0.10.0"
var buildDate = "2021-12-27T21:13:44Z"

func Start(in io.Reader, out io.Writer) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 "🚀 \033[31m»\033[0m ",
		HistoryFile:            homeDir + "/.rocket_history",
		InterruptPrompt:        "^C",
		DisableAutoSaveHistory: true,
	})

	if err != nil {
		panic(err)
	}

	defer rl.Close()

	env := object.NewEnvironment()
	imports := make(map[string]struct{})
	var cmds []string

	fmt.Println(SplashScreen())

	for {
		//source, err := line.Prompt("🚀 > ")
		line, err := rl.Readline()

		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		cmds = append(cmds, line)

		//if !strings.HasSuffix(line, ";") {
		//	rl.SetPrompt("🚀 \033[31m»»»\033[0m ")
		//	continue
		//}

		cmd := strings.Join(cmds, " ")
		rl.SetPrompt("🚀 \033[31m»\033[0m ")
		rl.SaveHistory(cmd)

		l := lexer.New(cmd)
		p := parser.New(l, imports)

		object.AddEvaluator(evaluator.Eval)

		var program *ast.Program

		program, imports = p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Println("❌ " + evaluated.Inspect())
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

func printParserErrors(errors []string) {
	fmt.Println("🔥 Great, you broke it!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
