package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Index struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *Index) expressionNode()      {}
func (ie *Index) TokenLiteral() string { return ie.Token.Literal }
func (ie *Index) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
