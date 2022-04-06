package my_parser

import (
	"monkey/my_ast"
	token "monkey/my_token"
)

// parseStatement parse until curToken is ; or EOF
func (p *Parser) parseStatement() my_ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseReturnStatement: return <EXPR>
// example: return a
func (p *Parser) parseReturnStatement() *my_ast.ReturnStatement {
	stmt := &my_ast.ReturnStatement{}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *my_ast.ExpressionStatement {
	stmt := &my_ast.ExpressionStatement{
		Expression: p.parseExpression(LOWEST),
	}
	// NOTE: if next token is ;, then consume it;
	// in repl, expressions without ; is also legal, so no error here;
	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseBlockStatement: only called by parsingIfExpression()
func (p *Parser) parseBlockStatement() *my_ast.BlockStatement {
	if !p.isCurToken(token.LBRACE) {
		p.appendTokenError(token.LBRACE, p.curToken)
		return nil
	}
	p.nextToken()
	bs := &my_ast.BlockStatement{}
	for !p.isCurToken(token.RBRACE) && !p.isCurToken(token.EOF) {
		bs.Statements = append(bs.Statements, p.parseStatement())
		p.nextToken()
	}
	// TODO: what if an empty block statement?
	return bs
}
