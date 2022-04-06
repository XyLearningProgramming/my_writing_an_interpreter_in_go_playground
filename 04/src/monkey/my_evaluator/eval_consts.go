package my_evaluator

import (
	"monkey/my_object"
)

var (
	TRUE             = &my_object.Boolean{Value: true}
	FALSE            = &my_object.Boolean{Value: false}
	TRUE_AS_ONE      = &my_object.Integer{Value: 1}
	FALSE_AS_ZERO    = &my_object.Integer{Value: 0}
	TRUE_AS_ONE_FL   = &my_object.Float{Value: 1}
	FALSE_AS_ZERO_FL = &my_object.Float{Value: 0}
	NULL             = &my_object.Null{}
)
