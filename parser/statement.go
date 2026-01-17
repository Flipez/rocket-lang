package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.RETURN:
		return p.parseReturn()
	case token.BREAK:
		return p.parseBreak()
	case token.NEXT:
		return p.parseNext()
	case token.IDENT:
		// Check if this might be a multiple assignment (a, b, c = ...)
		if p.isMultipleAssignment() {
			return p.parseMultipleAssignment()
		}
		return p.parseExpressionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// isMultipleAssignment checks if the current position is the start of a multiple assignment
// by looking ahead for the pattern: IDENT COMMA
func (p *Parser) isMultipleAssignment() bool {
	// Check pattern: current is IDENT and next is COMMA
	return p.curToken.Type == token.IDENT && p.peekToken.Type == token.COMMA
}

// parseMultipleAssignment parses a, b, c = expression
func (p *Parser) parseMultipleAssignment() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	assign := &ast.Assign{Token: p.curToken}

	// Parse first identifier
	names := []ast.Expression{&ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}}

	// Parse remaining identifiers separated by commas
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		p.nextToken() // move to next identifier

		if p.curToken.Type != token.IDENT {
			p.errors = append(p.errors, "expected identifier in multiple assignment")
			return stmt
		}

		names = append(names, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
	}

	// Expect assignment operator
	if !p.expectPeek(token.ASSIGN) {
		return stmt
	}

	assign.Names = names
	assign.Token = p.curToken

	// Parse the value
	p.nextToken()
	assign.Value = p.parseExpression(LOWEST)

	stmt.Expression = assign
	return stmt
}
