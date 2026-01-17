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
		if p.isMultipleAssignment() {
			return p.parseMultipleAssignment()
		}
		return p.parseExpressionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) isMultipleAssignment() bool {
	return p.curToken.Type == token.IDENT && p.peekToken.Type == token.COMMA
}

func (p *Parser) parseMultipleAssignment() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	assign := &ast.Assign{Token: p.curToken}

	names := []ast.Expression{&ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		if p.curToken.Type != token.IDENT {
			p.errors = append(p.errors, "expected identifier in multiple assignment")
			return stmt
		}

		names = append(names, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
	}

	if !p.expectPeek(token.ASSIGN) {
		return stmt
	}

	assign.Names = names
	assign.Token = p.curToken

	p.nextToken()
	assign.Value = p.parseExpression(LOWEST)

	stmt.Expression = assign
	return stmt
}
