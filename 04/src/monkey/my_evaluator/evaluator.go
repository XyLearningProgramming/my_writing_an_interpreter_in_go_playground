package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func Eval(node my_ast.Node, env *my_object.Environment) my_object.Object {
	switch node := node.(type) {
	case *my_ast.Program:
		return evalProgram(node.Statements, env)
	case *my_ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)
	case *my_ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *my_ast.ReturnStatement:
		return &my_object.ReturnValue{
			Value: Eval(node.Value, env),
		}
	case *my_ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Ident.Value, val)
	case *my_ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		functionObj, ok := function.(*my_object.Function)
		if !ok {
			return newError("not a function: %s", function.Type())
		}
		args := evalArguments(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		// extend env var now to create new set of bindings
		return evalFunction(functionObj, args)

	case *my_ast.Function:
		return &my_object.Function{Parameters: node.Parameters, Env: env, Body: node.Body}
	case *my_ast.IfExpression:
		return evalIfExpression(node, env)
	case *my_ast.PrefixExpression:
		return evalPrefixNode(node, env)
	case *my_ast.InfixExpression:
		return evalInfixNode(node, env)
	case *my_ast.Identifier:
		return evalIdentifier(node, env)
	case *my_ast.Boolean:
		return booleanNodeToObject(node)
	case *my_ast.Integer:
		// NOTE: Risk of overflow here
		// but, introducing an unsigned int will quickly get messy
		// when applying infix calculation involving overflow
		return &my_object.Integer{Value: int64(node.Value)}
	case *my_ast.Float:
		return &my_object.Float{Value: node.Value}
	}
	return nil
}
