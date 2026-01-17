package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseReturn() *ast.Return {
	stmt := &ast.Return{Token: p.curToken}

	p.nextToken()

	firstExpr := p.parseExpression(LOWEST)

	if p.peekTokenIs(token.COMMA) {
		elements := []ast.Expression{firstExpr}

		for p.peekTokenIs(token.COMMA) {
			p.nextToken()
			p.nextToken()
			elements = append(elements, p.parseExpression(LOWEST))
		}

		stmt.ReturnValue = &ast.Array{
			Token:    stmt.Token,
			Elements: elements,
		}
	} else {
		stmt.ReturnValue = firstExpr
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
