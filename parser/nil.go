package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseNil() ast.Expression {
	return &ast.Nil{Token: p.curToken}
}
