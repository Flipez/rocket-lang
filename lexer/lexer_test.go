package lexer

import (
	"testing"

	"github.com/flipez/rocket-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `five = 5;
	ten = 10;

	add = def(x, y) {
		x + y;
		};

	result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"}

	5 <= 10 >= 5;
	4 % 3;
	break("test")

	true and false
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "def"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.INT, "5"},
		{token.LT_EQ, "<="},
		{token.INT, "10"},
		{token.GT_EQ, ">="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "4"},
		{token.PERCENT, "%"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.BREAK, "break"},
		{token.LPAREN, "("},
		{token.STRING, "test"},
		{token.RPAREN, ")"},
		{token.TRUE, "true"},
		{token.AND, "and"},
		{token.FALSE, "false"},
		{token.EOF, ""},
	}

	l := New(input, "test")

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestNonASCIIStrings covers what a double-quoted literal outside ASCII used to
// become. l.ch is a byte, and appending string(l.ch) converts its numeric value
// to a rune and encodes that -- so every byte of a multi-byte character became a
// character of its own, and "тест" arrived as eight characters holding the
// values of its eight UTF-8 bytes.
//
// A single-quoted literal was always right, because it slices the input rather
// than rebuilding it. That the two disagreed is what gives the game away.
func TestNonASCIIStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"cyrillic", `"тест"`, "тест"},
		{"cjk", `"こんにちは"`, "こんにちは"},
		{"accents", `"café"`, "café"},
		{"an emoji inside a string", `"a👍b"`, "a👍b"},
		{"mixed with escapes", `"т\tе"`, "т\tе"},
		{"escapes still work", `"a\tb\nc\"d\\e"`, "a\tb\nc\"d\\e"},
		{"ascii is unchanged", `"plain"`, "plain"},
		{"empty", `""`, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := New(tt.input, "test").NextToken()

			if tok.Type != token.STRING {
				t.Fatalf("token type %s, want STRING", tok.Type)
			}
			if tok.Literal != tt.want {
				t.Errorf("literal %q, want %q", tok.Literal, tt.want)
			}
			if len([]rune(tok.Literal)) != len([]rune(tt.want)) {
				t.Errorf("%d characters, want %d", len([]rune(tok.Literal)), len([]rune(tt.want)))
			}
		})
	}
}

// TestQuotingStylesAgree checks the two kinds of literal produce the same
// string. They did not: '...' sliced the input and "..." rebuilt it wrongly.
func TestQuotingStylesAgree(t *testing.T) {
	for _, content := range []string{"тест", "こんにちは", "café", "plain", "a👍b"} {
		double := New(`"`+content+`"`, "test").NextToken()
		single := New(`'`+content+`'`, "test").NextToken()

		if double.Literal != single.Literal {
			t.Errorf("%q: double-quoted gave %q, single-quoted gave %q", content, double.Literal, single.Literal)
		}
	}
}

// TestNonASCIIIdentifier covers the same mistake in readIdentifier. It only ever
// worked because a name and its uses were mangled identically, so it matched
// itself while being wrong.
func TestNonASCIIIdentifier(t *testing.T) {
	tok := New("föö", "test").NextToken()

	if tok.Type != token.IDENT {
		t.Fatalf("token type %s, want IDENT", tok.Type)
	}
	if tok.Literal != "föö" {
		t.Errorf("literal %q, want %q", tok.Literal, "föö")
	}
}

// TestEmojiTokensStillWork guards the emoji handling, which assembles a
// character from its bytes and is the one place that was already deliberate
// about multi-byte input.
func TestEmojiTokensStillWork(t *testing.T) {
	tests := []struct {
		input string
		want  token.TokenType
	}{
		{"👍", token.TRUE},
		{"👎", token.FALSE},
		{"➕", token.PLUS},
	}

	for _, tt := range tests {
		if tok := New(tt.input, "test").NextToken(); tok.Type != tt.want {
			t.Errorf("%s lexed as %s, want %s", tt.input, tok.Type, tt.want)
		}
	}
}

// TestTokenPositions pins the line and column of every token in a small
// multi-line program. Positions used to be taken after the token had been read,
// so a token at the end of a line was tagged with the line after it -- which
// made a line break invisible to the parser.
func TestTokenPositions(t *testing.T) {
	// 1: a
	// 2: [1]
	// 3: foo(b
	// 4:  - 1)
	input := "a\n[1]\nfoo(b\n - 1)"

	expected := []struct {
		expectedType     token.TokenType
		expectedLiteral  string
		expectedLine     int
		expectedPosition int
	}{
		{token.IDENT, "a", 1, 1},
		{token.LBRACKET, "[", 2, 1},
		{token.INT, "1", 2, 2},
		{token.RBRACKET, "]", 2, 3},
		{token.IDENT, "foo", 3, 1},
		{token.LPAREN, "(", 3, 4},
		{token.IDENT, "b", 3, 5},
		{token.MINUS, "-", 4, 2},
		{token.INT, "1", 4, 4},
		{token.RPAREN, ")", 4, 5},
	}

	l := New(input, "test")

	for i, tt := range expected {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("token %d: type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("token %d: literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.LineNumber != tt.expectedLine || tok.LinePosition != tt.expectedPosition {
			t.Errorf("token %d (%q): position wrong. expected=%d:%d, got=%d:%d",
				i, tok.Literal, tt.expectedLine, tt.expectedPosition, tok.LineNumber, tok.LinePosition)
		}
	}
}

// TestPositionsAfterComment checks that a comment does not carry its own
// position over to the token that follows it. The comment used to be skipped by
// recursing into NextToken, which is easy to break when the position is
// recorded by the caller.
func TestPositionsAfterComment(t *testing.T) {
	// 1: # a comment
	// 2: x = 1 # trailing
	// 3: # another
	// 4: y
	input := "# a comment\nx = 1 # trailing\n# another\ny"

	expected := []struct {
		literal  string
		line     int
		position int
	}{
		{"x", 2, 1},
		{"=", 2, 3},
		{"1", 2, 5},
		{"y", 4, 1},
	}

	l := New(input, "test")

	for i, tt := range expected {
		tok := l.NextToken()

		if tok.Literal != tt.literal {
			t.Fatalf("token %d: literal wrong. expected=%q, got=%q", i, tt.literal, tok.Literal)
		}
		if tok.LineNumber != tt.line || tok.LinePosition != tt.position {
			t.Errorf("token %d (%q): position wrong. expected=%d:%d, got=%d:%d",
				i, tok.Literal, tt.line, tt.position, tok.LineNumber, tok.LinePosition)
		}
	}
}
