package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseImport() ast.Expression {
	expression := &ast.Import{Token: p.curToken}

	if !p.expectPeek(token.STRING) {
		return nil
	}
	expression.Path = p.curToken.Literal

	if p.peekTokenIs(token.AS) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		expression.Alias = p.curToken.Literal
	}

	if p.peekTokenIs(token.ONLY) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		expression.Only = append(expression.Only, p.curToken.Literal)

		for p.peekTokenIs(token.COMMA) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			expression.Only = append(expression.Only, p.curToken.Literal)
		}
	}

	return expression
}
