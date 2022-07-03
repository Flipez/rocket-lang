package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseNext() *ast.Next {
	stmt := &ast.Next{Token: p.curToken}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		if p.peekTokenIs(token.RPAREN) {
			stmt.NextValue = p.createNil()
			p.nextToken()
		} else {
			stmt.NextValue = p.parseExpression(LOWEST)
		}

		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
	} else {
		stmt.NextValue = p.createNil()
	}

	return stmt
}
