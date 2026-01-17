package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseReturn() *ast.Return {
	stmt := &ast.Return{Token: p.curToken}

	p.nextToken()

	// Parse first expression
	firstExpr := p.parseExpression(LOWEST)

	// Check if there are comma-separated values (return 1, 2, 3)
	if p.peekTokenIs(token.COMMA) {
		// Multiple return values - wrap in array literal
		elements := []ast.Expression{firstExpr}

		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // consume comma
			p.nextToken() // move to next expression
			elements = append(elements, p.parseExpression(LOWEST))
		}

		// Create an array literal with all the expressions
		stmt.ReturnValue = &ast.Array{
			Token:    stmt.Token,
			Elements: elements,
		}
	} else {
		// Single return value
		stmt.ReturnValue = firstExpr
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
