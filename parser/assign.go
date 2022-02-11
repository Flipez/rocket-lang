package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseAssignExpression(name ast.Expression) ast.Expression {
	stmt := &ast.Assign{Token: p.curToken}
	if n, ok := name.(*ast.Identifier); ok {
		stmt.Name = n
	} else if index, ok := name.(*ast.Index); ok {
		stmt.Name = index
	} else {
		msg := fmt.Sprintf("%d:%d: expected assign token to be IDENT, got %s instead", p.curToken.LineNumber, p.curToken.LinePosition, name.TokenLiteral())
		p.errors = append(p.errors, msg)
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}
