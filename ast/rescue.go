package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type Rescue struct {
	Token      token.Token // the { token
	ErrorIdent token.Token
	Block      *Block
}

func (r *Rescue) TokenLiteral() string { return r.Token.Literal }
func (r *Rescue) String() string {
	return r.Block.String()
}
