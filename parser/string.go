package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Token: p.curToken, Value: p.curToken.Literal}
}
