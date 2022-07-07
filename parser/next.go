package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseNext() *ast.Next {
	stmt := &ast.Next{Token: p.curToken}

	return stmt
}
