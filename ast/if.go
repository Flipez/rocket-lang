package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type If struct {
	Token       token.Token // the if token
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (ie *If) TokenLiteral() string { return ie.Token.Literal }
func (ie *If) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(ie.Condition.String())
	out.WriteString(")\n  ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("\nelse\n  ")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString("\nend")

	return out.String()
}
