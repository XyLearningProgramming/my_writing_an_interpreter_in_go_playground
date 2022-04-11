package my_evaluator

import (
	"fmt"
	"monkey/my_object"
)

var builtins = map[string]*my_object.Builtin{
	"len": {
		Fn: func(args ...my_object.Object) my_object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments: got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *my_object.String:
				return &my_object.Integer{Value: int64(len(arg.Value))}
			case *my_object.Array:
				return &my_object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to len not supported: got %s", arg.Type())
			}
		},
	},
	"append": {
		Fn: func(args ...my_object.Object) my_object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments: got=%d, want=2", len(args))
			}
			if args[0].Type() != my_object.ARRAY_OBJ {
				return newError("first argument to `append` must be ARRAY: got=%s", args[0].Type())
			}
			newElements := make([]my_object.Object, len(args)+1, len(args)+1)
			copy(newElements, args[0].(*my_object.Array).Elements)
			newElements[len(args)] = args[1]
			return &my_object.Array{Elements: newElements}
		},
	},
	"put": {
		Fn: func(args ...my_object.Object) my_object.Object {
			fmt.Println()
			for _, arg := range args {
				fmt.Print(arg.String())
			}
			return nil
		},
	},
}
