package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNewArrayWithObjects(t *testing.T) {
	arr := object.NewArrayWithObjects(object.NewString("a"))
	if v := arr.Type(); v != object.ARRAY_OBJ {
		t.Errorf("array.Type() return wrong type: %s", v)
	}

	if v := arr.Elements[0].Type(); v != object.STRING_OBJ {
		t.Errorf("first array element should be a string object")
	}
}

func TestArrayObject(t *testing.T) {
	tests := []inputTestCase{
		{"[1] == [1]", true},
		{"[1] == [true]", false},
		{"[1] == [true, 1]", false},
	}

	testInput(t, tests)
}

func TestArrayObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3][0]`, 1},
		{`[1,2,3].size()`, 3},
		{`[1,2,3].remove_last()`, "[1, 2]"},
		{`[1,2,3].type()`, "ARRAY"},
		{`a = []; a.append!(1); a`, "[1]"},
		{`[].nope()`, "test:1:3: undefined method `.nope()` for ARRAY"},
		{"a = [\"a\", \"b\"]; b = []; foreach i, item in a \n b.append!(item) \nend; b.size()", 2},
		{`[1,2,3].index_of(4)`, -1},
		{`[1,2,3].index_of(3)`, 2},
		{`[1,2,3].index_of(true)`, -1},
		{`[1,2,3].index_of()`, "too few arguments: got=0, want=1"},
		{"a = []; b = []; foreach i in a \n b.append!(a[i]) \nend; a.size()==b.size()", true},
		{`[1,1,2].unique().size()`, 2},
		{`[true,true,2].unique().size()`, 2},
		{`["test","test",2].unique().size()`, 2},
		// nil is not hashable, so it cannot be de-duplicated. Written with a
		// literal rather than relying on what a !-method hands back.
		{`[nil].unique()`, "element 0 is not HASHABLE, got NIL"},
		{"[].first()", nil},
		{"[1,2,3].first()", 1},
		{"[].last()", nil},
		{"[1,2,3].last()", 3},
		{"[1,2,3].to_json()", "[1,2,3]"},
		{`["test",true,3].to_json()`, `["test",true,3]`},
		{`[3.4, 3.1, 2.0].sort()`, `[2.0, 3.1, 3.4]`},
		{`[3, 1, 4].sort()`, `[1, 3, 4]`},
		{`["Gopher", "Go", "Alpha"].sort()`, `["Alpha", "Go", "Gopher"]`},
		{`["Gopher", 1, "Alpha"].sort()`, "elements must all be one COMPARABLE type, got STRING at 0 and INTEGER at 1"},
		{`[1, "Go", 1].sort()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},
		{`[2.0, "Go", 2.0].sort()`, "elements must all be one COMPARABLE type, got FLOAT at 0 and STRING at 1"},
		{`[true, "Go", true].sort()`, "element 0 is not COMPARABLE, got BOOLEAN"},
		{`[].sort()`, `[]`},
		{`["a", "b", 1, 2].reverse()`, `[2, 1, "b", "a"]`},
		{`[1,2,3].contains?(4)`, false},
		{`[1,2,3].contains?(3)`, true},
		{`[1,2,3].contains?(true)`, false},
		{`[1,2,3].contains?()`, "too few arguments: got=0, want=1"},
		{`[1,2,3,4,5,6,7,8,9].chunks(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8, 9]]`},
		{`[1,2,3,4,5,6,7,8].chunks(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8]]`},
		{`[1,2].chunks(3)`, `[[1, 2]]`},
		{`[1,2].chunks(0)`, `invalid slice size, needs to be > 0`},
		// A hash used to be rejected here because Hash had no to_string. Now it
		// renders like it does everywhere else, and only a type that genuinely
		// has no string form -- a function -- is refused.
		{"[1,2,3,{}].join()", "123{}"},
		{"[1,2,3,[4]].join()", "123[4]"},
		{"[def() end].join()", "element 0 is not STRINGABLE, got FUNCTION"},
		{"[1,2,3].join()", "123"},
		{"[1,2,3].join('-')", "1-2-3"},
		{"['1',2, 2.5,{}].sum()", "element 3 is not INTEGERABLE, got HASH"},
		{"['1', 2, 2.5].sum()", 5},
		// A type can be convertible in principle and still fail to convert:
		// nil and a non-numeric string both used to contribute a silent 0.
		{"[1, nil].sum()", "element 1 is not INTEGERABLE, got NIL"},
		{"['abc'].sum()", "element 0 does not convert to a number, got STRING"},
		{`[[1, 2], [3, 4]].to_matrix()`, "2x2 matrix\n┌          ┐\n│ 1.0  2.0 │\n│ 3.0  4.0 │\n└          ┘"},
		{`[[1, 2], [3, 4]].to_matrix().to_array()`, "[[1.0, 2.0], [3.0, 4.0]]"},
		{`[1, 2].to_matrix()`, "failed to convert array to matrix: matrix must be created from 2D array"},
		{`[[1, 2], [3]].to_matrix()`, "failed to convert array to matrix: row 1 has inconsistent length (expected 2, got 1)"},
	}

	testInput(t, tests)
}

