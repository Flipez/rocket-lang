package ast

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Hash struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *Hash) TokenLiteral() string { return hl.Token.Literal }
func (hl *Hash) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
