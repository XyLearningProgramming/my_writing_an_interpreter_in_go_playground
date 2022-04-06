package my_ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	prog := &Program{
		Statements: []Statement{
			&LetStatement{
				Ident: &Identifier{
					Value: "x",
				},
				Value: &Identifier{
					Value: "y",
				},
			},
		},
	}
	assert.Equal(t, "let x = y;", prog.String())
}
