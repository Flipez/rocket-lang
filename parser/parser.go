package parser

import (
	"github.com/flipez/monkey/ast"
	"github.com/flipez/monkey/lexer"
	"github.com/flipez/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (*Parser) ParseProgram() *ast.Program {
	return nil
}
