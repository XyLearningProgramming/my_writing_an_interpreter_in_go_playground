package my_evaluator

import (
	"monkey/my_ast"
	"monkey/my_object"
	"strings"
)

func evalIndexExpression(left my_object.Object, indexNode *my_ast.IndexExpression, env *my_object.Environment) my_object.Object {
	switch left := left.(type) {
	case *my_object.String:
		elements := []my_object.Object{}
		for _, char := range left.Value {
			elements = append(elements, &my_object.String{Value: string(char)})
		}
		stringArr := &my_object.Array{Elements: elements}
		returnedStringArr := evalArrayIndexExpression(stringArr, indexNode, env)
		if isError(returnedStringArr) {
			return returnedStringArr
		}
		if returnedStringArr, rok := returnedStringArr.(*my_object.String); rok {
			return returnedStringArr
		}
		sb := &strings.Builder{}
		for _, relem := range returnedStringArr.(*my_object.Array).Elements {
			sb.WriteString(relem.(*my_object.String).Value)
		}
		return &my_object.String{Value: sb.String()}

	case *my_object.Array:
		return evalArrayIndexExpression(left, indexNode, env)
	case *my_object.Hash:
		return evalHashIndexExpression(left, indexNode.StartIndex, env)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array *my_object.Array, indexNode *my_ast.IndexExpression, env *my_object.Environment) my_object.Object {
	// shortcut: if no start or end index or stride, return error
	if !indexNode.IsSetStartIndex {
		return newError("array-like indexing with empty expression")
	}
	// parse start index
	startIdx := int64(0)
	if indexNode.StartIndex != nil {
		startIdxEvalObj := Eval(indexNode.StartIndex, env)
		startIdxEvalObjInt, iok := startIdxEvalObj.(*my_object.Integer)
		if !iok {
			return newError("array-like indexing expecting INT, but got %s", startIdxEvalObj.Type())
		}
		startIdx = startIdxEvalObjInt.Value
	}
	if startIdx >= int64(len(array.Elements)) {
		return newError("index %d out of array with length %d", startIdx, len(array.Elements))
	}
	if startIdx < 0 {
		if (-startIdx) >= int64(len(array.Elements)) {
			return newError("index %d out of array with length %d", startIdx, len(array.Elements))
		}
		startIdx = int64(len(array.Elements)) + startIdx
	}
	// shortcut: if no end index or stride, return element immediately
	if !indexNode.IsSetEndIndex && !indexNode.IsSetStride {
		return array.Elements[startIdx]
	}
	// parse end index
	endIdx := int64(len(array.Elements))
	if indexNode.EndIndex != nil {
		endIdxEvalObj := Eval(indexNode.EndIndex, env)
		endIdxEvalObjInt, eok := endIdxEvalObj.(*my_object.Integer)
		if !eok {
			return newError("array-like indexing expecting INT, but got %s", endIdxEvalObj.Type())
		}
		endIdx = endIdxEvalObjInt.Value
	}
	if endIdx >= int64(len(array.Elements)) {
		endIdx = int64(len(array.Elements))
	}
	if endIdx < 0 {
		// return empty if end index out of boundary without error
		if (-endIdx) >= int64(len(array.Elements)) {
			return EMPTY_ARRAY
		}
		endIdx = int64(len(array.Elements)) + endIdx
	}
	// parse stride
	stride := int64(1)
	if indexNode.Stride != nil {
		strideEvalObj := Eval(indexNode.Stride, env)
		strideEvalObjInt, sok := strideEvalObj.(*my_object.Integer)
		if !sok {
			return newError("array-like indexing expecting INT, but got %s", strideEvalObj.Type())
		}
		stride = strideEvalObjInt.Value
	}
	if stride == 0 {
		return newError("array-like indexing expecting non-zero stride")
	}
	// calculate
	if stride > 0 {
		results := &my_object.Array{Elements: []my_object.Object{}}
		for idx := startIdx; idx < endIdx; idx += stride {
			results.Elements = append(results.Elements, array.Elements[idx])
		}
		return results
	}
	if stride < 0 {
		if indexNode.StartIndex == nil {
			startIdx = int64(len(array.Elements)) - 1
		}
		if indexNode.EndIndex == nil {
			endIdx = -1
		}
		results := &my_object.Array{Elements: []my_object.Object{}}
		for idx := startIdx; idx > endIdx; idx += stride {
			results.Elements = append(results.Elements, array.Elements[idx])
		}
		return results
	}
	panic("unreachable path")
}

func evalHashIndexExpression(hash *my_object.Hash, index my_ast.Expression, env *my_object.Environment) my_object.Object {
	indexObj := Eval(index, env)
	if isError(indexObj) {
		return indexObj
	}
	key, ok := indexObj.(my_object.HashableObject)
	if !ok {
		return newError("key type not hashable: %s", indexObj.Type())
	}
	pair, ok := hash.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}
