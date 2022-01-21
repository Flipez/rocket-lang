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
}

func New(input string) *Lexer {
	l := &Lexer{input: input, currentLine: 0, positionInLine: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.positionInLine += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok.Type = token.EQ
			tok.Literal = literal
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
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		} else {
			tok.Type = token.SLASH
			tok.Literal = string(l.ch)
		}
	case '+':
		tok.Type = token.PLUS
		tok.Literal = string(l.ch)
	case '-':
		tok.Type = token.MINUS
		tok.Literal = string(l.ch)
	case '*':
		tok.Type = token.ASTERISK
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
		tok.Literal = l.readString()
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

	tok.LineNumber = l.currentLine
	tok.LinePosition = l.positionInLine
	l.readChar()
	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	id := ""

	position := l.position
	rposition := l.readPosition

	for l.isIdentifier(l.ch) {
		id += string(l.ch)
		l.readChar()
	}

	if strings.Contains(id, ".") {
		offset := strings.Index(id, ".")
		id = id[:offset]

		l.position = position
		l.readPosition = rposition
		for offset > 0 {
			l.readChar()
			offset--
		}
	}

	return id
}

func (l *Lexer) isNewline() bool {
	if l.ch == '\n' {
		l.currentLine += 1
		l.positionInLine = 0
		return true
	}

	return false
}

func (l *Lexer) isIdentifier(ch byte) bool {
	return !isDigit(ch) && !l.isWhitespace(ch) && !isBrace(ch) && !isOperator(ch) && !isComparison(ch) && !isCompound(ch) && !isBrace(ch) && !isParen(ch) && !isBracket(ch) && !isEmpty(ch)
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
	return ch == ',' || ch == ':' || ch == '"' || ch == ';'
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
