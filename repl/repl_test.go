//go:build !wasm

package repl

import (
	"bytes"
	"strings"
	"testing"
)

// run feeds lines to one session and returns everything it wrote.
func run(lines ...string) string {
	var out bytes.Buffer

	session := NewSession(&out)
	for _, line := range lines {
		session.Eval(line)
	}

	return out.String()
}

// TestParseErrorDoesNotEndTheSession is the reason this file exists. Start used
// to return when a line failed to parse, so a single typo ended the session and
// took every variable defined in it -- while a runtime error carried on.
func TestParseErrorDoesNotEndTheSession(t *testing.T) {
	output := run(`a = 21`, `)`, `a * 2`)

	if !strings.Contains(output, "parser errors:") {
		t.Errorf("expected the parse error to be reported, got:\n%s", output)
	}

	if !strings.Contains(output, "» 42") {
		t.Errorf("expected the line after the parse error to still be evaluated, got:\n%s", output)
	}
}

func TestRuntimeErrorDoesNotEndTheSession(t *testing.T) {
	output := run(`nope()`, `1 + 1`)

	if !strings.Contains(output, "identifier not found: nope") {
		t.Errorf("expected the runtime error to be reported, got:\n%s", output)
	}

	if !strings.Contains(output, "» 2") {
		t.Errorf("expected the next line to still be evaluated, got:\n%s", output)
	}
}

// TestEnvironmentPersistsAcrossLines covers the whole point of a session: one
// environment, shared by every line.
func TestEnvironmentPersistsAcrossLines(t *testing.T) {
	output := run(`a = 2`, `b = 3`, `a * b`)

	if !strings.Contains(output, "» 6") {
		t.Errorf("expected a and b to survive to the third line, got:\n%s", output)
	}
}

// TestReassignmentWorks is not what AllowRebind governs -- plain reassignment
// was never restricted -- but it is worth pinning, since it is what a reader
// does constantly in a REPL.
func TestReassignmentWorks(t *testing.T) {
	output := run(`a = 1`, `a = 2`, `a`)

	if strings.Contains(output, "ERROR") {
		t.Errorf("expected reassignment to work, got:\n%s", output)
	}

	if !strings.Contains(output, "» 2") {
		t.Errorf("expected a to be 2 after reassignment, got:\n%s", output)
	}
}

// TestReimportingTheSameNameIsAllowed is what NewSession's AllowRebind is for.
// Outside the REPL, binding a module to a name already in use is an error;
// re-entering an import line is normal at a prompt.
//
// The first version of this test asserted plain reassignment instead, and a
// mutation that deleted AllowRebind entirely left it passing.
func TestReimportingTheSameNameIsAllowed(t *testing.T) {
	output := run(
		`import "../fixtures/module" as M`,
		`import "../fixtures/module" as M`,
		`M.Sum(1, 2)`,
	)

	if strings.Contains(output, "name already in use") {
		t.Errorf("expected the second import to be allowed in a session, got:\n%s", output)
	}

	if !strings.Contains(output, "» 3") {
		t.Errorf("expected the module to still be usable, got:\n%s", output)
	}
}

func TestLinesThatProduceNoOutput(t *testing.T) {
	tests := []struct {
		name string
		line string
	}{
		{"empty", ""},
		{"whitespace only", "   \t  "},
		{"a comment", "# just a comment"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if output := run(tt.line); output != "" {
				t.Errorf("expected no output for %q, got %q", tt.line, output)
			}
		})
	}
}

// TestResultsAreInspected pins the shape of a result line, which is what a
// reader sees for every expression they type.
func TestResultsAreInspected(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{`1 + 1`, "» 2\n"},
		{`"abc"`, "» \"abc\"\n"},
		{`[1, "a"]`, "» [1, \"a\"]\n"},
		{`nil`, "» nil\n"},
		{`print("hi")`, "» nil\n"},
		{`1.5`, "» 1.5\n"},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			// print writes to stdout rather than to the session, so only the
			// result line is compared.
			if output := run(tt.line); output != tt.expected {
				t.Errorf("for %q expected %q, got %q", tt.line, tt.expected, output)
			}
		})
	}
}

