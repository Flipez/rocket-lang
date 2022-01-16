package lexer

import (
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
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
			tok = token.NewToken(token.EQ, literal)
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.NewToken(token.NOT_EQ, literal)
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			tok = l.NextToken()
		} else {
			tok = token.NewToken(token.SLASH, l.ch)
		}
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '.':
		tok = token.NewToken(token.PERIOD, l.ch)
	case ':':
		tok = token.NewToken(token.COLON, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '[':
		tok = token.NewToken(token.LBRACKET, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACKET, l.ch)
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
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

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

	for isIdentifier(l.ch) {
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

func isIdentifier(ch byte) bool {
	return !isDigit(ch) && !isWhitespace(ch) && !isBrace(ch) && !isOperator(ch) && !isComparison(ch) && !isCompound(ch) && !isBrace(ch) && !isParen(ch) && !isBracket(ch) && !isEmpty(ch)
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
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
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
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