func TestArrayInspect(t *testing.T) {
	arr1 := object.NewArray(nil)

	if arr1.Type() != object.ARRAY_OBJ {
		t.Errorf("array.Type() returns wrong type")
	}
}

func TestArrayHashKey(t *testing.T) {
	arr1 := &object.Array{Elements: []object.Object{}}
	arr2 := &object.Array{Elements: []object.Object{}}
	diff1 := &object.Array{Elements: []object.Object{&object.String{Value: "Hello World"}}}
	diff2 := &object.Array{Elements: []object.Object{&object.String{Value: "Hello Another World"}}}

	if arr1.HashKey() != arr2.HashKey() {
		t.Errorf("arrays with same content have different hash keys")
	}

	if diff1.HashKey() == diff2.HashKey() {
		t.Errorf("arrays with different content have same hash keys")
	}
}

func TestArrayAdd(t *testing.T) {
	array := object.NewArray(nil)
	array.Add("a")
	if len(array.Elements) != 1 || array.Elements[0].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 0")
	}

	array.Add(object.NewString("b"))
	if len(array.Elements) != 2 || array.Elements[1].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 1")
	}

	array.Add(object.NewString("c"), 1)
	if len(array.Elements) != 4 || array.Elements[2].Type() != object.STRING_OBJ {
		t.Errorf("expected array to have a string value on index 2")
	}
	if len(array.Elements) != 4 || array.Elements[3].Type() != object.INTEGER_OBJ {
		t.Errorf("expected array to have an integer value on index 3")
	}
}

