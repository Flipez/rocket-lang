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

	matches, err := fs.Glob(os.DirFS(testDir), "*.rl")
	if err != nil {
		t.Errorf("failed to read test dir 'tests/': %s", err)
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

		fakeStdout, err := os.CreateTemp("", strings.TrimSuffix(match, ".rl"))
		if err != nil {
			t.Errorf("%s: %s", match, err)
			continue
		}
		defer os.Remove(fakeStdout.Name())

		os.Stdout = fakeStdout
		runProgram(string(rlCode))
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
