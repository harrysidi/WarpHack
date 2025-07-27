package main

import (
	"log"

	"warphack/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame(12345) // Use a seed

	ebiten.SetWindowSize(game.MapWidth*game.TileSize, game.MapHeight*game.TileSize)
	ebiten.SetWindowTitle("WarpHack Dungeon")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
