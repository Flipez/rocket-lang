package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
		return nil
	}
	leftExp := prefix()

	if p.peekTokenIs(token.COLON) {
		return leftExp
	}

	for !p.peekTokenIs(token.SEMICOLON) &&
		!p.peekTokenIs(token.COLON) &&
		precedence < p.peekPrecedence() {
		if p.peekStartsStatement() {
			return leftExp
		}

		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// ambiguousPrefixInfix are the tokens the parser registers both ways: `[` opens
// an array literal or indexes, `(` groups or calls, `-` negates or subtracts.
// Every other infix token carries only the one meaning, so a line break in
// front of it cannot mean anything but continuation -- `* 3` is not an
// expression on its own, whereas `[1, 2]` and `-3` are.
//
// TestAmbiguousPrefixInfixIsComplete derives this set from the parser's own
// registration tables, so a token registered both ways later is not silently
// left out of the decision.
var ambiguousPrefixInfix = map[token.TokenType]bool{
	token.LBRACKET: true,
	token.LPAREN:   true,
	token.MINUS:    true,
}

// peekStartsStatement reports whether the next token should be read as the
// start of a new statement rather than as an operator continuing this one.
//
// Nothing terminates a statement in RocketLang, so a line break is the only
// separator there is, and it is only worth consulting for a token that could go
// either way. Without this, `puts("a")` on one line followed by
// `[1].each(puts)` on the next indexed the result of puts -- silently, since
// indexing nil is a runtime error rather than a syntax one.
func (p *Parser) peekStartsStatement() bool {
	return p.bracketDepth == 0 &&
		ambiguousPrefixInfix[p.peekToken.Type] &&
		p.peekToken.LineNumber > p.curToken.LineNumber
}
