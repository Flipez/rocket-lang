package stdlib

import (
	"bufio"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// TestIOWriteAddsNoNewline is the whole point of IO.write existing beside
// print, so it is what gets pinned.
func TestIOWriteAddsNoNewline(t *testing.T) {
	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write

	Modules["IO"].Functions["write"].Call(
		[]object.Object{object.NewString("ab")},
		*object.NewEnvironment(),
	)

	write.Close()
	os.Stdout = original

	// io.ReadAll, not strings.Builder.ReadFrom -- strings.Builder is a Writer
	// and has no ReadFrom method.
	out, err := io.ReadAll(read)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "ab" {
		t.Errorf("IO.write should emit exactly %q, got %q", "ab", string(out))
	}
}

func TestIOReadLineIsRegistered(t *testing.T) {
	if _, exists := Modules["IO"].Functions["read_line"]; !exists {
		t.Error("IO.read_line should be registered")
	}
}

// withStdinReader swaps the package-level stdinReader for the duration of a
// test and restores it afterwards.
func withStdinReader(t *testing.T, r *bufio.Reader) {
	t.Helper()

	original := stdinReader
	t.Cleanup(func() { stdinReader = original })
	stdinReader = r
}

// TestIOReadLineReadsMultipleLinesInSequence is the test that would have
// caught read_line's original defect: building a fresh *bufio.Reader on
// every call meant the first call swallowed everything still on the pipe
// (bufio reads ahead into its own buffer), and every call after it saw a
// descriptor already drained past the point the discarded buffer held. A
// test that only reads once, like TestIOReadLineIsRegistered, cannot see
// that -- the loop read_line exists for is what has to be exercised.
func TestIOReadLineReadsMultipleLinesInSequence(t *testing.T) {
	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	withStdinReader(t, bufio.NewReader(read))

	if _, err := write.WriteString("one\ntwo\nthree\n"); err != nil {
		t.Fatal(err)
	}
	write.Close()

	for _, want := range []string{"one", "two", "three"} {
		result := Modules["IO"].Functions["read_line"].Call(nil, *object.NewEnvironment())

		str, ok := result.(*object.String)
		if !ok {
			t.Fatalf("expected a String for %q, got %s", want, result.Type())
		}

		if str.Value != want {
			t.Errorf("expected %q, got %q", want, str.Value)
		}
	}

	// The pipe is exhausted and closed: the next read is EOF, not a fault.
	result := Modules["IO"].Functions["read_line"].Call(nil, *object.NewEnvironment())
	if result != object.NIL {
		t.Errorf("expected nil at end of input, got %s", result.Inspect())
	}
}

// errReader always fails with something other than io.EOF, standing in for
// a genuine read failure on the descriptor.
type errReader struct{}

func (errReader) Read(_ []byte) (int, error) {
	return 0, errors.New("boom: simulated read failure")
}

// TestIOReadLineDistinguishesRealErrorsFromEOF pins the fix for treating any
// error as "end of input": read_line's whole justification for returning nil
// is that EOF is not a fault, but a real read failure is, and a caller needs
// to be able to tell those apart.
func TestIOReadLineDistinguishesRealErrorsFromEOF(t *testing.T) {
	withStdinReader(t, bufio.NewReader(errReader{}))

	result := Modules["IO"].Functions["read_line"].Call(nil, *object.NewEnvironment())

	if !object.IsError(result) {
		t.Fatalf("expected an error for a non-EOF read failure, got %s", result.Inspect())
	}
}

// TestIOWriteConcatenatesMultipleArguments pins write's ArgPattern of
// OverloadArg(ANY): several arguments are written back to back with no
// separator, the same as a single call would be.
func TestIOWriteConcatenatesMultipleArguments(t *testing.T) {
	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write

	Modules["IO"].Functions["write"].Call(
		[]object.Object{object.NewString("a"), object.NewString("b")},
		*object.NewEnvironment(),
	)

	write.Close()
	os.Stdout = original

	out, err := io.ReadAll(read)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "ab" {
		t.Errorf("IO.write should concatenate arguments with no separator, got %q", string(out))
	}
}
