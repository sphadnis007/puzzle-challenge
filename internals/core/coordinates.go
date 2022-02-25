package core

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(row, col int) *Coordinates {
	return &Coordinates{X: row, Y: col}
}

func GetCoordinateCopy(c *Coordinates) *Coordinates {
	return &Coordinates{X: c.X, Y: c.Y}
}

func (c *Coordinates) CheckIfInPuzzleBoundry(p IPuzzle) bool {
	if c.X >= 0 && c.X < p.GetRows() && c.Y >= 0 && c.Y < p.GetCols() {
		return true
	}
	return false
}

func (c *Coordinates) Update(u []int) {
	c.X += u[0]
	c.Y += u[1]
}
