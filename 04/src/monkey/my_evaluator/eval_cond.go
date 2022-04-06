package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalIfExpression(ie *my_ast.IfExpression, env *my_object.Environment) my_object.Object {
	cond := Eval(ie.Condition, env)
	if isError(cond) {
		return cond
	}
	if isTruthy(cond) {
		return Eval(ie.Consequence, env)
	}
	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}
	return NULL
}

func isTruthy(obj my_object.Object) bool {
	switch obj {
	case NULL:
		fallthrough
	case FALSE:
		return false
	default:
		return true
	}
}
