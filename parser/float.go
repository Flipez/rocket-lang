package parser

import (
	"fmt"
	"strconv"

	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseFloat() ast.Expression {
	lit := &ast.Float{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		token := p.curToken
		msg := fmt.Sprintf("%d:%d: could not parse %q as float", token.LineNumber, token.LinePosition, token.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
