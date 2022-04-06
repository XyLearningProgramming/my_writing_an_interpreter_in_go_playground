package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalPrefixNode(node *my_ast.PrefixExpression, env *my_object.Environment) my_object.Object {
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	switch node.Operator {
	case my_ast.PREOP_BANG:
		return evalPrefixOperatorBang(right)
	case my_ast.PREOP_MINUS:
		return evalPrefixOperatorMinus(right)
	}
	return newError("unknown operator: %s%s", node.Operator, right.Type())
}

func evalPrefixOperatorBang(right my_object.Object) my_object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalPrefixOperatorMinus(right my_object.Object) my_object.Object {
	switch right := right.(type) {
	case *my_object.Float:
		return &my_object.Float{Value: -right.Value}
	case *my_object.Integer:
		return &my_object.Integer{Value: -right.Value}
	// case *my_object.UnsignedInteger:
	// 	if right.Value > math.MaxInt64 {
	// 		// NOTE: Cannot convert correctly!
	// 		return NULL
	// 	}
	// 	return &my_object.Integer{Value: -int64(right.Value)}
	default:
		switch right {
		case FALSE:
			return FALSE
		case TRUE:
			return &my_object.Integer{Value: -1}
		default:
			return newError("unknown operator: %s%s", my_ast.PREOP_MINUS, right.Type())
		}
	}
}
