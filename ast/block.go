package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Block struct {
	Token      token.Token // the { token
	Statements []Statement
	Rescue     *Block
}

func (bs *Block) TokenLiteral() string { return bs.Token.Literal }
func (bs *Block) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	if bs.Rescue != nil {
		out.WriteString("\nrescue\n")
		for _, s := range bs.Rescue.Statements {
			out.WriteString(s.String())
		}
	}

	return out.String()
}
