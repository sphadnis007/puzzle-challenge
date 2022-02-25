package core

import (
	"log"
)

type IPuzzle interface {
	DisplayPuzzle() []string
	GetRows() int
	GetCols() int
	GetName() string
	GetContents(row, col int) byte
}

func NewPuzzle(arr [][]byte, s string) IPuzzle {
	return &Puzzle{Rows: len(arr), Cols: len(arr[0]), Array: arr, Name: s}
}

type Puzzle struct {
	Name  string
	Rows  int
	Cols  int
	Array [][]byte
}

func (p *Puzzle) DisplayPuzzle() []string {
	var res []string
	var spaceByte string = " "

	for i := 0; i < p.Rows; i++ {
		s := ""
		for j := 0; j < p.Cols; j++ {
			s += string(p.GetContents(i, j)) + spaceByte
		}
		res = append(res, s)
	}

	return res
}

func (p *Puzzle) RotatePuzzle() {
	log.Println("Implement me!")
}

func (p *Puzzle) GetRows() int {
	return p.Rows
}

func (p *Puzzle) GetCols() int {
	return p.Cols
}

func (p *Puzzle) GetName() string {
	return p.Name
}

func (p *Puzzle) GetContents(row, col int) byte {
	return p.Array[row][col]
}
