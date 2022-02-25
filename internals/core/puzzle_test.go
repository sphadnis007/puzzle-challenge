package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sphadnis007/puzzle/internals/utils"
)

func TestPuzzleFunctions(t *testing.T) {
	p := NewPuzzle(utils.P1, "p1")

	n := p.GetRows()
	assert.Equal(t, 3, n)

	n = p.GetCols()
	assert.Equal(t, 3, n)

	var output byte = 'a'
	b := p.GetContents(0, 2)
	assert.Equal(t, output, b)

	name := p.GetName()
	assert.Equal(t, "p1", name)

	s := p.DisplayPuzzle()
	assert.NotNil(t, s)
}
