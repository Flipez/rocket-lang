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
			"if (true)\n  puts(true)\nelse\n  puts(false)\nend",
			"if (true)\n  puts(true)\nelse\n  puts(false)\nend",
		},
		{
			`true ? puts(true) : puts(false)`,
			`true ? puts(true) : puts(false)`,
		},
		{
			"foreach i, e in [1, 2, 3] {\n  puts(i)\n}",
			"foreach i, e in [1, 2, 3] {\n  puts(i)\n}",
		},
		{
			"if (true)\n  return (true)\nelse\n  puts(false)\nend",
			"if (true)\n  return (true)\nelse\n  puts(false)\nend",
		},
		{
			"while (true)\n  puts(true)\nend",
			"while (true)\n  puts(true)\nend",
		},
		{
			"while (true) {\n  puts(true)\n}",
			"while (true)\n  puts(true)\nend",
		},
		{
			"while {\n  puts(true)\n}",
			"",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l, make(map[string]struct{}))

		program, _ := p.ParseProgram()

		if program.String() != tt.expected {
			t.Errorf("program.String() wrong.\ngot=\t\t`%q`,\nexpected=\t`%q`",
				program.String(),
				tt.expected)
		}
	}
}
