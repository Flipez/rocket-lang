//go:build !wasm

package repl

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

var buildVersion = "v0.10.0"
var buildDate = "2021-12-27T21:13:44Z"

const prompt = "🚀 \033[31m»\033[0m "

// Session is what a REPL session accumulates: one environment, shared by every
// line, so a variable defined on one line is there on the next.
//
// It exists so the read-eval-print part can be driven without a terminal.
// Everything used to sit inside Start, wired to readline and fmt.Println, and
// the package had no tests at all -- which is how a parse error came to end the
// session.
type Session struct {
	env *object.Environment
	out io.Writer
}

// NewSession starts a session that writes to out.
func NewSession(out io.Writer) *Session {
	env := object.NewEnvironment()

	// The REPL lets a name be reassigned, which a module may not do.
	env.AllowRebind()

	return &Session{env: env, out: out}
}

// Eval runs one line and writes whatever the user should see.
//
// A line that does not parse is reported and the session carries on. It used to
// return from Start instead, so a single typo ended the session and took every
// variable defined in it -- while a runtime error, the far more common mistake,
// correctly carried on.
func (s *Session) Eval(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	p := parser.New(lexer.New(line, ""))
	program := p.ParseProgram()

	if errors := p.Errors(); len(errors) > 0 {
		s.printParserErrors(errors)

		return
	}

	if evaluated := evaluator.Eval(program, s.env); evaluated != nil {
		fmt.Fprintln(s.out, "» "+evaluated.Inspect())
	}
}

func (s *Session) printParserErrors(errors []string) {
	fmt.Fprintln(s.out, "🔥 Great, you broke it!")
	fmt.Fprintln(s.out, " parser errors:")

	for _, msg := range errors {
		fmt.Fprintf(s.out, "\t %s\n", msg)
	}
}

// Start reads lines from in and writes to out until in is exhausted.
func Start(in io.Reader, out io.Writer) {
	config := &readline.Config{
		Prompt:                 prompt,
		InterruptPrompt:        "^C",
		DisableAutoSaveHistory: true,
		Stdin:                  io.NopCloser(in),
		Stdout:                 out,
	}

	// History is a convenience, not a requirement: without a home directory to
	// put the file in, the session runs without one rather than refusing to
	// start.
	if homeDir, err := os.UserHomeDir(); err == nil {
		config.HistoryFile = homeDir + "/.rocket_history"
	}

	rl, err := readline.NewEx(config)
	if err != nil {
		panic(err)
	}

	defer rl.Close()

	session := NewSession(out)

	fmt.Fprintln(out, SplashScreen())

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		rl.SaveHistory(line)
		session.Eval(line)
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
