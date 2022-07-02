package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}

	p.nextToken()

	stmt.BreakValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
