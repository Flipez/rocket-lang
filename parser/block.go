package parser

import (
	"fmt"
	"strings"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.blockDepth++

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.END) && !p.curTokenIs(token.ELSE) && !p.curTokenIs(token.ELIF) && !p.curTokenIs(token.RESCUE) {

		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	p.blockDepth--

	// Every caller of parseBlock expects an explicit terminator. Reaching EOF
	// means the block was never closed, so report it instead of silently
	// treating end-of-file as the end of the block -- which used to let an
	// unterminated block swallow the rest of the source with no diagnostic.
	if p.curTokenIs(token.EOF) {
		// Don't pile a second diagnostic onto a position that already failed:
		// an unterminated block is usually the knock-on effect of the earlier
		// error, and reporting both just obscures the real cause.
		position := fmt.Sprintf("%d:%d:", block.Token.LineNumber, block.Token.LinePosition)
		if len(p.errors) == 0 || !strings.HasPrefix(p.errors[len(p.errors)-1], position) {
			p.errors = append(p.errors, position+" unexpected end of file, expected `end` to close the block opened here")
		}

		return block
	}

	if p.curTokenIs(token.RESCUE) {
		block.Rescue = &ast.Rescue{Token: p.curToken}
		p.expectPeek(token.IDENT)
		block.Rescue.ErrorIdent = p.curToken
		block.Rescue.Block = p.parseBlock()
	}

	return block
}
