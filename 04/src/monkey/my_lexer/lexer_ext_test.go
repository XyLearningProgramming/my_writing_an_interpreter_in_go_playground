package my_lexer

import (
	token "monkey/my_token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatToken(t *testing.T) {
	input := "1.123;1,2;1."
	expect := []*token.Token{
		{Type: token.FLOAT, Literal: "1.123"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INT, Literal: "1"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.INT, Literal: "2"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INT, Literal: "1"},
		{Type: token.ILLEGAL, Literal: "."},
	}
	testTokensWithInput(t, input, expect)
}

func TestStringToken(t *testing.T) {
	input := `
"boo"
'foo'
"boo\"foo"
'boo""foo'
'boo\\"foo'
'\"'
"\n"
`
	expects := []*token.Token{
		{Type: token.STRING, Literal: "boo"},
		{Type: token.STRING, Literal: "foo"},
		{Type: token.STRING, Literal: `boo"foo`},
		{Type: token.STRING, Literal: `boo""foo`},
		{Type: token.STRING, Literal: `boo\"foo`},
		{Type: token.STRING, Literal: `"`},
		{Type: token.STRING, Literal: "\n"},
	}
	testTokensWithInput(t, input, expects)
}

func TestStringTokenWithQuotes(t *testing.T) {
	input := "\"Hello\tWorld\n\""
	expects := []*token.Token{
		{Type: token.STRING, Literal: "Hello\tWorld\n"},
	}
	testTokensWithInput(t, input, expects)
}

func testTokensWithInput(t *testing.T, input string, expects []*token.Token) {
	lexer := New(input)
	for _, exp := range expects {
		tok := lexer.NextToken()
		if tok.Type == token.EOF {
			break
		}
		// fmt.Printf("t: type: %s: literal: %s\n", t.Type, t.Literal)
		assert.Equal(t, *exp, tok)
	}
}
