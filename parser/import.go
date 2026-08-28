package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

// parseImportAsExpression is registered as the prefix parser for `import` so
// that using it where a value is expected produces a message naming the real
// problem, rather than the generic "no prefix parse function" error. A well
// placed import never reaches this: parseStatement handles token.IMPORT
// before any expression parsing begins.
func (p *Parser) parseImportAsExpression() ast.Expression {
	p.errors = append(p.errors, fmt.Sprintf("%d:%d: `import` is a statement and cannot be used as an expression", p.curToken.LineNumber, p.curToken.LinePosition))
	return nil
}

func (p *Parser) parseImport() ast.Statement {
	expression := &ast.Import{Token: p.curToken}

	if !p.expectPeek(token.STRING) {
		return nil
	}
	expression.Path = p.curToken.Literal

	if p.peekTokenIs(token.AS) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		expression.Alias = p.curToken.Literal
	}

	if p.peekTokenIs(token.ONLY) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		expression.Only = append(expression.Only, p.curToken.Literal)

		for p.peekTokenIs(token.COMMA) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			expression.Only = append(expression.Only, p.curToken.Literal)
		}
	}

	return expression
}
