package my_ast

import (
	"monkey/token"
	"strconv"
	"strings"
)

type Node interface {
	// DebugString: for debugging purposes, usually literal value of a token;
	// if an expression, ususally literal value of the first token;
	// if an statement, usually call DebugString() on its first expression;
	// if an ident, literal value of its `Value` field;
	DebugString() string
	// String: output the formatted codes if legal;
	// it has the following rules:
	// one statement per line;
	// each line will end with semicolon;
	// each token has a space in between;
	String() string
}

const (
	NodeStringNewLine    = "\n"
	NodeStringSemiColon  = ";"
	NodeStringTokenSpace = " "
)

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

// root node

type Program struct {
	Statements []Statement
}

func (p *Program) DebugString() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].DebugString()
	}
	return ""
}

func (p *Program) String() string {
	sb := strings.Builder{}
	for idx, s := range p.Statements {
		sb.WriteString(s.String())
		if idx != len(p.Statements)-1 {
			sb.WriteString(NodeStringNewLine)
		}
	}
	return sb.String()
}

// statements

type LetStatement struct {
	Ident *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) DebugString() string {
	return l.Ident.DebugString()
}

func (l *LetStatement) String() string {
	sb := strings.Builder{}
	sb.WriteString(token.LookupKeywords(token.LET))
	sb.WriteString(NodeStringTokenSpace)
	sb.WriteString(l.Ident.Value)
	sb.WriteString(NodeStringTokenSpace)
	sb.WriteString(token.ASSIGN)
	sb.WriteString(NodeStringTokenSpace)
	sb.WriteString(l.Value.String())
	sb.WriteString(NodeStringSemiColon)
	return sb.String()
}

type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) DebugString() string {
	return r.Value.DebugString()
}

func (r *ReturnStatement) String() string {
	sb := strings.Builder{}
	sb.WriteString(token.LookupKeywords(token.RETURN))
	sb.WriteString(NodeStringTokenSpace)
	sb.WriteString(r.Value.String())
	sb.WriteString(NodeStringSemiColon)
	return sb.String()
}

type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) DebugString() string {
	if e.Expression == nil {
		return ""
	}
	return e.Expression.DebugString()
}

func (e *ExpressionStatement) String() string {
	if e.Expression == nil {
		return ""
	}
	return e.Expression.String()
}

// expressions

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) DebugString() string {
	return i.Value
}

func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Value uint64
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) DebugString() string {
	return strconv.FormatUint(i.Value, 10)
}

func (i *IntegerLiteral) String() string {
	return strconv.FormatUint(i.Value, 10)
}

type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) expressionNode() {}

func (f *FloatLiteral) DebugString() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

func (f *FloatLiteral) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

type PrefixOperator string

const (
	PREOP_MINUS PrefixOperator = token.MINUS
	PREOP_BANG  PrefixOperator = token.BANG
)

type PrefixExpression struct {
	Operator PrefixOperator
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) DebugString() string {
	return string(p.Operator)
}
func (p *PrefixExpression) String() string {
	sb := strings.Builder{}
	sb.WriteRune('(')
	sb.WriteString(string(p.Operator))
	sb.WriteString(p.Right.String())
	sb.WriteRune(')')
	return sb.String()
}

type InfixOperator string

const (
	INOP_MINUS    InfixOperator = token.MINUS
	INOP_PLUS     InfixOperator = token.PLUS
	INOP_ASTERISK InfixOperator = token.ASTERISK
	INOP_SLASH    InfixOperator = token.SLASH
	INOP_LT       InfixOperator = token.LT
	INOP_GT       InfixOperator = token.GT
	INOP_EQ       InfixOperator = token.EQ
	INOP_NOT_EQ   InfixOperator = token.NOT_EQ
)

type InfixExpression struct {
	Operator InfixOperator
	Left     Expression
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) DebugString() string {
	return string(i.Operator)
}

func (i *InfixExpression) String() string {
	sb := strings.Builder{}
	sb.WriteRune('(')
	sb.WriteString(i.Left.String())
	sb.WriteString(string(i.Operator))
	sb.WriteString(i.Right.String())
	sb.WriteRune(')')
	return sb.String()
}
