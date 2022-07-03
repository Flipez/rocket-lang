package parser

import (
	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/token"
)

func (p *Parser) parseNil() ast.Expression {
	return &ast.Nil{Token: p.curToken}
}

func (p *Parser) createNil() ast.Expression {
	exp := &ast.Nil{Token: p.curToken}
	exp.Token.Literal = "nil"
	exp.Token.Type = token.NIL

	return exp
}
