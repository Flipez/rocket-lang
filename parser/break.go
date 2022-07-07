package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}

	return stmt
}
