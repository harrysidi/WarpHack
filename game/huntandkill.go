package game

import (
	"fmt"
	"math/rand/v2"
)

type cell struct {
	x,y int
}

// HuntAndKill implements the Hunt and Kill maze generation algorithm.
func HuntAndKill(dungeonSeed uint64, width int, height int) *DungeonLevel {

	maze := [][]bool{}
	for y := range height {
		row := []cell{}
		for x := range width {
			row = append(row, false)
		}
		maze = append(maze, row)
	}
	fmt.Print("Generating maze with seed: ", dungeonSeed, "\n")
	fmt.Println(maze)

	rng := rand.New(rand.NewPCG(dungeonSeed, uint64(dungeonSeed/2+1)))
	// Start with a random cell
	startX := rng.IntN(width)
	startY := rng.IntN(height)
	maze[startY][startX]= true
	currentX := startX
	currentY := startY
	//kill phase
	for {
		unvisited := []cell{}
		wallToRemove := []cell{}
		if x+2 < width && !maze[y][x+2] {
			unvisited = append(unvisited, cell{y: y, x: x + 2})
			wallToRemove = append(wallToRemove, cell{y: y, x: x + 1})
		}
		if x-2 >= 0 && !maze[y][x-2]{
			unvisited = append(unvisited, cell{y: y, x: x - 2})
			wallToRemove = append(wallToRemove, cell{y: y, x: x - 1})
		}
		if y+2 < height && !maze[y+2][x] {
			unvisited = append(unvisited, cell{y: y + 2, x: x})
			wallToRemove = append(wallToRemove, cell{y: y + 1, x: x})
		
		}
		if y-2 >= 0 && !maze[y-2][x] {
			unvisited = append(unvisited, cell{y: y - 2, x: x})
			wallToRemove = append(wallToRemove, cell{y: y - 1, x: x})
		}
		
		if len(unvisited) >0 {
			index:= rng.IntN(len(unvisited))
			remove_wall := unvisited[rng.IntN(len(unvisited))]

			DungeonLevel[]
x
	}

}
