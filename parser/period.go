package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parsePeriod(obj ast.Expression) ast.Expression {
	if _, ok := obj.(*ast.This); ok {
		p.nextToken()
		fmt.Printf("%#v", p.curToken)
		prop := &ast.Property{Token: p.curToken, Left: obj}
		prop.Property = p.parseExpression(p.curPrecedence())

		return prop
	}

	if _, ok := p.imports[obj.String()]; ok {
		p.expectPeek(token.IDENT)
		index := &ast.String{Token: p.curToken, Value: p.curToken.Literal}
		return &ast.Index{Left: obj, Index: index}
	}

	p.nextToken()
	name := p.parseIdentifier()

	if ok := p.peekTokenIs(token.LPAREN); !ok {
		index := &ast.String{Token: p.curToken, Value: p.curToken.Literal}
		return &ast.Index{Left: obj, Index: index}
	}

	p.nextToken()
	return &ast.ObjectCall{Token: p.curToken, Object: obj, Call: p.parseCall(name)}
}
