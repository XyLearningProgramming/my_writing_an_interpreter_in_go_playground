package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalInfixNode(node *my_ast.InfixExpression, env *my_object.Environment) my_object.Object {
	leftObj := Eval(node.Left, env)
	if isError(leftObj) {
		return leftObj
	}
	rightObj := Eval(node.Right, env)
	if isError(rightObj) {
		return rightObj
	}
	switch leftObj := leftObj.(type) {
	case *my_object.Integer:
		switch rightObj := rightObj.(type) {
		case *my_object.Integer:
			return evalIntegerInfixExpression(node.Operator, leftObj, rightObj)
		case *my_object.Boolean:
			return evalIntegerInfixExpression(node.Operator, leftObj, booleanToIntObject(rightObj))
		case *my_object.Float:
			return evalFloatInfixExpression(node.Operator, integerToFloatObject(leftObj), rightObj)
		case *my_object.Null:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		default:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		}
	case *my_object.Boolean:
		switch rightObj := rightObj.(type) {
		case *my_object.Integer:
			return evalIntegerInfixExpression(node.Operator, booleanToIntObject(leftObj), rightObj)
		case *my_object.Boolean:
			return evalIntegerInfixExpression(node.Operator, booleanToIntObject(leftObj), booleanToIntObject(rightObj))
		case *my_object.Float:
			return evalFloatInfixExpression(node.Operator, booleanToFloatObject(leftObj), rightObj)
		case *my_object.Null:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		default:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		}
	case *my_object.Float:
		switch rightObj := rightObj.(type) {
		case *my_object.Integer:
			return evalFloatInfixExpression(node.Operator, leftObj, integerToFloatObject(rightObj))
		case *my_object.Boolean:
			return evalFloatInfixExpression(node.Operator, leftObj, booleanToFloatObject(rightObj))
		case *my_object.Float:
			return evalFloatInfixExpression(node.Operator, leftObj, rightObj)
		case *my_object.Null:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		default:
			return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
		}
	case *my_object.Null:
		// TODO: NULL==NULL? NULL>=1 yields false or NULL?
		if _, ok := rightObj.(*my_object.Null); ok {
			switch node.Operator {
			case "<":
				fallthrough
			case "!=":
				fallthrough
			case ">":
				return FALSE
			case "==":
				return TRUE
			default:
				return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
			}
		}
		// an error?
		return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
	default:
		return newError("unknown operator: %s%s%s", leftObj.Type(), node.Operator, rightObj.Type())
	}
}

func evalIntegerInfixExpression(
	operator my_ast.InfixOperator, left, right *my_object.Integer,
) my_object.Object {
	leftVal := left.Value
	rightVal := right.Value
	switch operator {
	case "+":
		return &my_object.Integer{Value: leftVal + rightVal}
	case "-":
		return &my_object.Integer{Value: leftVal - rightVal}
	case "*":
		return &my_object.Integer{Value: leftVal * rightVal}
	case "/":
		return &my_object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s%s%s", left.Type(), operator, right.Type())
	}
}

func evalFloatInfixExpression(
	operator my_ast.InfixOperator, left, right *my_object.Float,
) my_object.Object {
	leftVal := left.Value
	rightVal := right.Value
	switch operator {
	case "+":
		return &my_object.Float{Value: leftVal + rightVal}
	case "-":
		return &my_object.Float{Value: leftVal - rightVal}
	case "*":
		return &my_object.Float{Value: leftVal * rightVal}
	case "/":
		return &my_object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s%s%s", left.Type(), operator, right.Type())
	}
}
