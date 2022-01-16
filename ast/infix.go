package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Infix struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *Infix) expressionNode()      {}
func (ie *Infix) TokenLiteral() string { return ie.Token.Literal }
func (ie *Infix) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
