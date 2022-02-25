package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sphadnis007/puzzle/internals/core"
	"sphadnis007/puzzle/internals/utils"
)

func TestFindStringInRotatedPuzzleFound(t *testing.T) {
	// found in 1 rotate scenario
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P2, "p2"), "abc")
	c, direction, num, err := sp.FindStringInRotatedPuzzle()

	assert.NoError(t, err)
	assert.Equal(t, 1, num)
	assert.Equal(t, utils.Vertical, direction)
	assert.Equal(t, c.X, 0)
	assert.Equal(t, c.Y, 2)
}

func TestFindStringInRotatedPuzzleNotFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P2, "p2"), "axc")
	c, direction, num, err := sp.FindStringInRotatedPuzzle()

	assert.Error(t, err)
	assert.Equal(t, -1, num)
	assert.Equal(t, "", direction)
	assert.Nil(t, c)
}

func TestFindStringInPuzzleFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P1, "p1"), "xob")
	c, direction, found := sp.FindStringInPuzzle()

	assert.True(t, found)
	assert.Equal(t, c.X, 1)
	assert.Equal(t, c.Y, 0)
	assert.Equal(t, utils.Horizontal, direction)
}

func TestFindStringInPuzzleNotFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P1, "p1"), "xab")
	c, direction, found := sp.FindStringInPuzzle()

	assert.False(t, found)
	assert.Nil(t, c)
	assert.Equal(t, "", direction)
}

func TestFindStringFromCoordinatesFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P1, "p1"), "abc")
	found, direction := sp.FindStringFromCoordinates(core.NewCoordinates(0, 2))

	assert.True(t, found)
	assert.Equal(t, utils.Vertical, direction)
}

func TestFindStringFromCoordinatesNotFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P2, "p2"), "abc")
	found, direction := sp.FindStringFromCoordinates(core.NewCoordinates(0, 2))

	assert.False(t, found)
	assert.Equal(t, "", direction)
}

func TestFindInDirectionFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P1, "p1"), "abc")
	found := sp.FindInDirection(core.NewCoordinates(0, 2), utils.VerticalDown)

	assert.True(t, found)
}

func TestFindInDirectionNotFound(t *testing.T) {
	sp := NewSearchPuzzle(core.NewPuzzle(utils.P2, "p2"), "abc")
	found := sp.FindInDirection(core.NewCoordinates(0, 2), utils.VerticalDown)

	assert.False(t, found)
}
