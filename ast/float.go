package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type Float struct {
	Token token.Token
	Value float64
}

func (fl *Float) expressionNode()      {}
func (fl *Float) TokenLiteral() string { return fl.Token.Literal }
func (fl *Float) String() string       { return fl.Token.Literal }
