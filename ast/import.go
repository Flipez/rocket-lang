package ast

import (
	"bytes"
	"fmt"

	"github.com/flipez/rocket-lang/token"
)

type Import struct {
	Token    token.Token
	Location Expression
	Name     Expression
}

func (ie *Import) TokenLiteral() string { return ie.Token.Literal }
func (ie *Import) String() string {
	var out bytes.Buffer

	out.WriteString(ie.TokenLiteral())
	out.WriteString("(")
	out.WriteString(fmt.Sprintf("\"%s\"", ie.Location))
	if ie.Name != nil {
		out.WriteString(", ")
		out.WriteString(fmt.Sprintf("\"%s\"", ie.Name))
	}
	out.WriteString(")")

	return out.String()
}
