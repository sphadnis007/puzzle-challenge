package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sphadnis007/puzzle/internals/utils"
)

func TestGetCoordinateCopy(t *testing.T) {
	c := NewCoordinates(1, 1)
	n := GetCoordinateCopy(c)

	assert.Equal(t, c.X, n.X)
	assert.Equal(t, c.Y, n.Y)
}

func TestCoordinateUpdate(t *testing.T) {
	c := NewCoordinates(1, 1)

	c.Update(utils.HorizontalRight)
	assert.Equal(t, 1, c.X)
	assert.Equal(t, 2, c.Y)

	c.Update(utils.VerticalDown)
	assert.Equal(t, 2, c.X)
	assert.Equal(t, 2, c.Y)
}

func TestCheckIfInPuzzleBoundry(t *testing.T) {
	c := NewCoordinates(1, 1)
	p := NewPuzzle(utils.P1, "p1")

	cross := c.CheckIfInPuzzleBoundry(p)
	assert.True(t, cross)

	c.Update([]int{10, 10})
	cross = c.CheckIfInPuzzleBoundry(p)
	assert.False(t, cross)
}
