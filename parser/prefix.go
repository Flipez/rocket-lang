package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parsePrefix() ast.Expression {
	expression := &ast.Prefix{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
