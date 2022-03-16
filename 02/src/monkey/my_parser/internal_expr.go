package my_parser

import (
	"fmt"
	"monkey/my_ast"
	"monkey/token"
	"strconv"
)

type PrecedenceLevel int

const (
	_ PrecedenceLevel = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PREFIX      // -X or !X
	PRODUCT     // *
	CALL        // myFunction(X)
)

var InfixOperatorToPrecedences = map[my_ast.InfixOperator]PrecedenceLevel{
	my_ast.INOP_MINUS:    SUM,
	my_ast.INOP_PLUS:     SUM,
	my_ast.INOP_ASTERISK: PRODUCT,
	my_ast.INOP_SLASH:    PRODUCT,
	my_ast.INOP_LT:       LESSGREATER,
	my_ast.INOP_GT:       LESSGREATER,
	my_ast.INOP_EQ:       EQUALS,
	my_ast.INOP_NOT_EQ:   EQUALS,
}

func tokenPrecedenceLevel(t *token.Token) PrecedenceLevel {
	if pl, pok :=
		InfixOperatorToPrecedences[my_ast.InfixOperator(t.Type)]; pok {
		return pl
	}
	return LOWEST
}

type (
	prefixParseFn func() my_ast.Expression
	infixParseFn  func(my_ast.Expression) my_ast.Expression
)

func (p *Parser) parseExpression(precedence PrecedenceLevel) my_ast.Expression {
	prefixExpr, pok := p.prefixParseFns[p.curToken.Type]
	if !pok {
		p.appendExprFuncError(p.curToken, true)
		return nil
	}
	leftExpr := prefixExpr()

	// NOTE: consume to semicolon or EOF
	// or when meet a higher precedence with current token
	for p.peekToken.Type != token.SEMICOLON &&
		p.peekToken.Type != token.EOF &&
		precedence < tokenPrecedenceLevel(&p.peekToken) {
		infixFn, iok := p.infixParseFns[p.peekToken.Type]
		if !iok {
			return leftExpr
		}
		p.nextToken()
		leftExpr = infixFn(leftExpr)
	}
	return leftExpr
}

func (p *Parser) parseIdentifier() my_ast.Expression {
	return &my_ast.Identifier{
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() my_ast.Expression {
	val, err := strconv.ParseUint(p.curToken.Literal, 10, 64)
	if err != nil {
		p.appendError(fmt.Sprintf("cannot parse %s as uint :%v", p.curToken.Literal, err))
		return nil
	}
	return &my_ast.Integer{Value: val}
}

func (p *Parser) parseFloatLiteral() my_ast.Expression {
	val, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		p.appendError(fmt.Sprintf("cannot parse %s as float: %v", p.curToken.Literal, err))
		return nil
	}
	return &my_ast.Float{Value: val}
}

func (p *Parser) parseBooleanLiteral() my_ast.Expression {
	if p.curToken.Type == token.TRUE {
		return &my_ast.Boolean{Value: true}
	} else if p.curToken.Type == token.FALSE {
		return &my_ast.Boolean{Value: false}
	} else {
		p.appendError(fmt.Sprintf(
			"cannot parse %s with type %s as boolean", p.curToken.Literal, p.curToken.Type,
		))
		return nil
	}
}

func (p *Parser) parsePrefixExpression() my_ast.Expression {
	expr := &my_ast.PrefixExpression{
		Operator: my_ast.PrefixOperator(p.curToken.Type),
	}
	p.nextToken()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfixExpression(left my_ast.Expression) my_ast.Expression {
	exp := &my_ast.InfixExpression{
		Left:     left,
		Operator: my_ast.InfixOperator(p.curToken.Type),
	}
	precedence := tokenPrecedenceLevel(&p.curToken)
	p.nextToken()
	exp.Right = p.parseExpression(precedence)
	return exp
}

func (p *Parser) parseGroupedExpression() my_ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.isPeekToken(token.RPAREN) {
		p.appendTokenError(token.RPAREN, p.peekToken)
		p.nextToken()
		return nil
	}
	p.nextToken()
	return exp
}

func (p *Parser) parseIfExpression() my_ast.Expression {
	// parse if condition as expression
	p.nextToken()
	if !p.isCurToken(token.LPAREN) {
		p.appendTokenError(token.LPAREN, p.curToken)
		return nil
	}
	p.nextToken()
	ie := &my_ast.IfExpression{Condition: p.parseExpression(LOWEST)}
	p.nextToken()
	if !p.isCurToken(token.RPAREN) {
		p.appendTokenError(token.RPAREN, p.curToken)
		return nil
	}
	// parse if consequence as block statement
	p.nextToken()

	ie.Consequence = p.parseBlockStatement()
	if !p.isCurToken(token.RBRACE) {
		p.appendTokenError(token.RBRACE, p.curToken)
		return nil
	}
	p.nextToken()
	// no "else" token is legal, return immediately
	if !p.isCurToken(token.ELSE) {
		return ie
	}
	// parse if alternative as block statement
	p.nextToken()
	ie.Alternative = p.parseBlockStatement()
	if !p.isCurToken(token.RBRACE) {
		p.appendTokenError(token.RBRACE, p.curToken)
		return nil
	}
	return ie
}

func (p *Parser) parseFunction() my_ast.Expression {
	p.nextToken()
	// TODO: no function name after fn?
	fe := &my_ast.Function{
		Parameters: p.parseFunctionParameters(),
	}
	if !p.isCurToken(token.RPAREN) {
		p.appendTokenError(token.RPAREN, p.curToken)
		return nil
	}
	p.nextToken()
	fe.Body = p.parseBlockStatement()
	return fe
}

func (p *Parser) parseFunctionParameters() []*my_ast.Identifier {
	if !p.isCurToken(token.LPAREN) {
		p.appendTokenError(token.LPAREN, p.curToken)
		return nil
	}
	params := []*my_ast.Identifier{}
	p.nextToken()
	if p.isPeekToken(token.RPAREN) {
		p.nextToken()
		return params
	}
	params = append(params, &my_ast.Identifier{Value: p.curToken.Literal})
	for p.isPeekToken(token.COMMA) {
		p.nextToken()
		p.nextToken()
		params = append(params, &my_ast.Identifier{Value: p.curToken.Literal})
	}
	p.nextToken()
	return params
}
