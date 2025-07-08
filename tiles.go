package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSet struct {
	filePath string
	tiles    [][]float32
	tileSize float32
}

func (TS *TileSet) loadTiles(filePath string, tileSize float32) {
	TS.tileSize = tileSize
	TS.filePath = filePath
	fmt.Println(TS.filePath)
	TS.tiles = [][]float32{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
}

func (TS *TileSet) drawTiles() {
	tileSize := rl.Vector2{X: TS.tileSize, Y: TS.tileSize}
	for x, tiles := range TS.tiles {
		fmt.Println("X", x)
		for y := range tiles {
			tileLocation := rl.Vector2{X: float32(x) * TS.tileSize, Y: float32(y) * TS.tileSize}
			switch y {
			case 1:
				rl.DrawRectangleV(tileLocation, tileSize, rl.Color{R: 0, G: 0, B: 0, A: 255})
			case 0:
				rl.DrawRectangleV(tileLocation, tileSize, rl.Color{R: 0, G: 0, B: 0, A: 0})
			}
		}
	}
}
