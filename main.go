package main

import (
	"fmt"
	"time"
	"warphack/game"
	"warphack/utils"
)

func main() {
	fmt.Println("Hello, Adventurer!")
	randomSeed := generateTimeSeed()
	fmt.Printf("Random Seed: %08x\n", randomSeed)

	dungeonRng := utils.CreateRNG(randomSeed)
	for i := 0; i < 10; i++ {
		fmt.Printf("Rolling a 6-sided die: %d\n", utils.RollDice(20, dungeonRng))
	}
	game.GenerateDungeon(randomSeed, 64, 64)
	fmt.Println("\n Dungeon generated successfully!")
	fmt.Println("\n Enjoy your adventure!")

}
func generateTimeSeed() uint64 {
	return uint64(time.Now().UnixNano() & 0xFFFFFFFF)
}
