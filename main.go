package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenWidth = 800
	const screenHeight = 800
	rl.InitWindow(screenWidth, screenHeight, "snake-go")
	defer rl.CloseWindow()

	apple := NewApple(200, 200)
	snake := NewSnake(400, 400)
	camera := rl.NewCamera2D(
		rl.NewVector2(400+20.0, 400+20.0),
		rl.NewVector2(screenWidth/2.0, screenHeight/2.0),
		0.0, 1.0)

	var elapsedTime float32
	moveInterval := float32(0.2) // Move every 0.2 seconds
	isGameOver := false
	score := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		elapsedTime += rl.GetFrameTime()

		// Player movement
		if rl.IsKeyReleased(rl.KeyRight) && snake.Direction != Left {
			snake.SetDirectionRight()
		} else if rl.IsKeyReleased(rl.KeyLeft) && snake.Direction != Right {
			snake.SetDirectionLeft()
		} else if rl.IsKeyReleased(rl.KeyUp) && snake.Direction != Down {
			snake.SetDirectionUp()
		} else if rl.IsKeyReleased(rl.KeyDown) && snake.Direction != Up {
			snake.SetDirectionDown()
		}

		// Check if one second has passed
		if elapsedTime >= moveInterval {
			snake.Update(apple, &isGameOver, &score)
			elapsedTime = 0.0
		}

		// Draw
		rl.BeginDrawing()
		if isGameOver {
			rl.ClearBackground(rl.Black)
			rl.DrawText("Game Over!", 350, 400, 20, rl.White)
			rl.EndDrawing()
			continue
		}

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 10, 20, rl.Black)
		rl.BeginMode2D(camera)
		apple.Draw()
		snake.Draw()

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
