package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth int32 = 1000
	screenHeight int32 = 480
)

var (
	running bool = true
)

func drawScene() {

}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()
	
}

func render() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		drawScene()

		rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

}

func quit() {
	defer rl.CloseWindow()

}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
