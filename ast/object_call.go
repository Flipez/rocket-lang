package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type ObjectCall struct {
	StartToken token.Token
	Token      token.Token
	Object     Expression
	Call       Expression
}

func (oce *ObjectCall) TokenLiteral() string { return oce.Token.Literal }
func (oce *ObjectCall) String() string {
	var out bytes.Buffer
	out.WriteString(oce.Object.String())
	out.WriteString(".")
	out.WriteString(oce.Call.String())

	return out.String()
}
