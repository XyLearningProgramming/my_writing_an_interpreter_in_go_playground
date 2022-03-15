package lexer

import (
	"monkey/token"
	"strings"
)

func (l *Lexer) readNumberWithDot() *token.Token {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if isDot(l.ch) && isDigit(l.peekChar()) {
		sb := strings.Builder{}
		sb.WriteString(l.input[position:l.position])
		sb.WriteRune(floatDot)
		l.readChar()
		position = l.position
		for isDigit(l.ch) {
			l.readChar()
		}
		sb.WriteString(l.input[position:l.position])
		return &token.Token{
			Literal: sb.String(),
			Type:    token.FLOAT,
		}
	} else {
		return &token.Token{
			Literal: l.input[position:l.position],
			Type:    token.INT,
		}
	}
}
