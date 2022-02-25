package services

import (
	"errors"
	"log"
	"sphadnis007/puzzle/internals/core"
	"sphadnis007/puzzle/internals/utils"
)

type ISearchPuzzle interface {
	FindStringInRotatedPuzzle() (*core.Coordinates, string, int, error)
	FindStringInPuzzle() (*core.Coordinates, string, bool)
	FindStringFromCoordinates(c *core.Coordinates) (bool, string)
	FindInDirection(c *core.Coordinates, direction []int) bool
}

func NewSearchPuzzle(pz core.IPuzzle, str string) ISearchPuzzle {
	log.Println("Creating search puzzle", str)
	return &SearchPuzzle{p: pz, s: []byte(str)}
}

type SearchPuzzle struct {
	p core.IPuzzle // puzzle source
	s []byte       // string to search
}

func (sp *SearchPuzzle) GetPuzzle() core.IPuzzle {
	return sp.p
}

func (sp *SearchPuzzle) GetStr() []byte {
	return sp.s
}

func (sp *SearchPuzzle) FindStringInRotatedPuzzle() (*core.Coordinates, string, int, error) {
	if c, direction, found := sp.FindStringInPuzzle(); found {
		return c, direction, 0, nil
	}

	log.Println("Going to rotate string now")
	oldStr := sp.GetStr()
	// we can rotate string maximum of 25 times, 26th time onwards, same pattern will repeat
	// instead of rotating the whole puzzle, we are rotating the string and searching it
	for i := -1; i > -26; i-- {
		sp.s = utils.Rotate1DArray(oldStr, i)
		//log.Println(string(sp.s))
		if c, direction, found := sp.FindStringInPuzzle(); found {
			return c, direction, -1 * i, nil
		}
	}

	return nil, "", -1, errors.New("not found")
}

func (sp *SearchPuzzle) FindStringInPuzzle() (*core.Coordinates, string, bool) {
	p := sp.GetPuzzle()
	for row := 0; row < p.GetRows(); row++ {
		for col := 0; col < p.GetCols(); col++ {
			// if first byte/char matches
			//	- we search for the complete string in puzzle
			//log.Println("Found starting char", row, col)
			if p.GetContents(row, col) == sp.s[0] {
				c := core.NewCoordinates(row, col)
				if found, direction := sp.FindStringFromCoordinates(c); found {
					//log.Println("Found at", row, col)
					return c, direction, true
				}
			}
		}
	}

	return nil, "", false
}

func (sp *SearchPuzzle) FindStringFromCoordinates(c *core.Coordinates) (bool, string) {
	// sending coordinate copy as we update the coordinates while searching
	if sp.FindInDirection(core.GetCoordinateCopy(c), utils.HorizontalRight) {
		return true, utils.Horizontal
	}
	if sp.FindInDirection(core.GetCoordinateCopy(c), utils.VerticalDown) {
		return true, utils.Vertical
	}

	return false, ""
}

func (sp *SearchPuzzle) FindInDirection(c *core.Coordinates, direction []int) bool {
	p := sp.GetPuzzle()
	str := sp.s
	// string iterator starts from 1 as 0th element is already matched
	strItr := 1
	strSize := len(str)
	// updating the coordinates in a specific direction
	c.Update(direction)

	//log.Println("out: Coordinates", c, "strItr", strItr)
	for c.CheckIfInPuzzleBoundry(p) && strItr < strSize {
		//log.Println("in: Coordinates", c, "strItr", strItr)
		if p.GetContents(c.X, c.Y) == str[strItr] {
			strItr++
			c.Update(direction)
		} else {
			return false
		}
	}

	return strItr == strSize
}
