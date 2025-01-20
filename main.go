package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "ProjectGalateus MVP")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Black)

		rl.EndDrawing()
	}
}
