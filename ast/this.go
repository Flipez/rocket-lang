package ast

import (
	"github.com/flipez/rocket-lang/token"
)

type This struct {
	Expression
	Token token.Token
}
