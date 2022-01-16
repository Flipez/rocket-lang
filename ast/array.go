package ast

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Array struct {
	Token    token.Token
	Elements []Expression
}

func (al *Array) expressionNode()      {}
func (al *Array) TokenLiteral() string { return al.Token.Literal }
func (al *Array) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
