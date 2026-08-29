//go:build !wasm

package exercises_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/flipez/rocket-lang/evaluator"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
)

// update records what each reference solution prints, rather than checking it.
// The generator and the checker are the same code on purpose: an expectation
// produced by different code from the one that verifies it can disagree with it.
//
//	go test ./exercises -update
var update = flag.Bool("update", false, "record the expected output from the solutions")

// Exercise is one task, as the browser playground needs it.
type Exercise struct {
	ID       string `json:"id"`
	Topic    string `json:"topic"`
	Name     string `json:"name"`
	Task     string `json:"task"`
	Expected string `json:"expected"`
}

// TestUpdateExpectations writes exercises/expected/ and wasm/exercises.json
// from the solutions. It does nothing unless -update is given, so an ordinary
// test run cannot paper over a real difference by rewriting the expectation.
func TestUpdateExpectations(t *testing.T) {
	if !*update {
		t.Skip("run with -update to record the expected output")
	}

	solutions, err := filepath.Glob(filepath.Join("solutions", "*", "*.rl"))
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(solutions)

	bundle := make([]Exercise, 0, len(solutions))

	for _, solutionPath := range solutions {
		id := strings.TrimSuffix(strings.TrimPrefix(filepath.ToSlash(solutionPath), "solutions/"), ".rl")

		task, err := os.ReadFile(id + ".rl")
		if err != nil {
			t.Fatalf("solution %s has no exercise at %s.rl", solutionPath, id)
		}

		output := runFile(t, solutionPath)

		expectedPath := filepath.Join("expected", id+".txt")
		if err := os.MkdirAll(filepath.Dir(expectedPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(expectedPath, []byte(output), 0o644); err != nil {
			t.Fatal(err)
		}

		topic, name, _ := strings.Cut(id, "/")
		bundle = append(bundle, Exercise{ID: id, Topic: topic, Name: name, Task: string(task), Expected: output})

		t.Logf("recorded %s", id)
	}

	encoded, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join("..", "wasm", "exercises.json"), append(encoded, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Logf("recorded %d exercises into expected/ and wasm/exercises.json", len(bundle))
}

// TestBundleMatchesExpectations checks that wasm/exercises.json, which the
// playground serves, still agrees with the expectations on disk. Without this
// the browser could be grading against a stale bundle.
func TestBundleMatchesExpectations(t *testing.T) {
	encoded, err := os.ReadFile(filepath.Join("..", "wasm", "exercises.json"))
	if err != nil {
		t.Fatalf("the playground bundle is missing; run `go test ./exercises -update`: %s", err)
	}

	var bundle []Exercise
	if err := json.Unmarshal(encoded, &bundle); err != nil {
		t.Fatal(err)
	}

	expectations, err := filepath.Glob(filepath.Join("expected", "*", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if len(bundle) != len(expectations) {
		t.Fatalf("the bundle has %d exercises, expected/ has %d; run `go test ./exercises -update`", len(bundle), len(expectations))
	}

	for _, exercise := range bundle {
		want, err := os.ReadFile(filepath.Join("expected", exercise.ID+".txt"))
		if err != nil {
			t.Errorf("%s is in the bundle but has no expectation", exercise.ID)
			continue
		}
		if exercise.Expected != string(want) {
			t.Errorf("%s: the bundle and expected/ disagree; run `go test ./exercises -update`", exercise.ID)
		}

		task, err := os.ReadFile(exercise.ID + ".rl")
		if err != nil || string(task) != exercise.Task {
			t.Errorf("%s: the bundled task is not the file on disk; run `go test ./exercises -update`", exercise.ID)
		}
	}
}

// TestExercises checks the two properties that make an exercise set worth
// having, for every exercise: the reference solution produces exactly the
// recorded output, and the exercise as shipped does not.
//
// The first means an exercise can never promise something impossible, and that
// a language change which breaks an exercise fails the build instead of being
// discovered by a learner. The second means every exercise actually asks for
// something -- an exercise that already passes teaches nothing.
func TestExercises(t *testing.T) {
	expectations, err := filepath.Glob(filepath.Join("expected", "*", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if len(expectations) == 0 {
		t.Fatal("no expectations found; run `go test ./exercises -update`")
	}

	for _, expectedPath := range expectations {
		id := strings.TrimSuffix(strings.TrimPrefix(filepath.ToSlash(expectedPath), "expected/"), ".txt")

		t.Run(id, func(t *testing.T) {
			want, err := os.ReadFile(expectedPath)
			if err != nil {
				t.Fatal(err)
			}

			solution := runFile(t, filepath.Join("solutions", id+".rl"))
			if solution != string(want) {
				t.Errorf("the solution does not produce the recorded output.\n got: %q\nwant: %q\n"+
					"if the language changed, re-run `go test ./exercises -update`", solution, want)
			}

			task := runFile(t, id+".rl")
			if task == string(want) {
				t.Errorf("the exercise already produces the expected output, so it asks for nothing")
			}
		})
	}
}

// TestEveryExerciseHasASolution guards the other direction: an exercise without
// a solution would never be checked by the test above, because it walks the
// expectations, which are generated from the solutions.
func TestEveryExerciseHasASolution(t *testing.T) {
	tasks, err := filepath.Glob(filepath.Join("*", "*.rl"))
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range tasks {
		if strings.HasPrefix(task, "solutions"+string(filepath.Separator)) {
			continue
		}

		solution := filepath.Join("solutions", task)
		if _, err := os.Stat(solution); err != nil {
			t.Errorf("%s has no solution at %s", task, solution)
		}
	}
}

// runFile evaluates a file and returns what it printed, with the program's own
// final value dropped when it is nil -- the rule generate.go, run.sh and the
// playground all apply, so all four agree on what "the output" is.
func runFile(t *testing.T, path string) string {
	t.Helper()

	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: %s", path, err)
	}

	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write

	collected := make(chan string, 1)
	go func() {
		printed, _ := io.ReadAll(read)
		collected <- string(printed)
	}()

	l := lexer.New(string(source), path)
	p := parser.New(l)
	program := p.ParseProgram()

	var evaluated object.Object
	if len(p.Errors()) == 0 {
		evaluated = evaluator.Eval(program, object.NewEnvironment())
	}
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}

	write.Close()
	os.Stdout = original

	if len(p.Errors()) > 0 {
		t.Fatalf("%s does not parse: %s", path, strings.Join(p.Errors(), "; "))
	}

	printed := <-collected
	lines := strings.Split(strings.TrimSuffix(printed, "\n"), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "nil" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return ""
	}

	return strings.Join(lines, "\n") + "\n"
}
