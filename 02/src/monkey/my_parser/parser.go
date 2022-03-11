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
}

var ErrParseError = errors.New("parse error")

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}

	// TODO: Pratt parsing: link each type of tokens to one or multiple parsing functions

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
