package my_parser

import (
	"monkey/lexer"
	"monkey/my_ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpressionStatement(t *testing.T) {
	input := "a;b; let a=b"
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 3, len(prog.Statements))
	assert.Nil(t, p.err)
	assert.Equal(t, "a", prog.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.Identifier).Value)
	assert.Equal(t, "b", prog.Statements[1].(*my_ast.ExpressionStatement).Expression.(*my_ast.Identifier).Value)
	letStmt, lok := prog.Statements[2].(*my_ast.LetStatement)
	assert.True(t, lok)
	assert.Equal(t, "a", letStmt.Ident.Value)
	assert.Equal(t, "b", letStmt.Value.(*my_ast.Identifier).String())
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
	assert.EqualValues(t, 1, prog.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.Integer).Value)
	assert.EqualValues(t, 1.234, prog.Statements[1].(*my_ast.ExpressionStatement).Expression.(*my_ast.Float).Value)
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
	assert.EqualValues(t, 5, prefixNode.Right.(*my_ast.Integer).Value)
	prefixNode, pok = prog.Statements[1].(*my_ast.ExpressionStatement).
		Expression.(*my_ast.PrefixExpression)
	assert.True(t, pok)
	assert.EqualValues(t, my_ast.PREOP_MINUS, prefixNode.Operator)
	assert.EqualValues(t, 15.5, prefixNode.Right.(*my_ast.Float).Value)
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
		assert.EqualValues(t, test.leftExpr, infixNode.Left.(*my_ast.Integer).Value)
		assert.EqualValues(t, test.leftExpr, infixNode.Right.(*my_ast.Integer).Value)
		assert.EqualValues(t, test.operator, infixNode.Operator)
	}
}

type TestWithExpect struct {
	input  string
	expect string
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []TestWithExpect{
		{"-b*c", "(-(b*c));"},
		{"a*b-c", "((a*b)-c);"},
		{"!-c", "(!(-c));"},
		{"-1+2", "((-1)+2);"},
	}
	testStringedStatements(t, tests)
}

func TestBooleanExpression(t *testing.T) {
	tests := []TestWithExpect{
		{"return true", "return true;"},
		{"true + false", "(true+false);"},
	}
	testStringedStatements(t, tests)
}

func TestGroupedExpression(t *testing.T) {
	tests := []TestWithExpect{
		{"1+(2+ 3) +4", "((1+(2+3))+4);"},
		{"(5 +5 )/2", "((5+5)/2);"},
		{"-(\t5+ \t5)", "(-(5+5));"},
		{"!(true == true)", "(!(true==true));"},
	}
	testStringedStatements(t, tests)
}
func TestIfExpression(t *testing.T) {
	input := "if (x< y) { x} else\n\n{x;return y;}"
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.Nil(t, p.Error())
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 1, len(prog.Statements))
	assert.Equal(t, "if((x<y)){x;}else{x;return y;};", prog.Statements[0].String())
	es, eok := prog.Statements[0].(*my_ast.ExpressionStatement)
	assert.True(t, eok)
	is, iok := es.Expression.(*my_ast.IfExpression)
	assert.True(t, iok)
	assert.Equal(t, "(x<y)", is.Condition.String())
	assert.Equal(t, "x", is.Consequence.Statements[0].(*my_ast.ExpressionStatement).Expression.(*my_ast.Identifier).Value)
	assert.Equal(t, "y", is.Alternative.Statements[1].(*my_ast.ReturnStatement).Value.(*my_ast.Identifier).Value)
}

func testStringedStatements(t *testing.T, tests []TestWithExpect) {
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

func TestParseLetReturn(t *testing.T) {
	input := `
	let a = b; return c
	`
	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()
	assert.NotNil(t, prog)
	assert.NotNil(t, prog.Statements)
	assert.Equal(t, 2, len(prog.Statements))
	assert.Nil(t, p.err)
	assert.Equal(t, "a", prog.Statements[0].(*my_ast.LetStatement).Ident.Value)
	assert.Equal(t, "b", prog.Statements[0].(*my_ast.LetStatement).Value.(*my_ast.Identifier).Value)
	assert.Equal(t, "c", prog.Statements[1].(*my_ast.ReturnStatement).Value.(*my_ast.Identifier).Value)
}

func TestParseFunctionExpression(t *testing.T) {
	tests := []TestWithExpect{
		{"fn(x,y)\n{x+y;}", "fn(x,y){(x+y);};"},
	}
	testStringedStatements(t, tests)
}
