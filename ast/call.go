package ast

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Call struct {
	Token     token.Token // The ( token
	Callable  Expression
	Arguments []Expression
}

func (ce *Call) expressionNode()      {}
func (ce *Call) TokenLiteral() string { return ce.Token.Literal }
func (ce *Call) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Callable.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
