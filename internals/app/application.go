package app

import (
	"sphadnis007/puzzle/internals/core"
	"sphadnis007/puzzle/internals/servers"
	"sphadnis007/puzzle/internals/utils"
)

func StartApplication() {
	ps := core.NewPuzzleStore()
	ps.AddPuzzle(utils.P1, "p1")
	ps.AddPuzzle(utils.P2, "p2")
	ps.AddPuzzle(utils.P3, "p3")
	ps.AddPuzzle(utils.P4, "p4")
	ps.AddPuzzle(utils.P5, "p5")
	ps.AddPuzzle(utils.P6, "p6")

	servers.StartNewRestServer(ps)

	select {}
}
