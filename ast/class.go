package ast

import (
	"github.com/flipez/rocket-lang/token"
)

func (c *Class) TokenLiteral() string { return c.Token.Literal }
func (c *Class) String() string       { return c.TokenLiteral() }

type Class struct {
	Expression
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
