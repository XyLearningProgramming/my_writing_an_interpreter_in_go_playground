package my_evaluator

import (
	"math"
	lexer "monkey/my_lexer"
	"monkey/my_object"
	"monkey/my_parser"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  string
	expect interface{}
}

type testCaseTyped struct {
	input   string
	expect  interface{}
	refType reflect.Type
}

func TestEvalInteger(t *testing.T) {
	cases := []testCase{
		{"5", 5},
		{"10", 10},
	}
	for _, c := range cases {
		evaluated := testEval(t, c.input)
		assert.EqualValues(t, c.expect, evaluated.(*my_object.Integer).Value)
	}
}

func TestEvalBoolean(t *testing.T) {
	cases := []testCase{
		{"true", true},
		{"false", false},
	}
	for _, c := range cases {
		evaluated := testEval(t, c.input)
		assert.EqualValues(t, c.expect, evaluated.(*my_object.Boolean).Value)
	}
}

func TestPrefixBangOperator(t *testing.T) {
	cases := []testCase{
		{"!true", false},
		{"!5", false},
		{"!!a", true},
	}
	for _, c := range cases {
		evaluated := testEval(t, c.input)
		assert.EqualValues(t, c.expect, evaluated.(*my_object.Boolean).Value)
	}
}

func TestPrefixMinusOperator(t *testing.T) {
	cases := []testCase{
		{"-5", "-5"},
		{"5", "5"},
		{"-" + strconv.FormatUint(math.MaxUint64, 10), "1"},
		{"-" + strconv.FormatUint(math.MaxInt64, 10), "-" + strconv.FormatUint(math.MaxInt64, 10)},
		{"--" + strconv.FormatUint(math.MaxInt64, 10), strconv.FormatUint(math.MaxInt64, 10)},
		{strconv.FormatFloat(1.234, 'f', -1, 64), "1.234"},
		{strconv.FormatFloat(-1.234, 'f', -1, 64), "-1.234"},
	}
	for _, c := range cases {
		evaluated := testEval(t, c.input)
		assert.NotNil(t, evaluated, "case is: %+v", c)
		assert.EqualValues(t, c.expect, evaluated.String())
	}
}

var (
	mf        *my_object.Float
	mi        *my_object.Integer
	mb        *my_object.Boolean
	mn        *my_object.Null
	me        *my_object.Error
	floatType = reflect.TypeOf(mf)
	intType   = reflect.TypeOf(mi)
	boolType  = reflect.TypeOf(mb)
	nullType  = reflect.TypeOf(mn)
	errType   = reflect.TypeOf(me)
)

func TestInfixOperator(t *testing.T) {
	tests := []*testCaseTyped{
		{"1", 1, intType},
		{"2 * 2 * 2 * 2 * 2", 32, intType},
		{"-50 + 100 + -50", 0, intType},
		{"5 * 2 + 10", 20, intType},
		{"5 + 2 * 10", 25, intType},
		{"20 + 2 * -10", 0, intType},
		{"50 / 2 * 2 + 10", 60, intType},
		{"2 * (5 + 10)", 30, intType},
		{"3 * 3 * 3 + 10", 37, intType},
		{"3 * (3 * 3) + 10", 37, intType},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50, intType},
		{"5*5.0+1", 26, floatType},
		{"1.5+2-1", 2.5, floatType},
		{"true != \n\tfalse", true, boolType},
	}
	testCaseWithStruct(t, tests)
}

func TestIfElseStatements(t *testing.T) {
	tests := []*testCaseTyped{
		{"if (true) { 10 }", 10, intType},
		{"if (false) { 10 }", nil, nullType},
		{"if (1) { 10 }", 10, intType},
		{"if (1 < 2) { 10 }", 10, intType},
		{"if (1 > 2) { 10 }", nil, nullType},
		{"if (1 > 2) { 10 } else { 20 }", 20, intType},
		{"if (1 < 2) { 10 } else { 20 }", 10, intType},
	}
	testCaseWithStruct(t, tests)
}

func TestReturnStatements(t *testing.T) {
	tests := []*testCaseTyped{
		{"return 10", 10, intType},
		{"return 2*5; 9", 10, intType},
		{"9; return 2*5; 9;", 10, intType},
		{"if(true){if(10>1){return 10;} return 1;}", 10, intType},
	}
	testCaseWithStruct(t, tests)
}

func TestErrorHandling(t *testing.T) {
	tests := []*testCaseTyped{
		{"return -if(false){10}", "unknown operator: -NULL", errType},
		{"return NULL", "identifier not found: NULL", errType},
	}
	testCaseWithStruct(t, tests)
}

func TestLetStatements(t *testing.T) {
	tests := []*testCaseTyped{
		{"let a = 5; a;", 5, intType},
		{"let a = 5*5; 5*a", 5 * 5 * 5, intType},
		{"let a = 5; let b = a +5; b;", 10, intType},
	}
	testCaseWithStruct(t, tests)
}

func TestEvalFunctionObject(t *testing.T) {
	input := "fn(x){x+2};"
	evaluated := testEval(t, input)
	fn, ok := evaluated.(*my_object.Function)
	assert.True(t, ok)
	assert.Equal(t, 1, len(fn.Parameters))
	assert.Equal(t, "x", fn.Parameters[0].String())
	assert.Equal(t, "{(x+2);}", fn.Body.String())
}

func TestFunctionEvaluation(t *testing.T) {
	tests := []*testCaseTyped{
		{"let identity = fn(x) { x; }; identity(5);", 5, intType},
		{"let identity = fn(x) { return x; }; identity(5);", 5, intType},
		{"let double = fn(x) { x * 2; }; double(5);", 10, intType},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10, intType},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20, intType},
		{"fn(x) { x; }(5)", 5, intType},
	}
	testCaseWithStruct(t, tests)
}

func testCaseWithStruct(t *testing.T, tests []*testCaseTyped) {
	for _, c := range tests {
		evaluated := testEval(t, c.input)
		assert.NotNil(t, evaluated, "case is: %+v", c)
		ev := reflect.ValueOf(evaluated)
		assert.EqualValues(t, c.refType, ev.Type())
		switch c.refType {
		case floatType:
			assert.EqualValues(t, c.expect, ev.Elem().Field(0).Float())
		case intType:
			assert.EqualValues(t, c.expect, ev.Elem().Field(0).Int())
		case boolType:
			assert.EqualValues(t, c.expect, ev.Elem().Field(0).Bool())
		case nullType:
			assert.EqualValues(t, 0, ev.Elem().NumField())
		case errType:
			assert.EqualValues(t, c.expect, ev.Elem().Field(0).String())
		default:
			assert.Fail(t, "unknown ev: %T: %v: c: %+v", ev, ev, c)
		}
	}
}

func testEval(t *testing.T, input string) my_object.Object {
	l := lexer.New(input)
	p := my_parser.New(l)
	prog := p.Parse()
	assert.NoError(t, p.Error())
	// fmt.Println(prog.String())
	return Eval(prog, my_object.NewEnvironment())
}
