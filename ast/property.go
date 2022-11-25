package ast

import "github.com/flipez/rocket-lang/token"

type Property struct {
	Expression
	Assign

	Token    token.Token
	Left     Expression
	Property Expression
}

func (p *Property) TokenLiteral() string { return p.Token.Literal }
func (p *Property) String() string       { return p.Property.String() }
