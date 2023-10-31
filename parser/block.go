package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.END) && !p.curTokenIs(token.ELSE) && !p.curTokenIs(token.ELIF) && !p.curTokenIs(token.RESCUE) {

		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	if p.curTokenIs(token.RESCUE) {
		block.Rescue = &ast.Rescue{Token: p.curToken}
		p.expectPeek(token.IDENT)
		block.Rescue.ErrorIdent = p.curToken
		block.Rescue.Block = p.parseBlock()
	}

	return block
}
