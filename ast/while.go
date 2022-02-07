package ast

import (
	"fmt"

	"github.com/flipez/rocket-lang/token"
)

type While struct {
	Token     token.Token
	Condition Expression
	Body      *Block
}

func (w *While) TokenLiteral() string { return w.Token.Literal }
func (w *While) String() string {
	return fmt.Sprintf("%s (%s)\n  %s\nend", w.TokenLiteral(), w.Condition, w.Body)
}
