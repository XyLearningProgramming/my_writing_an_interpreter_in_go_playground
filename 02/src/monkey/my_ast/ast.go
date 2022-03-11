package my_ast

type Node interface {
	// DebugString: for debugging purposes;
	// if an expression, ususally literal value of the first token;
	// if an statement, usually call DebugString() on its first expression;
	// if an ident, literal value of its `Value` field;
	DebugString() string
}

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

// identifier

type Identifier struct {
	Value string
}

func (i *Identifier) DebugString() string {
	return i.Value
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

type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) DebugString() string {
	return r.Value.DebugString()
}

// expressions
