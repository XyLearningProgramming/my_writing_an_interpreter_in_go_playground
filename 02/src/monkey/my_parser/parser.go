package my_parser

import (
	"errors"
	"monkey/lexer"
	"monkey/my_ast"
	"monkey/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	err       error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

var ErrParseError = errors.New("parse error")

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:          l,
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	// NOTE: Pratt parsing: link each type of tokens to one or multiple parsing functions
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)

	// lexer.NextToken() will continue to produce EOF if finished without error
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() *my_ast.Program {
	prog := &my_ast.Program{
		Statements: []my_ast.Statement{},
	}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		// Error?
		p.nextToken()
	}
	return prog
}

func (p *Parser) Error() error {
	return p.err
}

func (p *Parser) registerPrefix(t token.TokenType, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *Parser) registerInfix(t token.TokenType, f infixParseFn) {
	p.infixParseFns[t] = f
}
