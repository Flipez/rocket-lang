package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type String struct {
	Token token.Token
	Value string
}

func (sl *String) TokenLiteral() string { return sl.Token.Literal }
func (sl *String) String() string       { return sl.TokenLiteral() }
