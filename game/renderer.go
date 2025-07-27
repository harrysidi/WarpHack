package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

// You'll need to move Game struct to game package or make fields public
func DrawMap(screen *ebiten.Image, g *Game) {
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			screenX := float32(x * TileSize)
			screenY := float32(y * TileSize)
			if g.GameMap[x][y] == 1 {
				vector.DrawFilledRect(screen, screenX, screenY, TileSize, TileSize, colornames.Darkgray, false)
			} else {
				vector.DrawFilledRect(screen, screenX, screenY, TileSize, TileSize, colornames.Lightgray, false)
			}
		}
	}
}

func DrawPlayer(screen *ebiten.Image, playerX, playerY float32) {
	playerScreenX := float32(playerX * TileSize)
	playerScreenY := float32(playerY * TileSize)
	vector.DrawFilledRect(screen, playerScreenX, playerScreenY, TileSize, TileSize, colornames.Red, false)
}
