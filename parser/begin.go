package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseBegin() ast.Expression {
	expression := &ast.Begin{Token: p.curToken}
	p.nextToken()
	expression.Block = p.parseBlock()
	return expression
}
