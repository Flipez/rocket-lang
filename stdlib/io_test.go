package stdlib

import (
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
