package my_object

import (
	"fmt"
	"hash/fnv"
	"math"
	"monkey/my_ast"
	"strconv"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ          = "INT"
	UNSIGNED_INTEGER_OBJ = "UINT"
	FLOAT_OBJ            = "FLOAT"
	BOOLEAN_OBJ          = "BOOLEAN"
	NULL_OBJ             = "NULL"
	RETURN_VALUE_OBJ     = "RETURN_VALUE"
	ERROR_OBJ            = "ERROR"
	FUNCTION_OBJ         = "FUNCTION"
	STRING_OBJ           = "STRING"
	BUILTIN_OBJ          = "BUILTIN"
	ARRAY_OBJ            = "ARRAY"
	HASH_OBJ             = "HASH"
)

type Object interface {
	Type() ObjectType
	String() string
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashableObject interface {
	Object
	HashKey() HashKey
}

type UnsignedInteger struct {
	Value uint64
}

func (i *UnsignedInteger) Type() ObjectType {
	return UNSIGNED_INTEGER_OBJ
}

func (i *UnsignedInteger) String() string {
	return strconv.FormatUint(i.Value, 10)
}

func (i *UnsignedInteger) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: i.Value}
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

func (f *Float) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

func (f *Float) HashKey() HashKey {
	return HashKey{Type: f.Type(), Value: math.Float64bits(f.Value)}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) String() string {
	return strconv.FormatBool(b.Value)
}

func (b *Boolean) HashKey() HashKey {
	if b.Value {
		return HashKey{Type: b.Type(), Value: 1}
	}
	return HashKey{Type: b.Type(), Value: 0}
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) String() string { return "null" }

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

func (r *ReturnValue) String() string {
	return r.Value.String()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }

func (e *Error) String() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*my_ast.Identifier
	Body       *my_ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

func (f *Function) String() string {
	sb := &strings.Builder{}
	sb.WriteString("fn(")
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	sb.WriteString(strings.Join(params, ","))
	sb.WriteString(")")
	sb.WriteString(f.Body.String())
	return sb.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) String() string {
	// return "\"" + s.Value + "\""
	return s.Value
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

func (b *Builtin) String() string {
	return "builtin function"
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }

func (a *Array) String() string {
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.String())
	}
	return "[" + strings.Join(elements, ",") + "]"
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) String() string {
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s:%s", pair.Key.String(), pair.Value.String()))
	}
	return "{" + strings.Join(pairs, ",") + "}"
}
