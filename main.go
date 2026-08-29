package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/planet"
	"github.com/flipez/rocket-lang/repl"
	"github.com/flipez/rocket-lang/utilities"
)

// subcommands are matched against the first argument before it is treated as a
// program file. A file whose name collides with one is reached as ./name, or by
// giving it the usual .rl extension.
var subcommands = map[string]func([]string, io.Writer, io.Writer) int{
	"planet": planet.Command,
}

func main() {
	if len(os.Args) > 1 {
		if run, ok := subcommands[os.Args[1]]; ok {
			os.Exit(run(os.Args[2:], os.Stdout, os.Stderr))
		}
	}

	version := flag.BoolP("version", "v", false, "Prints the version and build date.")
	exec := flag.StringP("exec", "e", "", "Runs the given code.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: rocket-lang [flags] [program file] [arguments]\n")
		fmt.Fprintf(os.Stderr, "       rocket-lang planet <command>\n\nAvailable flags:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	configureSearchPath()

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

// configureSearchPath appends the current project's planet directory to the
// module search path. It goes after ROCKETLANGPATH and the working directory,
// so a project's own modules always win a name clash with an installed planet.
func configureSearchPath() {
	utilities.InitSearchPaths()

	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	root, ok := planet.FindRoot(cwd)
	if !ok {
		return
	}

	planetsDir := filepath.Join(root, planet.DirName)
	if !utilities.Exists(planetsDir) {
		return
	}

	if err := utilities.AddPath(planetsDir); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not add %s to the search path: %s\n", planetsDir, err)
	}
}

func runProgram(input string, file string) {
	env := object.NewEnvironment()
	l := lexer.New(input, file)
	p := parser.New(l)

	object.AddEvaluator(evaluator.Eval)

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

func printParserErrors(errors []string) {
	fmt.Println("🔥 Great, you broke it!")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Printf("\t %s\n", msg)
	}
}
