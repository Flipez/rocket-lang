package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseHash() ast.Expression {
	hash := &ast.Hash{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.peekTokenIs(token.COLON) {
			// `{` only ever opens a hash literal: curly braces stopped being
			// block delimiters in #89. Someone writing `foreach i in 3 { ... }`
			// lands here, and "expected next token to be :" tells them nothing,
			// so say what `{` means when the key is followed by the closing
			// brace rather than by a value.
			if p.peekTokenIs(token.RBRACE) {
				p.errors = append(p.errors, fmt.Sprintf(
					"%d:%d: expected `:` after a hash key; `{` opens a hash literal, not a block -- a block needs no braces and closes with `end`",
					p.peekToken.LineNumber, p.peekToken.LinePosition))

				return nil
			}

			p.peekError(p.curToken, token.COLON)

			return nil
		}

		p.nextToken()

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}

	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}
