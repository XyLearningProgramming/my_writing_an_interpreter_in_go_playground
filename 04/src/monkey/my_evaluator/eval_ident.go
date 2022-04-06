package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalIdentifier(node *my_ast.Identifier, env *my_object.Environment) my_object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: %s", node.Value)
	}
	return val
}
