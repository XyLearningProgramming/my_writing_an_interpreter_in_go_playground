package my_parser

import (
	"fmt"
	"monkey/token"
)

// nextToken:
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) isCurToken(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) appendError(msg string) {
	if p.err == nil {
		p.err = ErrParseError
	}
	p.err = fmt.Errorf("%s: %w", msg, p.err)
}

func (p *Parser) appendTokenError(expect token.TokenType, value token.Token) {
	p.appendError(
		fmt.Sprintf(
			"expecting token %s, but got %s with literal %s instead",
			string(expect), string(value.Type), value.Literal,
		),
	)
}
