package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type ObjectCallExpression struct {
	Token  token.Token
	Object Expression
	Call   Expression
}

func (oce *ObjectCallExpression) expressionNode() {}

func (oce *ObjectCallExpression) TokenLiteral() string {
	return oce.Token.Literal
}

func (oce *ObjectCallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(oce.Object.String())
	out.WriteString(".")
	out.WriteString(oce.Call.String())

	return out.String()
}
