package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseArray() ast.Expression {
	array := &ast.Array{Token: p.curToken}

	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}
