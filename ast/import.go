package ast

import (
	"bytes"
	"fmt"

	"github.com/flipez/rocket-lang/token"
)

type Import struct {
	Token token.Token
	Name  Expression
}

func (ie *Import) expressionNode() {}

func (ie *Import) TokenLiteral() string { return ie.Token.Literal }

func (ie *Import) String() string {
	var out bytes.Buffer

	out.WriteString(ie.TokenLiteral())
	out.WriteString("(")
	out.WriteString(fmt.Sprintf("\"%s\"", ie.Name))
	out.WriteString(")")

	return out.String()
}
