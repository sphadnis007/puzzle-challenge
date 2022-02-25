package handlers

import (
	"log"

	"github.com/gin-gonic/gin"

	"sphadnis007/puzzle/internals/core"
	"sphadnis007/puzzle/internals/services"
)

type IRestHandler interface {
	GetPuzzleNames(c *gin.Context)
	FindInPuzzle(c *gin.Context)
	DisplayPuzzle(c *gin.Context)
}

func NewRestHandlers(ps core.IPuzzleStore) IRestHandler {
	return &RestHandler{PS: ps}
}

type RestHandler struct {
	PS core.IPuzzleStore
}

func (rs *RestHandler) GetPuzzleNames(c *gin.Context) {
	c.JSON(200, rs.PS.GetAllPuzzles())
}

func (rs *RestHandler) FindInPuzzle(c *gin.Context) {
	fname := c.Param("fname")
	pname := c.Param("pname")

	if pname == "" || fname == "" {
		c.JSON(400, "Puzzle name OR string to find is empty")
		return
	}

	log.Println("Received request to find", fname, "int puzzle", pname)

	p, err := rs.PS.GetPuzzleByName(pname)
	if err != nil {
		c.JSON(400, "Puzzle Does't exist!")
		return
	}

	s := services.NewSearchPuzzle(p, fname)
	// returns coordinate, direction, no. of rotates and error
	coordinates, direction, n, err := s.FindStringInRotatedPuzzle()
	if err != nil {
		c.JSON(404, "Word is not present in the puzzle!")
		return
	}

	response := map[string]interface{}{
		"Coordinates": map[string]interface{}{
			"Row":    coordinates.X,
			"Column": coordinates.Y,
		},
		"Direction":           direction,
		"Number of Rotations": n,
	}

	c.JSON(200, response)
}

func (rs *RestHandler) DisplayPuzzle(c *gin.Context) {
	pname := c.Param("pname")

	if pname == "" {
		c.JSON(400, "Puzzle name is empty")
		return
	}

	p, err := rs.PS.GetPuzzleByName(pname)
	if err != nil {
		c.JSON(400, "Puzzle Does't exist!")
		return
	}

	result := p.DisplayPuzzle()
	c.JSON(200, result)
}
