package main

import (
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"soup/drawer"
	"soup/game"
	"soup/player"
)

var worldMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 2, 2, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

var wallColors = map[int]color.RGBA{
	1: {R: 255, G: 165, B: 0, A: 255}, // Orange
	2: {R: 0, G: 0, B: 255, A: 255},   // Blue
}

func main() {
	game := &game.Game{
		Player: player.Player{
			X:     3.5,
			Y:     3.5,
			Angle: 0,
			Speed: 0.2,
			Rotation: 0.1,
		},
		Drawer: drawer.Drawer{
			Step:    0.05,
			MaxDist: 20,
		},
		WorldMap: worldMap,
		MapColors: wallColors,
		ScreenWidth: 640,
		ScreenHeight: 480,
	}

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

