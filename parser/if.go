package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseIf() ast.Expression {
	expression := &ast.If{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if p.peekTokenIs(token.LBRACE) {
		p.nextToken()
	}
	expression.Consequence = p.parseBlock()

	if p.curTokenIs(token.RBRACE) {
		p.nextToken()
	}

	if p.curTokenIs(token.ELSE) {

		if p.peekTokenIs(token.LBRACE) {
			p.nextToken()
		}
		expression.Alternative = p.parseBlock()
	}
	return expression
}
