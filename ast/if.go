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

func (ie *If) expressionNode()      {}
func (ie *If) TokenLiteral() string { return ie.Token.Literal }
func (ie *If) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString(" end")

	return out.String()
}
