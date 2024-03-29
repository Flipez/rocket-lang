package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseMethodCall(obj ast.Expression) ast.Expression {
	startToken := p.curToken
	if _, ok := p.imports[obj.String()]; ok {
		p.expectPeek(token.IDENT)
		index := &ast.String{Token: p.curToken, Value: p.curToken.Literal}
		return &ast.Index{Left: obj, Index: index}
	}

	p.nextToken()
	name := p.parseIdentifier()

	if ok := !p.peekTokenIs(token.LPAREN); ok {
		index := &ast.String{Token: p.curToken, Value: p.curToken.Literal}
		return &ast.Index{Left: obj, Index: index}
	}

	p.nextToken()
	return &ast.ObjectCall{StartToken: startToken, Token: p.curToken, Object: obj, Call: p.parseCall(name)}
}
