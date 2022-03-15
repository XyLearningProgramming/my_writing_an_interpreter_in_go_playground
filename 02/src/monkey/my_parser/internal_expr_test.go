package my_parser

import (
	"monkey/lexer"
	"monkey/my_ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpressionStatement(t *testing.T) {
	input := "a;b"
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 2, len(prog.Statements))
	assert.Nil(t, p.err)
	assert.Equal(t, "a", prog.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.Identifier).Value)
	assert.Equal(t, "b", prog.Statements[1].(*my_ast.ExpressionStatement).Expression.(*my_ast.Identifier).Value)
}

func TestParseNumberStatement(t *testing.T) {
	input := "1;1.234"
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 2, len(prog.Statements))
	assert.Nil(t, p.err)
	assert.EqualValues(t, 1, prog.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.IntegerLiteral).Value)
	assert.EqualValues(t, 1.234, prog.Statements[1].(*my_ast.ExpressionStatement).Expression.(*my_ast.FloatLiteral).Value)
}

func TestPrefixExpressionStatement(t *testing.T) {
	input := "!5;\n-15.5;"
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 2, len(prog.Statements))
	assert.Nil(t, p.err)
	prefixNode, pok := prog.Statements[0].(*my_ast.ExpressionStatement).
		Expression.(*my_ast.PrefixExpression)
	assert.True(t, pok)
	assert.EqualValues(t, my_ast.PREOP_BANG, prefixNode.Operator)
	assert.EqualValues(t, 5, prefixNode.Right.(*my_ast.IntegerLiteral).Value)
	prefixNode, pok = prog.Statements[1].(*my_ast.ExpressionStatement).
		Expression.(*my_ast.PrefixExpression)
	assert.True(t, pok)
	assert.EqualValues(t, my_ast.PREOP_MINUS, prefixNode.Operator)
	assert.EqualValues(t, 15.5, prefixNode.Right.(*my_ast.FloatLiteral).Value)
}

func TestInfixExpressionStatement(t *testing.T) {
	tests := []struct {
		input     string
		leftExpr  interface{}
		rightExpr interface{}
		operator  string
	}{
		{"5+5", 5, 5, "+"},
		{"5-5", 5, 5, "-"},
		{"5 * 5", 5, 5, "*"},
		{"5/\r\n5", 5, 5, "/"},
		{"5>   5", 5, 5, ">"},
		{"5\t<5", 5, 5, "<"},
		{"5== 5;", 5, 5, "=="},
		{"5 !=5", 5, 5, "!="},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		prog := p.Parse()
		assert.Nil(t, p.Error())
		assert.NotNil(t, prog)
		assert.NotNil(t, prog.Statements)
		assert.Equal(t, 1, len(prog.Statements))
		infixNode := prog.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.InfixExpression)
		assert.NotNil(t, infixNode)
		assert.EqualValues(t, test.leftExpr, infixNode.Left.(*my_ast.IntegerLiteral).Value)
		assert.EqualValues(t, test.leftExpr, infixNode.Right.(*my_ast.IntegerLiteral).Value)
		assert.EqualValues(t, test.operator, infixNode.Operator)
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"-b*c", "(-(b*c))"},
		{"a*b-c", "((a*b)-c)"},
		{"!-c", "(!(-c))"},
		{"-1+2", "((-1)+2)"},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		prog := p.Parse()
		assert.Nil(t, p.Error())
		assert.NotNil(t, prog)
		assert.NotNil(t, prog.Statements)
		assert.Equal(t, 1, len(prog.Statements))
		assert.Equal(t, test.expect, prog.Statements[0].String())
	}
}

// func TestParseLetReturn(t *testing.T) {
// 	input := `
// 	let a = b; return c
// 	`
// 	l := lexer.New(input)
// 	p := New(l)
// 	prog := p.Parse()
// 	assert.NotNil(t, prog)
// 	assert.NotNil(t, prog.Statements)
// 	assert.Equal(t, 2, len(prog.Statements))
// 	assert.Nil(t, p.err)
// 	assert.Equal(t, "a", prog.Statements[0].(*my_ast.LetStatement).Ident.Value)
// 	assert.Equal(t, "b", prog.Statements[0].(*my_ast.LetStatement).Value.(*my_ast.Identifier).Value)
// 	assert.Equal(t, "c", prog.Statements[1].(*my_ast.ReturnStatement).Value.(*my_ast.Identifier).Value)
// }
