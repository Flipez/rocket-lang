package ast

import (
	"bytes"
	"fmt"

	"github.com/flipez/rocket-lang/token"
)

type Index struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *Index) TokenLiteral() string { return ie.Token.Literal }
func (ie *Index) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type RangeIndex struct {
	Token       token.Token
	Left        Expression
	FirstIndex  Expression
	SecondIndex Expression
}

func (rie *RangeIndex) TokenLiteral() string { return rie.Token.Literal }
func (rie *RangeIndex) String() string {
	str := fmt.Sprintf("(%s[", rie.Left)
	if rie.FirstIndex != nil {
		str += rie.FirstIndex.String()
	}
	str += ":"
	if rie.SecondIndex != nil {
		str += rie.SecondIndex.String()
	}
	return str + "])"
}