// TestArrayBangConvention covers the mutation convention #231 introduced: a
// plain method leaves the receiver alone and returns a new array, a !-method
// changes the receiver and hands it back so calls chain. Before this, reverse
// and sort mutated without a bang while unique did not, so the same name meant
// different things on Array and String.
func TestArrayBangConvention(t *testing.T) {
	tests := []inputTestCase{
		// Plain methods are pure: the original is untouched.
		{`a = [3,1,2]; a.reverse(); a.to_json()`, "[3,1,2]"},
		{`a = [3,1,2]; a.sort(); a.to_json()`, "[3,1,2]"},
		{`a = [1,1,2]; a.unique(); a.to_json()`, "[1,1,2]"},
		{`[3,1,2].reverse().to_json()`, "[2,1,3]"},
		{`[3,1,2].sort().to_json()`, "[1,2,3]"},
		{`[1,1,2].unique().to_json()`, "[1,2]"},

		// !-methods change the receiver.
		{`a = [3,1,2]; a.reverse!(); a.to_json()`, "[2,1,3]"},
		{`a = [3,1,2]; a.sort!(); a.to_json()`, "[1,2,3]"},
		{`a = [1,1,2]; a.unique!(); a.to_json()`, "[1,2]"},

		// ...and hand it back, so they chain.
		{`[3,1,2].sort!().reverse!().to_json()`, "[3,2,1]"},
		{`[1].append(2).append(3).to_json()`, "[1,2,3]"},

		// unique keeps the order of first appearance. Building the result from a
		// map made this vary between runs.
		{`[5,3,1,4,2,3].unique().to_json()`, "[5,3,1,4,2]"},
		{`[5,3,1,4,2,3].unique!().to_json()`, "[5,3,1,4,2]"},

		// A sort that cannot compare its elements errors...
		{`[1,"a"].sort()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},
		// ...and leaves the array alone rather than half-ordered. The error has
		// to be caught, or it would propagate before to_json runs and the
		// assertion would pass without checking anything.
		{`a = [1,"a"]
		  begin
		    a.sort!()
		  rescue e
		  end
		  a.to_json()`, `[1,"a"]`},
	}

	testInput(t, tests)
}

// TestArrayRubyMethods covers the methods added to close the gap with Ruby's
// Array. Only the ones that need no user-supplied function are here; map,
// filter and reject are not possible until an object method can call one.
func TestArrayRubyMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[].empty?()`, true},
		{`[1].empty?()`, false},

		// count(x) is how often x occurs, which index_of() cannot tell you.
		// size() (no argument) is the one that answers "how many are there".
		{`[1,2,2,3].count(2)`, 2},
		{`[1,2,2,3].count(9)`, 0},

		// -1 when absent, the answer index_of() already gave.
		{`[1,2,2,3].last_index_of(2)`, 2},
		{`[1,2,3].last_index_of(9)`, -1},

		{`[3,1,2].min()`, 1},
		{`[3,1,2].max()`, 3},
		{`["b","a"].min()`, "a"},
		{`[2.5,1.5].max()`, 2.5},
		{`[].min()`, nil},
		{`[].max()`, nil},
		// min and max borrow sort's rule about what can be compared, so they
		// fail the same way and with the same words.
		{`[1,"a"].min()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},

		{`[1,2,3].first(2)`, "[1, 2]"},
		{`[1,2,3].first(9)`, "[1, 2, 3]"},
		{`[1,2,3].first(0)`, "[]"},
		{`[1,2,3].last(2)`, "[2, 3]"},
		{`[1,2,3].first(0 - 1)`, "negative count -1"},
		// Without a count they still answer the element, not a one-item array.
		{`[1,2,3].first()`, 1},
		{`[1,2,3].last()`, 3},

		{`[1,2,3].skip(2)`, "[3]"},
		{`[1,2,3].skip(9)`, "[]"},
		{`[1,2,3].skip(0 - 1)`, "negative count -1"},

		{`[1,nil,2,nil].compact()`, "[1, 2]"},
		{`[1,[2,[3]]].flatten()`, "[1, 2, 3]"},
		{`[1,[2,[3]]].flatten(1)`, "[1, 2, [3]]"},
		{`[1,[2,[3]]].flatten(0)`, "[1, [2, [3]]]"},
		{`[1,2,3].rotate()`, "[2, 3, 1]"},
		{`[1,2,3].rotate(2)`, "[3, 1, 2]"},
		{`[1,2,3].rotate(0 - 1)`, "[3, 1, 2]"},
		// The count wraps, so rotating three elements by four is by one.
		{`[1,2,3].rotate(4)`, "[2, 3, 1]"},
		{`[].rotate()`, "[]"},

		// remove_first mirrors remove_last: a new array without the first
		// element, leaving the receiver untouched. An empty array has no first
		// element to drop, so it comes back unchanged.
		{`a = [1,2,3]; a.remove_first().to_json()`, "[2,3]"},
		{`[].remove_first().to_json()`, "[]"},

		{`a = [2,3]; a.prepend(1).to_json()`, "[1,2,3]"},

		{`a = [1,2]; a.insert(1, 9).to_json()`, "[1,9,2]"},
		{`a = [1,2]; a.insert(0, 9).to_json()`, "[9,1,2]"},
		{`a = [1,2]; a.insert(2, 9).to_json()`, "[1,2,9]"},
		// A negative index counts back from the end, so -1 appends.
		{`a = [1,2]; a.insert(0 - 1, 9).to_json()`, "[1,2,9]"},
		{`a = [1,2]; a.insert(3, 9)`, "index out of range, got 3 but array has only 2 elements"},

		// remove takes out every occurrence. A target that is not there comes
		// back unchanged rather than erroring.
		{`a = [1,2,1]; a.remove(1).to_json()`, "[2]"},
		{`a = [1,2]; a.remove(9).to_json()`, "[1,2]"},

		{`a = [1,2,3]; a.remove_at(1).to_json()`, "[1,3]"},
		{`a = [1,2,3]; a.remove_at(0 - 1).to_json()`, "[1,2]"},
		// A position that is not there comes back unchanged rather than
		// erroring, the same tolerance remove() has for a miss.
		{`[1,2].remove_at(9).to_json()`, "[1,2]"},

		{`a = [1]; a.concat([2,3]).to_json()`, "[1,2,3]"},
		{`a = [1]; a.concat([]).to_json()`, "[1]"},

		// A method handing back part of an array must hand back a copy. Without
		// one the result shares memory with the original, and writing to the
		// original changes the result underneath the caller.
		{`a = [1,2,3]; b = a.skip(1); a[1] = 9; b.to_json()`, "[2,3]"},
		{`a = [1,2,3]; b = a.first(2); a[0] = 9; b.to_json()`, "[1,2]"},
		{`a = [1,2,3]; b = a.last(2); a[2] = 9; b.to_json()`, "[2,3]"},
	}
	testInput(t, tests)
}

func TestArraySkip(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3,4,5].skip(2)`, "[3, 4, 5]"},
		{`[1,2,3,4,5].skip_last(2)`, "[1, 2, 3]"},
		{`[1,2,3].skip(0)`, "[1, 2, 3]"},
		{`[1,2,3].skip(99)`, "[]"},
		{`[1,2,3].skip_last(99)`, "[]"},
		// skip_last's ReturnPattern has to say ERROR too, or this path renders
		// as returning ARRAY in the generated docs while actually erroring.
		{`[1,2,3].skip_last(0 - 1)`, "skip_last needs a count of zero or more, got -1"},
	}

	testInput(t, tests)
}

