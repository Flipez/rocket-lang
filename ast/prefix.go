package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Prefix struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *Prefix) expressionNode()      {}
func (pe *Prefix) TokenLiteral() string { return pe.Token.Literal }
func (pe *Prefix) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
