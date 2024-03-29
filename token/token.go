package token

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type         TokenType
	Literal      string
	LineNumber   int
	LinePosition int
	File         string
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
	PERCENT  = "%"

	QUESTION = "?"

	LT    = "<"
	LT_EQ = "<="
	GT    = ">"
	GT_EQ = ">="
	AND   = "AND"
	OR    = "OR"

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
	ELIF     = "ELIF"
	ELSE     = "ELSE"
	END      = "END"
	RETURN   = "RETURN"
	BREAK    = "BREAK"
	NEXT     = "NEXT"

	EQ     = "=="
	NOT_EQ = "!="

	PERIOD = "."

	FOREACH = "FOREACH"
	IN      = "IN"

	WHILE = "WHILE"

	EXPORT = "EXPORT"
	IMPORT = "IMPORT"

	NIL = "NIL"

	BEGIN  = "BEGIN"
	RESCUE = "RESCUE"

	RANGE_ROCKET_E = "->"
	RANGE_ROCKET_I = "=>"
	RANGE_STEPPER  = "^"
)

var keywords = map[string]TokenType{
	"def":     FUNCTION,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"elif":    ELIF,
	"end":     END,
	"else":    ELSE,
	"return":  RETURN,
	"break":   BREAK,
	"next":    NEXT,
	"foreach": FOREACH,
	"in":      IN,
	"while":   WHILE,
	"export":  EXPORT,
	"import":  IMPORT,
	"nil":     NIL,
	"and":     AND,
	"or":      OR,
	"begin":   BEGIN,
	"rescue":  RESCUE,
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

func NewToken(tokenType TokenType, literal interface{}, line int, position int) Token {
	byteLiteral, ok := literal.(byte)
	if ok {
		return Token{Type: tokenType, Literal: string(byteLiteral), LineNumber: line, LinePosition: position}
	}
	stringLiteral, ok := literal.(string)
	if ok {
		return Token{Type: tokenType, Literal: stringLiteral, LineNumber: line, LinePosition: position}
	}

	return Token{Type: ILLEGAL, Literal: fmt.Sprintf("%v", literal)}
}
