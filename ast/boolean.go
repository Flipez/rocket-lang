package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.TokenLiteral() }
