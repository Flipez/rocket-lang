package ast_test

import (
	"testing"

	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/parser"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`a = "test"`,
			"a = test",
		},
		{
			`a = [1,2,3, true]`,
			"a = [1, 2, 3, true]",
		},
		{
			"if (true)\n  print(true)\nelse\n  print(false)\nend",
			"if (true)\n  print(true)\nelse\n  print(false)\nend",
		},
		{
			`true ? print(true) : print(false)`,
			`true ? print(true) : print(false)`,
		},
		{
			"for i, e in [1, 2, 3] \n  print(i)\nend",
			"for i, e in [1, 2, 3] \n  print(i)\nend",
		},
		{
			"if (true)\n  return (true)\nelif (true)\n  return (true)\nelse\n  print(false)\nend",
			"if (true)\n  return (true)\nelif (true)\n  return (true)\nelse\n  print(false)\nend",
		},
		{
			"while (true)\n  print(true)\nend",
			"while (true)\n  print(true)\nend",
		},
		{
			"while (true)\n  print(true)\nend",
			"while (true)\n  print(true)\nend",
		},
		{
			"while\n  print(true)\n",
			"while (print(true))\n  \nend",
		},
		{
			"next",
			"next",
		},
		{
			"break",
			"break",
		},
		{
			"begin\ntrue\nrescue e\nfalse\nend",
			"begin\ntrue\nrescue e\nfalse\nend",
		},
		{
			"a, b, c = [1, 2, 3]",
			"a, b, c = [1, 2, 3]",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input, "test")
		p := parser.New(l)

		program := p.ParseProgram()

		if program.String() != tt.expected {
			t.Errorf("program.String() wrong.\ngot=\t\t`%q`,\nexpected=\t`%q`",
				program.String(),
				tt.expected)
		}
	}
}
