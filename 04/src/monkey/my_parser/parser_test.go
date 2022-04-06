package my_parser

import (
	"monkey/my_ast"
	lexer "monkey/my_lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	input := `
	let aa= 2; let b=1;
	let c = 1;
	`
	inputIdents := []string{
		"aa", "b", "c",
	}
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 3, len(prog.Statements))
	for idx, identValue := range inputIdents {
		assert.Equal(
			t,
			identValue,
			prog.Statements[idx].(*my_ast.LetStatement).Ident.Value)
	}
	assert.Nil(t, p.err)
}

func TestLetStatementError(t *testing.T) {
	input := `
	let 2;
	`
	l := lexer.New(input)
	p := New(l)
	p.Parse()
	assert.ErrorIs(t, p.err, ErrParseError)
}

func TestReturnStatement(t *testing.T) {
	input := `
	return a; return 123;
	`
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 2, len(prog.Statements))
	assert.Nil(t, p.err)
}
