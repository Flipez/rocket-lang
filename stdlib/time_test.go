package stdlib

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// TestTimeNowIsPlausible avoids pinning a clock value: it checks the result is
// an integer in a range that cannot be produced by a stub returning zero.
func TestTimeNowIsPlausible(t *testing.T) {
	result := Modules["Time"].Functions["now"].Call(nil, *object.NewEnvironment())

	seconds, ok := result.(*object.Integer)
	if !ok {
		t.Fatalf("Time.now should return an Integer, got %s", result.Type())
	}

	// 1 January 2020. Any real clock is past it; a stub is not.
	if seconds.Value < 1577836800 {
		t.Errorf("Time.now returned %d, which is not a current unix timestamp", seconds.Value)
	}
}
