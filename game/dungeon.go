package game

import (
	"fmt"
)

type DungeonLevel struct {
	Width  int
	Height int
	Cells  [][]uint16
	seed   uint64
}

func GenerateDungeon(dungeonSeed uint64, width int, height int) *DungeonLevel {

	level := DungeonLevel{
		Width:  64,
		Height: 64,
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

	for y := 0; y < level.Height; y++ {
		for x := 0; x < level.Width; x++ {
			fmt.Printf("%d", level.Cells[x][y])
		}
		fmt.Println()
	}

	fmt.Print("Generating dungeon with seed: ", dungeonSeed, "\n")
	return &level

}
