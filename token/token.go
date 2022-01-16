package token

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT" // add, foobar, x, y
	INT    = "INT"   // 123456
	FLOAT  = "FLOAT" // 123.456
	STRING = "STRING"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LBRACKET = "["
	RBRACKET = "]"

	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	END      = "END"
	RETURN   = "RETURN"

	EQ     = "=="
	NOT_EQ = "!="

	PERIOD = "."

	FOREACH = "FOREACH"
	IN      = "IN"

	EXPORT = "EXPORT"
	IMPORT = "IMPORT"
)

var keywords = map[string]TokenType{
	"def":     FUNCTION,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"end":     END,
	"else":    ELSE,
	"return":  RETURN,
	"foreach": FOREACH,
	"in":      IN,
	"export":  EXPORT,
	"import":  IMPORT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

var emojis = map[string]TokenType{
	"👍": TRUE,
	"👎": FALSE,
	"➕": PLUS,
}

func LookupEmoji(ident string) TokenType {
	if tok, ok := emojis[ident]; ok {
		return tok
	}

	return IDENT
}

var emojiLiterals = map[string]string{
	"👍": "true",
	"👎": "false",
	"➕": "+",
}

func LookupLiteral(ident string) string {
	if lit, ok := emojiLiterals[ident]; ok {
		return lit
	}

	return IDENT
}

func NewToken(tokenType TokenType, literal interface{}) Token {
	byteLiteral, ok := literal.(byte)
	if ok {
		return Token{Type: tokenType, Literal: string(byteLiteral)}
	}
	stringLiteral, ok := literal.(string)
	if ok {
		return Token{Type: tokenType, Literal: stringLiteral}
	}

	return Token{Type: ILLEGAL, Literal: fmt.Sprintf("%v", literal)}
}
