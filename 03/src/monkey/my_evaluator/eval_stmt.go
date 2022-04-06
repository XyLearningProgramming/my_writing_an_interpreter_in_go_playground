package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalProgram(stmts []my_ast.Statement, env *my_object.Environment) my_object.Object {
	var result my_object.Object
	for _, stmt := range stmts {
		result = Eval(stmt, env)
		switch result := result.(type) {
		case *my_object.ReturnValue:
			return result.Value
		case *my_object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(stmts []my_ast.Statement, env *my_object.Environment) my_object.Object {
	var result my_object.Object
	for _, stmt := range stmts {
		result = Eval(stmt, env)
		// to keep track of return value with its type in the block statement
		if result != nil {
			if rt := result.Type(); rt == my_object.ERROR_OBJ || rt == my_object.RETURN_VALUE_OBJ {
				return result
			}
		}
	}
	return result
}
