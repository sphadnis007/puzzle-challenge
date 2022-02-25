package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sphadnis007/puzzle/internals/utils"
)

func TestPuzzleStoreAdd(t *testing.T) {
	ps := NewPuzzleStore()

	// currently no puzzles are present in puzzle_store
	s := ps.GetAllPuzzles()
	assert.Equal(t, []string{}, s)

	ps.AddPuzzle(utils.P1, "p1")

	// check if puzzle present
	p, err := ps.GetPuzzleByName("p1")
	assert.Nil(t, err)
	assert.NotNil(t, p)

	ps.AddPuzzle(utils.P2, "p2")

	// get all puzzle names
	s = ps.GetAllPuzzles()
	assert.Equal(t, []string{"p1", "p2"}, s)
}

func TestGetPuzzleByNameError(t *testing.T) {
	ps := NewPuzzleStore()
	p, err := ps.GetPuzzleByName("p1")
	assert.Error(t, err)
	assert.Nil(t, p)
}
