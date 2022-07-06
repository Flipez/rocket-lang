package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Break struct {
	Token token.Token
}

func (b *Break) TokenLiteral() string { return b.Token.Literal }
func (b *Break) String() string {
	var out bytes.Buffer

	out.WriteString(b.TokenLiteral())
	return out.String()
}
