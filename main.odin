package main

import "core:fmt"
import rl "vendor:raylib"

Window :: struct {
    Width, Height, TargtFPS: i32,
    Running: bool,
    Title: cstring,
    BkgColour: rl.Color
}

Ball :: struct {
    radius: f32,
    pos, dir: rl.Vector2, 
    speed: i32,
    colour: rl.Color
}

window: Window = Window{1280, 720, 60, true, "Game Game", rl.SKYBLUE}
ball: Ball = Ball{20, {f32(window.Width/2), f32(window.Height/2)}, {0, 0}, 7, rl.RAYWHITE}

main :: proc() {
    rl.InitWindow(window.Width, window.Height, window.Title)
    rl.SetTargetFPS(window.TargtFPS)

    defer rl.CloseWindow()
    
    for window.Running {
        update()
        draw()
    }
} 

update :: proc() {
    window.Running = !rl.WindowShouldClose()   

    ball.dir = {0, 0} 
    if (rl.IsKeyDown(.W)) {
        ball.dir[1] -= 1
    }
    if (rl.IsKeyDown(.S)) {
        ball.dir[1] += 1
    }
    if (rl.IsKeyDown(.A)) {
        ball.dir[0] -= 1
    }
    if (rl.IsKeyDown(.D)) {
        ball.dir[0] += 1
    }

    ball.dir = rl.Vector2Normalize(ball.dir)
    ball.pos += ball.dir * f32(ball.speed)
    //fmt.println(ball.dir)
}

draw :: proc() {
    rl.BeginDrawing()

    rl.ClearBackground(window.BkgColour)
    rl.DrawCircleV(ball.pos, ball.radius, ball.colour)

    rl.EndDrawing()
}