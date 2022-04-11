package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
)

func evalExpressions(exps []my_ast.Expression, env *my_object.Environment) []my_object.Object {
	args := []my_object.Object{}
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []my_object.Object{evaluated}
		}
		args = append(args, evaluated)
	}
	return args
}

func evalFunction(fn *my_object.Function, args []my_object.Object) my_object.Object {
	env := my_object.NewEnclosedEnvironment(fn.Env)
	for idx, param := range fn.Parameters {
		env.Set(param.Value, args[idx])
	}
	return tryUnwrapReturnValue(Eval(fn.Body, env))
}
