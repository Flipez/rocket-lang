package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Begin struct {
	Token token.Token
	Block *Block
}

func (b *Begin) TokenLiteral() string { return b.Token.Literal }
func (b *Begin) String() string {
	var out bytes.Buffer

	out.WriteString("begin\n")
	out.WriteString(b.Block.String())
	out.WriteString("\nend")

	return out.String()
}
