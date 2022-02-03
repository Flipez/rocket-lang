package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseWhile() ast.Expression {
	expression := &ast.While{Token: p.curToken}

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
	expression.Body = p.parseBlock()

	if p.curTokenIs(token.RBRACE) {
		p.nextToken()
	}

	return expression
}
