package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseIf() ast.Expression {
	expression := &ast.If{Token: p.curToken, ConConPairs: make([]ast.ConditionConsequencePair, 1)}

	// Check if parentheses are used
	hasParens := p.peekTokenIs(token.LPAREN)
	if hasParens {
		p.nextToken() // consume LPAREN
		p.nextToken()
		expression.ConConPairs[0].Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.nextToken()
		expression.ConConPairs[0].Condition = p.parseExpression(LOWEST)
	}

	expression.ConConPairs[0].Consequence = p.parseBlock()

	for p.curTokenIs(token.ELIF) {
		// Check if parentheses are used for elif
		hasParens := p.peekTokenIs(token.LPAREN)
		if hasParens {
			p.nextToken() // consume LPAREN
			p.nextToken()

			var pair ast.ConditionConsequencePair
			pair.Condition = p.parseExpression(LOWEST)

			if !p.expectPeek(token.RPAREN) {
				return nil
			}

			pair.Consequence = p.parseBlock()
			expression.ConConPairs = append(expression.ConConPairs, pair)
		} else {
			p.nextToken()

			var pair ast.ConditionConsequencePair
			pair.Condition = p.parseExpression(LOWEST)

			pair.Consequence = p.parseBlock()
			expression.ConConPairs = append(expression.ConConPairs, pair)
		}
	}

	if p.curTokenIs(token.ELSE) {
		expression.Alternative = p.parseBlock()
	}

	return expression
}
