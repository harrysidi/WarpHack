package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	MapWidth  = 35
	MapHeight = 20
	TileSize  = 16
)

type Game struct {
	PlayerX, PlayerY float32
	GameMap          [MapWidth][MapHeight]int
	Seed             uint64
}

func NewGame(seed uint64) *Game {
	g := &Game{
		PlayerX: 1,
		PlayerY: 1,
		Seed:    seed,
	}

	// Generate initial map
	g.generateInitialMap()

	// Generate rooms
	rooms := GenerateLevel(seed, MapWidth, MapHeight)
	ApplyLevelToMap(&g.GameMap, rooms)

	return g
}

func ApplyLevelToMap(gameMap *[MapWidth][MapHeight]int, level *DungeonLevel) {
	for y := 0; y < level.Height; y++ {
		for x := 0; x < level.Width; x++ {
			if x < MapWidth && y < MapHeight {
				if level.Cells[x][y] == 1 {
					gameMap[x][y] = 1 // Wall
				} else {
					gameMap[x][y] = 0 // Floor
				}
			}
		}
	}
}

func (g *Game) generateInitialMap() {
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			if x == 0 || x == MapWidth-1 || y == 0 || y == MapHeight-1 {
				g.GameMap[x][y] = 1 // Wall
			} else {
				g.GameMap[x][y] = 0 // Floor
			}
		}
	}
}

func (g *Game) Update() error {
	newX, newY := g.PlayerX, g.PlayerY

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		newY--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		newY++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		newX--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		newX++
	}

	// Check bounds and walls
	if newX >= 0 && newX < MapWidth && newY >= 0 && newY < MapHeight {
		if g.GameMap[int(newX)][int(newY)] == 0 { // Only move if it's a floor tile
			g.PlayerX, g.PlayerY = newX, newY
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{25, 27, 8, 255})
	DrawMap(screen, g)
	DrawPlayer(screen, g.PlayerX, g.PlayerY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MapWidth * TileSize, MapHeight * TileSize
}
