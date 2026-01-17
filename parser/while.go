package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseWhile() ast.Expression {
	expression := &ast.While{Token: p.curToken}

	// Check if parentheses are used
	hasParens := p.peekTokenIs(token.LPAREN)
	if hasParens {
		p.nextToken() // consume LPAREN
		p.nextToken()
		expression.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.nextToken()
		expression.Condition = p.parseExpression(LOWEST)
	}

	expression.Body = p.parseBlock()

	return expression
}
