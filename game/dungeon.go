package game

import (
	"fmt"
	"math/rand/v2"
)

type DungeonLevel struct {
	Width  int
	Height int
	Cells  [][]uint16
	seed   uint64
}

func GenerateLevel(dungeonSeed uint64, width int, height int) *DungeonLevel {

	fmt.Print("Generating dungeon with seed: ", dungeonSeed, "\n")
	level := DungeonLevel{
		Width:  width,
		Height: height,
		Cells:  make([][]uint16, 64),
		seed:   dungeonSeed,
	}
	for i := range level.Cells {
		level.Cells[i] = make([]uint16, 64)
	}
	for y := 0; y < level.Height; y++ {
		for x := 0; x < level.Width; x++ {
			level.Cells[x][y] = 0
		}
	}

	rng := rand.New(rand.NewPCG(dungeonSeed, uint64(dungeonSeed/2+1)))

	for y := 0; y < level.Height; y++ {
		for x := 0; x < level.Width; x++ {
			if x == 0 || x == level.Width-1 || y == 0 || y == level.Height-1 {
				level.Cells[x][y] = 1 // Wall
			} else {
				if rng.IntN(100) < 20 { // 20% chance to place a wall
					level.Cells[x][y] = 1 // Wall
				} else {
					level.Cells[x][y] = 0 // Floor
				}
			}
		}
	}

	return &level

}
