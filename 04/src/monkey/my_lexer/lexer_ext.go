package my_lexer

import (
	token "monkey/my_token"
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

func (l *Lexer) readString(startQuote byte) string {
	sb := &strings.Builder{}
	for {
		l.readChar()
		if l.ch == startQuote || l.ch == 0 {
			break
		}
		// all back slash rules go here
		switch l.ch {
		case '\\':
			l.readChar()
			switch l.ch {
			case 'r':
				sb.WriteByte('\r')
			case 't':
				sb.WriteByte('\t')
			case 'n':
				sb.WriteByte('\n')
			default:
				sb.WriteByte(l.ch)
			}
			continue
		}
		sb.WriteByte(l.ch)
	}
	return sb.String()
}
