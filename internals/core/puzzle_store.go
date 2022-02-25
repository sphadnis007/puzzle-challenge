package core

import "errors"

// In-memory puzzle store
// Future enhancement - we can use database to store puzzles

type IPuzzleStore interface {
	AddPuzzle(arr [][]byte, name string)
	GetPuzzleByName(s string) (IPuzzle, error)
	GetAllPuzzles() []string
}

// creates empty puzzle store
func NewPuzzleStore() IPuzzleStore {
	return &PuzzleStore{Num: 0, DB: []IPuzzle{}}
}

type PuzzleStore struct {
	Num int       // total number of puzzles stored
	DB  []IPuzzle // puzzle array
}

func (ps *PuzzleStore) AddPuzzle(arr [][]byte, name string) {
	ps.DB = append(ps.DB, NewPuzzle(arr, name))
}

func (ps *PuzzleStore) GetPuzzleByName(s string) (IPuzzle, error) {
	for _, p := range ps.DB {
		if p.GetName() == s {
			return p, nil
		}
	}
	return nil, errors.New("not found")
}

func (ps *PuzzleStore) GetAllPuzzles() []string {
	s := make([]string, 0)
	for _, p := range ps.DB {
		s = append(s, p.GetName())
	}
	return s
}
