package ast

import (
	"bytes"
	"strings"

	"github.com/flipez/rocket-lang/token"
)

type Assign struct {
	Token token.Token
	Names []Expression // Support multiple assignment targets
	Value Expression
}

func (as *Assign) TokenLiteral() string { return as.Token.Literal }
func (as *Assign) String() string {
	var out bytes.Buffer

	// Handle multiple names
	if len(as.Names) > 1 {
		names := make([]string, len(as.Names))
		for i, name := range as.Names {
			names[i] = name.String()
		}
		out.WriteString(strings.Join(names, ", "))
	} else if len(as.Names) == 1 {
		out.WriteString(as.Names[0].String())
	}

	out.WriteString(" = ")
	out.WriteString(as.Value.String())
	return out.String()
}
