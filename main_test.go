package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRocketlangCode(t *testing.T) {
	origStdout := os.Stdout
	defer func() {
		os.Stdout = origStdout
	}()

	testDir := "tests"

	var matches []string
	err := fs.WalkDir(os.DirFS(testDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".rl") {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk test dir 'tests/': %s", err)
	}
	for _, match := range matches {
		filename := filepath.Join(testDir, match)
		rlCode, err := os.ReadFile(filename)
		if err != nil {
			t.Errorf("%s: %s", match, err)
			continue
		}

		expectedStdout, err := os.ReadFile(strings.ReplaceAll(filename, ".rl", ".expected"))
		if err != nil {
			t.Errorf("%s: %s", match, err)
			continue
		}

		// CreateTemp's pattern arg rejects path separators, and match now carries
		// one for fixtures grouped in subdirectories (e.g. "lang/01_print.rl").
		tempPattern := strings.ReplaceAll(strings.TrimSuffix(match, ".rl"), "/", "_")
		fakeStdout, err := os.CreateTemp("", tempPattern)
		if err != nil {
			t.Errorf("%s: %s", match, err)
			continue
		}
		defer os.Remove(fakeStdout.Name())

		os.Stdout = fakeStdout
		runProgram(string(rlCode), filename)
		os.Stdout = origStdout

		resultStdout, err := os.ReadFile(fakeStdout.Name())
		if err != nil {
			t.Errorf("%s: %s", match, err)
			continue
		}

		if string(resultStdout) != string(expectedStdout) {
			fmt.Printf("--- stdout ---\n%s--- expected ---\n%s", resultStdout, expectedStdout)
			t.Errorf("%s: stdout does not match expected", match)
		}
	}
}

// TestExitCodes covers what the process reports. Nothing did before: an
// uncaught error, a parse error and a missing file all exited 0, so
// `rocket-lang build.rl && deploy` ran the deploy after a crash.
func TestExitCodes(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   int
	}{
		{"a program that runs", `print("fine")`, exitOK},
		{"a program with no output", `a = 1`, exitOK},
		// An empty program evaluates to nothing at all, which is a different
		// branch from evaluating to a value.
		{"an empty program", ``, exitOK},
		{"only a comment", `// nothing here`, exitOK},
		// An error aborts the rest of the program and becomes its result, so
		// the result being an error means nothing handled it.
		{"an uncaught error", `nil.no_such_method()`, exitFailure},
		{"an uncaught error part way through", `print("before")` + "\n" + `nil.nope()`, exitFailure},
		{"a parse error", `def f(`, exitFailure},
		// A rescued error is handled, so the program succeeded.
		{"a rescued error", "begin\n  1 / 0\nrescue e\n  print(\"handled\")\nend", exitOK},
		// A failed conversion answers nil rather than erroring, so it is not a
		// failure -- which is the distinction nil was introduced for.
		{"a failed conversion", `print("abc".to_integer())`, exitOK},
		// An error stored in a variable was handled by the program deciding to
		// keep it, and is not the final value.
		{"an error that is not the last value", "begin\n  nil.nope()\nrescue e\nend\nprint(\"after\")", exitOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureRun(t, func() int { return runProgram(tt.source, "test") })

			if got != tt.want {
				t.Errorf("exit code %d, want %d", got, tt.want)
			}
		})
	}
}

// TestRunFileReportsAMissingFile covers the case that was silent: no output at
// all and a successful exit.
func TestRunFileReportsAMissingFile(t *testing.T) {
	code := captureRun(t, func() int { return runFile(filepath.Join(t.TempDir(), "nope.rl")) })

	if code != exitFailure {
		t.Errorf("a missing file should fail, got exit %d", code)
	}
}

// TestRunFileRunsAFile checks the ordinary path still works through runFile.
func TestRunFileRunsAFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ok.rl")
	if err := os.WriteFile(path, []byte(`print("hello")`), 0o644); err != nil {
		t.Fatal(err)
	}

	if code := captureRun(t, func() int { return runFile(path) }); code != exitOK {
		t.Errorf("a readable program should succeed, got exit %d", code)
	}
}

// captureRun runs fn with stdout and stderr sent to a temporary file, so a test
// does not print the diagnostics it is provoking, and returns fn's exit code.
func captureRun(t *testing.T, fn func() int) int {
	t.Helper()

	sink, err := os.CreateTemp(t.TempDir(), "output")
	if err != nil {
		t.Fatal(err)
	}

	originalOut, originalErr := os.Stdout, os.Stderr
	os.Stdout, os.Stderr = sink, sink

	defer func() {
		os.Stdout, os.Stderr = originalOut, originalErr
		sink.Close()
	}()

	return fn()
}
