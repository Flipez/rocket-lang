package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		if p.peekTokenIs(token.RPAREN) {
			stmt.BreakValue = p.createNil()
			p.nextToken()
		} else {
			stmt.BreakValue = p.parseExpression(LOWEST)
		}

		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
	} else {
		stmt.BreakValue = p.createNil()
	}

	return stmt
}
