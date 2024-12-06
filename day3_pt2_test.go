package aoc

import (
	"testing"

	"github.com/matryer/is"
)

func TestLexMul(t *testing.T) {
	is := is.New(t)
	l := newLexer("mul(253,978)mul(723,145)")
	go l.run()

	items := []item{}
	for i := range l.items {
		items = append(items, i)
	}

	expected := []item{
		{
			typ:   itemMul,
			value: "mul(253,978)",
		},
		{
			typ:   itemMul,
			value: "mul(723,145)",
		},
	}

	is.Equal(expected, items)
}
