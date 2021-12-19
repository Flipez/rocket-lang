package token

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
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	EQ     = "=="
	NOT_EQ = "!="

	PERIOD = "."

	FOREACH = "FOREACH"
	IN      = "IN"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"let":     LET,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"else":    ELSE,
	"return":  RETURN,
	"foreach": FOREACH,
	"in":      IN,
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
