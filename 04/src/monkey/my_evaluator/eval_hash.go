package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalHashExpression(node *my_ast.HashExpression, env *my_object.Environment) my_object.Object {
	pairs := make(map[my_object.HashKey]my_object.HashPair)
	for kn, vn := range node.Pairs {
		if ksn, kok := kn.(*my_ast.Identifier); kok {
			kn = &my_ast.StringExpression{Value: ksn.Value}
		}
		key := Eval(kn, env)
		hashableKey, hok := key.(my_object.HashableObject)
		if !hok {
			return newError("key type not hashable: %s", key.Type())
		}
		value := Eval(vn, env)
		pairs[hashableKey.HashKey()] = my_object.HashPair{Key: key, Value: value}
	}
	return &my_object.Hash{Pairs: pairs}
}
