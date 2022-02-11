package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseIndex(left ast.Expression) ast.Expression {
	exp := &ast.Index{Token: p.curToken, Left: left}

	p.nextToken()

	if p.curTokenIs(token.COLON) {
		p.nextToken()
		return p.parseRangeIndex(exp)
	}

	exp.Index = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.COLON) {
		p.nextToken()
		return p.parseRangeIndex(exp)
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

func (p *Parser) parseRangeIndex(index *ast.Index) ast.Expression {
	exp := &ast.RangeIndex{Token: index.Token, Left: index.Left, FirstIndex: index.Index}

	if p.curTokenIs(token.COLON) {
		p.nextToken()
	}

	if p.curTokenIs(token.RBRACKET) {
		return exp
	}

	exp.SecondIndex = p.parseExpression(LOWEST)

	if p.curTokenIs(token.COLON) {
		p.nextToken()
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}
