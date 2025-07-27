package game

import (
	"fmt"
	"math/rand/v2"
)

type cell struct {
	x, y    int
	visited bool
}

// HuntAndKill implements the Hunt and Kill maze generation algorithm.
func HuntAndKill(dungeonSeed uint64, width int, height int) *DungeonLevel {

	maze := [][]cell{}
	for y := range height {
		row := []cell{}
		for x := range width {
			row = append(row, cell{x: x, y: y, visited: false})
		}
		maze = append(maze, row)
	}
	fmt.Print("Generating maze with seed: ", dungeonSeed, "\n")
	rng := rand.New(rand.NewPCG(dungeonSeed, uint64(dungeonSeed/2+1)))
	// Start with a random cell
	startX := rng.IntN(width)
	startY := rng.IntN(height)
	maze[startY][startX].visited = true

	//kill phase

}
