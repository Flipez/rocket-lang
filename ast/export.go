package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

// Export marks a top-level binding as part of a module's public surface.
// Value is nil for the `export Name` form, which exports an existing binding.
type Export struct {
	Token token.Token
	Name  string
	Value Expression
}

func (e *Export) TokenLiteral() string { return e.Token.Literal }
func (e *Export) String() string {
	var out bytes.Buffer

	out.WriteString(e.TokenLiteral())
	out.WriteString(" ")

	if fn, ok := e.Value.(*Function); ok {
		out.WriteString(fn.String())
		return out.String()
	}

	out.WriteString(e.Name)

	if e.Value != nil {
		out.WriteString(" = ")
		out.WriteString(e.Value.String())
	}

	return out.String()
}
