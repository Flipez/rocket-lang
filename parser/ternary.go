package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseTernary(left ast.Expression) ast.Expression {
	expression := &ast.Ternary{Token: p.curToken, Condition: left}

	p.nextToken()

	expression.Consequence = p.parseExpression(p.curPrecedence())

	if p.expectPeek(token.COLON) {
		p.nextToken()
		expression.Alternative = p.parseExpression(p.curPrecedence())
	}
	return expression
}
