package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	Width, Height int32
	Running       bool
	Title         string
	BkgColour     rl.Color
	dt            float32
}

type Ball struct {
	radius   float32
	pos, dir rl.Vector2
	speed    float32
	colour   rl.Color
}

var window Window = Window{480, 360, true, "Game Game", rl.SkyBlue, 0}
var ball Ball = Ball{20, rl.Vector2{X: float32(window.Width / 2), Y: float32(window.Height / 2)}, rl.Vector2{X: 0, Y: 0}, 600, rl.RayWhite}
var tileSetOne TileSet

func init() {
	rl.InitWindow(window.Width, window.Height, window.Title)
	//rl.SetTargetFPS(window.TargtFPS)

	ball = Ball{20, rl.Vector2{X: float32(window.Width / 2), Y: float32(window.Height / 2)}, rl.Vector2{X: 0, Y: 0}, 600, rl.RayWhite}
	window = Window{1280, 720, true, "Game Game", rl.SkyBlue, 0}
	tileSetOne.loadTiles("loading file path", 32)
}

func main() {
	defer rl.CloseWindow()

	for window.Running {
		update()
		draw()
	}
}

func update() {
	window.Running = !rl.WindowShouldClose()

	window.dt = rl.GetFrameTime()
	//fmt.Println(window.dt)

	ball.dir = rl.Vector2{X: 0, Y: 0}
	if rl.IsKeyDown(rl.KeyW) {
		ball.dir.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		ball.dir.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		ball.dir.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		ball.dir.X += 1
	}

	ball.dir = rl.Vector2Normalize(ball.dir)
	ball.pos.X += ball.dir.X * ball.speed * window.dt
	ball.pos.Y += ball.dir.Y * ball.speed * window.dt
	//fmt.println(ball.dir)
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(window.BkgColour)
	rl.DrawFPS(0, 0)
	rl.DrawCircleV(ball.pos, ball.radius, ball.colour)
	tileSetOne.drawTiles()
	//rl.DrawRectangleV(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 32, Y: 32}, rl.Black)

	rl.EndDrawing()
}
