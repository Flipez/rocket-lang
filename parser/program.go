package parser

import (
	"path/filepath"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) ParseProgram() (*ast.Program, map[string]struct{}) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)

			if expStmt, ok := stmt.(*ast.ExpressionStatement); ok {
				if importExpr, ok := expStmt.Expression.(*ast.Import); ok {
					implicitVarName := filepath.Base(importExpr.Name.String())
					p.imports[implicitVarName] = struct{}{}
				}
			}
		}
		p.nextToken()
	}

	return program, p.imports
}
