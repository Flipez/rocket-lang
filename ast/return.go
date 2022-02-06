package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Return struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *Return) TokenLiteral() string { return rs.Token.Literal }
func (rs *Return) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())

	if rs.ReturnValue != nil {
		out.WriteString(" (")
		out.WriteString(rs.ReturnValue.String())
		out.WriteString(")")
	}
	return out.String()
}
