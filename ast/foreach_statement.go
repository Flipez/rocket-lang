package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type ForeachStatement struct {
	Token token.Token
	Index string
	Ident string
	Value Expression
	Body  *BlockStatement
}

func (fes *ForeachStatement) expressionNode()      {}
func (fes *ForeachStatement) TokenLiteral() string { return fes.Token.Literal }
func (fes *ForeachStatement) String() string {
	var out bytes.Buffer
	out.WriteString("foreach ")
	out.WriteString(fes.Ident)
	out.WriteString(" ")
	out.WriteString(fes.Value.String())
	out.WriteString(fes.Body.String())
	return out.String()
}
