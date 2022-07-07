package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Next struct {
	Token token.Token
}

func (n *Next) TokenLiteral() string { return n.Token.Literal }
func (n *Next) String() string {
	var out bytes.Buffer

	out.WriteString(n.TokenLiteral())
	return out.String()
}
