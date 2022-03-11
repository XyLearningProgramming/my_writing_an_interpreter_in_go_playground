package my_parser

import (
	"monkey/my_ast"
	"monkey/token"
)

func (p *Parser) parseStatement() my_ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement: let <IDENT> = <EXPR>
// example: let a = 1 + 2
func (p *Parser) parseLetStatement() *my_ast.LetStatement {
	stmt := &my_ast.LetStatement{}
	if !p.isPeekToken(token.IDENT) {
		p.appendTokenError(token.IDENT, p.peekToken)
		return nil
	}
	p.nextToken()
	stmt.Ident = &my_ast.Identifier{
		Value: p.curToken.Literal,
	}
	if !p.isPeekToken(token.ASSIGN) {
		p.appendTokenError(token.ASSIGN, p.peekToken)
		return nil
	}
	p.nextToken()
	// TODO: parse following expressions from peekToken
	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseReturnStatement: return <EXPR>
// example: return a
func (p *Parser) parseReturnStatement() *my_ast.ReturnStatement {
	stmt := &my_ast.ReturnStatement{}
	p.nextToken()
	// TODO: parse expressions from curToken
	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
