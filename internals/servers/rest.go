package servers

import (
	"log"

	"github.com/gin-gonic/gin"

	"sphadnis007/puzzle/internals/core"
	"sphadnis007/puzzle/internals/handlers"
)

func StartNewRestServer(ps core.IPuzzleStore) {
	e := gin.Default()
	h := handlers.NewRestHandlers(ps)

	// get list of puzzles
	e.GET("/api/v1/puzzles", h.GetPuzzleNames)

	// find string "fname" in puzzle "pname"
	e.GET("/api/v1/puzzles/:pname/:fname", h.FindInPuzzle)

	// display contents of puzzle with name "pname"
	e.GET("/api/v1/puzzles/:pname", h.DisplayPuzzle)

	// Todo: POST call to add puzzle

	if err := e.Run(":8080"); err != nil {
		log.Println("Failed to start REST server")
	}
}
