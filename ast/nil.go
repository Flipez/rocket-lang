package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type Nil struct {
	Token token.Token
}

func (n *Nil) TokenLiteral() string { return n.Token.Literal }
func (n *Nil) String() string       { return n.TokenLiteral() }
