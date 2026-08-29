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
		{`[1,2,3].pop()`, 3},
		{`[1,2,3].type()`, "ARRAY"},
		{`a = []; a.push(1); a`, "[1]"},
		{`[].nope()`, "test:1:3: undefined method `.nope()` for ARRAY"},
		{"a = [\"a\", \"b\"]; b = []; foreach i, item in a \n b.push(item) \nend; b.size()", 2},
		{`[1,2,3].index(4)`, -1},
		{`[1,2,3].index(3)`, 2},
		{`[1,2,3].index(true)`, -1},
		{`[1,2,3].index()`, "to few arguments: got=0, want=1"},
		{"a = []; b = []; foreach i in a \n b.push(a[i]) \nend; a.size()==b.size()", true},
		{`[1,1,2].uniq().size()`, 2},
		{`[true,true,2].uniq().size()`, 2},
		{`["test","test",2].uniq().size()`, 2},
		// nil is not hashable, so it cannot be de-duplicated. Written with a
		// literal rather than relying on what a !-method hands back.
		{`[nil].uniq()`, "failed because element NIL is not hashable"},
		{"[].first()", nil},
		{"[1,2,3].first()", 1},
		{"[].last()", nil},
		{"[1,2,3].last()", 3},
		{"[1,2,3].to_json()", "[1,2,3]"},
		{`["test",true,3].to_json()`, `["test",true,3]`},
		{`[3.4, 3.1, 2.0].sort()`, `[2.0, 3.1, 3.4]`},
		{`[3, 1, 4].sort()`, `[1, 3, 4]`},
		{`["Gopher", "Go", "Alpha"].sort()`, `["Alpha", "Go", "Gopher"]`},
		{`["Gopher", 1, "Alpha"].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[1, "Go", 1].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[2.0, "Go", 2.0].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[true, "Go", true].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
		{`[].sort()`, `[]`},
		{`["a", "b", 1, 2].reverse()`, `[2, 1, "b", "a"]`},
		{`[1,2,3].include?(4)`, false},
		{`[1,2,3].include?(3)`, true},
		{`[1,2,3].include?(true)`, false},
		{`[1,2,3].include?()`, "to few arguments: got=0, want=1"},
		{`[1,2,3,4,5,6,7,8,9].slices(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8, 9]]`},
		{`[1,2,3,4,5,6,7,8].slices(3)`, `[[1, 2, 3], [4, 5, 6], [7, 8]]`},
		{`[1,2].slices(3)`, `[[1, 2]]`},
		{`[1,2].slices(0)`, `invalid slice size, needs to be > 0`},
		// A hash used to be rejected here because Hash had no to_s. Now it
		// renders like it does everywhere else, and only a type that genuinely
		// has no string form -- a function -- is refused.
		{"[1,2,3,{}].join()", "123{}"},
		{"[1,2,3,[4]].join()", "123[4]"},
		{"[def() end].join()", "Found non stringable element FUNCTION on index 0"},
		{"[1,2,3].join()", "123"},
		{"[1,2,3].join('-')", "1-2-3"},
		{"['1',2, 2.5,{}].sum()", "Found non number element HASH on index 3"},
		{"['1', 2, 2.5].sum()", 5},
		// A type can be convertible in principle and still fail to convert:
		// nil and a non-numeric string both used to contribute a silent 0.
		{"[1, nil].sum()", "Found non number element NIL on index 1"},
		{"['abc'].sum()", "Found non number element STRING on index 0"},
		{`[[1, 2], [3, 4]].to_m()`, "2x2 matrix\n┌          ┐\n│ 1.0  2.0 │\n│ 3.0  4.0 │\n└          ┘"},
		{`[[1, 2], [3, 4]].to_m().to_a()`, "[[1.0, 2.0], [3.0, 4.0]]"},
		{`[1, 2].to_m()`, "failed to convert array to matrix: matrix must be created from 2D array"},
		{`[[1, 2], [3]].to_m()`, "failed to convert array to matrix: row 1 has inconsistent length (expected 2, got 1)"},
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
// and sort mutated without a bang while uniq did not, so the same name meant
// different things on Array and String.
func TestArrayBangConvention(t *testing.T) {
	tests := []inputTestCase{
		// Plain methods are pure: the original is untouched.
		{`a = [3,1,2]; a.reverse(); a.to_json()`, "[3,1,2]"},
		{`a = [3,1,2]; a.sort(); a.to_json()`, "[3,1,2]"},
		{`a = [1,1,2]; a.uniq(); a.to_json()`, "[1,1,2]"},
		{`[3,1,2].reverse().to_json()`, "[2,1,3]"},
		{`[3,1,2].sort().to_json()`, "[1,2,3]"},
		{`[1,1,2].uniq().to_json()`, "[1,2]"},

		// !-methods change the receiver.
		{`a = [3,1,2]; a.reverse!(); a.to_json()`, "[2,1,3]"},
		{`a = [3,1,2]; a.sort!(); a.to_json()`, "[1,2,3]"},
		{`a = [1,1,2]; a.uniq!(); a.to_json()`, "[1,2]"},

		// ...and hand it back, so they chain.
		{`[3,1,2].sort!().reverse!().to_json()`, "[3,2,1]"},
		{`[1].push(2).push(3).to_json()`, "[1,2,3]"},

		// uniq keeps the order of first appearance. Building the result from a
		// map made this vary between runs.
		{`[5,3,1,4,2,3].uniq().to_json()`, "[5,3,1,4,2]"},
		{`[5,3,1,4,2,3].uniq!().to_json()`, "[5,3,1,4,2]"},

		// pop still mutates and hands back the element, as in Ruby: a pure pop
		// would just be last().
		{`a = [1,2,3]; a.pop()`, 3},
		{`a = [1,2,3]; a.pop(); a.to_json()`, "[1,2]"},

		// A sort that cannot compare its elements errors...
		{`[1,"a"].sort()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},
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
// select and reject are not possible until an object method can call one.
func TestArrayRubyMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[].empty?()`, true},
		{`[1].empty?()`, false},

		// count() is size(); count(x) is how often x occurs, which index()
		// cannot tell you.
		{`[1,2,2,3].count()`, 4},
		{`[1,2,2,3].count(2)`, 2},
		{`[1,2,2,3].count(9)`, 0},

		// -1 when absent, the answer index() already gave.
		{`[1,2,2,3].rindex(2)`, 2},
		{`[1,2,3].rindex(9)`, -1},

		{`[3,1,2].min()`, 1},
		{`[3,1,2].max()`, 3},
		{`["b","a"].min()`, "a"},
		{`[2.5,1.5].max()`, 2.5},
		{`[].min()`, nil},
		{`[].max()`, nil},
		// min and max borrow sort's rule about what can be compared, so they
		// fail the same way and with the same words.
		{`[1,"a"].min()`, "Array does contain either an object not INTEGER, FLOAT or STRING or is mixed"},

		{`[1,2,3].first(2)`, "[1, 2]"},
		{`[1,2,3].first(9)`, "[1, 2, 3]"},
		{`[1,2,3].first(0)`, "[]"},
		{`[1,2,3].last(2)`, "[2, 3]"},
		{`[1,2,3].first(0 - 1)`, "negative count -1"},
		// Without a count they still answer the element, not a one-item array.
		{`[1,2,3].first()`, 1},
		{`[1,2,3].last()`, 3},

		{`[1,2,3].take(2)`, "[1, 2]"},
		{`[1,2,3].drop(2)`, "[3]"},
		{`[1,2,3].take(9)`, "[1, 2, 3]"},
		{`[1,2,3].drop(9)`, "[]"},
		{`[1,2,3].take(0 - 1)`, "negative count -1"},

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

		// shift mirrors pop: it changes the array and hands back the element.
		{`a = [1,2,3]; a.shift()`, 1},
		{`a = [1,2,3]; a.shift(); a.to_json()`, "[2,3]"},
		{`[].shift()`, nil},
		{`a = [2,3]; a.unshift(1).to_json()`, "[1,2,3]"},

		{`a = [1,2]; a.insert(1, 9).to_json()`, "[1,9,2]"},
		{`a = [1,2]; a.insert(0, 9).to_json()`, "[9,1,2]"},
		{`a = [1,2]; a.insert(2, 9).to_json()`, "[1,2,9]"},
		// A negative index counts back from the end, so -1 appends.
		{`a = [1,2]; a.insert(0 - 1, 9).to_json()`, "[1,2,9]"},
		{`a = [1,2]; a.insert(3, 9)`, "index out of range, got 3 but array has only 2 elements"},

		// delete takes out every occurrence and reports the element, or nil
		// when there was nothing to take out.
		{`a = [1,2,1]; a.delete(1)`, 1},
		{`a = [1,2,1]; a.delete(1); a.to_json()`, "[2]"},
		{`a = [1,2]; a.delete(9)`, nil},
		{`a = [1,2]; a.delete(9); a.to_json()`, "[1,2]"},

		{`a = [1,2,3]; a.delete_at(1)`, 2},
		{`a = [1,2,3]; a.delete_at(1); a.to_json()`, "[1,3]"},
		{`a = [1,2,3]; a.delete_at(0 - 1)`, 3},
		// A position that is not there gives nil, as first() and pop() do on an
		// empty array.
		{`[1,2].delete_at(9)`, nil},

		{`a = [1,2]; a.clear().to_json()`, "[]"},
		{`a = [1]; a.concat([2,3]).to_json()`, "[1,2,3]"},
		{`a = [1]; a.concat([]).to_json()`, "[1]"},

		// A method handing back part of an array must hand back a copy. Without
		// one the result shares memory with the original, and writing to the
		// original changes the result underneath the caller.
		{`a = [1,2,3]; b = a.drop(1); a[1] = 9; b.to_json()`, "[2,3]"},
		{`a = [1,2,3]; b = a.take(2); a[0] = 9; b.to_json()`, "[1,2]"},
		{`a = [1,2,3]; b = a.first(2); a[0] = 9; b.to_json()`, "[1,2]"},
		{`a = [1,2,3]; b = a.last(2); a[2] = 9; b.to_json()`, "[2,3]"},
	}
	testInput(t, tests)
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
	}
	testInput(t, tests)
}
