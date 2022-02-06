package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type Integer struct {
	Token token.Token
	Value int64
}

func (il *Integer) TokenLiteral() string { return il.Token.Literal }
func (il *Integer) String() string       { return il.TokenLiteral() }
