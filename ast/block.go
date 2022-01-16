package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Block struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *Block) statementNode()       {}
func (bs *Block) TokenLiteral() string { return bs.Token.Literal }
func (bs *Block) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
