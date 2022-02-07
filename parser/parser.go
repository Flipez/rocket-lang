package parser

import (
	"fmt"

	"github.com/flipez/rocket-lang/ast"
	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN      //=
	TERNARY     // ? :
	EQUALS      //==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	MODULO      // %
	PREFIX      // -X or !X
	CALL        // myFcuntion(X)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGN,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.LT_EQ:    LESSGREATER,
	token.GT:       LESSGREATER,
	token.GT_EQ:    LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.PERCENT:  MODULO,
	token.QUESTION: TERNARY,
	token.LPAREN:   CALL,
	token.PERIOD:   CALL,
	token.LBRACKET: INDEX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn

	imports map[string]struct{}
}

func New(l *lexer.Lexer, imports map[string]struct{}) *Parser {
	p := &Parser{
		l:       l,
		errors:  []string{},
		imports: imports,
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.FLOAT, p.parseFloat)
	p.registerPrefix(token.BANG, p.parsePrefix)
	p.registerPrefix(token.MINUS, p.parsePrefix)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIf)
	p.registerPrefix(token.FOREACH, p.parseForEach)
	p.registerPrefix(token.WHILE, p.parseWhile)
	p.registerPrefix(token.FUNCTION, p.parseFunction)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.LBRACKET, p.parseArray)
	p.registerPrefix(token.LBRACE, p.parseHash)
	p.registerPrefix(token.IMPORT, p.parseImport)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.ASSIGN, p.parseAssignExpression)
	p.registerInfix(token.PLUS, p.parseInfix)
	p.registerInfix(token.MINUS, p.parseInfix)
	p.registerInfix(token.SLASH, p.parseInfix)
	p.registerInfix(token.ASTERISK, p.parseInfix)
	p.registerInfix(token.PERCENT, p.parseInfix)
	p.registerInfix(token.EQ, p.parseInfix)
	p.registerInfix(token.NOT_EQ, p.parseInfix)
	p.registerInfix(token.PERIOD, p.parseMethodCall)
	p.registerInfix(token.LT, p.parseInfix)
	p.registerInfix(token.LT_EQ, p.parseInfix)
	p.registerInfix(token.GT, p.parseInfix)
	p.registerInfix(token.GT_EQ, p.parseInfix)
	p.registerInfix(token.LPAREN, p.parseCall)
	p.registerInfix(token.LBRACKET, p.parseIndex)
	p.registerInfix(token.QUESTION, p.parseTernary)

	// read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Token) {
	msg := fmt.Sprintf("%d:%d: expected next token to be %s, got %s instead",
		t.LineNumber, t.LinePosition, t.Type, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.Token) {
	msg := fmt.Sprintf("%d:%d: no prefix parse function for %s found", t.LineNumber, t.LinePosition, t.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(p.curToken)
		return false
	}
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}
