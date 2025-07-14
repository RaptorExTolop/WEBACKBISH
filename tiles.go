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
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
}

func (TS *TileSet) drawTiles() {
	for y, tiles := range TS.tiles {
		for x, tile := range tiles {
			if tile == 1 {
				rl.DrawRectangle(int32(x*int(TS.tileSize)), int32(y*int(TS.tileSize)), int32(TS.tileSize), int32(TS.tileSize), rl.Black)
			}
		}
	}
}
