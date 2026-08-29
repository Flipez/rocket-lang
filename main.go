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

// The process reports whether the program ran. Nothing did before: an uncaught
// error, a parse error and a missing file all exited 0, so anything scripting
// the interpreter -- a shell pipeline, CI -- could not tell success from
// failure, and `rocket-lang typo.rl` said nothing at all.
const (
	exitOK      = 0
	exitFailure = 1
)

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
		os.Exit(runProgram(*exec, ""))
	}

	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)

		return
	}

	os.Exit(runFile(os.Args[1]))
}

// runFile runs a program from disk. A file that cannot be read is reported
// rather than passed over in silence, which is what the missing else did.
func runFile(path string) int {
	file, err := utilities.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rocket-lang: cannot read %s: %s\n", path, err)

		return exitFailure
	}

	return runProgram(string(file), path)
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

func runProgram(input string, file string) int {
	env := object.NewEnvironment()
	l := lexer.New(input, file)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printParserErrors(p.Errors())

		return exitFailure
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated == nil {
		return exitOK
	}

	// The program's final value is printed whatever it is, because that is what
	// the interpreter does with a value. An error being that value means nothing
	// handled it -- an error aborts the rest of the program and becomes its
	// result -- so the process reports failure.
	fmt.Println(evaluated.Inspect())

	if object.IsError(evaluated) {
		return exitFailure
	}

	return exitOK
}

func printParserErrors(errors []string) {
	fmt.Fprintln(os.Stderr, "🔥 Great, you broke it!")
	fmt.Fprintln(os.Stderr, " parser errors:")
	for _, msg := range errors {
		fmt.Fprintf(os.Stderr, "\t %s\n", msg)
	}
}
