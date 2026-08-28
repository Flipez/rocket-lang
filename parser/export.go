package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseExport() ast.Statement {
	stmt := &ast.Export{Token: p.curToken}

	if p.blockDepth > 0 {
		p.errors = append(p.errors, fmt.Sprintf("%d:%d: `export` is only valid at the top level of a module", p.curToken.LineNumber, p.curToken.LinePosition))
		return stmt
	}

	// export def Name(...) ... end
	if p.peekTokenIs(token.FUNCTION) {
		p.nextToken()

		fn, ok := p.parseFunction().(*ast.Function)
		if !ok || fn.Name == "" {
			p.errors = append(p.errors, fmt.Sprintf("%d:%d: exported function must have a name", p.curToken.LineNumber, p.curToken.LinePosition))
			return stmt
		}

		stmt.Name = fn.Name
		stmt.Value = fn
		return stmt
	}

	if !p.expectPeek(token.IDENT) {
		return stmt
	}
	stmt.Name = p.curToken.Literal

	// export Name = expr
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		stmt.Value = p.parseExpression(LOWEST)
	}

	// otherwise: export Name, exporting an existing binding
	return stmt
}
