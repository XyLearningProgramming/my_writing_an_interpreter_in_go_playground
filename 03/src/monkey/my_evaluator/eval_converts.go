package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func booleanNodeToObject(b *my_ast.Boolean) *my_object.Boolean {
	if b.Value {
		return TRUE
	}
	return FALSE
}

func booleanToIntObject(b *my_object.Boolean) *my_object.Integer {
	if b == TRUE {
		return TRUE_AS_ONE
	}
	return FALSE_AS_ZERO
}

func booleanToFloatObject(b *my_object.Boolean) *my_object.Float {
	if b == TRUE {
		return TRUE_AS_ONE_FL
	}
	return FALSE_AS_ZERO_FL
}

func integerToFloatObject(i *my_object.Integer) *my_object.Float {
	return &my_object.Float{Value: float64(i.Value)}
}

func nativeBoolToBooleanObject(nb bool) *my_object.Boolean {
	if nb {
		return TRUE
	}
	return FALSE
}
