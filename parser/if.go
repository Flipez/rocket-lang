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
		if p.peekTokenIs(token.IF) {
			// "else if" (two keywords) shares its terminating `end` with the
			// enclosing if/else. Treat it like `elif`: parse the nested if as
			// a nested expression (which will itself end on that shared
			// `end`) rather than handing off to parseBlock, which would
			// unconditionally step past that `end` and keep consuming
			// whatever statements follow.
			elseToken := p.curToken
			p.nextToken() // advance onto the nested IF token

			nested := p.parseIf()

			expression.Alternative = &ast.Block{
				Token: elseToken,
				Statements: []ast.Statement{
					&ast.ExpressionStatement{Token: elseToken, Expression: nested},
				},
			}
		} else {
			expression.Alternative = p.parseBlock()
		}
	}

	return expression
}
