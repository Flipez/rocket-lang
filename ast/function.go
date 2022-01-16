package ast

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Function struct {
	Token      token.Token
	Name       string
	Parameters []*Identifier
	Body       *Block
}

func (fl *Function) expressionNode()      {}
func (fl *Function) TokenLiteral() string { return fl.Token.Literal }
func (fl *Function) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
