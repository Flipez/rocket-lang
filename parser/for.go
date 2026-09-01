package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseForEach() ast.Expression {
	expression := &ast.Foreach{Token: p.curToken}

	p.nextToken()
	expression.Ident = p.curToken.Literal

	if p.peekTokenIs(token.COMMA) {

		p.nextToken()

		if !p.peekTokenIs(token.IDENT) {
			p.errors = append(p.errors, fmt.Sprintf(
				"%d:%d: second argument to foreach must be ident, got %v",
				p.peekToken.LineNumber,
				p.peekToken.LinePosition,
				p.peekToken))
			return nil
		}
		p.nextToken()

		expression.Index = expression.Ident
		expression.Ident = p.curToken.Literal

	}

	if !p.expectPeek(token.IN) {
		p.errors = append(p.errors, fmt.Sprintf(
			"%d:%d: expected `in` after foreach arguments, got %v",
			p.peekToken.LineNumber,
			p.peekToken.LinePosition,
			p.peekToken))
		return nil
	}
	p.nextToken()

	if p.peekTokenIs(token.RANGE_ROCKET_E) || p.peekTokenIs(token.RANGE_ROCKET_I) {
		if p.peekTokenIs(token.RANGE_ROCKET_I) {
			expression.Inclusive = true
		}
		expression.Start = p.parseExpression(LOWEST)
		p.nextToken()
	}
	if p.curTokenIs(token.RANGE_ROCKET_E) || p.curTokenIs(token.RANGE_ROCKET_I) {
		if p.peekTokenIs(token.RANGE_ROCKET_I) {
			expression.Inclusive = true
		}
		p.nextToken()
		expression.Value = p.parseExpression(LOWEST)
	} else {
		expression.Value = p.parseExpression(LOWEST)
		if expression.Value == nil {
			return nil
		}

		// don't allow negative iterable integer
		if prefix, ok := expression.Value.(*ast.Prefix); ok && prefix.Operator == "-" {
			p.errors = append(p.errors, fmt.Sprintf(
				"%d:%d: expected positive value got %v",
				p.peekToken.LineNumber,
				p.peekToken.LinePosition,
				prefix))
			return nil
		}
	}

	if p.peekTokenIs(token.RANGE_STEPPER) {
		p.nextToken()
		p.nextToken()
		expression.Step = p.parseExpression(LOWEST)
	}

	expression.Body = p.parseBlock()

	return expression
}
