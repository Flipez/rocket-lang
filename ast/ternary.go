package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type Ternary struct {
	Expression
	Token       token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (t *Ternary) TokenLiteral() string { return t.Token.Literal }
func (t *Ternary) String() string {
	var out bytes.Buffer

	out.WriteString(t.Condition.String())
	out.WriteString(" ? ")
	out.WriteString(t.Condition.String())

	if t.Alternative != nil {
		out.WriteString(" : ")
		out.WriteString(t.Alternative.String())
	}

	return out.String()
}
