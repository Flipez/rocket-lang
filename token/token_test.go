package token

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	token := NewToken(RBRACE, "}")
	if token != (Token{Type: RBRACE, Literal: "}"}) {
		t.Fatalf("wrong token received, got type `%s` and literal `%s`", token.Type, token.Literal)
	}
	token = NewToken(RBRACE, true)
	if token != (Token{Type: ILLEGAL, Literal: "true"}) {
		t.Fatalf("wrong token received, got type `%s` and literal `%s`", token.Type, token.Literal)
	}
}
