package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseMethodCall(obj ast.Expression) ast.Expression {
	if _, ok := p.imports[obj.String()]; ok {
		p.expectPeek(token.IDENT)
		index := &ast.String{Token: p.curToken, Value: p.curToken.Literal}
		return &ast.Index{Left: obj, Index: index}
	}

	methodCall := &ast.ObjectCall{Token: p.curToken, Object: obj}
	p.nextToken()
	name := p.parseIdentifier()
	p.nextToken()
	methodCall.Call = p.parseCall(name)
	return methodCall
}
