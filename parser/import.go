package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseImport() ast.Expression {
	expression := &ast.Import{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	argList := p.parseExpressionList(token.RPAREN)

	expression.Location = argList[0]

	if len(argList) == 2 {
		expression.Name = argList[1]
	}

	return expression
}
