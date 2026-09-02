package lexer

import (
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Lexer struct {
	input          string
	position       int  // current position in input (points to current char)
	readPosition   int  // current reading position in input (after current char)
	ch             byte // current char under examination
	currentLine    int
	positionInLine int
	file           string
}

func New(input string, file string) *Lexer {
	l := &Lexer{input: input, currentLine: 1, positionInLine: 0, file: file}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// The newline is checked before it is overwritten, so the line advances as
	// the lexer steps off the newline rather than onto it. Incrementing on the
	// way in left currentLine pointing at the next line while the lexer was
	// still sitting on the newline that ended the previous one, which is how a
	// token at the end of a line came to be tagged with the line after it.
	if l.ch == '\n' {
		l.currentLine += 1
		l.positionInLine = 0
	}

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
	l.positionInLine += 1
}

// NextToken returns the next token, tagged with the position at which it
// starts.
//
// The position has to be taken before the token is read, because reading it
// moves the lexer past it: readIdentifier stops only once it has consumed the
// character that ended the identifier, so an identifier at the end of a line
// used to be tagged with the line *after* it. That made a line break invisible
// to the parser, and a line break is the only thing that separates two
// statements in a language without terminators.
func (l *Lexer) NextToken() token.Token {
	l.skipIgnored()

	line, position := l.currentLine, l.positionInLine

	tok := l.scanToken()
	tok.LineNumber = line
	tok.LinePosition = position
	tok.File = l.file

	return tok
}

// skipIgnored advances past anything that is not a token. A comment runs to the
// end of its line, and may be followed by more whitespace and more comments.
func (l *Lexer) skipIgnored() {
	l.skipWhitespace()

	for l.ch == '#' {
		l.skipComment()
		l.skipWhitespace()
	}
}

// scanToken reads one token. It does not set the position: NextToken does that,
// from where the token started.
func (l *Lexer) scanToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok.Type = token.AND
			tok.Literal = "and"
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok.Type = token.OR
			tok.Literal = "or"
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok.Type = token.EQ
			tok.Literal = literal
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok.Type = token.RANGE_ROCKET_I
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type = token.ASSIGN
			tok.Literal = string(l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok.Type = token.NOT_EQ
			tok.Literal = literal
		} else {
			tok.Type = token.BANG
			tok.Literal = string(l.ch)
		}
	case '/':
		tok.Type = token.SLASH
		tok.Literal = string(l.ch)
	case '+':
		tok.Type = token.PLUS
		tok.Literal = string(l.ch)
	case '-':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok.Type = token.RANGE_ROCKET_E
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type = token.MINUS
			tok.Literal = string(l.ch)
		}
	case '*':
		tok.Type = token.ASTERISK
		tok.Literal = string(l.ch)
	case '%':
		tok.Type = token.PERCENT
		tok.Literal = string(l.ch)
	case '?':
		tok.Type = token.QUESTION
		tok.Literal = string(l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.LT_EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type = token.LT
			tok.Literal = string(l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.GT_EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok.Type = token.GT
			tok.Literal = string(l.ch)
		}
	case ';':
		tok.Type = token.SEMICOLON
		tok.Literal = string(l.ch)
	case ',':
		tok.Type = token.COMMA
		tok.Literal = string(l.ch)
	case '.':
		tok.Type = token.PERIOD
		tok.Literal = string(l.ch)
	case ':':
		tok.Type = token.COLON
		tok.Literal = string(l.ch)
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = string(l.ch)
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = string(l.ch)
	case '(':
		tok.Type = token.LPAREN
		tok.Literal = string(l.ch)
	case ')':
		tok.Type = token.RPAREN
		tok.Literal = string(l.ch)
	case '[':
		tok.Type = token.LBRACKET
		tok.Literal = string(l.ch)
	case ']':
		tok.Type = token.RBRACKET
		tok.Literal = string(l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readDoubleQuoteString()
	case '\'':
		tok.Type = token.STRING
		tok.Literal = l.readSingleQuoteString()
	case '^':
		tok.Type = token.RANGE_STEPPER
		tok.Literal = "^"
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			return tok
		} else if i := isEmoji(l.ch); i > 0 {
			out := make([]byte, i)

			for i := 0; i < len(out); i++ {
				out[i] = l.ch
				l.readChar()
			}

			tok.Literal = token.LookupLiteral(string(out))
			tok.Type = token.LookupEmoji(string(out))

			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readDoubleQuoteString() string {
	// Collected as bytes, and turned into a string once at the end.
	//
	// Appending string(l.ch) instead looks equivalent and is not: l.ch is a
	// byte, and string(aByte) converts its numeric value to a rune and encodes
	// that. So each byte of a multi-byte character became a character of its
	// own -- "тест" arrived as eight characters holding the values of its eight
	// UTF-8 bytes, which is why size() answered 8 and reverse() and remove_last()
	// cut characters in half.
	var out []byte

	for {
		l.readChar()

		if l.ch == '\\' {
			nextCh := l.peekChar()
			switch nextCh {
			case '"':
				l.readChar()
				out = append(out, '"')
			case 'n':
				l.readChar()
				out = append(out, '\n')
			case 't':
				l.readChar()
				out = append(out, '\t')
			case 'r':
				l.readChar()
				out = append(out, '\r')
			case '\\':
				l.readChar()
				out = append(out, '\\')
			default:
				// Not an escape sequence anyone recognises, so keep what is
				// there, backslash included.
				out = append(out, l.ch)
			}
		} else if l.ch == '"' || l.ch == 0 {
			break
		} else {
			out = append(out, l.ch)
		}
	}

	return string(out)
}

func (l *Lexer) readSingleQuoteString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '\'' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	// Bytes, for the same reason as readDoubleQuoteString. A name outside ASCII
	// was mangled the same way, and only worked because its definition and its
	// uses were mangled identically -- so it matched itself while being wrong.
	var id []byte

	for l.isIdentifier(l.ch) {
		id = append(id, l.ch)
		l.readChar()
	}

	return string(id)
}

func (l *Lexer) isNewline() bool {
	return l.ch == '\n'
}

func (l *Lexer) isIdentifier(ch byte) bool {
	return !l.isWhitespace(ch) && !isBrace(ch) && !isOperator(ch) && !isComparison(ch) && !isCompound(ch) && !isBrace(ch) && !isParen(ch) && !isBracket(ch) && !isEmpty(ch)
}

func (l *Lexer) isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || l.isNewline() || ch == '\r'
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '%' || ch == '-' || ch == '/' || ch == '*'
}

func isComparison(ch byte) bool {
	return ch == '=' || ch == '>' || ch == '<'
}

func isCompound(ch byte) bool {
	return ch == ',' || ch == ':' || ch == '"' || ch == ';' || ch == '.'
}

func isBrace(ch byte) bool {
	return ch == '{' || ch == '}'
}

func isBracket(ch byte) bool {
	return ch == '[' || ch == ']'
}

// is parenthesis
func isParen(ch byte) bool {
	return ch == '(' || ch == ')'
}

func isEmpty(ch byte) bool {
	return ch == 0
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.isNewline() || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for !l.isNewline() && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isEmoji(ch byte) int {
	switch int(ch) {
	case 240:
		return 4
	case 226:
		return 3
	}

	return -1
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
