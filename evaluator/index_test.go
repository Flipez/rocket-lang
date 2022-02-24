package evaluator

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestEvalIndex(t *testing.T) {
	testcases := []struct {
		left, index object.Object
		colon       string
		expected    object.Object
	}{
		// "abcde"[2] => c
		{
			left:     object.NewString("abcde"),
			index:    object.NewInteger(2),
			expected: object.NewString("c"),
		},
		// "abcde"[-2] => d
		{
			left:     object.NewString("abcde"),
			index:    object.NewInteger(-2),
			expected: object.NewString("d"),
		},
		// "abcde"[10] => NULL
		{
			left:     object.NewString("abcde"),
			index:    object.NewInteger(10),
			expected: object.NULL,
		},

		// {"a": 1}["a"] => 1
		{
			left: object.NewHash(map[object.HashKey]object.HashPair{
				object.NewString("a").HashKey(): object.HashPair{
					Key:   object.NewString("a"),
					Value: object.NewInteger(1),
				},
			}),
			index:    object.NewString("a"),
			expected: object.NewInteger(1),
		},
		// {"a": 1}["b"] => NULL
		{
			left: object.NewHash(map[object.HashKey]object.HashPair{
				object.NewString("a").HashKey(): object.HashPair{
					Key:   object.NewString("a"),
					Value: object.NewInteger(1),
				},
			}),
			index:    object.NewString("b"),
			expected: object.NULL,
		},
		// {"a": 1}[NULL] => ERROR: unusable as hash key: NULL
		{
			left: object.NewHash(map[object.HashKey]object.HashPair{
				object.NewString("a").HashKey(): object.HashPair{
					Key:   object.NewString("a"),
					Value: object.NewInteger(1),
				},
			}),
			index:    object.NULL,
			expected: object.NewErrorFormat("unusable as hash key: NULL"),
		},

		// ["a","b","c","d","e"][2] => "c"
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			index:    object.NewInteger(2),
			expected: object.NewString("c"),
		},
		// ["a","b","c","d","e"][-2] => "d"
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			index:    object.NewInteger(-2),
			expected: object.NewString("d"),
		},

		// 12345[2] => ERROR: index operator not supported: INTEGER
		{
			left:     object.NewInteger(12345),
			index:    object.NewInteger(2),
			expected: object.NewErrorFormat("index operator not supported: INTEGER"),
		},
	}

	for _, tc := range testcases {
		obj := evalIndex(tc.left, tc.index)
		if obj.Type() != tc.expected.Type() {
			t.Errorf("expected object to be a %s, got %s", tc.expected.Type(), obj.Type())
			continue
		}
		if obj.Inspect() != tc.expected.Inspect() {
			t.Errorf("unexpected result, got=%s, want=%s", obj.Inspect(), tc.expected.Inspect())
			continue
		}
	}
}