// take was identical to first(n). Keeping both meant two names for one
// behaviour, so it is gone; this pins that it stays gone.
func TestArrayTakeIsGone(t *testing.T) {
	evaluated := testEval(`[1,2,3].take(2)`)

	if !object.IsError(evaluated) {
		t.Errorf("take should no longer exist, got %s", evaluated.Inspect())
	}
}

// clear returned an empty array, which carries no information of its own --
// [] already says the same thing. A bang-only clear! would have no pure
// counterpart, which is exactly what the pair rule forbids, so clear is gone
// rather than paired; this pins that it stays gone.
func TestArrayClearIsGone(t *testing.T) {
	evaluated := testEval(`[1,2].clear()`)

	if !object.IsError(evaluated) {
		t.Errorf("clear should no longer exist, got %s", evaluated.Inspect())
	}
}

// TestArrayBangPairsAreComplete checks that the new pairs follow the same rule
// as the old ones: the plain method leaves the receiver alone, the ! method
// changes it and hands it back so calls chain.
func TestArrayBangPairsAreComplete(t *testing.T) {
	tests := []inputTestCase{
		// Pure.
		{`a = [1,nil,2]; a.compact(); a.to_json()`, "[1,null,2]"},
		{`a = [1,[2]]; a.flatten(); a.to_json()`, "[1,[2]]"},
		{`a = [1,2,3]; a.rotate(); a.to_json()`, "[1,2,3]"},

		// In place.
		{`a = [1,nil,2]; a.compact!(); a.to_json()`, "[1,2]"},
		{`a = [1,[2]]; a.flatten!(); a.to_json()`, "[1,2]"},
		{`a = [1,2,3]; a.rotate!(); a.to_json()`, "[2,3,1]"},

		// ...and they chain, because each returns the array.
		{`[1,nil,[2,nil]].flatten!().compact!().to_json()`, "[1,2]"},
		{`[1,2,3].rotate!().rotate!().to_json()`, "[3,1,2]"},
		{`[1,2].compact!().type()`, "ARRAY"},

		// A ! method takes the same arguments as its plain counterpart.
		{`[1,[2,[3]]].flatten!(1).to_json()`, "[1,2,[3]]"},
		{`[1,2,3].rotate!(2).to_json()`, "[3,1,2]"},
		{`[1,2].flatten!(0 - 1)`, "negative depth -1"},

		// remove_last, remove_first, append, prepend, insert, remove, remove_at
		// and concat used to mutate without a bang -- the same shape that let
		// String#remove_last and Array#remove_last mean opposite things.
		// Paired now: pure leaves the receiver alone...
		{`a = [1,2,3]; a.remove_last(); a.to_json()`, "[1,2,3]"},
		{`a = [1,2,3]; a.remove_first(); a.to_json()`, "[1,2,3]"},
		{`a = [1,2,3]; a.append(4); a.to_json()`, "[1,2,3]"},
		{`a = [2,3]; a.prepend(1); a.to_json()`, "[2,3]"},
		{`a = [1,2]; a.insert(1, 9); a.to_json()`, "[1,2]"},
		{`a = [1,2,1]; a.remove(1); a.to_json()`, "[1,2,1]"},
		{`a = [1,2,3]; a.remove_at(1); a.to_json()`, "[1,2,3]"},
		{`a = [1]; a.concat([2,3]); a.to_json()`, "[1]"},

		// ...and ! mutates and hands back the receiver.
		{`a = [1,2,3]; a.remove_last!(); a.to_json()`, "[1,2]"},
		{`a = [1,2,3]; a.remove_first!(); a.to_json()`, "[2,3]"},
		{`a = [1,2,3]; a.append!(4); a.to_json()`, "[1,2,3,4]"},
		{`a = [2,3]; a.prepend!(1); a.to_json()`, "[1,2,3]"},
		{`a = [1,2]; a.insert!(1, 9); a.to_json()`, "[1,9,2]"},
		{`a = [1,2,1]; a.remove!(1); a.to_json()`, "[2]"},
		{`a = [1,2,3]; a.remove_at!(1); a.to_json()`, "[1,3]"},
		{`a = [1]; a.concat!([2,3]); a.to_json()`, "[1,2,3]"},

		// ...and they chain too.
		{`[1,2,3].append!(4).remove_first!().to_json()`, "[2,3,4]"},
		{`[1,2].insert!(1, 9).remove_at!(0).to_json()`, "[9,2]"},
		{`[1,2].append!(3).type()`, "ARRAY"},

		// A pop needs two calls now: peek with last()/first(), then mutate.
		// There is no return-the-removed-element shortcut any more.
		{`a = [1,2,3]; last = a.last(); a.remove_last!(); last`, 3},
		{`a = [1,2,3]; last = a.last(); a.remove_last!(); a.to_json()`, "[1,2]"},
	}
	testInput(t, tests)
}

