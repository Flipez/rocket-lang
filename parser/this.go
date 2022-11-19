package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseThis() ast.Expression {
	return &ast.This{Token: p.curToken}
}