func TestEvalRangeIndex(t *testing.T) {
	testcases := []struct {
		left, firstIndex, secondIndex object.Object
		expected                      object.Object
	}{
		// "abcde"[:] => abcde
		{
			left:     object.NewString("abcde"),
			expected: object.NewString("abcde"),
		},
		// "abcde"[2:] => cde
		{
			left:       object.NewString("abcde"),
			firstIndex: object.NewInteger(2),
			expected:   object.NewString("cde"),
		},
		// "abcde"[-2:] => de
		{
			left:       object.NewString("abcde"),
			firstIndex: object.NewInteger(-2),
			expected:   object.NewString("de"),
		},
		// "abcde"[:2] => ab
		{
			left:        object.NewString("abcde"),
			secondIndex: object.NewInteger(2),
			expected:    object.NewString("ab"),
		},
		// "abcde"[:-2] => abc
		{
			left:        object.NewString("abcde"),
			secondIndex: object.NewInteger(-2),
			expected:    object.NewString("abc"),
		},
		// "abcde"[1:4] => bcde
		{
			left:        object.NewString("abcde"),
			firstIndex:  object.NewInteger(1),
			secondIndex: object.NewInteger(4),
			expected:    object.NewString("bcd"),
		},
		// "abcde"[1:-2] => bc
		{
			left:        object.NewString("abcde"),
			firstIndex:  object.NewInteger(1),
			secondIndex: object.NewInteger(-2),
			expected:    object.NewString("bc"),
		},
		// "abcde"[-2:-1] => d
		{
			left:        object.NewString("abcde"),
			firstIndex:  object.NewInteger(-2),
			secondIndex: object.NewInteger(-1),
			expected:    object.NewString("d"),
		},
		// "abcde"[-2:-3] => NULL
		{
			left:        object.NewString("abcde"),
			firstIndex:  object.NewInteger(-2),
			secondIndex: object.NewInteger(-3),
			expected:    object.NULL,
		},

		// ["a","b","c","d","e"][:] => ["a", "b", "c", "d", "e"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			expected: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
		},
		// ["a","b","c","d","e"][2:] => ["c", "d", "e"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex: object.NewInteger(2),
			expected: object.NewArrayWithObjects(
				object.NewString("c"), object.NewString("d"), object.NewString("e"),
			),
		},
		// ["a","b","c","d","e"][-2:] => ["d", "e"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex: object.NewInteger(-2),
			expected: object.NewArrayWithObjects(
				object.NewString("d"), object.NewString("e"),
			),
		},
		// ["a","b","c","d","e"][:2] => ["a", "b"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			secondIndex: object.NewInteger(2),
			expected: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"),
			),
		},
		// ["a","b","c","d","e"][:-2] => ["a", "b", "c"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			secondIndex: object.NewInteger(-2),
			expected: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
			),
		},

		// ["a","b","c","d","e"][1:4] => ["b", "c", "d"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex:  object.NewInteger(1),
			secondIndex: object.NewInteger(4),
			expected: object.NewArrayWithObjects(
				object.NewString("b"), object.NewString("c"), object.NewString("d"),
			),
		},
		// ["a","b","c","d","e"][1:-2] => ["b", "c"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex:  object.NewInteger(1),
			secondIndex: object.NewInteger(-2),
			expected: object.NewArrayWithObjects(
				object.NewString("b"), object.NewString("c"),
			),
		},
		// ["a","b","c","d","e"][-2:-1] => ["d"]
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex:  object.NewInteger(-2),
			secondIndex: object.NewInteger(-1),
			expected: object.NewArrayWithObjects(
				object.NewString("d"),
			),
		},
		// ["a","b","c","d","e"][-2:-3] => NULL
		{
			left: object.NewArrayWithObjects(
				object.NewString("a"), object.NewString("b"), object.NewString("c"),
				object.NewString("d"), object.NewString("e"),
			),
			firstIndex:  object.NewInteger(-2),
			secondIndex: object.NewInteger(-3),
			expected:    object.NULL,
		},

		// 12345[2:] => ERROR: range index operator not supported: INTEGER
		{
			left:       object.NewInteger(12345),
			firstIndex: object.NewInteger(2),
			expected:   object.NewErrorFormat("range index operator not supported: INTEGER"),
		},

		// "abcde"[true:] => ERROR: invalid type for first index: BOOLEAN
		{
			left:       object.NewString("abcde"),
			firstIndex: object.TRUE,
			expected:   object.NewErrorFormat("invalid type for first index: BOOLEAN"),
		},
		// "abcde"[:true] => ERROR: invalid type for second index: BOOLEAN
		{
			left:        object.NewString("abcde"),
			secondIndex: object.TRUE,
			expected:    object.NewErrorFormat("invalid type for second index: BOOLEAN"),
		},
	}

	for _, tc := range testcases {
		obj := evalRangeIndex(tc.left, tc.firstIndex, tc.secondIndex)
		if obj.Type() != tc.expected.Type() {
			t.Errorf("expected object to be a %s, got %s", tc.expected.Type(), obj.Type())
			continue
		}
		if obj.Inspect() != tc.expected.Inspect() {
			t.Errorf("unexpected result, got=%s, want=%s", obj.Inspect(), tc.expected.Inspect())
			continue
		}
	}
}
