package ast

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Import struct {
	Token token.Token
	Path  string
	Alias string
	Only  []string
}

func (ie *Import) TokenLiteral() string { return ie.Token.Literal }
func (ie *Import) String() string {
	var out bytes.Buffer

	out.WriteString(ie.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(strconv.Quote(ie.Path))

	if ie.Alias != "" {
		out.WriteString(" as ")
		out.WriteString(ie.Alias)
	}

	if len(ie.Only) > 0 {
		out.WriteString(" only ")
		out.WriteString(strings.Join(ie.Only, ", "))
	}

	return out.String()
}
