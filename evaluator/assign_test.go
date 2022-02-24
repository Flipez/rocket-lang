package evaluator

import (
	"testing"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/object"
)

func newAstAssign(name, value ast.Expression) *ast.Assign {
	return &ast.Assign{
		Name:  name,
		Value: value,
	}
}

func prefilledEnv(objs map[string]object.Object) *object.Environment {
	env := object.NewEnvironment()
	for name, obj := range objs {
		env.Set(name, obj)
	}
	return env
}

func TestEvalAssign(t *testing.T) {
	testcases := []struct {
		env      *object.Environment
		a        *ast.Assign
		expected object.Object
	}{
		// valid string assignment
		{
			a: newAstAssign(
				&ast.Identifier{Value: "s"},
				&ast.String{Value: "abc"},
			),
			expected: object.NewString("abc"),
		},
		// assignment with an invalid expression
		{
			a: newAstAssign(
				&ast.Identifier{Value: "s"},
				&ast.Infix{
					Left:     &ast.String{Value: "abc"},
					Operator: "+",
					Right:    &ast.Integer{Value: 123},
				},
			),
			expected: object.NewErrorFormat("type mismatch: STRING + INTEGER"),
		},
		// valid array assignment with positive index
		{
			env: prefilledEnv(map[string]object.Object{
				"a": object.NewArrayWithObjects(
					object.NewString("a"),
					object.NewString("b"),
					object.NewString("c"),
				),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "a"},
					Index: &ast.Integer{Value: 0},
				},
				&ast.String{Value: "A"},
			),
			expected: object.NewString("A"),
		},
		// valid array assignment with negative index
		{
			env: prefilledEnv(map[string]object.Object{
				"a": object.NewArrayWithObjects(
					object.NewString("a"),
					object.NewString("b"),
					object.NewString("c"),
				),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "a"},
					Index: &ast.Integer{Value: -1},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewString("C"),
		},
		// array assignment with non integer index
		{
			env: prefilledEnv(map[string]object.Object{
				"a": object.NewArrayWithObjects(
					object.NewString("a"),
					object.NewString("b"),
					object.NewString("c"),
				),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "a"},
					Index: &ast.String{Value: "1"},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewErrorFormat("expected index to be an INTEGER, got STRING"),
		},
		// array assignment with integer index bigger than array size
		{
			env: prefilledEnv(map[string]object.Object{
				"a": object.NewArrayWithObjects(
					object.NewString("a"),
					object.NewString("b"),
					object.NewString("c"),
				),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "a"},
					Index: &ast.Integer{Value: 3},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewErrorFormat(
				"index out of range, got 3 but array has only 3 elements",
			),
		},
		// valid hash assignment
		{
			env: prefilledEnv(map[string]object.Object{
				"h": object.NewHash(map[object.HashKey]object.HashPair{
					object.NewString("a").HashKey(): object.HashPair{
						Key:   object.NewString("a"),
						Value: object.NewInteger(1),
					},
				}),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "h"},
					Index: &ast.String{Value: "a"},
				},
				&ast.Integer{Value: 2},
			),
			expected: object.NewInteger(2),
		},
		// hash assignment with invalid index
		{
			env: prefilledEnv(map[string]object.Object{
				"h": object.NewHash(map[object.HashKey]object.HashPair{
					object.NewString("a").HashKey(): object.HashPair{
						Key:   object.NewString("a"),
						Value: object.NewInteger(1),
					},
				}),
			}),
			a: newAstAssign(
				&ast.Index{
					Left: &ast.Identifier{Value: "h"},
					Index: &ast.Infix{
						Left:     &ast.String{Value: "abc"},
						Operator: "+",
						Right:    &ast.Integer{Value: 123},
					},
				},
				&ast.Integer{Value: 2},
			),
			expected: object.NewErrorFormat("expected index to be hashable"),
		},

		// valid string assignment with positive index
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.Integer{Value: 0},
				},
				&ast.String{Value: "A"},
			),
			expected: object.NewString("A"),
		},
		// valid string assignment with negative index
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.Integer{Value: -1},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewString("C"),
		},
		// string assignment with non integer index
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.String{Value: "1"},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewErrorFormat("expected index to be an INTEGER, got STRING"),
		},
		// string assignment with integer index bigger than string size
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.Integer{Value: 3},
				},
				&ast.String{Value: "C"},
			),
			expected: object.NewErrorFormat("index out of range, got 3 but string is only 3 long"),
		},
		// string assignment with valid index but invalid value type
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.Integer{Value: 0},
				},
				&ast.Integer{Value: 1},
			),
			expected: object.NewErrorFormat("expected STRING object, got INTEGER"),
		},
		// string assignment with valid index but invalid value size
		{
			env: prefilledEnv(map[string]object.Object{
				"s": object.NewString("abc"),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "s"},
					Index: &ast.Integer{Value: 0},
				},
				&ast.String{Value: "AA"},
			),
			expected: object.NewErrorFormat("expected STRING object to have a length of 1, got 2"),
		},
		// index assignment with non indexable object
		{
			env: prefilledEnv(map[string]object.Object{
				"i": object.NewInteger(123),
			}),
			a: newAstAssign(
				&ast.Index{
					Left:  &ast.Identifier{Value: "i"},
					Index: &ast.Integer{Value: 0},
				},
				&ast.Integer{Value: 4},
			),
			expected: object.NewErrorFormat("expected object to be indexable"),
		},
	}

	for _, tc := range testcases {
		if tc.env == nil {
			tc.env = object.NewEnvironment()
		}
		obj := evalAssign(tc.a, tc.env)
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
