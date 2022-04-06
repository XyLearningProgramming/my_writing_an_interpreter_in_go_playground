package my_lexer

import (
	token "monkey/my_token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatToken(t *testing.T) {
	input := "1.123;1,2;1."
	expect := []token.Token{
		{Type: token.FLOAT, Literal: "1.123"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INT, Literal: "1"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.INT, Literal: "2"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INT, Literal: "1"},
		{Type: token.ILLEGAL, Literal: "."},
	}
	lexer := New(input)
	for _, exp := range expect {
		tok := lexer.NextToken()
		if tok.Type == token.EOF {
			break
		}
		// fmt.Printf("t: type: %s: literal: %s\n", t.Type, t.Literal)
		assert.Equal(t, exp, tok)
	}
}