// TestElementGroupErrors covers the requirements a method places on the
// elements it is given. These were four checks with four unrelated messages,
// and sort's named neither the element at fault nor which of its two rules had
// been broken -- [1, nil].sort() and [1, 2.5].sort() said the same thing.
func TestElementGroupErrors(t *testing.T) {
	tests := []inputTestCase{
		// Each one names the group, the index and what was found instead.
		{`[def() end].join()`, "element 0 is not STRINGABLE, got FUNCTION"},
		{`[1, def() end].join()`, "element 1 is not STRINGABLE, got FUNCTION"},
		{`[nil].unique()`, "element 0 is not HASHABLE, got NIL"},
		{`[1, nil].unique()`, "element 1 is not HASHABLE, got NIL"},
		{`[1, nil].sum()`, "element 1 is not INTEGERABLE, got NIL"},
		{`[nil].sort()`, "element 0 is not COMPARABLE, got NIL"},
		{`[1, nil].sort()`, "element 1 is not COMPARABLE, got NIL"},
		{`[true].sort()`, "element 0 is not COMPARABLE, got BOOLEAN"},

		// Ordering has a second rule that no per-element group can state, so it
		// gets its own message -- and both types are named.
		{`[1, 2.5].sort()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and FLOAT at 1"},
		{`[1, "a"].sort()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},
		{`["a", 1].sort()`, "elements must all be one COMPARABLE type, got STRING at 0 and INTEGER at 1"},
		// min and max borrow the same rules, so they report the same way.
		{`[1, "a"].min()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},
		{`[1, "a"].max()`, "elements must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},
		{`[nil].min()`, "element 0 is not COMPARABLE, got NIL"},

		// One element of a COMPARABLE type is fine, and so is none.
		{`[1].sort().to_json()`, "[1]"},
		{`[].sort().to_json()`, "[]"},
		{`[2.5, 1.5].sort().to_json()`, "[1.5,2.5]"},
		{`["b", "a"].sort().to_json()`, `["a","b"]`},

		// INTEGERABLE is wider than NUMERIC on purpose: a string that parses
		// and a boolean both count, which is existing sum behaviour.
		{`["12"].sum()`, 12},
		{`[true].sum()`, 1},
		{`[1.9].sum()`, 1},
		// Being INTEGERABLE is not the same as converting, so that failure is
		// reported differently rather than blamed on the type.
		{`["abc"].sum()`, "element 0 does not convert to a number, got STRING"},
		{`[1, "abc"].sum()`, "element 1 does not convert to a number, got STRING"},
	}
	testInput(t, tests)
}

// TestArrayEach covers the first method that calls back into user code. Until
// the function applier was injected, a method could reach a callback's body but
// had no way to pass it an argument -- which is why HTTP.handle sets `request`
// as a variable instead of taking it as a parameter.
func TestArrayEach(t *testing.T) {
	tests := []inputTestCase{
		// The receiver comes back, so a walk chains. Returning nil would end
		// the chain, and each changes nothing that would be worth returning.
		{`a = [1,2,3]; a.each(def(x) end).to_json()`, "[1,2,3]"},
		{`[3,1,2].sort().each(def(x) end).size()`, 3},
		{`a = [1]; a.each(def(x) end).type()`, "ARRAY"},
		{`a = []; a.each(def(x) end).to_json()`, "[]"},

		// The callback actually receives the element, which is the whole point.
		{`out = []; a = [1,2,3]; a.each(def(x) out.append!(x * 2) end); out.to_json()`, "[2,4,6]"},
		{`out = []; a = ["x","y"]; a.each(def(s) out.append!(s.uppercase()) end); out.to_json()`, `["X","Y"]`},

		// break ends the walk, next moves it along. A function does not consume
		// either, so a callback can hand one back, and passing it through as a
		// value would be meaningless.
		{`out = []; a = [1,2,3,4]; a.each(def(x) if x == 3 break end out.append!(x) end); out.to_json()`, "[1,2]"},
		{`out = []; a = [1,2,3]; a.each(def(x) if x == 2 next end out.append!(x) end); out.to_json()`, "[1,3]"},
		// break still hands back the array, not the BREAK_VALUE.
		{`a = [1,2]; a.each(def(x) break end).type()`, "ARRAY"},

		// An error ends the walk and is handed on rather than swallowed.
		{`a = [1]; a.each(def(x) x.no_such_method() end)`, "test:1:25: undefined method `.no_such_method()` for INTEGER"},
		{`out = []; a = [1,2,3]; begin a.each(def(x) out.append!(x); x.nope() end) rescue e end out.to_json()`, "[1]"},

		// Arity is the applier's business, and it reports it the way a call
		// written out in full would.
		{`a = [1]; a.each(def(x, y) end)`, "too few arguments: got=1, want=2"},
		{`a = [1]; a.each(def() end)`, "too many arguments: got=1, want=0"},

		// A builtin is a value too, so it is callable.
		{`a = [1]; a.each(print).to_json()`, "[1]"},

		// CALLABLE refuses what cannot be called.
		{`a = [1]; a.each(1)`, "wrong argument type on position 1: got=INTEGER, want=CALLABLE"},
		{`a = [1]; a.each(nil)`, "wrong argument type on position 1: got=NIL, want=CALLABLE"},
		{`a = [1]; a.each()`, "too few arguments: got=0, want=1"},

		// A closure keeps its own scope, so the callback sees where it was
		// written rather than where it is called.
		{`factor = 10; out = []; a = [1,2]; a.each(def(x) out.append!(x * factor) end); out.to_json()`, "[10,20]"},
	}
	testInput(t, tests)
}

// TestArrayCallbackMethods covers the methods unlocked by the function applier.
// They all share one set of rules for what a callback's return means, so the
// cases that matter are the shared ones: break, next, an error, and arity.
func TestArrayCallbackMethods(t *testing.T) {
	tests := []inputTestCase{
		{`a = [1,2,3]; a.map(def(x) x * 2 end).to_json()`, "[2,4,6]"},
		{`a = [1,2,3,4]; a.filter(def(x) x % 2 == 0 end).to_json()`, "[2,4]"},
		{`a = [1,2,3,4]; a.reject(def(x) x % 2 == 0 end).to_json()`, "[1,3]"},
		// Only false and nil are false, so 0 and "" are yeses -- the language's
		// own truthiness, the same as if and while use.
		{`a = [1,2]; a.filter(def(x) 0 end).to_json()`, "[1,2]"},
		{`a = [1,2]; a.filter(def(x) "" end).to_json()`, "[1,2]"},
		{`a = [1,2]; a.filter(def(x) nil end).to_json()`, "[]"},

		{`a = [1,2,3]; a.reduce(0, def(sum, x) sum + x end)`, 6},
		{`a = [1,2,3]; a.reduce(1, def(p, x) p * x end)`, 6},
		{`a = [1,2]; a.reduce("", def(s, x) s + x.to_string() end)`, "12"},
		// An empty array answers with the starting value, which is why it is
		// required rather than taken from the first element.
		{`a = []; a.reduce(7, def(sum, x) sum + x end)`, 7},

		{`a = [1,2,3]; a.all?(def(x) x > 0 end)`, true},
		{`a = [1,2,3]; a.all?(def(x) x > 2 end)`, false},
		{`a = [1,2,3]; a.any?(def(x) x > 2 end)`, true},
		{`a = [1,2,3]; a.any?(def(x) x > 9 end)`, false},
		{`a = [1,2,3]; a.none?(def(x) x > 9 end)`, true},
		{`a = [1,2,3]; a.none?(def(x) x > 2 end)`, false},
		// An empty array is every element and no element at once.
		{`a = []; a.all?(def(x) false end)`, true},
		{`a = []; a.any?(def(x) true end)`, false},
		{`a = []; a.none?(def(x) true end)`, true},

		{`a = ["ccc","a","bb"]; a.sort_by(def(w) w.size() end).to_json()`, `["a","bb","ccc"]`},
		{`a = ["ccc","a","bb"]; a.min_by(def(w) w.size() end)`, "a"},
		{`a = ["ccc","a","bb"]; a.max_by(def(w) w.size() end)`, "ccc"},
		{`a = []; a.min_by(def(x) x end)`, nil},
		{`a = []; a.max_by(def(x) x end)`, nil},
		// The keys have to satisfy what sort requires of elements, and they are
		// reported the same way.
		{`a = ["a","b"]; a.sort_by(def(w) nil end)`, "element 0 is not COMPARABLE, got NIL"},
		{`a = ["a","b"]; a.sort_by(def(w) if w == "a" 1 else "x" end end)`, "keys must all be one COMPARABLE type, got INTEGER at 0 and STRING at 1"},

		// break ends the walk and the answer covers what was walked. next means
		// the element contributed nothing, which for map is a nil and for a
		// filter is a no.
		{`a = [1,2,3,4]; a.map(def(x) if x == 3 break end x end).to_json()`, "[1,2]"},
		{`a = [1,2,3]; a.map(def(x) if x == 2 next end x end).to_json()`, "[1,null,3]"},
		{`a = [1,2,3,4]; a.filter(def(x) if x == 3 break end true end).to_json()`, "[1,2]"},
		{`a = [1,2,3]; a.filter(def(x) if x == 2 next end true end).to_json()`, "[1,3]"},
		{`a = [1,2,3]; a.reduce(0, def(sum, x) if x == 3 break end sum + x end)`, 3},
		{`a = [1,2,3]; a.reduce(0, def(sum, x) if x == 2 next end sum + x end)`, 4},
		{`a = [1,2,3]; a.all?(def(x) if x == 3 break end x > 0 end)`, true},

		// The pure form leaves the receiver alone, the ! form changes it.
		{`a = [1,2]; a.map(def(x) x * 2 end); a.to_json()`, "[1,2]"},
		{`a = [1,2]; a.map!(def(x) x * 2 end); a.to_json()`, "[2,4]"},
		{`a = [1,2,3]; a.filter!(def(x) x > 1 end); a.to_json()`, "[2,3]"},
		{`a = [1,2,3]; a.reject!(def(x) x > 1 end); a.to_json()`, "[1]"},
		{`a = ["ccc","a"]; a.sort_by!(def(w) w.size() end); a.to_json()`, `["a","ccc"]`},
		{`a = [1,2]; a.map!(def(x) x end).type()`, "ARRAY"},

		// An error from a callback ends the walk and is passed on.
		{`a = [1]; a.map(def(x) x.nope() end)`, "test:1:24: undefined method `.nope()` for INTEGER"},
		{`a = [1]; a.filter(def(x) x.nope() end)`, "test:1:27: undefined method `.nope()` for INTEGER"},
		{`a = [1]; a.reduce(0, def(s, x) x.nope() end)`, "test:1:33: undefined method `.nope()` for INTEGER"},

		// Arity and callability are reported the same way everywhere.
		{`a = [1]; a.map(def(x, y) end)`, "too few arguments: got=1, want=2"},
		{`a = [1]; a.reduce(0, def(x) end)`, "too many arguments: got=2, want=1"},
		{`a = [1]; a.map(1)`, "wrong argument type on position 1: got=INTEGER, want=CALLABLE"},
		{`a = [1]; a.reduce(0, "x")`, "wrong argument type on position 2: got=STRING, want=CALLABLE"},

		// Chaining is the point of having them as methods.
		{`a = [1,2,3,4,5]; a.filter(def(x) x % 2 == 1 end).map(def(x) x * x end).sum()`, 35},
		{`a = [1,2,3]; a.map(def(x) x * 2 end).filter(def(x) x > 2 end).reduce(0, def(s, x) s + x end)`, 10},
	}
	testInput(t, tests)
}