func TestParserErrorsAreAllReported(t *testing.T) {
	output := run(`[1,`)

	if !strings.Contains(output, "🔥 Great, you broke it!") {
		t.Errorf("expected the banner, got:\n%s", output)
	}

	// Each error is on its own indented line.
	if !strings.Contains(output, "\n\t ") {
		t.Errorf("expected the errors to be indented, got:\n%s", output)
	}
}

func TestSplashScreenCarriesTheBuildInformation(t *testing.T) {
	splash := SplashScreen()

	for _, want := range []string{buildVersion, buildDate} {
		if !strings.Contains(splash, want) {
			t.Errorf("expected the splash screen to mention %q, got:\n%s", want, splash)
		}
	}

	// The name is ASCII art, so there is no literal "RocketLang" to look for.
	// The nose cone is the one part that is a fixed string.
	if !strings.Contains(splash, `/\`) {
		t.Errorf("expected the splash screen to draw the rocket, got:\n%s", splash)
	}

	if lines := strings.Count(splash, "\n"); lines < 6 {
		t.Errorf("expected the art to span several lines, got %d:\n%s", lines, splash)
	}
}

func TestSplashVersion(t *testing.T) {
	version := SplashVersion()

	if !strings.Contains(version, buildVersion) || !strings.Contains(version, buildDate) {
		t.Errorf("expected the version line to carry both, got %q", version)
	}

	if !strings.HasSuffix(version, "\n") {
		t.Errorf("expected the version line to end in a newline, got %q", version)
	}
}

// startOutput drives Start with the given input and returns what it wrote.
//
// HOME is redirected to a temporary directory because Start saves history: run
// against the real one, the tests would append their input to the reader's
// ~/.rocket_history.
func startOutput(t *testing.T, input string) string {
	t.Helper()
	t.Setenv("HOME", t.TempDir())

	var out bytes.Buffer
	Start(strings.NewReader(input), &out)

	return out.String()
}

// TestStartWritesToItsWriter covers Start honouring its arguments. It used to
// take in and out and ignore both, reading through readline's own stdin and
// writing with fmt.Println -- which is why nothing here could be tested.
func TestStartWritesToItsWriter(t *testing.T) {
	output := startOutput(t, "a = 21\na * 2\n")

	if !strings.Contains(output, buildVersion) {
		t.Errorf("expected the splash screen on the given writer, got:\n%s", output)
	}

	if !strings.Contains(output, "» 42") {
		t.Errorf("expected the result on the given writer, got:\n%s", output)
	}
}

func TestStartStopsAtEndOfInput(t *testing.T) {
	// No trailing newline on the last line, and nothing after it.
	if output := startOutput(t, "1 + 1"); !strings.Contains(output, "» 2") {
		t.Errorf("expected the final line to be evaluated, got:\n%s", output)
	}
}

func TestStartSurvivesAParseError(t *testing.T) {
	output := startOutput(t, "a = 1\n)\na + 1\n")

	if !strings.Contains(output, "parser errors:") {
		t.Errorf("expected the parse error to be reported, got:\n%s", output)
	}

	if !strings.Contains(output, "» 2") {
		t.Errorf("expected Start to carry on past the parse error, got:\n%s", output)
	}
}

func TestStartSkipsBlankLines(t *testing.T) {
	output := startOutput(t, "\n   \n1 + 1\n")

	if got := strings.Count(output, "»"); got != 1 {
		t.Errorf("expected exactly one result line, got %d:\n%s", got, output)
	}
}

// TestStartRunsWithoutAHomeDirectory checks the history file is a convenience.
// A missing home directory used to be log.Fatal, so the REPL refused to start.
func TestStartRunsWithoutAHomeDirectory(t *testing.T) {
	t.Setenv("HOME", "")

	var out bytes.Buffer
	Start(strings.NewReader("1 + 1\n"), &out)

	if !strings.Contains(out.String(), "» 2") {
		t.Errorf("expected the session to run without a home directory, got:\n%s", out.String())
	}
}
