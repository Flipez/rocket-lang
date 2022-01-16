package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseCall(callable ast.Expression) ast.Expression {
	exp := &ast.Call{Token: p.curToken, Callable: callable}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}
