package ast

import (
	"bytes"

	"github.com/flipez/rocket-lang/token"
)

type ConditionConsequencePair struct {
	Condition   Expression
	Consequence *Block
}

type If struct {
	Token       token.Token // the if token
	ConConPairs []ConditionConsequencePair
	Alternative *Block
}

func (ie *If) TokenLiteral() string { return ie.Token.Literal }
func (ie *If) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(ie.ConConPairs[0].Condition.String())
	out.WriteString(")\n  ")
	out.WriteString(ie.ConConPairs[0].Consequence.String())

	for _, pair := range ie.ConConPairs[1:] {
		out.WriteString("\nelif (")
		out.WriteString(pair.Condition.String())
		out.WriteString(")\n  ")
		out.WriteString(pair.Consequence.String())
	}

	if ie.Alternative != nil {
		out.WriteString("\nelse\n  ")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString("\nend")

	return out.String()
}
