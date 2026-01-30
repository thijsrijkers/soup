package game

import (
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"soup/drawer"
	"soup/player"
)

type Game struct {
	Player player.Player
	Drawer drawer.Drawer
	WorldMap [][]int
	MapColors map[int]color.RGBA
	ScreenWidth int
	ScreenHeight int
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		game.Player.X += math.Cos(game.Player.Angle) * game.Player.Speed
		game.Player.Y += math.Sin(game.Player.Angle) * game.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		game.Player.X -= math.Cos(game.Player.Angle) * game.Player.Speed
		game.Player.Y -= math.Sin(game.Player.Angle) * game.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.Angle -= game.Player.Rotation
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.Angle += game.Player.Rotation
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	fov := math.Pi / 3
	numRays := game.ScreenWidth

	for horizontalPosition := 0; horizontalPosition < numRays; horizontalPosition++ {
		rayAngle := game.Player.Angle - fov/2 + fov*float64(horizontalPosition)/float64(numRays)

		distance, wallColor := game.Drawer.CastRay(&game.Player, rayAngle, game.WorldMap, game.MapColors)
		distance *= math.Cos(game.Player.Angle - rayAngle)

		wallHeight := int(float64(game.ScreenHeight) / distance)

		drawer.DrawVerticalLine(screen, horizontalPosition, wallHeight, game.ScreenHeight, wallColor)
	}
}

func (game *Game) Layout(_, _ int) (int, int) {
	return game.ScreenWidth, game.ScreenHeight
}
