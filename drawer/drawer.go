package drawer

import (
	"image/color"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
	"soup/player"
)

type Drawer struct {
	Step    float64
	MaxDist float64
}

func (drawer *Drawer) CastRay(
	player *player.Player,
	rayAngle float64,
	worldMap [][]int,
	wallColors map[int]color.RGBA,
) (float64, color.Color) {
	var wallColor color.Color
	distance := 0.0

	for distance < drawer.MaxDist {
		x := player.X + math.Cos(rayAngle)*distance
		y := player.Y + math.Sin(rayAngle)*distance

		mapX := int(x)
		mapY := int(y)

		if mapY < 0 || mapY >= len(worldMap) ||
			mapX < 0 || mapX >= len(worldMap[0]) {
			return drawer.MaxDist, wallColor
		}
		
		mapPosition := worldMap[mapY][mapX]
		if mapPosition > 0 {
			return distance, wallColors[mapPosition]
		}

		distance += drawer.Step
	}

	return drawer.MaxDist, wallColor
}

func DrawVerticalLine(
	screen *ebiten.Image,
	horizontalPosition int,
	height int,
	screenHeight int,
	lineColor color.Color,
) {
	start := (screenHeight - height) / 2
	for y := 0; y < height; y++ {
		screen.Set(horizontalPosition, start+y, lineColor)
	}
}
