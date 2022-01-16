package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type ObjectCall struct {
	Token  token.Token
	Object Expression
	Call   Expression
}

func (oce *ObjectCall) expressionNode()      {}
func (oce *ObjectCall) TokenLiteral() string { return oce.Token.Literal }
func (oce *ObjectCall) String() string {
	var out bytes.Buffer
	out.WriteString(oce.Object.String())
	out.WriteString(".")
	out.WriteString(oce.Call.String())

	return out.String()
}
