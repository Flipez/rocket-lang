package parser

import (
	"github.com/flipez/rocket-lang/ast"
)

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expression := &ast.Infix{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}
