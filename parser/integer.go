package parser

import (
	"fmt"
	"strconv"

	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseInteger() ast.Expression {
	lit := &ast.Integer{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		token := p.curToken
		msg := fmt.Sprintf("%d:%d: could not parse `%q` as integer", token.LineNumber, token.LinePosition, token.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
