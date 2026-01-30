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
) float64 {

	distance := 0.0

	for distance < drawer.MaxDist {
		x := player.X + math.Cos(rayAngle)*distance
		y := player.Y + math.Sin(rayAngle)*distance

		mapX := int(x)
		mapY := int(y)

		if mapY < 0 || mapY >= len(worldMap) ||
			mapX < 0 || mapX >= len(worldMap[0]) {
			return drawer.MaxDist
		}

		if worldMap[mapY][mapX] == 1 {
			return distance
		}

		distance += drawer.Step
	}

	return drawer.MaxDist
}

func DrawVerticalLine(
	screen *ebiten.Image,
	horizontalPosition int,
	height int,
	screenHeight int,
) {
	start := (screenHeight - height) / 2
	for y := 0; y < height; y++ {
		screen.Set(horizontalPosition, start+y, color.White)
	}
}
